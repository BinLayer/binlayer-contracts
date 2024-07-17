import { task } from 'hardhat/config';
import { eNetwork, FORK } from '../../helpers';
import {
  POOL_PROXY_ID,
  LISTA_GATEWAY_ID,
  POOL_CONTROLLER_PROXY_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';

task(`deploy-wrapped-gateway`, `Deploys the WrappedTokenGateway contract`).setAction(
  async (_, hre) => {
    if (!hre.network.config.chainId) {
      throw new Error('INVALID_CHAIN_ID');
    }

    await hre.run('compile');

    console.log(`\n- WrappedTokenGateway deployment`);
    const { deployer } = await hre.getNamedAccounts();
    const network = (FORK ? FORK : hre.network.name) as eNetwork;
    const owner = getParamPerNetwork(Configs.Owner, network);

    const wrappedTokenAddress = getParamPerNetwork(Configs.WrappedTokenAddress, network);

    const { address: poolControllerAddress } = await hre.deployments.get(POOL_CONTROLLER_PROXY_ID);
    const { address: delegationControllerAddress } = await hre.deployments.get(
      DELEGATION_CONTROLLER_PROXY_ID
    );
    const { address: wbnbPoolProxyAddress } = await hre.deployments.get(`WBNB${POOL_PROXY_ID}`);

    const wrappedTokenGateway = await hre.deployments.deploy(`V2-${WRAPPED_TOKEN_GATEWAY_ID}`, {
      contract: 'WrappedTokenGateway',
      from: deployer,
      args: [
        wrappedTokenAddress,
        owner,
        wbnbPoolProxyAddress,
        poolControllerAddress,
        delegationControllerAddress,
      ],
    });
    console.log(`[Deployment][INFO] WrappedTokenGateway deployed ${wrappedTokenGateway.address}`);

    // DelegationController Owner
    const delegationController = await hre.ethers.getContractAt(
      'DelegationController',
      delegationControllerAddress
    );
    await delegationController.updateWrappedTokenGateway(
      (
        await hre.deployments.get(`V2-${WRAPPED_TOKEN_GATEWAY_ID}`)
      ).address
    );
    console.log(`[Deployment][INFO] DelegationControoler gatewary config successful`);
  }
);
