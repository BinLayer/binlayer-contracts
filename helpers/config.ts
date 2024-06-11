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
      slisBNB: {
        tokenAddress: '0x96F124Ce690F082f469066aFE90AF633F93d94d8',
        tokenName: 'Staked Lista BNB',
        tokenSymbol: 'slisBNB',
        maxPerDeposit: '10000000000000000000000',
        maxDeposits: '10000000000000000000000000',
        withdrawalDelay: 180,
      },
    },
    [eBscNetwork.bsc]: {
      slisBNB: {
        tokenAddress: '0xB0b84D294e0C75A6abe60171b70edEb2EFd14A1B',
        tokenName: 'Staked Lista BNB',
        tokenSymbol: 'slisBNB',
        maxPerDeposit: '10000000000000000000000',
        maxDeposits: '10000000000000000000000000',
        withdrawalDelay: 604800,
      },
    },
  },
  WrappedTokenAddress: {
    [eBscNetwork.bscTestnet]: '0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd',
    [eBscNetwork.bsc]: '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c',
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
    [eBscNetwork.bsc]: '0xb3ff6E47f9D29F53a26154C7A9d1A3cAeD52A4bE',
  },
  WhiteLister: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '0xfaAB033C2B84Ebb3c7d9092d3A53bFBe12f9f65C',
  },
  Pauser: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '0xFF27eed0E8237e6EecE694788235f3a0175213Bc',
  },
  Unpauser: {
    [eBscNetwork.bscTestnet]: '0x8F2CfED2bac25D698bC66BE1E67d6A96DEEDca76',
    [eBscNetwork.bsc]: '0x9C2384e52b1708a358AC71920f11e7906F2f3770',
  },
};
