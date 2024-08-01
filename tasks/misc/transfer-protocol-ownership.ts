import { FORK } from './../../helpers/hardhat-config-helpers';
import { task } from 'hardhat/config';
import { getContract, waitForTx } from '../../helpers/utilities/tx';
import { exit } from 'process';
import { eNetwork } from '../../helpers';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { Configs } from '../../helpers/config';
import {
  DELEGATION_CONTROLLER_PROXY_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
} from '../../helpers/deploy-ids';

task(`transfer-protocol-ownership`, `Transfer the ownership of protocol from deployer`).setAction(
  async (_, hre) => {
    // Deployer admins
    const { deployer } = await hre.getNamedAccounts();

    const network = (FORK ? FORK : hre.network.name) as eNetwork;
    const owner = getParamPerNetwork(Configs.Owner, network);

    if (!owner) {
      console.error(
        'The constant owner is undefined. Check missing admin address at owner constant'
      );
      exit(403);
    }

    console.log('--- CURRENT DEPLOYER ADDRESSES ---');
    console.table({
      deployer,
    });
    console.log('--- DESIRED MULTISIG OWNER ---');
    console.log(owner);

    const proxyAdmin = await getContract(PROXY_ADMIN_ID);
    const poolController = await hre.ethers.getContractAt(
      'PoolController',
      (
        await getContract(POOL_CONTROLLER_PROXY_ID)
      ).address
    );
    const delegationController = await hre.ethers.getContractAt(
      'DelegationController',
      (
        await getContract(DELEGATION_CONTROLLER_PROXY_ID)
      ).address
    );
    const slasherController = await hre.ethers.getContractAt(
      'Slasher',
      (
        await getContract(SLASHER_PROXY_ID)
      ).address
    );

    const proxyAdminOwner = await proxyAdmin.owner();
    if (proxyAdminOwner === deployer) {
      await waitForTx(await proxyAdmin.transferOwnership(owner));
      console.log('- Transfered of ProxyAdmin');
    }

    const delegationControllerOwner = await delegationController.owner();
    if (delegationControllerOwner === deployer) {
      await waitForTx(await delegationController.transferOwnership(owner));
      console.log('- Transfered of DelegationController.sol');
    }

    const poolControllerOwner = await poolController.owner();
    if (poolControllerOwner === deployer) {
      await waitForTx(await poolController.transferOwnership(owner));
      console.log('- Transfered of PoolController.sol');
    }

    /** Output of results*/
    const result = [
      {
        role: 'ProxyAdmin owner',
        address: await proxyAdmin.owner(),
        assert: (await proxyAdmin.owner()) === owner,
      },
      {
        role: 'DelegationController.sol owner',
        address: await delegationController.owner(),
        assert: (await delegationController.owner()) === owner,
      },
      {
        role: 'PoolController.sol owner',
        address: await poolController.owner(),
        assert: (await poolController.owner()) === owner,
      },
      {
        role: 'Slasher owner',
        address: await slasherController.owner(),
        assert: (await slasherController.owner()) === owner,
      },
    ];

    console.table(result);

    return;
  }
);
