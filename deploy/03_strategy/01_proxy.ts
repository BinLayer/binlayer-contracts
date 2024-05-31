import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork, waitForTx } from '../../helpers';
import {
  PAUSER_REGISTRY_ID,
  PROXY_ADMIN_ID,
  STRATEGY_IMPL_ID,
  STRATEGY_PROXY_ID,
} from '../../helpers/deploy-ids';
import { ethers } from 'ethers';
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

  const { address: pauserRegistryAddress } = await deployments.get(PAUSER_REGISTRY_ID);
  const StrategyImplArtifact = await deployments.get(STRATEGY_IMPL_ID);
  const ProxyAdminArtifact = await deployments.get(PROXY_ADMIN_ID);

  const ifaceStrategy = new ethers.utils.Interface(StrategyImplArtifact.abi);

  const configs = getParamPerNetwork(Configs.StrategyConfigs, network);

  for (let key in configs) {
    const config = configs[key];
    await deploy(`${key}${STRATEGY_PROXY_ID}`, {
      from: deployer,
      contract: 'TransparentUpgradeableProxy',
      args: [
        StrategyImplArtifact.address,
        ProxyAdminArtifact.address,
        ifaceStrategy.encodeFunctionData('initialize', [
          config.maxPerDeposit,
          config.maxDeposits,
          config.tokenAddress,
          pauserRegistryAddress,
        ]),
      ],
      ...COMMON_DEPLOY_PARAMS,
    });
  }

  return true;
};

func.id = 'StrategyInitialize';
func.tags = ['strategy', 'add-strategy'];

export default func;
