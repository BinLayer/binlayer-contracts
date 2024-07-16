import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork, getContract, waitForTx } from '../../helpers';
import {
  DELEGATION_CONTROLLER_IMPL_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  EMPTY_CONTRANCT_ID,
  PAUSER_REGISTRY_ID,
  PROXY_ADMIN_ID,
  SLASHER_IMPL_ID,
  SLASHER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
  POOL_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { ethers } from 'ethers';
import { getParamPerNetwork, isProduction } from '../../helpers/config-helpers';
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

  const DelegationControllerProxyArtifact = await deployments.get(DELEGATION_CONTROLLER_PROXY_ID);
  const PoolControllerProxyArtifact = await deployments.get(POOL_CONTROLLER_PROXY_ID);
  const SlasherProxyArtifact = await deployments.get(SLASHER_PROXY_ID);
  const DelegationControllerImplArtifact = await deployments.get(DELEGATION_CONTROLLER_IMPL_ID);
  const PoolControllerImplArtifact = await deployments.get(POOL_CONTROLLER_IMPL_ID);
  const SlasherImplArtifact = await deployments.get(SLASHER_IMPL_ID);

  const ifaceDelegation = new ethers.utils.Interface(DelegationControllerImplArtifact.abi);
  const ifacePool = new ethers.utils.Interface(PoolControllerImplArtifact.abi);
  const ifaceSlasher = new ethers.utils.Interface(SlasherImplArtifact.abi);

  const proxyAdmin = await getContract(PROXY_ADMIN_ID);

  const owner = isProduction() ? getParamPerNetwork(Configs.Owner, network) : deployer;
  const minWithdrawalDelay = getParamPerNetwork(Configs.MinWithdrawalDelay, network);
  const configs = getParamPerNetwork(Configs.PoolConfigs, network);
  const pools = [];
  const withdrawalDelays = [];
  const thirdPartyTransfersForbiddenValues = [];
  for (let key in configs) {
    const config = configs[key];
    const { address: poolAddress } = await deployments.get(`${key}${POOL_PROXY_ID}`);
    pools.push(poolAddress);
    withdrawalDelays.push(config.withdrawalDelay);
    thirdPartyTransfersForbiddenValues.push(false);
  }

  if (pools.length == 0 || pools.length != withdrawalDelays.length) {
    throw '[Deployment][Error] pools withdrawalDeplays not match';
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
        pools,
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
      await delegationControllerInstance.updateWrappedTokenGateway(
        wrappedTokenGatewayArtifact.address
      )
    );
    console.log(`[Deployment][INFO] Config gateway ${pools}`);
  }

  const whiteLister = isProduction() ? getParamPerNetwork(Configs.WhiteLister, network) : deployer;
  const poolControllerPausedStatus = getParamPerNetwork(
    Configs.PoolControllerPausedStatus,
    network
  );

  await waitForTx(
    await proxyAdmin.upgradeAndCall(
      PoolControllerProxyArtifact.address,
      PoolControllerImplArtifact.address,
      ifacePool.encodeFunctionData('initialize', [
        deployer, // need transfer ownership
        deployer, // need transfer whitelister
        pauserRegistryAddress,
        poolControllerPausedStatus,
      ])
    )
  );
  console.log('[Deployment][INFO] Upgraded PoolController.sol');

  const poolControllerInstance = await hre.ethers.getContractAt(
    'PoolController',
    PoolControllerProxyArtifact.address
  );
  await waitForTx(
    await poolControllerInstance.addPoolsToDepositWhitelist(
      pools,
      thirdPartyTransfersForbiddenValues
    )
  );
  console.log(`[Deployment][INFO] Config pool manager white list pool ${pools}`);

  await waitForTx(await poolControllerInstance.setPoolWhitelister(whiteLister!));
  console.log(`[Deployment][INFO] Transfered pool white lister to ${whiteLister}`);

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
