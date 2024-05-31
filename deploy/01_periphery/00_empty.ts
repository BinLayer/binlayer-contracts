import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork } from '../../helpers';
import { EMPTY_CONTRANCT_ID } from '../../helpers/deploy-ids';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  await deploy(EMPTY_CONTRANCT_ID, {
    from: deployer,
    contract: 'EmptyContract',
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'EmptyContract';
func.tags = ['periphery'];

export default func;
