import { FORK } from './../../helpers/hardhat-config-helpers';
import { task } from 'hardhat/config';
import { getContract, waitForTx } from '../../helpers/utilities/tx';
import { exit } from 'process';
import { DelegationManager, eNetwork, Slasher, StrategyManager } from '../../helpers';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { Configs } from '../../helpers/config';
import {
  DELEGATION_MANAGER_PROXY_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  STRATEGY_MANAGER_IMPL_ID,
  STRATEGY_MANAGER_PROXY_ID,
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
    const strategyManager = await hre.ethers.getContractAt(
      'StrategyManager',
      (
        await getContract(STRATEGY_MANAGER_PROXY_ID)
      ).address
    );
    const delegationManager = await hre.ethers.getContractAt(
      'DelegationManager',
      (
        await getContract(DELEGATION_MANAGER_PROXY_ID)
      ).address
    );
    const slasherManager = await hre.ethers.getContractAt(
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

    const delegationManagerOwner = await delegationManager.owner();
    if (delegationManagerOwner === deployer) {
      await waitForTx(await delegationManager.transferOwnership(owner));
      console.log('- Transfered of DelegationManager');
    }

    const strategyManagerOwner = await strategyManager.owner();
    if (strategyManagerOwner === deployer) {
      await waitForTx(await strategyManager.transferOwnership(owner));
      console.log('- Transfered of StrategyManager');
    }

    /** Output of results*/
    const result = [
      {
        role: 'ProxyAdmin owner',
        address: await proxyAdmin.owner(),
        assert: (await proxyAdmin.owner()) === owner,
      },
      {
        role: 'DelegationManager owner',
        address: await delegationManager.owner(),
        assert: (await delegationManager.owner()) === owner,
      },
      {
        role: 'StrategyManager owner',
        address: await strategyManager.owner(),
        assert: (await strategyManager.owner()) === owner,
      },
      {
        role: 'Slasher owner',
        address: await slasherManager.owner(),
        assert: (await slasherManager.owner()) === owner,
      },
    ];

    console.table(result);

    return;
  }
);
