import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';

/**
 * The following script runs after the deployment starts
 */

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  console.log('=== Post deployment hook ===');

  await hre.run('print-deployments');
  // await hre.run('transfer-protocol-ownership');
  await hre.run('export', { exportAll: 'deployments/deployments_info.ts' });
};

func.tags = ['after-deploy'];
func.runAtTheEnd = true;
export default func;
