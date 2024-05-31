import { iParamsPerNetwork, eNetwork } from './types';

export const getParamPerNetwork = <T>(
  param: iParamsPerNetwork<T> | undefined,
  network: eNetwork
): T | undefined => {
  if (!param) return undefined;

  return param[network];
};
