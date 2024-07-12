import { task } from 'hardhat/config';
import { eNetwork, FORK, getContract, waitForTx, ZERO_ADDRESS } from '../../helpers';
import {
  DELEGATION_CONTROLLER_IMPL_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  DELEGATION_MANAGER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  STRATEGY_MANAGER_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';

task(`upgrade-core`, `Upgrade PoolController & DelegationController`).setAction(async (_, hre) => {
  if (!hre.network.config.chainId) {
    throw new Error('INVALID_CHAIN_ID');
  }

  await hre.run('compile');

  console.log(`\n- Core Upgrading`);

  const { deployer } = await hre.getNamedAccounts();
  const network = (FORK ? FORK : hre.network.name) as eNetwork;

  const { address: strategyManagerProxyAddress } = await hre.deployments.get(
    STRATEGY_MANAGER_PROXY_ID
  );
  const { address: delegationManagerProxyAddress } = await hre.deployments.get(
    DELEGATION_MANAGER_PROXY_ID
  );
  const { address: slasherAddress } = await hre.deployments.get(SLASHER_PROXY_ID);

  const poolControllerImpl = await hre.deployments.deploy(POOL_CONTROLLER_IMPL_ID, {
    contract: 'PoolController',
    from: deployer,
    args: [delegationManagerProxyAddress, slasherAddress],
  });
  console.log(`[Deployment][INFO] PoolController impl deployed ${poolControllerImpl.address}`);

  const delegationControllerImpl = await hre.deployments.deploy(DELEGATION_CONTROLLER_IMPL_ID, {
    contract: 'DelegationController',
    from: deployer,
    args: [strategyManagerProxyAddress, slasherAddress],
  });
  console.log(
    `[Deployment][INFO] DelegationController impl deployed ${delegationControllerImpl.address}`
  );

  // MultiSig call DelegationManagerProxy updateWrappedTokenGateway

  const proxyAdmin = await getContract(PROXY_ADMIN_ID);
  await waitForTx(
    await proxyAdmin.upgrade(
      strategyManagerProxyAddress,
      (
        await hre.deployments.get(POOL_CONTROLLER_IMPL_ID)
      ).address
    )
  );
  console.log('StrategyManager upgrade successful!');

  await waitForTx(
    await proxyAdmin.upgrade(
      delegationManagerProxyAddress,
      (
        await hre.deployments.get(DELEGATION_CONTROLLER_IMPL_ID)
      ).address
    )
  );
  console.log('DelegationManager upgrade successful!');
});
