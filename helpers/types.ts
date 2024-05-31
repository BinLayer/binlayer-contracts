import { BigNumber, BigNumberish, BytesLike, Signer } from 'ethers/lib/ethers';
import { Deployment } from 'hardhat-deploy/types';

export type eNetwork = eEthereumNetwork | eArbitrumNetwork | eBscNetwork;

export enum eEthereumNetwork {
  main = 'main',
  hardhat = 'hardhat',
  goerli = 'goerli',
  sepolia = 'sepolia',
}

export enum eArbitrumNetwork {
  arbitrum = 'arbitrum',
  arbitrumSepolia = 'arbitrum-sepolia',
}

export enum eBscNetwork {
  bsc = 'bsc',
  bscTestnet = 'bsc-testnet',
}

export type tEthereumAddress = string;
export type tStringTokenBigUnits = string; // 1 ETH, or 10e6 USDC or 10e18 DAI
export type tBigNumberTokenBigUnits = BigNumber;
export type tStringTokenSmallUnits = string; // 1 wei, or 1 basic unit of USDC, or 1 basic unit of DAI
export type tBigNumberTokenSmallUnits = BigNumber;

export type iParamsPerNetwork<T> = {
  [k in eNetwork]?: T;
};

export type iParamsPerNetworkString<T> = {
  [k in string]?: T;
};

export type iParamsPerNetworkWithDefault<T> = {
  [k in eNetwork]?: T;
} & {
  default: T;
};

export interface iParamsPerNetworkAll<T> extends iEthereumParamsPerNetwork<T> {}

export interface iEthereumParamsPerNetwork<T> {
  [eEthereumNetwork.main]: T;
}
export interface iArbitrumParamsPerNetwork<T> {
  [eArbitrumNetwork.arbitrum]: T;
  [eArbitrumNetwork.arbitrumSepolia]: T;
}

export interface ITokenAddress {
  [token: string]: tEthereumAddress;
}

export interface SymbolMap<T> {
  [symbol: string]: T;
}

export interface StrategyConfig {
  tokenAddress: tEthereumAddress;
  tokenName: string;
  tokenSymbol: string;
  maxPerDeposit: BigNumberish;
  maxDeposits: BigNumberish;
  withdrawalDelay: number;
}

export interface IConfiguration {
  StrategyConfigs: iParamsPerNetwork<SymbolMap<StrategyConfig>>;
  MinWithdrawalDelay: iParamsPerNetwork<number>;
  StrategyManagerPausedStatus: iParamsPerNetwork<BigNumberish>;
  DelegationManagerPausedStatus: iParamsPerNetwork<BigNumberish>;
  SlasherPausedStatus: iParamsPerNetwork<BigNumberish>;
  WrappedTokenAddress: iParamsPerNetwork<tEthereumAddress>;
  Owner: iParamsPerNetwork<tEthereumAddress>;
  WhiteLister: iParamsPerNetwork<tEthereumAddress>;
  Pauser: iParamsPerNetwork<tEthereumAddress>;
  Unpauser: iParamsPerNetwork<tEthereumAddress>;
}
