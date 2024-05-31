import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork } from '../../helpers';
import { PAUSER_REGISTRY_ID } from '../../helpers/deploy-ids';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { Configs } from '../../helpers/config';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const network = (process.env.FORK ? process.env.FORK : hre.network.name) as eNetwork;

  const pauser = getParamPerNetwork(Configs.Pauser, network);
  const unpauser = getParamPerNetwork(Configs.Unpauser, network);

  await deploy(PAUSER_REGISTRY_ID, {
    from: deployer,
    contract: 'PauserRegistry',
    args: [[pauser], unpauser],
    ...COMMON_DEPLOY_PARAMS,
  });

  return true;
};

func.id = 'PauserRegistry';
func.tags = ['permissions'];

export default func;
