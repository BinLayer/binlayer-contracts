// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DelegationControllerStorage

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IDelegationControllerOperatorDetails is an auto generated low-level Go binding around an user-defined struct.
type IDelegationControllerOperatorDetails struct {
	EarningsReceiver   common.Address
	DelegationApprover common.Address
	StakerOptOutWindow uint32
}

// IDelegationControllerUnstakeParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationControllerUnstakeParams struct {
	Pools      []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

// IDelegationControllerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationControllerWithdrawal struct {
	Staker         common.Address
	DelegatedTo    common.Address
	Withdrawer     common.Address
	Nonce          *big.Int
	StartTimestamp uint32
	Pools          []common.Address
	Shares         []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// DelegationControllerStorageMetaData contains all meta data concerning the DelegationControllerStorage contract.
var DelegationControllerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_WITHDRAWAL_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakerNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateToBySignature\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"earningsReceiver\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalDelay\",\"inputs\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolWithdrawalDelay\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerOptOutWindow\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakes\",\"inputs\":[{\"name\":\"unstakeParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.UnstakeParams[]\",\"components\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraws\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrappedTokenGateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MinWithdrawalDelaySet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDetailsModified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolWithdrawalDelaySet\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateWrappedTokenGateway\",\"inputs\":[{\"name\":\"previousGateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currentGateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false}]",
}

// DelegationControllerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegationControllerStorageMetaData.ABI instead.
var DelegationControllerStorageABI = DelegationControllerStorageMetaData.ABI

// DelegationControllerStorage is an auto generated Go binding around an Ethereum contract.
type DelegationControllerStorage struct {
	DelegationControllerStorageCaller     // Read-only binding to the contract
	DelegationControllerStorageTransactor // Write-only binding to the contract
	DelegationControllerStorageFilterer   // Log filterer for contract events
}

// DelegationControllerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegationControllerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationControllerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegationControllerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationControllerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegationControllerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationControllerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegationControllerStorageSession struct {
	Contract     *DelegationControllerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DelegationControllerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegationControllerStorageCallerSession struct {
	Contract *DelegationControllerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// DelegationControllerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegationControllerStorageTransactorSession struct {
	Contract     *DelegationControllerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// DelegationControllerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegationControllerStorageRaw struct {
	Contract *DelegationControllerStorage // Generic contract binding to access the raw methods on
}

// DelegationControllerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegationControllerStorageCallerRaw struct {
	Contract *DelegationControllerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// DelegationControllerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegationControllerStorageTransactorRaw struct {
	Contract *DelegationControllerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegationControllerStorage creates a new instance of DelegationControllerStorage, bound to a specific deployed contract.
func NewDelegationControllerStorage(address common.Address, backend bind.ContractBackend) (*DelegationControllerStorage, error) {
	contract, err := bindDelegationControllerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorage{DelegationControllerStorageCaller: DelegationControllerStorageCaller{contract: contract}, DelegationControllerStorageTransactor: DelegationControllerStorageTransactor{contract: contract}, DelegationControllerStorageFilterer: DelegationControllerStorageFilterer{contract: contract}}, nil
}

// NewDelegationControllerStorageCaller creates a new read-only instance of DelegationControllerStorage, bound to a specific deployed contract.
func NewDelegationControllerStorageCaller(address common.Address, caller bind.ContractCaller) (*DelegationControllerStorageCaller, error) {
	contract, err := bindDelegationControllerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageCaller{contract: contract}, nil
}

// NewDelegationControllerStorageTransactor creates a new write-only instance of DelegationControllerStorage, bound to a specific deployed contract.
func NewDelegationControllerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegationControllerStorageTransactor, error) {
	contract, err := bindDelegationControllerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageTransactor{contract: contract}, nil
}

// NewDelegationControllerStorageFilterer creates a new log filterer instance of DelegationControllerStorage, bound to a specific deployed contract.
func NewDelegationControllerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegationControllerStorageFilterer, error) {
	contract, err := bindDelegationControllerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageFilterer{contract: contract}, nil
}

// bindDelegationControllerStorage binds a generic wrapper to an already deployed contract.
func bindDelegationControllerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegationControllerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationControllerStorage *DelegationControllerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegationControllerStorage.Contract.DelegationControllerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationControllerStorage *DelegationControllerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DelegationControllerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationControllerStorage *DelegationControllerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DelegationControllerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationControllerStorage *DelegationControllerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegationControllerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationControllerStorage *DelegationControllerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationControllerStorage *DelegationControllerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.DELEGATIONAPPROVALTYPEHASH(&_DelegationControllerStorage.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.DELEGATIONAPPROVALTYPEHASH(&_DelegationControllerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.DOMAINTYPEHASH(&_DelegationControllerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.DOMAINTYPEHASH(&_DelegationControllerStorage.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) MAXWITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _DelegationControllerStorage.Contract.MAXWITHDRAWALDELAY(&_DelegationControllerStorage.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _DelegationControllerStorage.Contract.MAXWITHDRAWALDELAY(&_DelegationControllerStorage.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) STAKERDELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "STAKER_DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.STAKERDELEGATIONTYPEHASH(&_DelegationControllerStorage.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.STAKERDELEGATIONTYPEHASH(&_DelegationControllerStorage.CallOpts)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) CalculateCurrentStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "calculateCurrentStakerDelegationDigestHash", staker, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateCurrentStakerDelegationDigestHash(&_DelegationControllerStorage.CallOpts, staker, operator, expiry)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateCurrentStakerDelegationDigestHash(&_DelegationControllerStorage.CallOpts, staker, operator, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, _delegationApprover, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateDelegationApprovalDigestHash(&_DelegationControllerStorage.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateDelegationApprovalDigestHash(&_DelegationControllerStorage.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) CalculateStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "calculateStakerDelegationDigestHash", staker, _stakerNonce, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateStakerDelegationDigestHash(&_DelegationControllerStorage.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateStakerDelegationDigestHash(&_DelegationControllerStorage.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) CalculateWithdrawalRoot(withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateWithdrawalRoot(&_DelegationControllerStorage.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	return _DelegationControllerStorage.Contract.CalculateWithdrawalRoot(&_DelegationControllerStorage.CallOpts, withdrawal)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.CumulativeWithdrawalsQueued(&_DelegationControllerStorage.CallOpts, arg0)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.CumulativeWithdrawalsQueued(&_DelegationControllerStorage.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) DelegatedTo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "delegatedTo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _DelegationControllerStorage.Contract.DelegatedTo(&_DelegationControllerStorage.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _DelegationControllerStorage.Contract.DelegatedTo(&_DelegationControllerStorage.CallOpts, arg0)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _DelegationControllerStorage.Contract.DelegationApprover(&_DelegationControllerStorage.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _DelegationControllerStorage.Contract.DelegationApprover(&_DelegationControllerStorage.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "delegationApproverSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _DelegationControllerStorage.Contract.DelegationApproverSaltIsSpent(&_DelegationControllerStorage.CallOpts, arg0, arg1)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _DelegationControllerStorage.Contract.DelegationApproverSaltIsSpent(&_DelegationControllerStorage.CallOpts, arg0, arg1)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageSession) DomainSeparator() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.DomainSeparator(&_DelegationControllerStorage.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) DomainSeparator() ([32]byte, error) {
	return _DelegationControllerStorage.Contract.DomainSeparator(&_DelegationControllerStorage.CallOpts)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) EarningsReceiver(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "earningsReceiver", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _DelegationControllerStorage.Contract.EarningsReceiver(&_DelegationControllerStorage.CallOpts, operator)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _DelegationControllerStorage.Contract.EarningsReceiver(&_DelegationControllerStorage.CallOpts, operator)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_DelegationControllerStorage *DelegationControllerStorageCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address, pools []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "getOperatorShares", operator, pools)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_DelegationControllerStorage *DelegationControllerStorageSession) GetOperatorShares(operator common.Address, pools []common.Address) ([]*big.Int, error) {
	return _DelegationControllerStorage.Contract.GetOperatorShares(&_DelegationControllerStorage.CallOpts, operator, pools)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) GetOperatorShares(operator common.Address, pools []common.Address) ([]*big.Int, error) {
	return _DelegationControllerStorage.Contract.GetOperatorShares(&_DelegationControllerStorage.CallOpts, operator, pools)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) GetWithdrawalDelay(opts *bind.CallOpts, pools []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "getWithdrawalDelay", pools)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) GetWithdrawalDelay(pools []common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.GetWithdrawalDelay(&_DelegationControllerStorage.CallOpts, pools)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) GetWithdrawalDelay(pools []common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.GetWithdrawalDelay(&_DelegationControllerStorage.CallOpts, pools)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageSession) IsDelegated(staker common.Address) (bool, error) {
	return _DelegationControllerStorage.Contract.IsDelegated(&_DelegationControllerStorage.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _DelegationControllerStorage.Contract.IsDelegated(&_DelegationControllerStorage.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageSession) IsOperator(operator common.Address) (bool, error) {
	return _DelegationControllerStorage.Contract.IsOperator(&_DelegationControllerStorage.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _DelegationControllerStorage.Contract.IsOperator(&_DelegationControllerStorage.CallOpts, operator)
}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) MinWithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "minWithdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) MinWithdrawalDelay() (*big.Int, error) {
	return _DelegationControllerStorage.Contract.MinWithdrawalDelay(&_DelegationControllerStorage.CallOpts)
}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) MinWithdrawalDelay() (*big.Int, error) {
	return _DelegationControllerStorage.Contract.MinWithdrawalDelay(&_DelegationControllerStorage.CallOpts)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationControllerStorage *DelegationControllerStorageCaller) OperatorDetails(opts *bind.CallOpts, operator common.Address) (IDelegationControllerOperatorDetails, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "operatorDetails", operator)

	if err != nil {
		return *new(IDelegationControllerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationControllerOperatorDetails)).(*IDelegationControllerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationControllerStorage *DelegationControllerStorageSession) OperatorDetails(operator common.Address) (IDelegationControllerOperatorDetails, error) {
	return _DelegationControllerStorage.Contract.OperatorDetails(&_DelegationControllerStorage.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) OperatorDetails(operator common.Address) (IDelegationControllerOperatorDetails, error) {
	return _DelegationControllerStorage.Contract.OperatorDetails(&_DelegationControllerStorage.CallOpts, operator)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) OperatorShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "operatorShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.OperatorShares(&_DelegationControllerStorage.CallOpts, arg0, arg1)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.OperatorShares(&_DelegationControllerStorage.CallOpts, arg0, arg1)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _DelegationControllerStorage.Contract.PendingWithdrawals(&_DelegationControllerStorage.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _DelegationControllerStorage.Contract.PendingWithdrawals(&_DelegationControllerStorage.CallOpts, arg0)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) PoolController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "poolController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageSession) PoolController() (common.Address, error) {
	return _DelegationControllerStorage.Contract.PoolController(&_DelegationControllerStorage.CallOpts)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) PoolController() (common.Address, error) {
	return _DelegationControllerStorage.Contract.PoolController(&_DelegationControllerStorage.CallOpts)
}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) PoolWithdrawalDelay(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "poolWithdrawalDelay", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) PoolWithdrawalDelay(arg0 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.PoolWithdrawalDelay(&_DelegationControllerStorage.CallOpts, arg0)
}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) PoolWithdrawalDelay(arg0 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.PoolWithdrawalDelay(&_DelegationControllerStorage.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageSession) Slasher() (common.Address, error) {
	return _DelegationControllerStorage.Contract.Slasher(&_DelegationControllerStorage.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) Slasher() (common.Address, error) {
	return _DelegationControllerStorage.Contract.Slasher(&_DelegationControllerStorage.CallOpts)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) StakerNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "stakerNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.StakerNonce(&_DelegationControllerStorage.CallOpts, arg0)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.StakerNonce(&_DelegationControllerStorage.CallOpts, arg0)
}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) StakerOptOutWindow(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "stakerOptOutWindow", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageSession) StakerOptOutWindow(operator common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.StakerOptOutWindow(&_DelegationControllerStorage.CallOpts, operator)
}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) StakerOptOutWindow(operator common.Address) (*big.Int, error) {
	return _DelegationControllerStorage.Contract.StakerOptOutWindow(&_DelegationControllerStorage.CallOpts, operator)
}

// WrappedTokenGateway is a free data retrieval call binding the contract method 0x1e119838.
//
// Solidity: function wrappedTokenGateway() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCaller) WrappedTokenGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationControllerStorage.contract.Call(opts, &out, "wrappedTokenGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenGateway is a free data retrieval call binding the contract method 0x1e119838.
//
// Solidity: function wrappedTokenGateway() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageSession) WrappedTokenGateway() (common.Address, error) {
	return _DelegationControllerStorage.Contract.WrappedTokenGateway(&_DelegationControllerStorage.CallOpts)
}

// WrappedTokenGateway is a free data retrieval call binding the contract method 0x1e119838.
//
// Solidity: function wrappedTokenGateway() view returns(address)
func (_DelegationControllerStorage *DelegationControllerStorageCallerSession) WrappedTokenGateway() (common.Address, error) {
	return _DelegationControllerStorage.Contract.WrappedTokenGateway(&_DelegationControllerStorage.CallOpts)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "decreaseDelegatedShares", staker, pool, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) DecreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DecreaseDelegatedShares(&_DelegationControllerStorage.TransactOpts, staker, pool, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) DecreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DecreaseDelegatedShares(&_DelegationControllerStorage.TransactOpts, staker, pool, shares)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DelegateTo(&_DelegationControllerStorage.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DelegateTo(&_DelegationControllerStorage.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) DelegateToBySignature(opts *bind.TransactOpts, staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "delegateToBySignature", staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DelegateToBySignature(&_DelegationControllerStorage.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.DelegateToBySignature(&_DelegationControllerStorage.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "increaseDelegatedShares", staker, pool, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) IncreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.IncreaseDelegatedShares(&_DelegationControllerStorage.TransactOpts, staker, pool, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) IncreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.IncreaseDelegatedShares(&_DelegationControllerStorage.TransactOpts, staker, pool, shares)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "modifyOperatorDetails", newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) ModifyOperatorDetails(newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.ModifyOperatorDetails(&_DelegationControllerStorage.TransactOpts, newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) ModifyOperatorDetails(newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.ModifyOperatorDetails(&_DelegationControllerStorage.TransactOpts, newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) RegisterAsOperator(opts *bind.TransactOpts, registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "registerAsOperator", registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) RegisterAsOperator(registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.RegisterAsOperator(&_DelegationControllerStorage.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) RegisterAsOperator(registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.RegisterAsOperator(&_DelegationControllerStorage.TransactOpts, registeringOperatorDetails, metadataURI)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoot)
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoot)
func (_DelegationControllerStorage *DelegationControllerStorageSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Undelegate(&_DelegationControllerStorage.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoot)
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Undelegate(&_DelegationControllerStorage.TransactOpts, staker)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) Unstakes(opts *bind.TransactOpts, unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "unstakes", unstakeParams)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_DelegationControllerStorage *DelegationControllerStorageSession) Unstakes(unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Unstakes(&_DelegationControllerStorage.TransactOpts, unstakeParams)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) Unstakes(unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Unstakes(&_DelegationControllerStorage.TransactOpts, unstakeParams)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "updateOperatorMetadataURI", metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.UpdateOperatorMetadataURI(&_DelegationControllerStorage.TransactOpts, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.UpdateOperatorMetadataURI(&_DelegationControllerStorage.TransactOpts, metadataURI)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) Withdraw(opts *bind.TransactOpts, withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "withdraw", withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) Withdraw(withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Withdraw(&_DelegationControllerStorage.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) Withdraw(withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Withdraw(&_DelegationControllerStorage.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactor) Withdraws(opts *bind.TransactOpts, withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationControllerStorage.contract.Transact(opts, "withdraws", withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationControllerStorage *DelegationControllerStorageSession) Withdraws(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Withdraws(&_DelegationControllerStorage.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationControllerStorage *DelegationControllerStorageTransactorSession) Withdraws(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationControllerStorage.Contract.Withdraws(&_DelegationControllerStorage.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// DelegationControllerStorageMinWithdrawalDelaySetIterator is returned from FilterMinWithdrawalDelaySet and is used to iterate over the raw logs and unpacked data for MinWithdrawalDelaySet events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageMinWithdrawalDelaySetIterator struct {
	Event *DelegationControllerStorageMinWithdrawalDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageMinWithdrawalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageMinWithdrawalDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageMinWithdrawalDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageMinWithdrawalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageMinWithdrawalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageMinWithdrawalDelaySet represents a MinWithdrawalDelaySet event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageMinWithdrawalDelaySet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinWithdrawalDelaySet is a free log retrieval operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterMinWithdrawalDelaySet(opts *bind.FilterOpts) (*DelegationControllerStorageMinWithdrawalDelaySetIterator, error) {

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "MinWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageMinWithdrawalDelaySetIterator{contract: _DelegationControllerStorage.contract, event: "MinWithdrawalDelaySet", logs: logs, sub: sub}, nil
}

// WatchMinWithdrawalDelaySet is a free log subscription operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchMinWithdrawalDelaySet(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageMinWithdrawalDelaySet) (event.Subscription, error) {

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "MinWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageMinWithdrawalDelaySet)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "MinWithdrawalDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinWithdrawalDelaySet is a log parse operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseMinWithdrawalDelaySet(log types.Log) (*DelegationControllerStorageMinWithdrawalDelaySet, error) {
	event := new(DelegationControllerStorageMinWithdrawalDelaySet)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "MinWithdrawalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageOperatorDetailsModifiedIterator is returned from FilterOperatorDetailsModified and is used to iterate over the raw logs and unpacked data for OperatorDetailsModified events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorDetailsModifiedIterator struct {
	Event *DelegationControllerStorageOperatorDetailsModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageOperatorDetailsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageOperatorDetailsModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageOperatorDetailsModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageOperatorDetailsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageOperatorDetailsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageOperatorDetailsModified represents a OperatorDetailsModified event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorDetailsModified struct {
	Operator           common.Address
	NewOperatorDetails IDelegationControllerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorDetailsModified is a free log retrieval operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterOperatorDetailsModified(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerStorageOperatorDetailsModifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageOperatorDetailsModifiedIterator{contract: _DelegationControllerStorage.contract, event: "OperatorDetailsModified", logs: logs, sub: sub}, nil
}

// WatchOperatorDetailsModified is a free log subscription operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchOperatorDetailsModified(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageOperatorDetailsModified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageOperatorDetailsModified)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDetailsModified is a log parse operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseOperatorDetailsModified(log types.Log) (*DelegationControllerStorageOperatorDetailsModified, error) {
	event := new(DelegationControllerStorageOperatorDetailsModified)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorMetadataURIUpdatedIterator struct {
	Event *DelegationControllerStorageOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageOperatorMetadataURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageOperatorMetadataURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerStorageOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageOperatorMetadataURIUpdatedIterator{contract: _DelegationControllerStorage.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageOperatorMetadataURIUpdated)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorMetadataURIUpdated is a log parse operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*DelegationControllerStorageOperatorMetadataURIUpdated, error) {
	event := new(DelegationControllerStorageOperatorMetadataURIUpdated)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorRegisteredIterator struct {
	Event *DelegationControllerStorageOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageOperatorRegistered represents a OperatorRegistered event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorRegistered struct {
	Operator        common.Address
	OperatorDetails IDelegationControllerOperatorDetails
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerStorageOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageOperatorRegisteredIterator{contract: _DelegationControllerStorage.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageOperatorRegistered)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseOperatorRegistered(log types.Log) (*DelegationControllerStorageOperatorRegistered, error) {
	event := new(DelegationControllerStorageOperatorRegistered)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorSharesDecreasedIterator struct {
	Event *DelegationControllerStorageOperatorSharesDecreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageOperatorSharesDecreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageOperatorSharesDecreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Pool     common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerStorageOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageOperatorSharesDecreasedIterator{contract: _DelegationControllerStorage.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageOperatorSharesDecreased)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSharesDecreased is a log parse operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseOperatorSharesDecreased(log types.Log) (*DelegationControllerStorageOperatorSharesDecreased, error) {
	event := new(DelegationControllerStorageOperatorSharesDecreased)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorSharesIncreasedIterator struct {
	Event *DelegationControllerStorageOperatorSharesIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageOperatorSharesIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageOperatorSharesIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Pool     common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerStorageOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageOperatorSharesIncreasedIterator{contract: _DelegationControllerStorage.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageOperatorSharesIncreased)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSharesIncreased is a log parse operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseOperatorSharesIncreased(log types.Log) (*DelegationControllerStorageOperatorSharesIncreased, error) {
	event := new(DelegationControllerStorageOperatorSharesIncreased)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStoragePoolWithdrawalDelaySetIterator is returned from FilterPoolWithdrawalDelaySet and is used to iterate over the raw logs and unpacked data for PoolWithdrawalDelaySet events raised by the DelegationControllerStorage contract.
type DelegationControllerStoragePoolWithdrawalDelaySetIterator struct {
	Event *DelegationControllerStoragePoolWithdrawalDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStoragePoolWithdrawalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStoragePoolWithdrawalDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStoragePoolWithdrawalDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStoragePoolWithdrawalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStoragePoolWithdrawalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStoragePoolWithdrawalDelaySet represents a PoolWithdrawalDelaySet event raised by the DelegationControllerStorage contract.
type DelegationControllerStoragePoolWithdrawalDelaySet struct {
	Pool          common.Address
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPoolWithdrawalDelaySet is a free log retrieval operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterPoolWithdrawalDelaySet(opts *bind.FilterOpts) (*DelegationControllerStoragePoolWithdrawalDelaySetIterator, error) {

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "PoolWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStoragePoolWithdrawalDelaySetIterator{contract: _DelegationControllerStorage.contract, event: "PoolWithdrawalDelaySet", logs: logs, sub: sub}, nil
}

// WatchPoolWithdrawalDelaySet is a free log subscription operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchPoolWithdrawalDelaySet(opts *bind.WatchOpts, sink chan<- *DelegationControllerStoragePoolWithdrawalDelaySet) (event.Subscription, error) {

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "PoolWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStoragePoolWithdrawalDelaySet)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "PoolWithdrawalDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolWithdrawalDelaySet is a log parse operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParsePoolWithdrawalDelaySet(log types.Log) (*DelegationControllerStoragePoolWithdrawalDelaySet, error) {
	event := new(DelegationControllerStoragePoolWithdrawalDelaySet)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "PoolWithdrawalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageStakerDelegatedIterator struct {
	Event *DelegationControllerStorageStakerDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageStakerDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageStakerDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageStakerDelegated represents a StakerDelegated event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationControllerStorageStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageStakerDelegatedIterator{contract: _DelegationControllerStorage.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageStakerDelegated)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakerDelegated is a log parse operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseStakerDelegated(log types.Log) (*DelegationControllerStorageStakerDelegated, error) {
	event := new(DelegationControllerStorageStakerDelegated)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageStakerForceUndelegatedIterator struct {
	Event *DelegationControllerStorageStakerForceUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageStakerForceUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageStakerForceUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageStakerForceUndelegated represents a StakerForceUndelegated event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationControllerStorageStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageStakerForceUndelegatedIterator{contract: _DelegationControllerStorage.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageStakerForceUndelegated)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakerForceUndelegated is a log parse operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseStakerForceUndelegated(log types.Log) (*DelegationControllerStorageStakerForceUndelegated, error) {
	event := new(DelegationControllerStorageStakerForceUndelegated)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageStakerUndelegatedIterator struct {
	Event *DelegationControllerStorageStakerUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageStakerUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageStakerUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageStakerUndelegated represents a StakerUndelegated event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationControllerStorageStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageStakerUndelegatedIterator{contract: _DelegationControllerStorage.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageStakerUndelegated)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakerUndelegated is a log parse operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseStakerUndelegated(log types.Log) (*DelegationControllerStorageStakerUndelegated, error) {
	event := new(DelegationControllerStorageStakerUndelegated)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageUpdateWrappedTokenGatewayIterator is returned from FilterUpdateWrappedTokenGateway and is used to iterate over the raw logs and unpacked data for UpdateWrappedTokenGateway events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageUpdateWrappedTokenGatewayIterator struct {
	Event *DelegationControllerStorageUpdateWrappedTokenGateway // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageUpdateWrappedTokenGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageUpdateWrappedTokenGateway)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageUpdateWrappedTokenGateway)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageUpdateWrappedTokenGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageUpdateWrappedTokenGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageUpdateWrappedTokenGateway represents a UpdateWrappedTokenGateway event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageUpdateWrappedTokenGateway struct {
	PreviousGateway common.Address
	CurrentGateway  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateWrappedTokenGateway is a free log retrieval operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterUpdateWrappedTokenGateway(opts *bind.FilterOpts) (*DelegationControllerStorageUpdateWrappedTokenGatewayIterator, error) {

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "UpdateWrappedTokenGateway")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageUpdateWrappedTokenGatewayIterator{contract: _DelegationControllerStorage.contract, event: "UpdateWrappedTokenGateway", logs: logs, sub: sub}, nil
}

// WatchUpdateWrappedTokenGateway is a free log subscription operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchUpdateWrappedTokenGateway(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageUpdateWrappedTokenGateway) (event.Subscription, error) {

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "UpdateWrappedTokenGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageUpdateWrappedTokenGateway)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "UpdateWrappedTokenGateway", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateWrappedTokenGateway is a log parse operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseUpdateWrappedTokenGateway(log types.Log) (*DelegationControllerStorageUpdateWrappedTokenGateway, error) {
	event := new(DelegationControllerStorageUpdateWrappedTokenGateway)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "UpdateWrappedTokenGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageWithdrawalCompletedIterator struct {
	Event *DelegationControllerStorageWithdrawalCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageWithdrawalCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageWithdrawalCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageWithdrawalCompleted represents a WithdrawalCompleted event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*DelegationControllerStorageWithdrawalCompletedIterator, error) {

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageWithdrawalCompletedIterator{contract: _DelegationControllerStorage.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageWithdrawalCompleted)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseWithdrawalCompleted(log types.Log) (*DelegationControllerStorageWithdrawalCompleted, error) {
	event := new(DelegationControllerStorageWithdrawalCompleted)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStorageWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the DelegationControllerStorage contract.
type DelegationControllerStorageWithdrawalQueuedIterator struct {
	Event *DelegationControllerStorageWithdrawalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegationControllerStorageWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStorageWithdrawalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegationControllerStorageWithdrawalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegationControllerStorageWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStorageWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStorageWithdrawalQueued represents a WithdrawalQueued event raised by the DelegationControllerStorage contract.
type DelegationControllerStorageWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     IDelegationControllerWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*DelegationControllerStorageWithdrawalQueuedIterator, error) {

	logs, sub, err := _DelegationControllerStorage.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStorageWithdrawalQueuedIterator{contract: _DelegationControllerStorage.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *DelegationControllerStorageWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _DelegationControllerStorage.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStorageWithdrawalQueued)
				if err := _DelegationControllerStorage.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationControllerStorage *DelegationControllerStorageFilterer) ParseWithdrawalQueued(log types.Log) (*DelegationControllerStorageWithdrawalQueued, error) {
	event := new(DelegationControllerStorageWithdrawalQueued)
	if err := _DelegationControllerStorage.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
