import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork, getContract, StrategyManager, waitForTx } from '../../helpers';
import {
  DELEGATION_MANAGER_IMPL_ID,
  DELEGATION_MANAGER_PROXY_ID,
  EMPTY_CONTRANCT_ID,
  PAUSER_REGISTRY_ID,
  PROXY_ADMIN_ID,
  SLASHER_IMPL_ID,
  SLASHER_PROXY_ID,
  STRATEGY_MANAGER_IMPL_ID,
  STRATEGY_MANAGER_PROXY_ID,
  STRATEGY_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { ethers } from 'ethers';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { Configs } from '../../helpers/config';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const network = (process.env.FORK ? process.env.FORK : hre.network.name) as eNetwork;

  const { address: pauserRegistryAddress } = await deployments.get(PAUSER_REGISTRY_ID);

  const DelegationControllerProxyArtifact = await deployments.get(DELEGATION_MANAGER_PROXY_ID);
  const StrategyManagerProxyArtifact = await deployments.get(STRATEGY_MANAGER_PROXY_ID);
  const SlasherProxyArtifact = await deployments.get(SLASHER_PROXY_ID);
  const DelegationControllerImplArtifact = await deployments.get(DELEGATION_MANAGER_IMPL_ID);
  const StrategyManagerImplArtifact = await deployments.get(STRATEGY_MANAGER_IMPL_ID);
  const SlasherImplArtifact = await deployments.get(SLASHER_IMPL_ID);

  const ifaceDelegation = new ethers.utils.Interface(DelegationControllerImplArtifact.abi);
  const ifaceStrategy = new ethers.utils.Interface(StrategyManagerImplArtifact.abi);
  const ifaceSlasher = new ethers.utils.Interface(SlasherImplArtifact.abi);

  const proxyAdmin = await getContract(PROXY_ADMIN_ID);

  const owner = getParamPerNetwork(Configs.Owner, network);
  const minWithdrawalDelay = getParamPerNetwork(Configs.MinWithdrawalDelay, network);
  const configs = getParamPerNetwork(Configs.StrategyConfigs, network);
  const strategies = [];
  const withdrawalDelays = [];
  const thirdPartyTransfersForbiddenValues = [];
  for (let key in configs) {
    const config = configs[key];
    const { address: strategyAddress } = await deployments.get(`${key}${STRATEGY_PROXY_ID}`);
    strategies.push(strategyAddress);
    withdrawalDelays.push(config.withdrawalDelay);
    thirdPartyTransfersForbiddenValues.push(false);
  }

  if (strategies.length == 0 || strategies.length != withdrawalDelays.length) {
    throw '[Deployment][Error] strategies withdrawalDeplays not match';
  }

  const delegationControllerPausedStatus = getParamPerNetwork(
    Configs.DelegationControllerPausedStatus,
    network
  );

  const wrappedTokenGatewayArtifact = await deployments.getOrNull(WRAPPED_TOKEN_GATEWAY_ID);

  await waitForTx(
    await proxyAdmin.upgradeAndCall(
      DelegationControllerProxyArtifact.address,
      DelegationControllerImplArtifact.address,
      ifaceDelegation.encodeFunctionData('initialize', [
        wrappedTokenGatewayArtifact == null ? owner : deployer,
        pauserRegistryAddress,
        delegationControllerPausedStatus,
        minWithdrawalDelay,
        strategies,
        withdrawalDelays,
      ])
    )
  );
  console.log('[Deployment][INFO] Upgraded DelegationController');

  const delegationControllerInstance = await hre.ethers.getContractAt(
    'DelegationController',
    DelegationControllerProxyArtifact.address
  );

  if (wrappedTokenGatewayArtifact) {
    await waitForTx(
      await delegationControllerInstance.updateWrappedTokenGateway(wrappedTokenGatewayArtifact.address)
    );
    console.log(`[Deployment][INFO] Config gateway ${strategies}`);
  }

  const whiteLister = getParamPerNetwork(Configs.WhiteLister, network);
  const strategyManagerPausedStatus = getParamPerNetwork(
    Configs.StrategyManagerPausedStatus,
    network
  );

  await waitForTx(
    await proxyAdmin.upgradeAndCall(
      StrategyManagerProxyArtifact.address,
      StrategyManagerImplArtifact.address,
      ifaceStrategy.encodeFunctionData('initialize', [
        deployer, // need transfer ownership
        deployer, // need transfer whitelister
        pauserRegistryAddress,
        strategyManagerPausedStatus,
      ])
    )
  );
  console.log('[Deployment][INFO] Upgraded StrategyManager');

  const strategyManagerInstance = await hre.ethers.getContractAt(
    'StrategyManager',
    StrategyManagerProxyArtifact.address
  );
  await waitForTx(
    await strategyManagerInstance.addStrategiesToDepositWhitelist(
      strategies,
      thirdPartyTransfersForbiddenValues
    )
  );
  console.log(`[Deployment][INFO] Config strategy manager white list strategy ${strategies}`);

  await waitForTx(await strategyManagerInstance.setStrategyWhitelister(whiteLister!));
  console.log(`[Deployment][INFO] Transfered strategy white lister to ${whiteLister}`);

  const slasherPausedStatus = getParamPerNetwork(Configs.SlasherPausedStatus, network);

  await waitForTx(
    await proxyAdmin.upgradeAndCall(
      SlasherProxyArtifact.address,
      SlasherImplArtifact.address,
      ifaceSlasher.encodeFunctionData('initialize', [
        owner,
        pauserRegistryAddress,
        slasherPausedStatus,
      ])
    )
  );
  console.log('[Deployment][INFO] Upgraded Slasher');

  return true;
};

func.id = 'CoreInitialize';
func.tags = ['core'];

export default func;
