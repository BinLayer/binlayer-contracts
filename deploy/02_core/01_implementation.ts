import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork } from '../../helpers';
import {
  DELEGATION_MANAGER_IMPL_ID,
  DELEGATION_MANAGER_PROXY_ID,
  EMPTY_CONTRANCT_ID,
  PROXY_ADMIN_ID,
  SLASHER_IMPL_ID,
  SLASHER_PROXY_ID,
  STRATEGY_MANAGER_IMPL_ID,
  STRATEGY_MANAGER_PROXY_ID,
} from '../../helpers/deploy-ids';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const DelegationControllerProxyArtifact = await deployments.get(DELEGATION_MANAGER_PROXY_ID);
  const StrategyManagerProxyArtifact = await deployments.get(STRATEGY_MANAGER_PROXY_ID);
  const SlasherProxyArtifact = await deployments.get(SLASHER_PROXY_ID);

  await deploy(DELEGATION_MANAGER_IMPL_ID, {
    from: deployer,
    contract: 'DelegationController',
    args: [StrategyManagerProxyArtifact.address, SlasherProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  await deploy(STRATEGY_MANAGER_IMPL_ID, {
    from: deployer,
    contract: 'StrategyManager',
    args: [DelegationControllerProxyArtifact.address, SlasherProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  await deploy(SLASHER_IMPL_ID, {
    from: deployer,
    contract: 'Slasher',
    args: [StrategyManagerProxyArtifact.address, DelegationControllerProxyArtifact.address],
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'CoreImpl';
func.tags = ['core'];

export default func;
