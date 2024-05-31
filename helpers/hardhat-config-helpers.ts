import fs from 'fs';
import path from 'path';
import { HardhatNetworkForkingUserConfig } from 'hardhat/types';
import {
  iParamsPerNetwork,
  eEthereumNetwork,
  eNetwork,
  eArbitrumNetwork,
  eBscNetwork,
} from './types';

require('dotenv').config();

export const DEFAULT_BLOCK_GAS_LIMIT = 12450000;
export const DEFAULT_GAS_PRICE = 8000000000;
export const INFURA_KEY = process.env.INFURA_KEY || '';
export const FORK = (process.env.FORK || '') as eNetwork;
export const FORK_BLOCK_NUMBER = process.env.FORK_BLOCK_NUMBER
  ? parseInt(process.env.FORK_BLOCK_NUMBER)
  : 0;
const MNEMONIC_PATH = "m/44'/60'/0'/0";
const MNEMONIC = process.env.MNEMONIC || '';

export const getInfuraKey = (net: eNetwork) => {
  switch (net) {
    case eEthereumNetwork.goerli:
      return process.env.GORLI_INFURA_KEY || INFURA_KEY;
    case eEthereumNetwork.sepolia:
      return process.env.SEPOLIA_INFURA_KEY || INFURA_KEY;
    default:
      return INFURA_KEY;
  }
};

export const NETWORKS_RPC_URL: iParamsPerNetwork<string> = {
  [eEthereumNetwork.main]: `https://mainnet.infura.io/v3/${getInfuraKey(eEthereumNetwork.main)}`,
  [eEthereumNetwork.hardhat]: 'http://localhost:8545',
  [eEthereumNetwork.goerli]: `https://goerli.infura.io/v3/${getInfuraKey(eEthereumNetwork.goerli)}`,
  [eEthereumNetwork.sepolia]: `https://sepolia.infura.io/v3/${getInfuraKey(
    eEthereumNetwork.sepolia
  )}`,
  [eArbitrumNetwork.arbitrum]: `https://arbitrum-mainnet.infura.io/v3/${getInfuraKey(
    eArbitrumNetwork.arbitrum
  )}`,
  [eArbitrumNetwork.arbitrumSepolia]: `https://arbitrum-sepolia.infura.io/v3/${getInfuraKey(
    eArbitrumNetwork.arbitrum
  )}`,
  [eBscNetwork.bsc]: `https://bsc-dataseed.bnbchain.org`,
  [eBscNetwork.bscTestnet]: `https://bsc-testnet-dataseed.bnbchain.org`,
};

export const LIVE_NETWORKS: iParamsPerNetwork<boolean> = {
  [eEthereumNetwork.sepolia]: true,
  [eEthereumNetwork.goerli]: true,
  [eEthereumNetwork.main]: true,
  [eArbitrumNetwork.arbitrum]: true,
  [eArbitrumNetwork.arbitrumSepolia]: true,
  [eBscNetwork.bsc]: true,
  [eBscNetwork.bscTestnet]: true,
};

const GAS_PRICE_PER_NET: iParamsPerNetwork<string | number> = {
  [eArbitrumNetwork.arbitrumSepolia]: 100000001,
};

export const buildForkConfig = (): HardhatNetworkForkingUserConfig | undefined => {
  let forkMode: HardhatNetworkForkingUserConfig | undefined;
  if (FORK && NETWORKS_RPC_URL[FORK]) {
    forkMode = {
      url: NETWORKS_RPC_URL[FORK] as string,
    };
    if (FORK_BLOCK_NUMBER) {
      forkMode.blockNumber = FORK_BLOCK_NUMBER;
    }
  }
  return forkMode;
};

export const loadTasks = (taskFolders: string[]): void =>
  taskFolders.forEach((folder) => {
    const tasksPath = path.join(__dirname, '../tasks', folder);
    fs.readdirSync(tasksPath)
      .filter((pth) => pth.includes('.ts') || pth.includes('.js'))
      .forEach((task) => {
        require(`${tasksPath}/${task}`);
      });
  });

export const getCommonNetworkConfig = (networkName: eNetwork, chainId?: number) => ({
  url: NETWORKS_RPC_URL[networkName] || '',
  blockGasLimit: DEFAULT_BLOCK_GAS_LIMIT,
  chainId,
  gasPrice: GAS_PRICE_PER_NET[networkName] || undefined,
  accounts: [`${process.env.ACCOUNT_SECRETKEY}`],
  live: LIVE_NETWORKS[networkName] || false,
});

export const hardhatNetworkSettings = {
  gasPrice: 'auto',
  initialBaseFeePerGas: '0',
  blockGasLimit: DEFAULT_BLOCK_GAS_LIMIT,
  throwOnTransactionFailures: true,
  throwOnCallFailures: true,
  chainId: 31337,
  forking: buildForkConfig(),
  saveDeployments: true,
  allowUnlimitedContractSize: true,
  tags: ['local'],
  accounts:
    FORK && !!MNEMONIC
      ? {
          mnemonic: MNEMONIC,
          path: MNEMONIC_PATH,
          initialIndex: 0,
          count: 10,
        }
      : undefined,
};

export const ETHERSCAN_KEY = process.env.ETHERSCAN_KEY || '';

export const ARBISCAN_KEY = process.env.ARBISCAN_KEY || '';

export const BASESCAN_KEY = process.env.BASESCAN_KEY || '';

export const BSCSCAN_KEY = process.env.BSCSCAN_KEY || '';

export const MANTLESCAN_KEY = process.env.MANTLESCAN_KEY || '';

export const MERLINSCAN_KEY = process.env.MERLINSCAN_KEY || '';
