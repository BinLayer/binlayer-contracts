// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import '@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol';
import '@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol';
import '@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol';
import '@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol';

import '../contracts/core/PoolController.sol';
import '../contracts/core/Slasher.sol';
import '../contracts/core/DelegationController.sol';
import '../contracts/core/AVSDirectory.sol';
import '../contracts/core/RewardsCoordinator.sol';

import '../contracts/pools/PoolBaseTVLLimits.sol';

import '../contracts/permissions/PauserRegistry.sol';

import '../contracts/mocks/EmptyContract.sol';

import 'forge-std/Script.sol';
import 'forge-std/Test.sol';

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/devnet/M2_Deploy_From_Scratch.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
contract BinLayerDeployer is Script, Test {
  // struct used to encode token info in config file
  struct PoolConfig {
    uint256 maxDeposits;
    uint256 maxPerDeposit;
    address tokenAddress;
    string tokenSymbol;
  }

  string public deployConfigPath;

  // BinLayer Contracts
  ProxyAdmin public binlayerProxyAdmin;
  PauserRegistry public binlayerPauserReg;
  Slasher public slasher;
  Slasher public slasherImplementation;
  DelegationController public delegation;
  DelegationController public delegationImplementation;
  PoolController public poolController;
  PoolController public poolControllerImplementation;
  AVSDirectory public avsDirectory;
  AVSDirectory public avsDirectoryImplementation;
  RewardsCoordinator public rewardsCoordinator;
  RewardsCoordinator public rewardsCoordinatorImplementation;
  PoolBase public basePoolImplementation;

  EmptyContract public emptyContract;

  address executorMultisig;
  address operationsMultisig;
  address pauserMultisig;

  // pools deployed
  PoolBaseTVLLimits[] public deployedPoolArray;

  // OTHER DEPLOYMENT PARAMETERS
  uint256 POOL_CONTROLLER_INIT_PAUSED_STATUS;
  uint256 SLASHER_INIT_PAUSED_STATUS;
  uint256 DELEGATION_INIT_PAUSED_STATUS;
  uint256 REWARDS_COORDINATOR_INIT_PAUSED_STATUS;

  // RewardsCoordinator
  uint32 REWARDS_COORDINATOR_MAX_REWARDS_DURATION;
  uint32 REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH;
  uint32 REWARDS_COORDINATOR_MAX_FUTURE_LENGTH;
  uint32 REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP;
  address REWARDS_COORDINATOR_UPDATER;
  uint32 REWARDS_COORDINATOR_ACTIVATION_DELAY;
  uint32 REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS;
  uint32 REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS;

  uint256 DELEGATION_WITHDRAWAL_DELAY_TIMESTAMP;

  function run(string memory configFileName) external {
    // read and log the chainID
    uint256 chainId = block.chainid;
    emit log_named_uint('You are deploying on ChainID', chainId);

    // READ JSON CONFIG DATA
    deployConfigPath = string(bytes(string.concat('script/configs/', configFileName)));
    string memory config_data = vm.readFile(deployConfigPath);
    // bytes memory parsedData = vm.parseJson(config_data);

    POOL_CONTROLLER_INIT_PAUSED_STATUS = stdJson.readUint(config_data, '.poolController.init_paused_status');
    SLASHER_INIT_PAUSED_STATUS = stdJson.readUint(config_data, '.slasher.init_paused_status');
    DELEGATION_INIT_PAUSED_STATUS = stdJson.readUint(config_data, '.delegation.init_paused_status');
    DELEGATION_WITHDRAWAL_DELAY_TIMESTAMP = stdJson.readUint(config_data, '.delegation.init_withdrawal_delay_blocks');
    REWARDS_COORDINATOR_INIT_PAUSED_STATUS = stdJson.readUint(config_data, '.rewardsCoordinator.init_paused_status');
    REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = uint32(
      stdJson.readUint(config_data, '.rewardsCoordinator.CALCULATION_INTERVAL_SECONDS')
    );
    REWARDS_COORDINATOR_MAX_REWARDS_DURATION = uint32(stdJson.readUint(config_data, '.rewardsCoordinator.MAX_REWARDS_DURATION'));
    REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH = uint32(stdJson.readUint(config_data, '.rewardsCoordinator.MAX_RETROACTIVE_LENGTH'));
    REWARDS_COORDINATOR_MAX_FUTURE_LENGTH = uint32(stdJson.readUint(config_data, '.rewardsCoordinator.MAX_FUTURE_LENGTH'));
    REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP = uint32(stdJson.readUint(config_data, '.rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP'));
    REWARDS_COORDINATOR_UPDATER = stdJson.readAddress(config_data, '.rewardsCoordinator.rewards_updater_address');
    REWARDS_COORDINATOR_ACTIVATION_DELAY = uint32(stdJson.readUint(config_data, '.rewardsCoordinator.activation_delay'));
    REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS = uint32(
      stdJson.readUint(config_data, '.rewardsCoordinator.calculation_interval_seconds')
    );
    REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS = uint32(
      stdJson.readUint(config_data, '.rewardsCoordinator.global_operator_commission_bips')
    );

    // tokens to deploy pools for
    PoolConfig[] memory poolConfigs;

    executorMultisig = stdJson.readAddress(config_data, '.multisig_addresses.executorMultisig');
    operationsMultisig = stdJson.readAddress(config_data, '.multisig_addresses.operationsMultisig');
    pauserMultisig = stdJson.readAddress(config_data, '.multisig_addresses.pauserMultisig');
    // load token list
    bytes memory poolConfigsRaw = stdJson.parseRaw(config_data, '.pools');
    poolConfigs = abi.decode(poolConfigsRaw, (PoolConfig[]));

    require(executorMultisig != address(0), 'executorMultisig address not configured correctly!');
    require(operationsMultisig != address(0), 'operationsMultisig address not configured correctly!');

    // START RECORDING TRANSACTIONS FOR DEPLOYMENT
    vm.startBroadcast();

    // deploy proxy admin for ability to upgrade proxy contracts
    binlayerProxyAdmin = new ProxyAdmin();

    //deploy pauser registry
    {
      address[] memory pausers = new address[](3);
      pausers[0] = executorMultisig;
      pausers[1] = operationsMultisig;
      pausers[2] = pauserMultisig;
      binlayerPauserReg = new PauserRegistry(pausers, executorMultisig);
    }

    /**
     * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
     * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
     */
    emptyContract = new EmptyContract();
    delegation = DelegationController(address(new TransparentUpgradeableProxy(address(emptyContract), address(binlayerProxyAdmin), '')));
    poolController = PoolController(address(new TransparentUpgradeableProxy(address(emptyContract), address(binlayerProxyAdmin), '')));
    avsDirectory = AVSDirectory(address(new TransparentUpgradeableProxy(address(emptyContract), address(binlayerProxyAdmin), '')));
    slasher = Slasher(address(new TransparentUpgradeableProxy(address(emptyContract), address(binlayerProxyAdmin), '')));
    rewardsCoordinator = RewardsCoordinator(
      address(new TransparentUpgradeableProxy(address(emptyContract), address(binlayerProxyAdmin), ''))
    );

    // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
    delegationImplementation = new DelegationController(poolController, slasher);
    poolControllerImplementation = new PoolController(delegation, slasher);
    avsDirectoryImplementation = new AVSDirectory(delegation);
    slasherImplementation = new Slasher(poolController, delegation);
    rewardsCoordinatorImplementation = new RewardsCoordinator(
      delegation,
      poolController,
      REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
      REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
      REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
      REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
      REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
    );

    // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
    {
      IPool[] memory _pools;
      uint256[] memory _withdrawalDelayBlocks;
      binlayerProxyAdmin.upgradeAndCall(
        ITransparentUpgradeableProxy(payable(address(delegation))),
        address(delegationImplementation),
        abi.encodeWithSelector(
          DelegationController.initialize.selector,
          executorMultisig,
          binlayerPauserReg,
          DELEGATION_INIT_PAUSED_STATUS,
          DELEGATION_WITHDRAWAL_DELAY_TIMESTAMP,
          _pools,
          _withdrawalDelayBlocks
        )
      );
    }
    binlayerProxyAdmin.upgradeAndCall(
      ITransparentUpgradeableProxy(payable(address(poolController))),
      address(poolControllerImplementation),
      abi.encodeWithSelector(
        PoolController.initialize.selector,
        executorMultisig,
        operationsMultisig,
        binlayerPauserReg,
        POOL_CONTROLLER_INIT_PAUSED_STATUS
      )
    );
    binlayerProxyAdmin.upgradeAndCall(
      ITransparentUpgradeableProxy(payable(address(slasher))),
      address(slasherImplementation),
      abi.encodeWithSelector(Slasher.initialize.selector, executorMultisig, binlayerPauserReg, SLASHER_INIT_PAUSED_STATUS)
    );
    binlayerProxyAdmin.upgradeAndCall(
      ITransparentUpgradeableProxy(payable(address(avsDirectory))),
      address(avsDirectoryImplementation),
      abi.encodeWithSelector(AVSDirectory.initialize.selector, executorMultisig, binlayerPauserReg, 0)
    );
    binlayerProxyAdmin.upgradeAndCall(
      ITransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
      address(rewardsCoordinatorImplementation),
      abi.encodeWithSelector(
        RewardsCoordinator.initialize.selector,
        executorMultisig,
        binlayerPauserReg,
        REWARDS_COORDINATOR_INIT_PAUSED_STATUS,
        REWARDS_COORDINATOR_UPDATER,
        REWARDS_COORDINATOR_ACTIVATION_DELAY,
        REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS
      )
    );

    // deploy PoolBaseTVLLimits contract implementation
    basePoolImplementation = new PoolBaseTVLLimits(poolController);
    // create upgradeable proxies that each point to the implementation and initialize them
    for (uint256 i = 0; i < poolConfigs.length; ++i) {
      deployedPoolArray.push(
        PoolBaseTVLLimits(
          address(
            new TransparentUpgradeableProxy(
              address(basePoolImplementation),
              address(binlayerProxyAdmin),
              abi.encodeWithSelector(
                PoolBaseTVLLimits.initialize.selector,
                poolConfigs[i].maxPerDeposit,
                poolConfigs[i].maxDeposits,
                IERC20(poolConfigs[i].tokenAddress),
                binlayerPauserReg
              )
            )
          )
        )
      );
    }

    binlayerProxyAdmin.transferOwnership(executorMultisig);

    // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
    vm.stopBroadcast();

    // CHECK CORRECTNESS OF DEPLOYMENT
    _verifyContractsPointAtOneAnother(
      delegationImplementation,
      poolControllerImplementation,
      slasherImplementation,
      rewardsCoordinatorImplementation
    );
    _verifyContractsPointAtOneAnother(delegation, poolController, slasher, rewardsCoordinator);
    _verifyImplementationsSetCorrectly();
    _verifyInitialOwners();
    _checkPauserInitializations();
    _verifyInitializationParams();

    // WRITE JSON DATA
    string memory parent_object = 'parent object';

    string memory deployed_pools = 'pools';
    for (uint256 i = 0; i < poolConfigs.length; ++i) {
      vm.serializeAddress(deployed_pools, poolConfigs[i].tokenSymbol, address(deployedPoolArray[i]));
    }
    string memory deployed_pools_output = poolConfigs.length == 0
      ? ''
      : vm.serializeAddress(
        deployed_pools,
        poolConfigs[poolConfigs.length - 1].tokenSymbol,
        address(deployedPoolArray[poolConfigs.length - 1])
      );

    string memory deployed_addresses = 'addresses';
    vm.serializeAddress(deployed_addresses, 'binlayerProxyAdmin', address(binlayerProxyAdmin));
    vm.serializeAddress(deployed_addresses, 'binlayerPauserReg', address(binlayerPauserReg));
    vm.serializeAddress(deployed_addresses, 'slasher', address(slasher));
    vm.serializeAddress(deployed_addresses, 'slasherImplementation', address(slasherImplementation));
    vm.serializeAddress(deployed_addresses, 'delegation', address(delegation));
    vm.serializeAddress(deployed_addresses, 'delegationImplementation', address(delegationImplementation));
    vm.serializeAddress(deployed_addresses, 'avsDirectory', address(avsDirectory));
    vm.serializeAddress(deployed_addresses, 'avsDirectoryImplementation', address(avsDirectoryImplementation));
    vm.serializeAddress(deployed_addresses, 'poolController', address(poolController));
    vm.serializeAddress(deployed_addresses, 'poolControllerImplementation', address(poolControllerImplementation));
    vm.serializeAddress(deployed_addresses, 'rewardsCoordinator', address(rewardsCoordinator));
    vm.serializeAddress(deployed_addresses, 'rewardsCoordinatorImplementation', address(rewardsCoordinatorImplementation));
    vm.serializeAddress(deployed_addresses, 'basePoolImplementation', address(basePoolImplementation));
    vm.serializeAddress(deployed_addresses, 'emptyContract', address(emptyContract));
    string memory deployed_addresses_output = vm.serializeString(deployed_addresses, 'pools', deployed_pools_output);

    string memory parameters = 'parameters';
    vm.serializeAddress(parameters, 'executorMultisig', executorMultisig);
    string memory parameters_output = vm.serializeAddress(parameters, 'operationsMultisig', operationsMultisig);

    string memory chain_info = 'chainInfo';
    vm.serializeUint(chain_info, 'deploymentBlock', block.number);
    string memory chain_info_output = vm.serializeUint(chain_info, 'chainId', chainId);

    // serialize all the data
    vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
    vm.serializeString(parent_object, chain_info, chain_info_output);
    string memory finalJson = vm.serializeString(parent_object, parameters, parameters_output);
    // TODO: should output to different file depending on configFile passed to run()
    //       so that we don't override mainnet output by deploying to goerli for eg.
    vm.writeJson(finalJson, 'script/output/binlayer_deployment_data.json');
  }

  function _verifyContractsPointAtOneAnother(
    DelegationController delegationContract,
    PoolController poolControllerContract,
    Slasher /*slasherContract*/,
    RewardsCoordinator rewardsCoordinatorContract
  ) internal view {
    require(delegationContract.slasher() == slasher, 'delegation: slasher address not set correctly');
    require(delegationContract.poolController() == poolController, 'delegation: poolController address not set correctly');

    require(poolControllerContract.slasher() == slasher, 'poolController: slasher address not set correctly');
    require(poolControllerContract.delegation() == delegation, 'poolController: delegation address not set correctly');
    require(rewardsCoordinatorContract.delegationController() == delegation, 'rewardsCoordinator: delegation address not set correctly');

    require(rewardsCoordinatorContract.poolController() == poolController, 'rewardsCoordinator: poolController address not set correctly');
  }

  function _verifyImplementationsSetCorrectly() internal view {
    require(
      binlayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(delegation)))) ==
        address(delegationImplementation),
      'delegation: implementation set incorrectly'
    );
    require(
      binlayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(poolController)))) ==
        address(poolControllerImplementation),
      'poolController: implementation set incorrectly'
    );
    require(
      binlayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(slasher)))) == address(slasherImplementation),
      'slasher: implementation set incorrectly'
    );
    require(
      binlayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(rewardsCoordinator)))) ==
        address(rewardsCoordinatorImplementation),
      'rewardsCoordinator: implementation set incorrectly'
    );

    for (uint256 i = 0; i < deployedPoolArray.length; ++i) {
      require(
        binlayerProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(payable(address(deployedPoolArray[i])))) ==
          address(basePoolImplementation),
        'pool: implementation set incorrectly'
      );
    }
  }

  function _verifyInitialOwners() internal view {
    require(poolController.owner() == executorMultisig, 'poolController: owner not set correctly');
    require(delegation.owner() == executorMultisig, 'delegation: owner not set correctly');
    // removing slasher requirements because there is no slasher as part of m2-mainnet release
    // require(slasher.owner() == executorMultisig, "slasher: owner not set correctly");

    require(binlayerProxyAdmin.owner() == executorMultisig, 'binlayerProxyAdmin: owner not set correctly');
  }

  function _checkPauserInitializations() internal view {
    require(delegation.pauserRegistry() == binlayerPauserReg, 'delegation: pauser registry not set correctly');
    require(poolController.pauserRegistry() == binlayerPauserReg, 'poolController: pauser registry not set correctly');
    // removing slasher requirements because there is no slasher as part of m2-mainnet release
    // require(slasher.pauserRegistry() == binlayerPauserReg, "slasher: pauser registry not set correctly");
    require(rewardsCoordinator.pauserRegistry() == binlayerPauserReg, 'rewardsCoordinator: pauser registry not set correctly');
    require(binlayerPauserReg.isPauser(operationsMultisig), 'pauserRegistry: operationsMultisig is not pauser');
    require(binlayerPauserReg.isPauser(executorMultisig), 'pauserRegistry: executorMultisig is not pauser');
    require(binlayerPauserReg.isPauser(pauserMultisig), 'pauserRegistry: pauserMultisig is not pauser');
    require(binlayerPauserReg.unpauser() == executorMultisig, 'pauserRegistry: unpauser not set correctly');

    for (uint256 i = 0; i < deployedPoolArray.length; ++i) {
      require(deployedPoolArray[i].pauserRegistry() == binlayerPauserReg, 'PoolBaseTVLLimits: pauser registry not set correctly');
      require(deployedPoolArray[i].paused() == 0, 'PoolBaseTVLLimits: init paused status set incorrectly');
    }

    // // pause *nothing*
    // uint256 POOL_CONTROLLER_INIT_PAUSED_STATUS = 0;
    // // pause *everything*
    // uint256 SLASHER_INIT_PAUSED_STATUS = type(uint256).max;
    // // pause *everything*
    // uint256 DELEGATION_INIT_PAUSED_STATUS = type(uint256).max;
    // // pause *nothing*
    // uint256 DELAYED_WITHDRAWAL_ROUTER_INIT_PAUSED_STATUS = 0;
    // require(poolController.paused() == 0, "poolController: init paused status set incorrectly");
    // require(slasher.paused() == type(uint256).max, "slasher: init paused status set incorrectly");
    // require(delegation.paused() == type(uint256).max, "delegation: init paused status set incorrectly");
  }

  function _verifyInitializationParams() internal {
    // // one week in blocks -- 50400
    // uint32 POOL_CONTROLLER_INIT_WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    // uint32 DELAYED_WITHDRAWAL_ROUTER_INIT_WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    // require(poolController.withdrawalDelayBlocks() == 7 days / 12 seconds,
    //     "poolController: withdrawalDelayBlocks initialized incorrectly");
    // require(delayedWithdrawalRouter.withdrawalDelayBlocks() == 7 days / 12 seconds,
    //     "delayedWithdrawalRouter: withdrawalDelayBlocks initialized incorrectly");
    // uint256 MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 32 ether;
    require(poolController.poolWhitelister() == operationsMultisig, 'poolController: poolWhitelister address not set correctly');

    require(basePoolImplementation.poolController() == poolController, 'basePoolImplementation: poolController set incorrectly');

    string memory config_data = vm.readFile(deployConfigPath);
    for (uint i = 0; i < deployedPoolArray.length; i++) {
      uint256 maxPerDeposit = stdJson.readUint(config_data, string.concat('.pools[', vm.toString(i), '].max_per_deposit'));
      uint256 maxDeposits = stdJson.readUint(config_data, string.concat('.pools[', vm.toString(i), '].max_deposits'));
      (uint256 setMaxPerDeposit, uint256 setMaxDeposits) = deployedPoolArray[i].getTVLLimits();
      require(setMaxPerDeposit == maxPerDeposit, 'setMaxPerDeposit not set correctly');
      require(setMaxDeposits == maxDeposits, 'setMaxDeposits not set correctly');
    }
  }
}
