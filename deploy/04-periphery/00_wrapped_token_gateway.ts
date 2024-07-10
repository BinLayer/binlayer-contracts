import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { DeployFunction } from 'hardhat-deploy/types';
import { COMMON_DEPLOY_PARAMS } from '../../helpers/env';
import { eNetwork, FORK } from '../../helpers';
import {
  DELEGATION_CONTROLLER_PROXY_ID,
  EMPTY_CONTRANCT_ID,
  PROXY_ADMIN_ID,
  SLASHER_PROXY_ID,
  STRATEGY_IMPL_ID,
  STRATEGY_MANAGER_PROXY_ID,
  STRATEGY_PROXY_ID,
  WRAPPED_TOKEN_GATEWAY_ID,
} from '../../helpers/deploy-ids';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { Configs } from '../../helpers/config';

const func: DeployFunction = async function ({
  getNamedAccounts,
  deployments,
  ...hre
}: HardhatRuntimeEnvironment) {
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const network = (FORK ? FORK : hre.network.name) as eNetwork;

  const wbnbStrategyArtifact = await deployments.getOrNull(`WBNB${STRATEGY_PROXY_ID}`);
  if (!wbnbStrategyArtifact) {
    console.log('[Deployment] Skipping wrapped token gateway');
    return true;
  }

  const owner = getParamPerNetwork(Configs.Owner, network);
  const wrappedTokenAddress = getParamPerNetwork(Configs.WrappedTokenAddress, network);
  if (!owner || !wrappedTokenAddress) {
    throw '[Deployment][Error] Owner or WrappedTokenAddress not config';
  }

  const { address: strategyManagerAddress } = await deployments.get(STRATEGY_MANAGER_PROXY_ID);
  const { address: delegationControllerAddress } = await deployments.get(DELEGATION_CONTROLLER_PROXY_ID);

  await deploy(WRAPPED_TOKEN_GATEWAY_ID, {
    contract: 'WrappedTokenGateway',
    from: deployer,
    args: [
      wrappedTokenAddress,
      owner,
      wbnbStrategyArtifact.address,
      strategyManagerAddress,
      delegationControllerAddress,
    ],
    ...COMMON_DEPLOY_PARAMS,
  });
  return true;
};

func.id = 'WrappedTokenGateway';
func.tags = ['periphery'];

export default func;
