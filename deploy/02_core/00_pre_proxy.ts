import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork } from '../../helpers';
import {
  DELEGATION_MANAGER_PROXY_ID,
  EMPTY_CONTRANCT_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  STRATEGY_MANAGER_PROXY_ID,
} from '../../helpers/deploy-ids';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const EmptyArtifact = await deployments.get(EMPTY_CONTRANCT_ID);

  const ProxyAdminArtifact = await deployments.get(PROXY_ADMIN_ID);

  await deploy(STRATEGY_MANAGER_PROXY_ID, {
    from: deployer,
    contract: 'TransparentUpgradeableProxy',
    args: [EmptyArtifact.address, ProxyAdminArtifact.address, '0x'],
    ...COMMON_DEPLOY_PARAMS,
  });

  await deploy(DELEGATION_MANAGER_PROXY_ID, {
    from: deployer,
    contract: 'TransparentUpgradeableProxy',
    args: [EmptyArtifact.address, ProxyAdminArtifact.address, '0x'],
    ...COMMON_DEPLOY_PARAMS,
  });

  await deploy(SLASHER_PROXY_ID, {
    from: deployer,
    contract: 'TransparentUpgradeableProxy',
    args: [EmptyArtifact.address, ProxyAdminArtifact.address, '0x'],
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'CoreProxy';
func.tags = ['core'];

export default func;
