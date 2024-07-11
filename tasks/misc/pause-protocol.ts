import { FORK } from '../../helpers/hardhat-config-helpers';
import { task } from 'hardhat/config';
import { getContract, waitForTx } from '../../helpers/utilities/tx';
import { exit } from 'process';
import { DelegationController, eNetwork, Slasher, PoolController } from '../../helpers';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { Configs } from '../../helpers/config';
import {
  DELEGATION_CONTROLLER_PROXY_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
} from '../../helpers/deploy-ids';

task(`pause-protocol`, `Pause protocol`).setAction(async (_, hre) => {
  // Deployer admins
  const { deployer } = await hre.getNamedAccounts();

  const network = (FORK ? FORK : hre.network.name) as eNetwork;
  const owner = getParamPerNetwork(Configs.Owner, network);

  const strategyController = await hre.ethers.getContractAt(
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

  console.log('Previous status: ', await delegationController['paused(uint8)'](0));
  console.log('Previous status: ', await delegationController['paused(uint8)'](1));
  console.log('Previous status: ', await delegationController['paused(uint8)'](2));

  await waitForTx(await delegationController.pause(1));

  console.log('After status: ', await delegationController['paused(uint8)'](0));
  console.log('After status: ', await delegationController['paused(uint8)'](1));
  console.log('After status: ', await delegationController['paused(uint8)'](2));

  return;
});
