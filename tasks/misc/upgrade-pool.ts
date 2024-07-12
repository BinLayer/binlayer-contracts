import { task } from 'hardhat/config';
import { eNetwork, FORK, getContract, waitForTx, ZERO_ADDRESS } from '../../helpers';
import {
  DELEGATION_MANAGER_PROXY_ID,
  POOL_IMPL_ID,
  PROXY_ADMIN_ID,
  STRATEGY_MANAGER_PROXY_ID,
  STRATEGY_PROXY_ID,
} from '../../helpers/deploy-ids';

task(`upgrade-pool`, `Upgrade Pool`)
  .addParam('symbol', 'Pool underlying token symbol')
  .setAction(async ({ symbol }, hre) => {
    if (!hre.network.config.chainId) {
      throw new Error('INVALID_CHAIN_ID');
    }

    await hre.run('compile');

    console.log(`\n- Pool Upgrading`);

    const { deployer } = await hre.getNamedAccounts();

    const network = (FORK ? FORK : hre.network.name) as eNetwork;

    const { address: strategyProxyAddress } = await hre.deployments.get(
      `${symbol}${STRATEGY_PROXY_ID}`
    );
    const { address: strategyManagerProxyAddress } = await hre.deployments.get(
      STRATEGY_MANAGER_PROXY_ID
    );

    const poolImpl = await hre.deployments.deploy(POOL_IMPL_ID, {
      from: deployer,
      contract: 'PoolBaseTVLLimits',
      args: [strategyManagerProxyAddress],
    });

    console.log(`[Deployment][INFO] Pool impl deployed ${poolImpl.address}`);

    // MultiSig call

    const proxyAdmin = await getContract(PROXY_ADMIN_ID);
    await waitForTx(
      await proxyAdmin.upgrade(
        strategyProxyAddress,
        (
          await hre.deployments.get(POOL_IMPL_ID)
        ).address
      )
    );
    console.log(`${symbol} Pool upgrade successful!'`);
  });
