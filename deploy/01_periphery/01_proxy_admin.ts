import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork } from '../../helpers';
import { PROXY_ADMIN_ID } from '../../helpers/deploy-ids';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  await deploy(PROXY_ADMIN_ID, {
    from: deployer,
    contract: 'ProxyAdmin',
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'ProxyAdmin';
func.tags = ['periphery'];

export default func;
