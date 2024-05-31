import { FORK } from '../../helpers/hardhat-config-helpers';
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

task(`pause-protocol`, `Pause protocol`).setAction(async (_, hre) => {
  // Deployer admins
  const { deployer } = await hre.getNamedAccounts();

  const network = (FORK ? FORK : hre.network.name) as eNetwork;
  const owner = getParamPerNetwork(Configs.Owner, network);

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

  console.log('Previous status: ', await delegationManager['paused(uint8)'](0));
  console.log('Previous status: ', await delegationManager['paused(uint8)'](1));
  console.log('Previous status: ', await delegationManager['paused(uint8)'](2));

  await waitForTx(await delegationManager.pause(1));

  console.log('After status: ', await delegationManager['paused(uint8)'](0));
  console.log('After status: ', await delegationManager['paused(uint8)'](1));
  console.log('After status: ', await delegationManager['paused(uint8)'](2));

  return;
});
