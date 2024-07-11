import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { POOL_IMPL_ID, POOL_CONTROLLER_PROXY_ID } from '../../helpers/deploy-ids';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const PoolControllerProxyArtifact = await deployments.get(POOL_CONTROLLER_PROXY_ID);

  await deploy(POOL_IMPL_ID, {
    from: deployer,
    contract: 'PoolBaseTVLLimits',
    args: [PoolControllerProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'PoolImpl';
func.tags = ['pool'];

export default func;
