import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { eNetwork } from '../../../helpers';
import { getParamPerNetwork, isProduction } from '../../../helpers/config-helpers';
import { TESTNET_TOKEN_SUFFIX } from '../../../helpers/deploy-ids';
import { COMMON_DEPLOY_PARAMS } from '../../../helpers/env';
import { Configs } from '../../../helpers/config';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  if (isProduction()) {
    console.log('[Deployment] Skipping testnet token setup at production');
    // Early exit if is not a testnet
    return;
  }

  const network = (process.env.FORK ? process.env.FORK : hre.network.name) as eNetwork;
  console.log(`- Setting up testnet tokens " at "${network}" network`);

  const configs = getParamPerNetwork(Configs.PoolConfigs, network);

  for (let key in configs) {
    const config = configs[key];
    await deploy(`${key}${TESTNET_TOKEN_SUFFIX}`, {
      from: deployer,
      contract: 'MintableERC20',
      args: [config.tokenName, config.tokenSymbol],
      ...COMMON_DEPLOY_PARAMS,
    });
  }

  return true;
};

func.id = 'TestnetLST';
func.tags = ['token-setup'];

export default func;
