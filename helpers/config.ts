import { ZERO_ADDRESS } from './constants';
import { IConfiguration, eBscNetwork } from './types';

export const Configs: IConfiguration = {
  StrategyConfigs: {
    [eBscNetwork.bscTestnet]: {
      WBNB: {
        tokenAddress: '0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd',
        tokenName: 'Wrapped BNB',
        tokenSymbol: 'WBNB',
        maxPerDeposit:
          '115792089237316195423570985008687907853269984665640564039457584007913129639935',
        maxDeposits:
          '115792089237316195423570985008687907853269984665640564039457584007913129639935',
        withdrawalDelay: 180,
      },
      stBNB: {
        tokenAddress: '0xED392fedA1539D48004a6e037c372f09f730ef29',
        tokenName: 'Staked BNB',
        tokenSymbol: 'stBNB',
        maxPerDeposit:
          '115792089237316195423570985008687907853269984665640564039457584007913129639935',
        maxDeposits:
          '115792089237316195423570985008687907853269984665640564039457584007913129639935',
        withdrawalDelay: 180,
      },
    },
    [eBscNetwork.bsc]: {
      stBNB: {
        tokenAddress: '',
        tokenName: 'Staked BNB',
        tokenSymbol: 'stBNB',
        maxPerDeposit: '100000000000000000000',
        maxDeposits: '3000000000000000000000',
        withdrawalDelay: 604800,
      },
    },
  },
  WrappedTokenAddress: {
    [eBscNetwork.bscTestnet]: '0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd',
    [eBscNetwork.bsc]: '',
  },
  MinWithdrawalDelay: {
    [eBscNetwork.bscTestnet]: 60,
    [eBscNetwork.bsc]: 259200,
  },
  StrategyManagerPausedStatus: {
    [eBscNetwork.bscTestnet]: 0,
    [eBscNetwork.bsc]: 0,
  },
  DelegationManagerPausedStatus: {
    [eBscNetwork.bscTestnet]: 1,
    [eBscNetwork.bsc]: 1,
  },
  SlasherPausedStatus: {
    [eBscNetwork.bscTestnet]:
      '115792089237316195423570985008687907853269984665640564039457584007913129639935',
    [eBscNetwork.bsc]:
      '115792089237316195423570985008687907853269984665640564039457584007913129639935',
  },
  Owner: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '',
  },
  WhiteLister: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '',
  },
  Pauser: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '',
  },
  Unpauser: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '',
  },
};
