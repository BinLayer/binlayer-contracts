import { task } from 'hardhat/config';
import { eNetwork, FORK, ZERO_ADDRESS } from '../../helpers';
import {
  POOL_PROXY_ID,
  LISTA_GATEWAY_ID,
  POOL_CONTROLLER_PROXY_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { BigNumberish } from 'ethers';

task(`register-as-operator`, `DelegationController register as operator`).setAction(
  async (_, hre) => {
    if (!hre.network.config.chainId) {
      throw new Error('INVALID_CHAIN_ID');
    }

    await hre.run('compile');

    console.log(`\n- Register As Operator`);
    const { deployer } = await hre.getNamedAccounts();
    const network = (FORK ? FORK : hre.network.name) as eNetwork;
    const { address: delegationControllerAddress } = await hre.deployments.get(
      DELEGATION_CONTROLLER_PROXY_ID
    );

    const delegationController = await hre.ethers.getContractAt(
      'DelegationController',
      delegationControllerAddress
    );
    const operatorDetailsParams: {
      __deprecated_earningsReceiver: string;
      delegationApprover: string;
      stakerOptOutWindow: BigNumberish;
    } = {
      __deprecated_earningsReceiver: ZERO_ADDRESS,
      delegationApprover: ZERO_ADDRESS,
      stakerOptOutWindow: 0,
    };
    const metadataURI =
      'https://raw.githubusercontent.com/matthew7251/Metadata/main/Matthew_Metadata.json';

    // https://raw.githubusercontent.com/matthew7251/Metadata/main/Bracle_Metadata.json

    await delegationController.registerAsOperator(operatorDetailsParams, metadataURI);
    console.log(`[Deployment][INFO] DelegationControoler gatewary config successful`);
  }
);
