import { task } from 'hardhat/config';
import {
  DelegationController,
  eNetwork,
  FORK,
  getContract,
  PauserRegistry,
  waitForTx,
  WrappedTokenGateway,
  ZERO_ADDRESS,
} from '../../helpers';
import {
  DELEGATION_CONTROLLER_IMPL_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  PAUSER_REGISTRY_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
  POOL_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { parseUnits } from 'ethers/lib/utils';

task(`config-pool`, `Config pool`).setAction(async (_, hre) => {
  if (!hre.network.config.chainId) {
    throw new Error('INVALID_CHAIN_ID');
  }

  const PoolManagerProxyArtifact = await hre.deployments.get(POOL_CONTROLLER_PROXY_ID);
  const poolManagerInstance = await hre.ethers.getContractAt(
    'PoolManager',
    PoolManagerProxyArtifact.address
  );
  const DelegationControllerProxyArtifact = await hre.deployments.get(DELEGATION_CONTROLLER_PROXY_ID);
  const delegationControllerInstance = await hre.ethers.getContractAt(
    'DelegationController',
    DelegationControllerProxyArtifact.address
  );
  const PoolProxyArtifact = await hre.deployments.get(`slisBNB${POOL_PROXY_ID}`);
  const poolInstance = await hre.ethers.getContractAt(
    'PoolBaseTVLLimits',
    PoolProxyArtifact.address
  );

  await waitForTx(
    await poolManagerInstance.addPoolsToDepositWhitelist(
      [PoolProxyArtifact.address],
      [false]
    )
  );
  await waitForTx(
    await delegationControllerInstance.setPoolWithdrawalDelay(
      [PoolProxyArtifact.address],
      [180]
    )
  );

  const pauserRegistry = (await getContract(PAUSER_REGISTRY_ID)) as PauserRegistry;
  console.log(await poolInstance.totalShares());
  console.log(await pauserRegistry.unpauser());
  console.log(await poolInstance.getTVLLimits());
  console.log(await poolInstance.underlyingToken());
  console.log(
    await poolManagerInstance.poolIsWhitelistedForDeposit(poolInstance.address)
  );
  console.log(await delegationControllerInstance.poolWithdrawalDelay(poolInstance.address));
  console.log(await delegationControllerInstance.getWithdrawalDelay([poolInstance.address]));
  console.log(await delegationControllerInstance.minWithdrawalDelay());
});
