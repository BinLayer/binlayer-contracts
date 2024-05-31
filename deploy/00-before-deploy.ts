import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { getWalletBalances } from '../helpers';

/**
 * The following script runs before the deployment starts
 */

const func: DeployFunction = async function ({}: HardhatRuntimeEnvironment) {
  // Print the accounts with their balance before the deployment script
  const balances = await getWalletBalances();
  console.log('\nAccounts');
  console.log('========');
  console.table(balances);
};

func.tags = ['before-deploy'];

export default func;
