import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { iParamsPerNetwork, eNetwork, eEthereumNetwork } from './types';
declare var hre: HardhatRuntimeEnvironment;

export const getParamPerNetwork = <T>(
  param: iParamsPerNetwork<T> | undefined,
  network: eNetwork
): T | undefined => {
  if (!param) return undefined;

  return param[network];
};

export const isProduction = (): boolean => {
  const network = (process.env.FORK ? process.env.FORK : hre.network.name) as eNetwork;

  return network != eEthereumNetwork.hardhat;
};
