import { task } from 'hardhat/config';
import { eNetwork, FORK, getContract, waitForTx, ZERO_ADDRESS } from '../../helpers';
import {
  DELEGATION_CONTROLLER_IMPL_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
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

  const { address: poolControllerProxyAddress } = await hre.deployments.get(
    POOL_CONTROLLER_PROXY_ID
  );
  const { address: delegationControllerProxyAddress } = await hre.deployments.get(
    DELEGATION_CONTROLLER_PROXY_ID
  );
  const { address: slasherAddress } = await hre.deployments.get(SLASHER_PROXY_ID);

  const poolControllerImpl = await hre.deployments.deploy(POOL_CONTROLLER_IMPL_ID, {
    contract: 'PoolController',
    from: deployer,
    args: [delegationControllerProxyAddress, slasherAddress],
  });
  console.log(`[Deployment][INFO] PoolController impl deployed ${poolControllerImpl.address}`);

  const delegationControllerImpl = await hre.deployments.deploy(DELEGATION_CONTROLLER_IMPL_ID, {
    contract: 'DelegationController',
    from: deployer,
    args: [poolControllerProxyAddress, slasherAddress],
  });
  console.log(
    `[Deployment][INFO] DelegationController impl deployed ${delegationControllerImpl.address}`
  );

  // MultiSig call DelegationControllerProxy updateWrappedTokenGateway

  const proxyAdmin = await getContract(PROXY_ADMIN_ID);
  await waitForTx(
    await proxyAdmin.upgrade(
      poolControllerProxyAddress,
      (
        await hre.deployments.get(POOL_CONTROLLER_IMPL_ID)
      ).address
    )
  );
  console.log('PoolController upgrade successful!');

  await waitForTx(
    await proxyAdmin.upgrade(
      delegationControllerProxyAddress,
      (
        await hre.deployments.get(DELEGATION_CONTROLLER_IMPL_ID)
      ).address
    )
  );
  console.log('DelegationController upgrade successful!');
});
