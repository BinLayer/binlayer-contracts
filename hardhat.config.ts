import {
  ARBISCAN_KEY,
  BSCSCAN_KEY,
  ETHERSCAN_KEY,
  getCommonNetworkConfig,
  hardhatNetworkSettings,
  loadTasks,
} from './helpers/hardhat-config-helpers';
import { eArbitrumNetwork, eBscNetwork, eEthereumNetwork } from './helpers/types';
import { DEFAULT_NAMED_ACCOUNTS } from './helpers/constants';

import '@nomicfoundation/hardhat-toolbox';
import '@openzeppelin/hardhat-upgrades';
import 'hardhat-deploy';
import 'hardhat-contract-sizer';
import 'hardhat-abi-exporter';
import 'hardhat-gas-reporter';
import '@nomiclabs/hardhat-ethers';
import '@nomicfoundation/hardhat-foundry';

const SKIP_LOAD = process.env.SKIP_LOAD === 'true';
const TASK_FOLDERS = ['misc'];

// Prevent to load tasks before compilation and typechain
if (!SKIP_LOAD) {
  loadTasks(TASK_FOLDERS);
}

export default {
  external: {
    contracts: [
      {
        artifacts: 'node_modules/@openzeppelin/upgrades-core/artifacts',
      },
    ],
  },
  abiExporter: {
    path: './abis', // path to ABI export directory (relative to Hardhat root)
    runOnCompile: true, // whether to automatically export ABIs during compilation
    clear: true, // whether to delete old ABI files in path on compilation
    flat: true, // whether to flatten output directory (may cause name collisions)
    pretty: false, // whether to use interface-style formatting of output for better readability
    except: ['@openzeppelin/contracts'],
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: false,
    disambiguatePaths: false,
  },
  solidity: {
    compilers: [
      {
        version: '0.8.9',
        settings: {
          optimizer: { enabled: true, runs: 1_500 },
          evmVersion: 'berlin',
        },
      },
      {
        version: '0.8.20',
        settings: {
          optimizer: { enabled: true, runs: 1_500 },
          evmVersion: 'berlin',
        },
      },
    ],
  },
  typechain: {
    outDir: 'typechain',
    target: 'ethers-v5',
  },
  gasReporter: {
    enabled: false,
  },
  networks: {
    hardhat: {
      ...hardhatNetworkSettings,
    },
    localhost: {
      url: 'http://127.0.0.1:8545',
      ...hardhatNetworkSettings,
    },
    [eEthereumNetwork.main]: getCommonNetworkConfig(eEthereumNetwork.main, 1),
    [eEthereumNetwork.goerli]: getCommonNetworkConfig(eEthereumNetwork.goerli, 5),
    [eEthereumNetwork.sepolia]: getCommonNetworkConfig(eEthereumNetwork.sepolia, 11155111),
    [eArbitrumNetwork.arbitrum]: {
      ...getCommonNetworkConfig(eArbitrumNetwork.arbitrum, 42161),
      verify: {
        etherscan: {
          apiKey: ARBISCAN_KEY,
        },
      },
    },
    [eArbitrumNetwork.arbitrumSepolia]: {
      ...getCommonNetworkConfig(eArbitrumNetwork.arbitrumSepolia, 421614),
      verify: {
        etherscan: {
          apiUrl: 'https://api-sepolia.arbiscan.io',
          apiKey: ARBISCAN_KEY,
        },
      },
    },
    [eBscNetwork.bsc]: {
      ...getCommonNetworkConfig(eBscNetwork.bsc, 56),
      verify: {
        etherscan: {
          apiUrl: 'https://api.bscscan.com',
          apiKey: BSCSCAN_KEY,
        },
      },
    },
    [eBscNetwork.bscTestnet]: {
      ...getCommonNetworkConfig(eBscNetwork.bscTestnet, 97),
      verify: {
        etherscan: {
          apiUrl: 'https://api-testnet.bscscan.com',
          apiKey: BSCSCAN_KEY,
        },
      },
    },
  },
  namedAccounts: {
    ...DEFAULT_NAMED_ACCOUNTS,
  },
  mocha: {
    timeout: 0,
  },
  verify: {
    etherscan: {
      apiKey: ETHERSCAN_KEY,
    },
  },
  etherscan: {
    apiKey: {
      mainnet: ETHERSCAN_KEY,
      sepolia: ETHERSCAN_KEY,
      goerli: ETHERSCAN_KEY,
      arbitrumOne: ARBISCAN_KEY,
      [eArbitrumNetwork.arbitrumSepolia]: ARBISCAN_KEY,
      [eBscNetwork.bsc]: BSCSCAN_KEY,
      [eBscNetwork.bscTestnet]: BSCSCAN_KEY,
    },
    customChains: [
      {
        network: 'arbitrum-sepolia',
        chainId: 421614,
        urls: {
          apiURL: 'https://api-sepolia.arbiscan.io/api',
          browserURL: 'https://sepolia.arbiscan.io/',
        },
      },
      {
        network: 'bsc',
        chainId: 56,
        urls: {
          apiURL: 'https://api.bscscan.com/api',
          browserURL: 'https://bscscan.com/',
        },
      },
      {
        network: 'bsc-testnet',
        chainId: 97,
        urls: {
          apiURL: 'https://api-testnet.bscscan.com/api',
          browserURL: 'https://testnet.bscscan.com/',
        },
      },
    ],
  },
};
