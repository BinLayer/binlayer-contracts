import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork, waitForTx } from '../../helpers';
import {
  PAUSER_REGISTRY_ID,
  PROXY_ADMIN_ID,
  POOL_IMPL_ID,
  POOL_PROXY_ID,
  TESTNET_TOKEN_SUFFIX,
} from '../../helpers/deploy-ids';
import { ethers } from 'ethers';
import { getParamPerNetwork, isProduction } from '../../helpers/config-helpers';
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
  const PoolImplArtifact = await deployments.get(POOL_IMPL_ID);
  const ProxyAdminArtifact = await deployments.get(PROXY_ADMIN_ID);

  const ifacePool = new ethers.utils.Interface(PoolImplArtifact.abi);

  const configs = getParamPerNetwork(Configs.PoolConfigs, network);

  for (let key in configs) {
    const config = configs[key];
    const tokenAddress = isProduction()
      ? config.tokenAddress
      : (await deployments.get(`${key}${TESTNET_TOKEN_SUFFIX}`)).address;
    await deploy(`${key}${POOL_PROXY_ID}`, {
      from: deployer,
      contract:
        '@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy',
      args: [
        PoolImplArtifact.address,
        ProxyAdminArtifact.address,
        ifacePool.encodeFunctionData('initialize', [
          config.maxPerDeposit,
          config.maxDeposits,
          tokenAddress,
          pauserRegistryAddress,
        ]),
      ],
      ...COMMON_DEPLOY_PARAMS,
    });
  }

  return true;
};

func.id = 'PoolInitialize';
func.tags = ['pool', 'add-pool'];

export default func;
