import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork } from '../../helpers';
import {
  DELEGATION_CONTROLLER_IMPL_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  EMPTY_CONTRANCT_ID,
  PROXY_ADMIN_ID,
  SLASHER_IMPL_ID,
  SLASHER_PROXY_ID,
  POOL_CONTROLLER_IMPL_ID,
  POOL_CONTROLLER_PROXY_ID,
} from '../../helpers/deploy-ids';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const DelegationControllerProxyArtifact = await deployments.get(DELEGATION_CONTROLLER_PROXY_ID);
  const PoolControllerProxyArtifact = await deployments.get(POOL_CONTROLLER_PROXY_ID);
  const SlasherProxyArtifact = await deployments.get(SLASHER_PROXY_ID);

  await deploy(DELEGATION_CONTROLLER_IMPL_ID, {
    from: deployer,
    contract: 'DelegationController',
    args: [PoolControllerProxyArtifact.address, SlasherProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  await deploy(POOL_CONTROLLER_IMPL_ID, {
    from: deployer,
    contract: 'PoolController',
    args: [DelegationControllerProxyArtifact.address, SlasherProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  await deploy(SLASHER_IMPL_ID, {
    from: deployer,
    contract: 'Slasher',
    args: [PoolControllerProxyArtifact.address, DelegationControllerProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'CoreImpl';
func.tags = ['core'];

export default func;
