import { task } from 'hardhat/config';
import { eNetwork, FORK } from '../../helpers';
import {
  STRATEGY_MANAGER_PROXY_ID,
  STRATEGY_PROXY_ID,
  LISTA_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';

task(`deploy-lista-gateway`, `Deploys the ListaGateway contract`).setAction(async (_, hre) => {
  if (!hre.network.config.chainId) {
    throw new Error('INVALID_CHAIN_ID');
  }

  await hre.run('compile');

  console.log(`\n- ListaGateway deployment`);

  const network = (FORK ? FORK : hre.network.name) as eNetwork;
  const owner = getParamPerNetwork(Configs.Owner, network);

  // bsc
  const listaStakeManagerAddress = '0x1adB950d8bB3dA4bE104211D5AB038628e477fE6';
  const slisBNBAddress = '0xB0b84D294e0C75A6abe60171b70edEb2EFd14A1B';

  // bsc-testnet
  // const listaStakeManagerAddress = '0x3A2231E023296bbc177d7587DdFd60B0A96F3A29';
  // const slisBNBAddress = '0x96F124Ce690F082f469066aFE90AF633F93d94d8';

  const { address: strategyAddress } = await hre.deployments.get(`slisBNB${STRATEGY_PROXY_ID}`);
  const { address: strategyManagerAddress } = await hre.deployments.get(STRATEGY_MANAGER_PROXY_ID);

  const { deployer } = await hre.getNamedAccounts();
  const listaGateway = await hre.deployments.deploy(LISTA_GATEWAY_ID, {
    contract: 'ListaGateway',
    from: deployer,
    args: [
      slisBNBAddress,
      owner,
      listaStakeManagerAddress,
      strategyAddress,
      strategyManagerAddress,
    ],
  });
  console.log(`[Deployment][INFO] ListaGateway deployed ${listaGateway.address}`);
});
