export default {
  "56": [
    {
      "name": "bsc",
      "chainId": "56",
      "contracts": {
        "DelegationController-Implementation": {
          "address": "0xa1f206f3dF721AF747D007AE135901240fb74fb0",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                },
                {
                  "internalType": "contract ISlasher",
                  "name": "_slasher",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MinWithdrawalDelaySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "newOperatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "OperatorDetailsModified",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "OperatorMetadataURIUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "operatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "OperatorRegistered",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "OperatorSharesDecreased",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "OperatorSharesIncreased",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerDelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerForceUndelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerUndelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "StrategyWithdrawalDelaySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousGateway",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "currentGateway",
                  "type": "address"
                }
              ],
              "name": "UpdateWrappedTokenGateway",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "bytes32",
                  "name": "withdrawalRoot",
                  "type": "bytes32"
                }
              ],
              "name": "WithdrawalCompleted",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "bytes32",
                  "name": "withdrawalRoot",
                  "type": "bytes32"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationManager.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                }
              ],
              "name": "WithdrawalQueued",
              "type": "event"
            },
            {
              "inputs": [],
              "name": "DELEGATION_APPROVAL_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "DOMAIN_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "MAX_STAKER_OPT_OUT_WINDOW",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "MAX_WITHDRAWAL_DELAY",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "STAKER_DELEGATION_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateCurrentStakerDelegationDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "_delegationApprover",
                  "type": "address"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateDelegationApprovalDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "_stakerNonce",
                  "type": "uint256"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateStakerDelegationDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                }
              ],
              "name": "calculateWithdrawalRoot",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "pure",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                },
                {
                  "internalType": "contract IERC20[]",
                  "name": "tokens",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256",
                  "name": "middlewareTimesIndex",
                  "type": "uint256"
                },
                {
                  "internalType": "bool",
                  "name": "receiveAsTokens",
                  "type": "bool"
                }
              ],
              "name": "completeQueuedWithdrawal",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal[]",
                  "name": "withdrawals",
                  "type": "tuple[]"
                },
                {
                  "internalType": "contract IERC20[][]",
                  "name": "tokens",
                  "type": "address[][]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "middlewareTimesIndexes",
                  "type": "uint256[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "receiveAsTokens",
                  "type": "bool[]"
                }
              ],
              "name": "completeQueuedWithdrawals",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "cumulativeWithdrawalsQueued",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "decreaseDelegatedShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "approverSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                }
              ],
              "name": "delegateTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "stakerSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "approverSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                }
              ],
              "name": "delegateToBySignature",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "delegatedTo",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "delegationApprover",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "name": "delegationApproverSaltIsSpent",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "domainSeparator",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "earningsReceiver",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "getDelegatableShares",
              "outputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategies",
                  "type": "address[]"
                }
              ],
              "name": "getOperatorShares",
              "outputs": [
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategies",
                  "type": "address[]"
                }
              ],
              "name": "getWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "increaseDelegatedShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "initialOwner",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "initialPausedStatus",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "_minWithdrawalDelay",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IStrategy[]",
                  "name": "_strategies",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "_withdrawalDelay",
                  "type": "uint256[]"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "isDelegated",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "isOperator",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "minWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "newOperatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "modifyOperatorDetails",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "operatorDetails",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "name": "pendingWithdrawals",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    }
                  ],
                  "internalType": "struct IDelegationManager.QueuedWithdrawalParams[]",
                  "name": "queuedWithdrawalParams",
                  "type": "tuple[]"
                }
              ],
              "name": "queueWithdrawals",
              "outputs": [
                {
                  "internalType": "bytes32[]",
                  "name": "",
                  "type": "bytes32[]"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "registeringOperatorDetails",
                  "type": "tuple"
                },
                {
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "registerAsOperator",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newMinWithdrawalDelay",
                  "type": "uint256"
                }
              ],
              "name": "setMinWithdrawalDelay",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategies",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "withdrawalDelay",
                  "type": "uint256[]"
                }
              ],
              "name": "setStrategyWithdrawalDelay",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "slasher",
              "outputs": [
                {
                  "internalType": "contract ISlasher",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "stakerNonce",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "stakerOptOutWindow",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyManager",
              "outputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "strategyWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "undelegate",
              "outputs": [
                {
                  "internalType": "bytes32[]",
                  "name": "withdrawalRoots",
                  "type": "bytes32[]"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "updateOperatorMetadataURI",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_newWrappedTokenGateway",
                  "type": "address"
                }
              ],
              "name": "updateWrappedTokenGateway",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "wrappedTokenGateway",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "DelegationManager-Proxy": {
          "address": "0xDF8843c6580BC4b92B203d63E0aA662536eD8B7E",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "EmptyContract": {
          "address": "0x42ca622595d971bFc2a971ebf2074Ed8F2f77D0f",
          "abi": [
            {
              "inputs": [],
              "name": "foo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "ListaGateway": {
          "address": "0x60D7870b6FA2D79e896aaEFD6b44929697527246",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_slisBNB",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "_owner",
                  "type": "address"
                },
                {
                  "internalType": "contract IListaStakeManager",
                  "name": "_listaStakeManager",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "_strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "depositNativeToken",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "emergencyNativeTransfer",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "emergencyTokenTransfer",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getListaContracts",
              "outputs": [
                {
                  "internalType": "contract IListaStakeManager",
                  "name": "_listaStakeManager",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "_slisBNB",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getStrategyAddress",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "PauserRegistry": {
          "address": "0x2be75b9E2504507631A576D7CaD6008638e6a069",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address[]",
                  "name": "_pausers",
                  "type": "address[]"
                },
                {
                  "internalType": "address",
                  "name": "_unpauser",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "pauser",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "bool",
                  "name": "canPause",
                  "type": "bool"
                }
              ],
              "name": "PauserStatusChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousUnpauser",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newUnpauser",
                  "type": "address"
                }
              ],
              "name": "UnpauserChanged",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "isPauser",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newPauser",
                  "type": "address"
                },
                {
                  "internalType": "bool",
                  "name": "canPause",
                  "type": "bool"
                }
              ],
              "name": "setIsPauser",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newUnpauser",
                  "type": "address"
                }
              ],
              "name": "setUnpauser",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "unpauser",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "ProxyAdmin": {
          "address": "0xaBfE3C2B5b353a6d130DE66af816E8869535494C",
          "abi": [
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeProxyAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                }
              ],
              "name": "getProxyAdmin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                }
              ],
              "name": "getProxyImplementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "upgrade",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            }
          ]
        },
        "Slasher-Implementation": {
          "address": "0x073402A2BA391fdBf8BC77f38e8F0b08db9D5668",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IDelegationManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previouslySlashedAddress",
                  "type": "address"
                }
              ],
              "name": "FrozenStatusReset",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "index",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "stalestUpdateTimestamp",
                  "type": "uint32"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "latestServeUntilTimestamp",
                  "type": "uint32"
                }
              ],
              "name": "MiddlewareTimesAdded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "slashedOperator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "slashingContract",
                  "type": "address"
                }
              ],
              "name": "OperatorFrozen",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                }
              ],
              "name": "OptedIntoSlashing",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "contractCanSlashOperatorUntilTimestamp",
                  "type": "uint32"
                }
              ],
              "name": "SlashingAbilityRevoked",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "canSlash",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "canWithdraw",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "contractCanSlashOperatorUntilTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "delegation",
              "outputs": [
                {
                  "internalType": "contract IDelegationManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "freezeOperator",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getCorrectValueForInsertAfter",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getMiddlewareTimesIndexServeUntilTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getMiddlewareTimesIndexStalestUpdateTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "isFrozen",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "latestUpdateTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "middlewareTimesLength",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "operatorToMiddlewareTimes",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "uint32",
                      "name": "stalestUpdateTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "latestServeUntilTimestamp",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct ISlasher.MiddlewareTimes",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorWhitelistedContractsLinkedListEntry",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorWhitelistedContractsLinkedListSize",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "optIntoSlashing",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "recordFirstStakeUpdate",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "recordLastStakeUpdateAndRevokeSlashingAbility",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "recordStakeUpdate",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address[]",
                  "name": "",
                  "type": "address[]"
                }
              ],
              "name": "resetFrozenStatus",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyManager",
              "outputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "whitelistedContractDetails",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "uint32",
                      "name": "registrationMayBeginAtTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "contractCanSlashOperatorUntilTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "latestUpdateTimestamp",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct ISlasher.MiddlewareDetails",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "Slasher-Proxy": {
          "address": "0x74D9d2324965688D0c712763CA01566a1a6c53A6",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "Strategy-Implementation": {
          "address": "0x821f24c381a67Ee6115CbBe0F9fECA5a16b1080E",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MaxPerDepositUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MaxTotalDepositsUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "deposit",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "newShares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "explanation",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "pure",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getTVLLimits",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "_maxPerDeposit",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "_maxTotalDeposits",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "_underlyingToken",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "_underlyingToken",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                }
              ],
              "name": "initializeBase",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "maxPerDeposit",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "maxTotalDeposits",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newMaxPerDeposit",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "newMaxTotalDeposits",
                  "type": "uint256"
                }
              ],
              "name": "setTVLLimits",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "shares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "sharesToUnderlying",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "sharesToUnderlyingView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyManager",
              "outputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "totalShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountUnderlying",
                  "type": "uint256"
                }
              ],
              "name": "underlyingToShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountUnderlying",
                  "type": "uint256"
                }
              ],
              "name": "underlyingToSharesView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "underlyingToken",
              "outputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "userUnderlying",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "userUnderlyingView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "recipient",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "withdraw",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "StrategyManager-Implementation": {
          "address": "0x5ac652dC32469a1F04a75DD8c05AAF77780234E5",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IDelegationManager",
                  "name": "_delegation",
                  "type": "address"
                },
                {
                  "internalType": "contract ISlasher",
                  "name": "_slasher",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "Deposit",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                }
              ],
              "name": "StrategyAddedToDepositWhitelist",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                }
              ],
              "name": "StrategyRemovedFromDepositWhitelist",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAddress",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAddress",
                  "type": "address"
                }
              ],
              "name": "StrategyWhitelisterChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "bool",
                  "name": "value",
                  "type": "bool"
                }
              ],
              "name": "UpdatedThirdPartyTransfersForbidden",
              "type": "event"
            },
            {
              "inputs": [],
              "name": "DEPOSIT_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "DOMAIN_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "addShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategiesToWhitelist",
                  "type": "address[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "thirdPartyTransfersForbiddenValues",
                  "type": "bool[]"
                }
              ],
              "name": "addStrategiesToDepositWhitelist",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "delegation",
              "outputs": [
                {
                  "internalType": "contract IDelegationManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "depositIntoStrategy",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                },
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                },
                {
                  "internalType": "bytes",
                  "name": "signature",
                  "type": "bytes"
                }
              ],
              "name": "depositIntoStrategyWithSignature",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "depositIntoStrategyWithStaker",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "domainSeparator",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "getDeposits",
              "outputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "initialOwner",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "initialStrategyWhitelister",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "initialPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "nonces",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "removeShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategiesToRemoveFromWhitelist",
                  "type": "address[]"
                }
              ],
              "name": "removeStrategiesFromDepositWhitelist",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newStrategyWhitelister",
                  "type": "address"
                }
              ],
              "name": "setStrategyWhitelister",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "bool",
                  "name": "value",
                  "type": "bool"
                }
              ],
              "name": "setThirdPartyTransfersForbidden",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "slasher",
              "outputs": [
                {
                  "internalType": "contract ISlasher",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "stakerStrategyList",
              "outputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "stakerStrategyListLength",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "stakerStrategyShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "strategyIsWhitelistedForDeposit",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyWhitelister",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "thirdPartyTransfersForbidden",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "recipient",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                }
              ],
              "name": "withdrawSharesAsTokens",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "StrategyManager-Proxy": {
          "address": "0xB36307B06beC1Fe002f08f98B223A62C1E3AFb88",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "slisBNB-Strategy-Proxy": {
          "address": "0x5040649097A4f4c84ce87E571b79D6Cac850d8e1",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        }
      }
    }
  ],
  "97": [
    {
      "name": "bsc-testnet",
      "chainId": "97",
      "contracts": {
        "DelegationManager-Implementation": {
          "address": "0xaD88B52FB6077B20903b15F8569Cf9B69A3F9DF7",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                },
                {
                  "internalType": "contract ISlasher",
                  "name": "_slasher",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MinWithdrawalDelaySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "newOperatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "OperatorDetailsModified",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "OperatorMetadataURIUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "operatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "OperatorRegistered",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "OperatorSharesDecreased",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "OperatorSharesIncreased",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerDelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerForceUndelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerUndelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "StrategyWithdrawalDelaySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousGateway",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "currentGateway",
                  "type": "address"
                }
              ],
              "name": "UpdateWrappedTokenGateway",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "bytes32",
                  "name": "withdrawalRoot",
                  "type": "bytes32"
                }
              ],
              "name": "WithdrawalCompleted",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "bytes32",
                  "name": "withdrawalRoot",
                  "type": "bytes32"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationManager.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                }
              ],
              "name": "WithdrawalQueued",
              "type": "event"
            },
            {
              "inputs": [],
              "name": "DELEGATION_APPROVAL_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "DOMAIN_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "MAX_STAKER_OPT_OUT_WINDOW",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "MAX_WITHDRAWAL_DELAY",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "STAKER_DELEGATION_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateCurrentStakerDelegationDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "_delegationApprover",
                  "type": "address"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateDelegationApprovalDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "_stakerNonce",
                  "type": "uint256"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateStakerDelegationDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                }
              ],
              "name": "calculateWithdrawalRoot",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "pure",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                },
                {
                  "internalType": "contract IERC20[]",
                  "name": "tokens",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256",
                  "name": "middlewareTimesIndex",
                  "type": "uint256"
                },
                {
                  "internalType": "bool",
                  "name": "receiveAsTokens",
                  "type": "bool"
                }
              ],
              "name": "completeQueuedWithdrawal",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal[]",
                  "name": "withdrawals",
                  "type": "tuple[]"
                },
                {
                  "internalType": "contract IERC20[][]",
                  "name": "tokens",
                  "type": "address[][]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "middlewareTimesIndexes",
                  "type": "uint256[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "receiveAsTokens",
                  "type": "bool[]"
                }
              ],
              "name": "completeQueuedWithdrawals",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "cumulativeWithdrawalsQueued",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "decreaseDelegatedShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "approverSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                }
              ],
              "name": "delegateTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "stakerSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "approverSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                }
              ],
              "name": "delegateToBySignature",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "delegatedTo",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "delegationApprover",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "name": "delegationApproverSaltIsSpent",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "domainSeparator",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "earningsReceiver",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "getDelegatableShares",
              "outputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategies",
                  "type": "address[]"
                }
              ],
              "name": "getOperatorShares",
              "outputs": [
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategies",
                  "type": "address[]"
                }
              ],
              "name": "getWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "increaseDelegatedShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "initialOwner",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "initialPausedStatus",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "_minWithdrawalDelay",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IStrategy[]",
                  "name": "_strategies",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "_withdrawalDelay",
                  "type": "uint256[]"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "isDelegated",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "isOperator",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "minWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "newOperatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "modifyOperatorDetails",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "operatorDetails",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "name": "pendingWithdrawals",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    }
                  ],
                  "internalType": "struct IDelegationManager.QueuedWithdrawalParams[]",
                  "name": "queuedWithdrawalParams",
                  "type": "tuple[]"
                }
              ],
              "name": "queueWithdrawals",
              "outputs": [
                {
                  "internalType": "bytes32[]",
                  "name": "",
                  "type": "bytes32[]"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationManager.OperatorDetails",
                  "name": "registeringOperatorDetails",
                  "type": "tuple"
                },
                {
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "registerAsOperator",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newMinWithdrawalDelay",
                  "type": "uint256"
                }
              ],
              "name": "setMinWithdrawalDelay",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategies",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "withdrawalDelay",
                  "type": "uint256[]"
                }
              ],
              "name": "setStrategyWithdrawalDelay",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "slasher",
              "outputs": [
                {
                  "internalType": "contract ISlasher",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "stakerNonce",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "stakerOptOutWindow",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyManager",
              "outputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "strategyWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "undelegate",
              "outputs": [
                {
                  "internalType": "bytes32[]",
                  "name": "withdrawalRoots",
                  "type": "bytes32[]"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "updateOperatorMetadataURI",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_newWrappedTokenGateway",
                  "type": "address"
                }
              ],
              "name": "updateWrappedTokenGateway",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "wrappedTokenGateway",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "DelegationManager-Proxy": {
          "address": "0xcC4b428124C226c87872d0932b7549DA0531e364",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "EmptyContract": {
          "address": "0xea4076BD0D1Ad6C4f921a3dd557D4b450027F21D",
          "abi": [
            {
              "inputs": [],
              "name": "foo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "ListaGateway": {
          "address": "0x4618ADb8A2439F498D292392FB97A4Cf7c9A964F",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_slisBNB",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "_owner",
                  "type": "address"
                },
                {
                  "internalType": "contract IListaStakeManager",
                  "name": "_listaStakeManager",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "_strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "depositNativeToken",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "emergencyNativeTransfer",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "emergencyTokenTransfer",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getListaContracts",
              "outputs": [
                {
                  "internalType": "contract IListaStakeManager",
                  "name": "_listaStakeManager",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "_slisBNB",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getStrategyAddress",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "PauserRegistry": {
          "address": "0xc442a3598947Da34baC42bE34586675C215f57dE",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address[]",
                  "name": "_pausers",
                  "type": "address[]"
                },
                {
                  "internalType": "address",
                  "name": "_unpauser",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "pauser",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "bool",
                  "name": "canPause",
                  "type": "bool"
                }
              ],
              "name": "PauserStatusChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousUnpauser",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newUnpauser",
                  "type": "address"
                }
              ],
              "name": "UnpauserChanged",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "isPauser",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newPauser",
                  "type": "address"
                },
                {
                  "internalType": "bool",
                  "name": "canPause",
                  "type": "bool"
                }
              ],
              "name": "setIsPauser",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newUnpauser",
                  "type": "address"
                }
              ],
              "name": "setUnpauser",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "unpauser",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "ProxyAdmin": {
          "address": "0xadb053b67dF6091C39F7c2E522CF6Fecd5d9BCF8",
          "abi": [
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeProxyAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                }
              ],
              "name": "getProxyAdmin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                }
              ],
              "name": "getProxyImplementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "upgrade",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract TransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            }
          ]
        },
        "Slasher-Implementation": {
          "address": "0x74759803bDC6DE7C39c452D652E4A55f8870Fc7e",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IDelegationManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previouslySlashedAddress",
                  "type": "address"
                }
              ],
              "name": "FrozenStatusReset",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "index",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "stalestUpdateTimestamp",
                  "type": "uint32"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "latestServeUntilTimestamp",
                  "type": "uint32"
                }
              ],
              "name": "MiddlewareTimesAdded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "slashedOperator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "slashingContract",
                  "type": "address"
                }
              ],
              "name": "OperatorFrozen",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                }
              ],
              "name": "OptedIntoSlashing",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "contractCanSlashOperatorUntilTimestamp",
                  "type": "uint32"
                }
              ],
              "name": "SlashingAbilityRevoked",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "canSlash",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "canWithdraw",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "contractCanSlashOperatorUntilTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "delegation",
              "outputs": [
                {
                  "internalType": "contract IDelegationManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "freezeOperator",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getCorrectValueForInsertAfter",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getMiddlewareTimesIndexServeUntilTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getMiddlewareTimesIndexStalestUpdateTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "isFrozen",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "latestUpdateTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "middlewareTimesLength",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "operatorToMiddlewareTimes",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "uint32",
                      "name": "stalestUpdateTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "latestServeUntilTimestamp",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct ISlasher.MiddlewareTimes",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorWhitelistedContractsLinkedListEntry",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorWhitelistedContractsLinkedListSize",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "optIntoSlashing",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "recordFirstStakeUpdate",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "recordLastStakeUpdateAndRevokeSlashingAbility",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "recordStakeUpdate",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address[]",
                  "name": "",
                  "type": "address[]"
                }
              ],
              "name": "resetFrozenStatus",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyManager",
              "outputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "whitelistedContractDetails",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "uint32",
                      "name": "registrationMayBeginAtTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "contractCanSlashOperatorUntilTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "latestUpdateTimestamp",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct ISlasher.MiddlewareDetails",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "Slasher-Proxy": {
          "address": "0x9Cf182755a5F974B896590e1DaD99219217C2Da8",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "Strategy-Implementation": {
          "address": "0xD5dbA71698E52271438B90E1fC658cB94F637aa1",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MaxPerDepositUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MaxTotalDepositsUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "deposit",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "newShares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "explanation",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "pure",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getTVLLimits",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "_maxPerDeposit",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "_maxTotalDeposits",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "_underlyingToken",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "_underlyingToken",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                }
              ],
              "name": "initializeBase",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "maxPerDeposit",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "maxTotalDeposits",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newMaxPerDeposit",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "newMaxTotalDeposits",
                  "type": "uint256"
                }
              ],
              "name": "setTVLLimits",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "shares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "sharesToUnderlying",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "sharesToUnderlyingView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyManager",
              "outputs": [
                {
                  "internalType": "contract IStrategyManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "totalShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountUnderlying",
                  "type": "uint256"
                }
              ],
              "name": "underlyingToShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountUnderlying",
                  "type": "uint256"
                }
              ],
              "name": "underlyingToSharesView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "underlyingToken",
              "outputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "userUnderlying",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "userUnderlyingView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "recipient",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "withdraw",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "StrategyManager-Implementation": {
          "address": "0xE2a60661B9675583F9446df0128E0cf4BFee963a",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IDelegationManager",
                  "name": "_delegation",
                  "type": "address"
                },
                {
                  "internalType": "contract ISlasher",
                  "name": "_slasher",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "Deposit",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                }
              ],
              "name": "StrategyAddedToDepositWhitelist",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                }
              ],
              "name": "StrategyRemovedFromDepositWhitelist",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAddress",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAddress",
                  "type": "address"
                }
              ],
              "name": "StrategyWhitelisterChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "bool",
                  "name": "value",
                  "type": "bool"
                }
              ],
              "name": "UpdatedThirdPartyTransfersForbidden",
              "type": "event"
            },
            {
              "inputs": [],
              "name": "DEPOSIT_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "DOMAIN_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "addShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategiesToWhitelist",
                  "type": "address[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "thirdPartyTransfersForbiddenValues",
                  "type": "bool[]"
                }
              ],
              "name": "addStrategiesToDepositWhitelist",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "delegation",
              "outputs": [
                {
                  "internalType": "contract IDelegationManager",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "depositIntoStrategy",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                },
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                },
                {
                  "internalType": "bytes",
                  "name": "signature",
                  "type": "bytes"
                }
              ],
              "name": "depositIntoStrategyWithSignature",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "depositIntoStrategyWithStaker",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "domainSeparator",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "getDeposits",
              "outputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "initialOwner",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "initialStrategyWhitelister",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "initialPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "nonces",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "removeShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy[]",
                  "name": "strategiesToRemoveFromWhitelist",
                  "type": "address[]"
                }
              ],
              "name": "removeStrategiesFromDepositWhitelist",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newStrategyWhitelister",
                  "type": "address"
                }
              ],
              "name": "setStrategyWhitelister",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "bool",
                  "name": "value",
                  "type": "bool"
                }
              ],
              "name": "setThirdPartyTransfersForbidden",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "slasher",
              "outputs": [
                {
                  "internalType": "contract ISlasher",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "stakerStrategyList",
              "outputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "stakerStrategyListLength",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "stakerStrategyShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "strategyIsWhitelistedForDeposit",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "strategyWhitelister",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IStrategy",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "thirdPartyTransfersForbidden",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "recipient",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "strategy",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                }
              ],
              "name": "withdrawSharesAsTokens",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "StrategyManager-Proxy": {
          "address": "0xd914b20BcFcdBb396697a957fc28a77947593bE5",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "WBNB-Strategy-Proxy": {
          "address": "0xE72A7b6bf9bEAA55449986e402329c262954a675",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "WrappedTokenGateway": {
          "address": "0xB12699f228cE32f74DCd46b5265D16B9e1B25DfD",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_wrappedToken",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "_owner",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategy",
                  "name": "_strategy",
                  "type": "address"
                },
                {
                  "internalType": "contract IStrategyManager",
                  "name": "_strategyManager",
                  "type": "address"
                },
                {
                  "internalType": "contract IDelegationManager",
                  "name": "_delegationManager",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "depositNativeToken",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "emergencyNativeTransfer",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "emergencyTokenTransfer",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getStrategyAddress",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getWrappedTokenAddress",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IStrategy[]",
                      "name": "strategies",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationManager.Withdrawal[]",
                  "name": "withdrawals",
                  "type": "tuple[]"
                },
                {
                  "internalType": "contract IERC20[][]",
                  "name": "tokens",
                  "type": "address[][]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "middlewareTimesIndexs",
                  "type": "uint256[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "receiveAsTokens",
                  "type": "bool[]"
                }
              ],
              "name": "withdrawNativeTokens",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "slisBNB-Strategy-Proxy": {
          "address": "0xFbA63a28d7CEbF397dC9D5F6a5fA44eE83597001",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "stBNB-Strategy-Proxy": {
          "address": "0xddE6B9eB6914AB0E75e800D14E5A72008ed6679E",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "inputs": [],
              "name": "admin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "implementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "implementation_",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                }
              ],
              "name": "upgradeTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newImplementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeToAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        }
      }
    }
  ],
  "31337": [
    {
      "name": "hardhat",
      "chainId": "31337",
      "contracts": {
        "PauserRegistry": {
          "address": "0x5FbDB2315678afecb367f032d93F642f64180aa3",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address[]",
                  "name": "_pausers",
                  "type": "address[]"
                },
                {
                  "internalType": "address",
                  "name": "_unpauser",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "pauser",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "bool",
                  "name": "canPause",
                  "type": "bool"
                }
              ],
              "name": "PauserStatusChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousUnpauser",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newUnpauser",
                  "type": "address"
                }
              ],
              "name": "UnpauserChanged",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "isPauser",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newPauser",
                  "type": "address"
                },
                {
                  "internalType": "bool",
                  "name": "canPause",
                  "type": "bool"
                }
              ],
              "name": "setIsPauser",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newUnpauser",
                  "type": "address"
                }
              ],
              "name": "setUnpauser",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "unpauser",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "EmptyContract": {
          "address": "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
          "abi": [
            {
              "inputs": [],
              "name": "foo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "ProxyAdmin": {
          "address": "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
          "abi": [
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "contract ITransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "changeProxyAdmin",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract ITransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                }
              ],
              "name": "getProxyAdmin",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract ITransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                }
              ],
              "name": "getProxyImplementation",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract ITransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "upgrade",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract ITransparentUpgradeableProxy",
                  "name": "proxy",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "name": "upgradeAndCall",
              "outputs": [],
              "stateMutability": "payable",
              "type": "function"
            }
          ]
        },
        "PoolController-Proxy": {
          "address": "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "DelegationController-Proxy": {
          "address": "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "Slasher-Proxy": {
          "address": "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "DelegationController-Implementation": {
          "address": "0x0165878A594ca255338adfa4d48449f69242Eb8F",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IPoolController",
                  "name": "_poolController",
                  "type": "address"
                },
                {
                  "internalType": "contract ISlasher",
                  "name": "_slasher",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MinWithdrawalDelaySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationController.OperatorDetails",
                  "name": "newOperatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "OperatorDetailsModified",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "OperatorMetadataURIUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationController.OperatorDetails",
                  "name": "operatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "OperatorRegistered",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "OperatorSharesDecreased",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "OperatorSharesIncreased",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "PoolWithdrawalDelaySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerDelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerForceUndelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "StakerUndelegated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousGateway",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "currentGateway",
                  "type": "address"
                }
              ],
              "name": "UpdateWrappedTokenGateway",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "bytes32",
                  "name": "withdrawalRoot",
                  "type": "bytes32"
                }
              ],
              "name": "WithdrawalCompleted",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "bytes32",
                  "name": "withdrawalRoot",
                  "type": "bytes32"
                },
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IPool[]",
                      "name": "pools",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "indexed": false,
                  "internalType": "struct IDelegationController.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                }
              ],
              "name": "WithdrawalQueued",
              "type": "event"
            },
            {
              "inputs": [],
              "name": "DELEGATION_APPROVAL_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "DOMAIN_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "MAX_STAKER_OPT_OUT_WINDOW",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "MAX_WITHDRAWAL_DELAY",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "STAKER_DELEGATION_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateCurrentStakerDelegationDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "_delegationApprover",
                  "type": "address"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateDelegationApprovalDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "_stakerNonce",
                  "type": "uint256"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                }
              ],
              "name": "calculateStakerDelegationDigestHash",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IPool[]",
                      "name": "pools",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationController.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                }
              ],
              "name": "calculateWithdrawalRoot",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "pure",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "cumulativeWithdrawalsQueued",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "decreaseDelegatedShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "approverSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                }
              ],
              "name": "delegateTo",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "stakerSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "components": [
                    {
                      "internalType": "bytes",
                      "name": "signature",
                      "type": "bytes"
                    },
                    {
                      "internalType": "uint256",
                      "name": "expiry",
                      "type": "uint256"
                    }
                  ],
                  "internalType": "struct ISignatureUtils.SignatureWithExpiry",
                  "name": "approverSignatureAndExpiry",
                  "type": "tuple"
                },
                {
                  "internalType": "bytes32",
                  "name": "approverSalt",
                  "type": "bytes32"
                }
              ],
              "name": "delegateToBySignature",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "delegatedTo",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "delegationApprover",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "name": "delegationApproverSaltIsSpent",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "domainSeparator",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "earningsReceiver",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "getDelegatableShares",
              "outputs": [
                {
                  "internalType": "contract IPool[]",
                  "name": "",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool[]",
                  "name": "pools",
                  "type": "address[]"
                }
              ],
              "name": "getOperatorShares",
              "outputs": [
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool[]",
                  "name": "pools",
                  "type": "address[]"
                }
              ],
              "name": "getWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "increaseDelegatedShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "initialOwner",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "initialPausedStatus",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "_minWithdrawalDelay",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IPool[]",
                  "name": "_pools",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "_withdrawalDelay",
                  "type": "uint256[]"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "isDelegated",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "isOperator",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "minWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationController.OperatorDetails",
                  "name": "newOperatorDetails",
                  "type": "tuple"
                }
              ],
              "name": "modifyOperatorDetails",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "operatorDetails",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationController.OperatorDetails",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "name": "pendingWithdrawals",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "poolController",
              "outputs": [
                {
                  "internalType": "contract IPoolController",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "poolWithdrawalDelay",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "earningsReceiver",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegationApprover",
                      "type": "address"
                    },
                    {
                      "internalType": "uint32",
                      "name": "stakerOptOutWindow",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct IDelegationController.OperatorDetails",
                  "name": "registeringOperatorDetails",
                  "type": "tuple"
                },
                {
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "registerAsOperator",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newMinWithdrawalDelay",
                  "type": "uint256"
                }
              ],
              "name": "setMinWithdrawalDelay",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool[]",
                  "name": "pools",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "withdrawalDelay",
                  "type": "uint256[]"
                }
              ],
              "name": "setPoolWithdrawalDelay",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "slasher",
              "outputs": [
                {
                  "internalType": "contract ISlasher",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "stakerNonce",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                }
              ],
              "name": "stakerOptOutWindow",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "undelegate",
              "outputs": [
                {
                  "internalType": "bytes32[]",
                  "name": "withdrawalRoots",
                  "type": "bytes32[]"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "contract IPool[]",
                      "name": "pools",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    }
                  ],
                  "internalType": "struct IDelegationController.UnstakeParams[]",
                  "name": "unstakeParams",
                  "type": "tuple[]"
                }
              ],
              "name": "unstakes",
              "outputs": [
                {
                  "internalType": "bytes32[]",
                  "name": "",
                  "type": "bytes32[]"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "metadataURI",
                  "type": "string"
                }
              ],
              "name": "updateOperatorMetadataURI",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_newWrappedTokenGateway",
                  "type": "address"
                }
              ],
              "name": "updateWrappedTokenGateway",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IPool[]",
                      "name": "pools",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationController.Withdrawal",
                  "name": "withdrawal",
                  "type": "tuple"
                },
                {
                  "internalType": "contract IERC20[]",
                  "name": "tokens",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256",
                  "name": "middlewareTimesIndex",
                  "type": "uint256"
                },
                {
                  "internalType": "bool",
                  "name": "receiveAsTokens",
                  "type": "bool"
                }
              ],
              "name": "withdraw",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "components": [
                    {
                      "internalType": "address",
                      "name": "staker",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "delegatedTo",
                      "type": "address"
                    },
                    {
                      "internalType": "address",
                      "name": "withdrawer",
                      "type": "address"
                    },
                    {
                      "internalType": "uint256",
                      "name": "nonce",
                      "type": "uint256"
                    },
                    {
                      "internalType": "uint32",
                      "name": "startTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "contract IPool[]",
                      "name": "pools",
                      "type": "address[]"
                    },
                    {
                      "internalType": "uint256[]",
                      "name": "shares",
                      "type": "uint256[]"
                    }
                  ],
                  "internalType": "struct IDelegationController.Withdrawal[]",
                  "name": "withdrawals",
                  "type": "tuple[]"
                },
                {
                  "internalType": "contract IERC20[][]",
                  "name": "tokens",
                  "type": "address[][]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "middlewareTimesIndexes",
                  "type": "uint256[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "receiveAsTokens",
                  "type": "bool[]"
                }
              ],
              "name": "withdraws",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "wrappedTokenGateway",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "PoolController-Implementation": {
          "address": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IDelegationController",
                  "name": "_delegation",
                  "type": "address"
                },
                {
                  "internalType": "contract ISlasher",
                  "name": "_slasher",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "Deposit",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                }
              ],
              "name": "PoolAddedToDepositWhitelist",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                }
              ],
              "name": "PoolRemovedFromDepositWhitelist",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAddress",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAddress",
                  "type": "address"
                }
              ],
              "name": "PoolWhitelisterChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "bool",
                  "name": "value",
                  "type": "bool"
                }
              ],
              "name": "UpdatedThirdPartyTransfersForbidden",
              "type": "event"
            },
            {
              "inputs": [],
              "name": "DEPOSIT_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "DOMAIN_TYPEHASH",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool[]",
                  "name": "poolsToWhitelist",
                  "type": "address[]"
                },
                {
                  "internalType": "bool[]",
                  "name": "thirdPartyTransfersForbiddenValues",
                  "type": "bool[]"
                }
              ],
              "name": "addPoolsToDepositWhitelist",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "addShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "delegation",
              "outputs": [
                {
                  "internalType": "contract IDelegationController",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "depositIntoPool",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                },
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "expiry",
                  "type": "uint256"
                },
                {
                  "internalType": "bytes",
                  "name": "signature",
                  "type": "bytes"
                }
              ],
              "name": "depositIntoPoolWithSignature",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "depositIntoPoolWithStaker",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "domainSeparator",
              "outputs": [
                {
                  "internalType": "bytes32",
                  "name": "",
                  "type": "bytes32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "getDeposits",
              "outputs": [
                {
                  "internalType": "contract IPool[]",
                  "name": "",
                  "type": "address[]"
                },
                {
                  "internalType": "uint256[]",
                  "name": "",
                  "type": "uint256[]"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "initialOwner",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "initialPoolWhitelister",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "initialPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "nonces",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "poolIsWhitelistedForDeposit",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "poolWhitelister",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool[]",
                  "name": "poolsToRemoveFromWhitelist",
                  "type": "address[]"
                }
              ],
              "name": "removePoolsFromDepositWhitelist",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                }
              ],
              "name": "removeShares",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newPoolWhitelister",
                  "type": "address"
                }
              ],
              "name": "setPoolWhitelister",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "bool",
                  "name": "value",
                  "type": "bool"
                }
              ],
              "name": "setThirdPartyTransfersForbidden",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "slasher",
              "outputs": [
                {
                  "internalType": "contract ISlasher",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "stakerPoolList",
              "outputs": [
                {
                  "internalType": "contract IPool",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "staker",
                  "type": "address"
                }
              ],
              "name": "stakerPoolListLength",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "stakerPoolShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPool",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "thirdPartyTransfersForbidden",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "recipient",
                  "type": "address"
                },
                {
                  "internalType": "contract IPool",
                  "name": "pool",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "shares",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                }
              ],
              "name": "withdrawSharesAsTokens",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "Slasher-Implementation": {
          "address": "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IPoolController",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IDelegationController",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previouslySlashedAddress",
                  "type": "address"
                }
              ],
              "name": "FrozenStatusReset",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "index",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "stalestUpdateTimestamp",
                  "type": "uint32"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "latestServeUntilTimestamp",
                  "type": "uint32"
                }
              ],
              "name": "MiddlewareTimesAdded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "slashedOperator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "slashingContract",
                  "type": "address"
                }
              ],
              "name": "OperatorFrozen",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                }
              ],
              "name": "OptedIntoSlashing",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "operator",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint32",
                  "name": "contractCanSlashOperatorUntilTimestamp",
                  "type": "uint32"
                }
              ],
              "name": "SlashingAbilityRevoked",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "canSlash",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "canWithdraw",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "contractCanSlashOperatorUntilTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "delegation",
              "outputs": [
                {
                  "internalType": "contract IDelegationController",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "freezeOperator",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getCorrectValueForInsertAfter",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getMiddlewareTimesIndexServeUntilTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "getMiddlewareTimesIndexStalestUpdateTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "isFrozen",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "latestUpdateTimestamp",
              "outputs": [
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "middlewareTimesLength",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "operatorToMiddlewareTimes",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "uint32",
                      "name": "stalestUpdateTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "latestServeUntilTimestamp",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct ISlasher.MiddlewareTimes",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorWhitelistedContractsLinkedListEntry",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "operatorWhitelistedContractsLinkedListSize",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "optIntoSlashing",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "poolController",
              "outputs": [
                {
                  "internalType": "contract IPoolController",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "recordFirstStakeUpdate",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                }
              ],
              "name": "recordLastStakeUpdateAndRevokeSlashingAbility",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint32",
                  "name": "",
                  "type": "uint32"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "name": "recordStakeUpdate",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address[]",
                  "name": "",
                  "type": "address[]"
                }
              ],
              "name": "resetFrozenStatus",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "whitelistedContractDetails",
              "outputs": [
                {
                  "components": [
                    {
                      "internalType": "uint32",
                      "name": "registrationMayBeginAtTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "contractCanSlashOperatorUntilTimestamp",
                      "type": "uint32"
                    },
                    {
                      "internalType": "uint32",
                      "name": "latestUpdateTimestamp",
                      "type": "uint32"
                    }
                  ],
                  "internalType": "struct ISlasher.MiddlewareDetails",
                  "name": "",
                  "type": "tuple"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            }
          ]
        },
        "Pool-Implementation": {
          "address": "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "contract IPoolController",
                  "name": "_poolController",
                  "type": "address"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint8",
                  "name": "version",
                  "type": "uint8"
                }
              ],
              "name": "Initialized",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MaxPerDepositUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "previousValue",
                  "type": "uint256"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newValue",
                  "type": "uint256"
                }
              ],
              "name": "MaxTotalDepositsUpdated",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Paused",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "pauserRegistry",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "PauserRegistrySet",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "Unpaused",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "deposit",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "newShares",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "explanation",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "pure",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "getTVLLimits",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "_maxPerDeposit",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "_maxTotalDeposits",
                  "type": "uint256"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "_underlyingToken",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                }
              ],
              "name": "initialize",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "_underlyingToken",
                  "type": "address"
                },
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "_pauserRegistry",
                  "type": "address"
                }
              ],
              "name": "initializeBase",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "maxPerDeposit",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "maxTotalDeposits",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "pause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauseAll",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint8",
                  "name": "index",
                  "type": "uint8"
                }
              ],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "paused",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "pauserRegistry",
              "outputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "poolController",
              "outputs": [
                {
                  "internalType": "contract IPoolController",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "contract IPauserRegistry",
                  "name": "newPauserRegistry",
                  "type": "address"
                }
              ],
              "name": "setPauserRegistry",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newMaxPerDeposit",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "newMaxTotalDeposits",
                  "type": "uint256"
                }
              ],
              "name": "setTVLLimits",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "shares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "sharesToUnderlying",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "sharesToUnderlyingView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "totalShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountUnderlying",
                  "type": "uint256"
                }
              ],
              "name": "underlyingToShares",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "amountUnderlying",
                  "type": "uint256"
                }
              ],
              "name": "underlyingToSharesView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "underlyingToken",
              "outputs": [
                {
                  "internalType": "contract IERC20",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "newPausedStatus",
                  "type": "uint256"
                }
              ],
              "name": "unpause",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "userUnderlying",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "user",
                  "type": "address"
                }
              ],
              "name": "userUnderlyingView",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "recipient",
                  "type": "address"
                },
                {
                  "internalType": "contract IERC20",
                  "name": "token",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amountShares",
                  "type": "uint256"
                }
              ],
              "name": "withdraw",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "stBNB-TestnetMintableERC20": {
          "address": "0x610178dA211FEF7D417bC0e6FeD39F05609AD788",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "name",
                  "type": "string"
                },
                {
                  "internalType": "string",
                  "name": "symbol",
                  "type": "string"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "owner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "Approval",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "from",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "Transfer",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "owner",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                }
              ],
              "name": "allowance",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "approve",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                }
              ],
              "name": "balanceOf",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "decimals",
              "outputs": [
                {
                  "internalType": "uint8",
                  "name": "",
                  "type": "uint8"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "subtractedValue",
                  "type": "uint256"
                }
              ],
              "name": "decreaseAllowance",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "addedValue",
                  "type": "uint256"
                }
              ],
              "name": "increaseAllowance",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "mint",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "mint",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "name",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "symbol",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "totalSupply",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "transfer",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "from",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "transferFrom",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "slisBNB-TestnetMintableERC20": {
          "address": "0xB7f8BC63BbcaD18155201308C8f3540b07f84F5e",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "name",
                  "type": "string"
                },
                {
                  "internalType": "string",
                  "name": "symbol",
                  "type": "string"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "owner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "Approval",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "previousOwner",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "OwnershipTransferred",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "from",
                  "type": "address"
                },
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "Transfer",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "owner",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                }
              ],
              "name": "allowance",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "approve",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "account",
                  "type": "address"
                }
              ],
              "name": "balanceOf",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "decimals",
              "outputs": [
                {
                  "internalType": "uint8",
                  "name": "",
                  "type": "uint8"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "subtractedValue",
                  "type": "uint256"
                }
              ],
              "name": "decreaseAllowance",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "spender",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "addedValue",
                  "type": "uint256"
                }
              ],
              "name": "increaseAllowance",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "mint",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "uint256",
                  "name": "value",
                  "type": "uint256"
                }
              ],
              "name": "mint",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "name",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "renounceOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "symbol",
              "outputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "totalSupply",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "transfer",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "from",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "to",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "transferFrom",
              "outputs": [
                {
                  "internalType": "bool",
                  "name": "",
                  "type": "bool"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "newOwner",
                  "type": "address"
                }
              ],
              "name": "transferOwnership",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]
        },
        "stBNB-Pool-Proxy": {
          "address": "0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        },
        "slisBNB-Pool-Proxy": {
          "address": "0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82",
          "abi": [
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "_logic",
                  "type": "address"
                },
                {
                  "internalType": "address",
                  "name": "admin_",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "_data",
                  "type": "bytes"
                }
              ],
              "stateMutability": "payable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "previousAdmin",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "newAdmin",
                  "type": "address"
                }
              ],
              "name": "AdminChanged",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "beacon",
                  "type": "address"
                }
              ],
              "name": "BeaconUpgraded",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": true,
                  "internalType": "address",
                  "name": "implementation",
                  "type": "address"
                }
              ],
              "name": "Upgraded",
              "type": "event"
            },
            {
              "stateMutability": "payable",
              "type": "fallback"
            },
            {
              "stateMutability": "payable",
              "type": "receive"
            }
          ]
        }
      }
    }
  ]
} as const;