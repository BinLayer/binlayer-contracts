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
  STRATEGY_MANAGER_IMPL_ID,
  STRATEGY_MANAGER_PROXY_ID,
  STRATEGY_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { parseUnits } from 'ethers/lib/utils';

task(`config-strategy`, `Config strategy`).setAction(async (_, hre) => {
  if (!hre.network.config.chainId) {
    throw new Error('INVALID_CHAIN_ID');
  }

  const StrategyManagerProxyArtifact = await hre.deployments.get(STRATEGY_MANAGER_PROXY_ID);
  const strategyManagerInstance = await hre.ethers.getContractAt(
    'StrategyManager',
    StrategyManagerProxyArtifact.address
  );
  const DelegationControllerProxyArtifact = await hre.deployments.get(DELEGATION_CONTROLLER_PROXY_ID);
  const delegationControllerInstance = await hre.ethers.getContractAt(
    'DelegationController',
    DelegationControllerProxyArtifact.address
  );
  const StrategyProxyArtifact = await hre.deployments.get(`slisBNB${STRATEGY_PROXY_ID}`);
  const strategyInstance = await hre.ethers.getContractAt(
    'StrategyBaseTVLLimits',
    StrategyProxyArtifact.address
  );

  await waitForTx(
    await strategyManagerInstance.addStrategiesToDepositWhitelist(
      [StrategyProxyArtifact.address],
      [false]
    )
  );
  await waitForTx(
    await delegationControllerInstance.setStrategyWithdrawalDelay(
      [StrategyProxyArtifact.address],
      [180]
    )
  );

  const pauserRegistry = (await getContract(PAUSER_REGISTRY_ID)) as PauserRegistry;
  console.log(await strategyInstance.totalShares());
  console.log(await pauserRegistry.unpauser());
  console.log(await strategyInstance.getTVLLimits());
  console.log(await strategyInstance.underlyingToken());
  console.log(
    await strategyManagerInstance.strategyIsWhitelistedForDeposit(strategyInstance.address)
  );
  console.log(await delegationControllerInstance.strategyWithdrawalDelay(strategyInstance.address));
  console.log(await delegationControllerInstance.getWithdrawalDelay([strategyInstance.address]));
  console.log(await delegationControllerInstance.minWithdrawalDelay());
});
