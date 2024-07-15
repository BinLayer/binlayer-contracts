import { task } from 'hardhat/config';
import { eNetwork, FORK, getContract, waitForTx, ZERO_ADDRESS } from '../../helpers';
import {
  POOL_CONTROLLER_PROXY_ID,
  POOL_IMPL_ID,
  POOL_PROXY_ID,
  PROXY_ADMIN_ID,
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

    const { address: poolProxyAddress } = await hre.deployments.get(`${symbol}${POOL_PROXY_ID}`);
    const { address: poolControllerProxyAddress } = await hre.deployments.get(
      POOL_CONTROLLER_PROXY_ID
    );

    const poolImpl = await hre.deployments.deploy(POOL_IMPL_ID, {
      from: deployer,
      contract: 'PoolBaseTVLLimits',
      args: [poolControllerProxyAddress],
    });

    console.log(`[Deployment][INFO] Pool impl deployed ${poolImpl.address}`);

    // MultiSig call

    const proxyAdmin = await getContract(PROXY_ADMIN_ID);
    await waitForTx(
      await proxyAdmin.upgrade(
        poolControllerProxyAddress,
        (
          await hre.deployments.get(POOL_IMPL_ID)
        ).address
      )
    );
    console.log(`${symbol} Pool upgrade successful!'`);
  });
