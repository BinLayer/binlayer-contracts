// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDelegationController

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

// IDelegationControllerMetaData contains all meta data concerning the IDelegationController contract.
var IDelegationControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakerNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateToBySignature\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"earningsReceiver\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalDelay\",\"inputs\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolWithdrawalDelay\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerNonce\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerOptOutWindow\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakes\",\"inputs\":[{\"name\":\"unstakeParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.UnstakeParams[]\",\"components\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraws\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MinWithdrawalDelaySet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDetailsModified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolWithdrawalDelaySet\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateWrappedTokenGateway\",\"inputs\":[{\"name\":\"previousGateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currentGateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false}]",
}

// IDelegationControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use IDelegationControllerMetaData.ABI instead.
var IDelegationControllerABI = IDelegationControllerMetaData.ABI

// IDelegationController is an auto generated Go binding around an Ethereum contract.
type IDelegationController struct {
	IDelegationControllerCaller     // Read-only binding to the contract
	IDelegationControllerTransactor // Write-only binding to the contract
	IDelegationControllerFilterer   // Log filterer for contract events
}

// IDelegationControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelegationControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegationControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelegationControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegationControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelegationControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegationControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelegationControllerSession struct {
	Contract     *IDelegationController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDelegationControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelegationControllerCallerSession struct {
	Contract *IDelegationControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IDelegationControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelegationControllerTransactorSession struct {
	Contract     *IDelegationControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IDelegationControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelegationControllerRaw struct {
	Contract *IDelegationController // Generic contract binding to access the raw methods on
}

// IDelegationControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelegationControllerCallerRaw struct {
	Contract *IDelegationControllerCaller // Generic read-only contract binding to access the raw methods on
}

// IDelegationControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelegationControllerTransactorRaw struct {
	Contract *IDelegationControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelegationController creates a new instance of IDelegationController, bound to a specific deployed contract.
func NewIDelegationController(address common.Address, backend bind.ContractBackend) (*IDelegationController, error) {
	contract, err := bindIDelegationController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelegationController{IDelegationControllerCaller: IDelegationControllerCaller{contract: contract}, IDelegationControllerTransactor: IDelegationControllerTransactor{contract: contract}, IDelegationControllerFilterer: IDelegationControllerFilterer{contract: contract}}, nil
}

// NewIDelegationControllerCaller creates a new read-only instance of IDelegationController, bound to a specific deployed contract.
func NewIDelegationControllerCaller(address common.Address, caller bind.ContractCaller) (*IDelegationControllerCaller, error) {
	contract, err := bindIDelegationController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerCaller{contract: contract}, nil
}

// NewIDelegationControllerTransactor creates a new write-only instance of IDelegationController, bound to a specific deployed contract.
func NewIDelegationControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelegationControllerTransactor, error) {
	contract, err := bindIDelegationController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerTransactor{contract: contract}, nil
}

// NewIDelegationControllerFilterer creates a new log filterer instance of IDelegationController, bound to a specific deployed contract.
func NewIDelegationControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelegationControllerFilterer, error) {
	contract, err := bindIDelegationController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerFilterer{contract: contract}, nil
}

// bindIDelegationController binds a generic wrapper to an already deployed contract.
func bindIDelegationController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDelegationControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegationController *IDelegationControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelegationController.Contract.IDelegationControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegationController *IDelegationControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegationController.Contract.IDelegationControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegationController *IDelegationControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegationController.Contract.IDelegationControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegationController *IDelegationControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelegationController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegationController *IDelegationControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegationController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegationController *IDelegationControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegationController.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _IDelegationController.Contract.DELEGATIONAPPROVALTYPEHASH(&_IDelegationController.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _IDelegationController.Contract.DELEGATIONAPPROVALTYPEHASH(&_IDelegationController.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _IDelegationController.Contract.DOMAINTYPEHASH(&_IDelegationController.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _IDelegationController.Contract.DOMAINTYPEHASH(&_IDelegationController.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) STAKERDELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "STAKER_DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _IDelegationController.Contract.STAKERDELEGATIONTYPEHASH(&_IDelegationController.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _IDelegationController.Contract.STAKERDELEGATIONTYPEHASH(&_IDelegationController.CallOpts)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) CalculateCurrentStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "calculateCurrentStakerDelegationDigestHash", staker, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateCurrentStakerDelegationDigestHash(&_IDelegationController.CallOpts, staker, operator, expiry)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateCurrentStakerDelegationDigestHash(&_IDelegationController.CallOpts, staker, operator, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, _delegationApprover, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateDelegationApprovalDigestHash(&_IDelegationController.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateDelegationApprovalDigestHash(&_IDelegationController.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) CalculateStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "calculateStakerDelegationDigestHash", staker, _stakerNonce, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateStakerDelegationDigestHash(&_IDelegationController.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateStakerDelegationDigestHash(&_IDelegationController.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) CalculateWithdrawalRoot(withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateWithdrawalRoot(&_IDelegationController.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	return _IDelegationController.Contract.CalculateWithdrawalRoot(&_IDelegationController.CallOpts, withdrawal)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.CumulativeWithdrawalsQueued(&_IDelegationController.CallOpts, staker)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.CumulativeWithdrawalsQueued(&_IDelegationController.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address)
func (_IDelegationController *IDelegationControllerCaller) DelegatedTo(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "delegatedTo", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address)
func (_IDelegationController *IDelegationControllerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _IDelegationController.Contract.DelegatedTo(&_IDelegationController.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address)
func (_IDelegationController *IDelegationControllerCallerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _IDelegationController.Contract.DelegatedTo(&_IDelegationController.CallOpts, staker)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_IDelegationController *IDelegationControllerCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_IDelegationController *IDelegationControllerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _IDelegationController.Contract.DelegationApprover(&_IDelegationController.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_IDelegationController *IDelegationControllerCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _IDelegationController.Contract.DelegationApprover(&_IDelegationController.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address _delegationApprover, bytes32 salt) view returns(bool)
func (_IDelegationController *IDelegationControllerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, _delegationApprover common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "delegationApproverSaltIsSpent", _delegationApprover, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address _delegationApprover, bytes32 salt) view returns(bool)
func (_IDelegationController *IDelegationControllerSession) DelegationApproverSaltIsSpent(_delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _IDelegationController.Contract.DelegationApproverSaltIsSpent(&_IDelegationController.CallOpts, _delegationApprover, salt)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address _delegationApprover, bytes32 salt) view returns(bool)
func (_IDelegationController *IDelegationControllerCallerSession) DelegationApproverSaltIsSpent(_delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _IDelegationController.Contract.DelegationApproverSaltIsSpent(&_IDelegationController.CallOpts, _delegationApprover, salt)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IDelegationController *IDelegationControllerSession) DomainSeparator() ([32]byte, error) {
	return _IDelegationController.Contract.DomainSeparator(&_IDelegationController.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IDelegationController *IDelegationControllerCallerSession) DomainSeparator() ([32]byte, error) {
	return _IDelegationController.Contract.DomainSeparator(&_IDelegationController.CallOpts)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_IDelegationController *IDelegationControllerCaller) EarningsReceiver(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "earningsReceiver", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_IDelegationController *IDelegationControllerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _IDelegationController.Contract.EarningsReceiver(&_IDelegationController.CallOpts, operator)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_IDelegationController *IDelegationControllerCallerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _IDelegationController.Contract.EarningsReceiver(&_IDelegationController.CallOpts, operator)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_IDelegationController *IDelegationControllerCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address, pools []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "getOperatorShares", operator, pools)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_IDelegationController *IDelegationControllerSession) GetOperatorShares(operator common.Address, pools []common.Address) ([]*big.Int, error) {
	return _IDelegationController.Contract.GetOperatorShares(&_IDelegationController.CallOpts, operator, pools)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_IDelegationController *IDelegationControllerCallerSession) GetOperatorShares(operator common.Address, pools []common.Address) ([]*big.Int, error) {
	return _IDelegationController.Contract.GetOperatorShares(&_IDelegationController.CallOpts, operator, pools)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) GetWithdrawalDelay(opts *bind.CallOpts, pools []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "getWithdrawalDelay", pools)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) GetWithdrawalDelay(pools []common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.GetWithdrawalDelay(&_IDelegationController.CallOpts, pools)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) GetWithdrawalDelay(pools []common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.GetWithdrawalDelay(&_IDelegationController.CallOpts, pools)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_IDelegationController *IDelegationControllerCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_IDelegationController *IDelegationControllerSession) IsDelegated(staker common.Address) (bool, error) {
	return _IDelegationController.Contract.IsDelegated(&_IDelegationController.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_IDelegationController *IDelegationControllerCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _IDelegationController.Contract.IsDelegated(&_IDelegationController.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_IDelegationController *IDelegationControllerCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_IDelegationController *IDelegationControllerSession) IsOperator(operator common.Address) (bool, error) {
	return _IDelegationController.Contract.IsOperator(&_IDelegationController.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_IDelegationController *IDelegationControllerCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _IDelegationController.Contract.IsOperator(&_IDelegationController.CallOpts, operator)
}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) MinWithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "minWithdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) MinWithdrawalDelay() (*big.Int, error) {
	return _IDelegationController.Contract.MinWithdrawalDelay(&_IDelegationController.CallOpts)
}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) MinWithdrawalDelay() (*big.Int, error) {
	return _IDelegationController.Contract.MinWithdrawalDelay(&_IDelegationController.CallOpts)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_IDelegationController *IDelegationControllerCaller) OperatorDetails(opts *bind.CallOpts, operator common.Address) (IDelegationControllerOperatorDetails, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "operatorDetails", operator)

	if err != nil {
		return *new(IDelegationControllerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationControllerOperatorDetails)).(*IDelegationControllerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_IDelegationController *IDelegationControllerSession) OperatorDetails(operator common.Address) (IDelegationControllerOperatorDetails, error) {
	return _IDelegationController.Contract.OperatorDetails(&_IDelegationController.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_IDelegationController *IDelegationControllerCallerSession) OperatorDetails(operator common.Address) (IDelegationControllerOperatorDetails, error) {
	return _IDelegationController.Contract.OperatorDetails(&_IDelegationController.CallOpts, operator)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address pool) view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) OperatorShares(opts *bind.CallOpts, operator common.Address, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "operatorShares", operator, pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address pool) view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) OperatorShares(operator common.Address, pool common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.OperatorShares(&_IDelegationController.CallOpts, operator, pool)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address pool) view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) OperatorShares(operator common.Address, pool common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.OperatorShares(&_IDelegationController.CallOpts, operator, pool)
}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address pool) view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) PoolWithdrawalDelay(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "poolWithdrawalDelay", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address pool) view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) PoolWithdrawalDelay(pool common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.PoolWithdrawalDelay(&_IDelegationController.CallOpts, pool)
}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address pool) view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) PoolWithdrawalDelay(pool common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.PoolWithdrawalDelay(&_IDelegationController.CallOpts, pool)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address staker) view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) StakerNonce(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "stakerNonce", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address staker) view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) StakerNonce(staker common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.StakerNonce(&_IDelegationController.CallOpts, staker)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address staker) view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) StakerNonce(staker common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.StakerNonce(&_IDelegationController.CallOpts, staker)
}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_IDelegationController *IDelegationControllerCaller) StakerOptOutWindow(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDelegationController.contract.Call(opts, &out, "stakerOptOutWindow", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_IDelegationController *IDelegationControllerSession) StakerOptOutWindow(operator common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.StakerOptOutWindow(&_IDelegationController.CallOpts, operator)
}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_IDelegationController *IDelegationControllerCallerSession) StakerOptOutWindow(operator common.Address) (*big.Int, error) {
	return _IDelegationController.Contract.StakerOptOutWindow(&_IDelegationController.CallOpts, operator)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_IDelegationController *IDelegationControllerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "decreaseDelegatedShares", staker, pool, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_IDelegationController *IDelegationControllerSession) DecreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDelegationController.Contract.DecreaseDelegatedShares(&_IDelegationController.TransactOpts, staker, pool, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) DecreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDelegationController.Contract.DecreaseDelegatedShares(&_IDelegationController.TransactOpts, staker, pool, shares)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_IDelegationController *IDelegationControllerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_IDelegationController *IDelegationControllerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _IDelegationController.Contract.DelegateTo(&_IDelegationController.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _IDelegationController.Contract.DelegateTo(&_IDelegationController.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_IDelegationController *IDelegationControllerTransactor) DelegateToBySignature(opts *bind.TransactOpts, staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "delegateToBySignature", staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_IDelegationController *IDelegationControllerSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _IDelegationController.Contract.DelegateToBySignature(&_IDelegationController.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _IDelegationController.Contract.DelegateToBySignature(&_IDelegationController.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_IDelegationController *IDelegationControllerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "increaseDelegatedShares", staker, pool, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_IDelegationController *IDelegationControllerSession) IncreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDelegationController.Contract.IncreaseDelegatedShares(&_IDelegationController.TransactOpts, staker, pool, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) IncreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IDelegationController.Contract.IncreaseDelegatedShares(&_IDelegationController.TransactOpts, staker, pool, shares)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_IDelegationController *IDelegationControllerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "modifyOperatorDetails", newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_IDelegationController *IDelegationControllerSession) ModifyOperatorDetails(newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _IDelegationController.Contract.ModifyOperatorDetails(&_IDelegationController.TransactOpts, newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) ModifyOperatorDetails(newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _IDelegationController.Contract.ModifyOperatorDetails(&_IDelegationController.TransactOpts, newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_IDelegationController *IDelegationControllerTransactor) RegisterAsOperator(opts *bind.TransactOpts, registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "registerAsOperator", registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_IDelegationController *IDelegationControllerSession) RegisterAsOperator(registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _IDelegationController.Contract.RegisterAsOperator(&_IDelegationController.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) RegisterAsOperator(registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _IDelegationController.Contract.RegisterAsOperator(&_IDelegationController.TransactOpts, registeringOperatorDetails, metadataURI)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoot)
func (_IDelegationController *IDelegationControllerTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoot)
func (_IDelegationController *IDelegationControllerSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _IDelegationController.Contract.Undelegate(&_IDelegationController.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoot)
func (_IDelegationController *IDelegationControllerTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _IDelegationController.Contract.Undelegate(&_IDelegationController.TransactOpts, staker)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_IDelegationController *IDelegationControllerTransactor) Unstakes(opts *bind.TransactOpts, unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "unstakes", unstakeParams)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_IDelegationController *IDelegationControllerSession) Unstakes(unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _IDelegationController.Contract.Unstakes(&_IDelegationController.TransactOpts, unstakeParams)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_IDelegationController *IDelegationControllerTransactorSession) Unstakes(unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _IDelegationController.Contract.Unstakes(&_IDelegationController.TransactOpts, unstakeParams)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_IDelegationController *IDelegationControllerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "updateOperatorMetadataURI", metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_IDelegationController *IDelegationControllerSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _IDelegationController.Contract.UpdateOperatorMetadataURI(&_IDelegationController.TransactOpts, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _IDelegationController.Contract.UpdateOperatorMetadataURI(&_IDelegationController.TransactOpts, metadataURI)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_IDelegationController *IDelegationControllerTransactor) Withdraw(opts *bind.TransactOpts, withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "withdraw", withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_IDelegationController *IDelegationControllerSession) Withdraw(withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IDelegationController.Contract.Withdraw(&_IDelegationController.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) Withdraw(withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _IDelegationController.Contract.Withdraw(&_IDelegationController.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_IDelegationController *IDelegationControllerTransactor) Withdraws(opts *bind.TransactOpts, withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _IDelegationController.contract.Transact(opts, "withdraws", withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_IDelegationController *IDelegationControllerSession) Withdraws(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _IDelegationController.Contract.Withdraws(&_IDelegationController.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_IDelegationController *IDelegationControllerTransactorSession) Withdraws(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _IDelegationController.Contract.Withdraws(&_IDelegationController.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// IDelegationControllerMinWithdrawalDelaySetIterator is returned from FilterMinWithdrawalDelaySet and is used to iterate over the raw logs and unpacked data for MinWithdrawalDelaySet events raised by the IDelegationController contract.
type IDelegationControllerMinWithdrawalDelaySetIterator struct {
	Event *IDelegationControllerMinWithdrawalDelaySet // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerMinWithdrawalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerMinWithdrawalDelaySet)
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
		it.Event = new(IDelegationControllerMinWithdrawalDelaySet)
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
func (it *IDelegationControllerMinWithdrawalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerMinWithdrawalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerMinWithdrawalDelaySet represents a MinWithdrawalDelaySet event raised by the IDelegationController contract.
type IDelegationControllerMinWithdrawalDelaySet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinWithdrawalDelaySet is a free log retrieval operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_IDelegationController *IDelegationControllerFilterer) FilterMinWithdrawalDelaySet(opts *bind.FilterOpts) (*IDelegationControllerMinWithdrawalDelaySetIterator, error) {

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "MinWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerMinWithdrawalDelaySetIterator{contract: _IDelegationController.contract, event: "MinWithdrawalDelaySet", logs: logs, sub: sub}, nil
}

// WatchMinWithdrawalDelaySet is a free log subscription operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_IDelegationController *IDelegationControllerFilterer) WatchMinWithdrawalDelaySet(opts *bind.WatchOpts, sink chan<- *IDelegationControllerMinWithdrawalDelaySet) (event.Subscription, error) {

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "MinWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerMinWithdrawalDelaySet)
				if err := _IDelegationController.contract.UnpackLog(event, "MinWithdrawalDelaySet", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseMinWithdrawalDelaySet(log types.Log) (*IDelegationControllerMinWithdrawalDelaySet, error) {
	event := new(IDelegationControllerMinWithdrawalDelaySet)
	if err := _IDelegationController.contract.UnpackLog(event, "MinWithdrawalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerOperatorDetailsModifiedIterator is returned from FilterOperatorDetailsModified and is used to iterate over the raw logs and unpacked data for OperatorDetailsModified events raised by the IDelegationController contract.
type IDelegationControllerOperatorDetailsModifiedIterator struct {
	Event *IDelegationControllerOperatorDetailsModified // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerOperatorDetailsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerOperatorDetailsModified)
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
		it.Event = new(IDelegationControllerOperatorDetailsModified)
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
func (it *IDelegationControllerOperatorDetailsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerOperatorDetailsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerOperatorDetailsModified represents a OperatorDetailsModified event raised by the IDelegationController contract.
type IDelegationControllerOperatorDetailsModified struct {
	Operator           common.Address
	NewOperatorDetails IDelegationControllerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorDetailsModified is a free log retrieval operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_IDelegationController *IDelegationControllerFilterer) FilterOperatorDetailsModified(opts *bind.FilterOpts, operator []common.Address) (*IDelegationControllerOperatorDetailsModifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerOperatorDetailsModifiedIterator{contract: _IDelegationController.contract, event: "OperatorDetailsModified", logs: logs, sub: sub}, nil
}

// WatchOperatorDetailsModified is a free log subscription operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_IDelegationController *IDelegationControllerFilterer) WatchOperatorDetailsModified(opts *bind.WatchOpts, sink chan<- *IDelegationControllerOperatorDetailsModified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerOperatorDetailsModified)
				if err := _IDelegationController.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseOperatorDetailsModified(log types.Log) (*IDelegationControllerOperatorDetailsModified, error) {
	event := new(IDelegationControllerOperatorDetailsModified)
	if err := _IDelegationController.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the IDelegationController contract.
type IDelegationControllerOperatorMetadataURIUpdatedIterator struct {
	Event *IDelegationControllerOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerOperatorMetadataURIUpdated)
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
		it.Event = new(IDelegationControllerOperatorMetadataURIUpdated)
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
func (it *IDelegationControllerOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the IDelegationController contract.
type IDelegationControllerOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_IDelegationController *IDelegationControllerFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*IDelegationControllerOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerOperatorMetadataURIUpdatedIterator{contract: _IDelegationController.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_IDelegationController *IDelegationControllerFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *IDelegationControllerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerOperatorMetadataURIUpdated)
				if err := _IDelegationController.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*IDelegationControllerOperatorMetadataURIUpdated, error) {
	event := new(IDelegationControllerOperatorMetadataURIUpdated)
	if err := _IDelegationController.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the IDelegationController contract.
type IDelegationControllerOperatorRegisteredIterator struct {
	Event *IDelegationControllerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerOperatorRegistered)
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
		it.Event = new(IDelegationControllerOperatorRegistered)
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
func (it *IDelegationControllerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerOperatorRegistered represents a OperatorRegistered event raised by the IDelegationController contract.
type IDelegationControllerOperatorRegistered struct {
	Operator        common.Address
	OperatorDetails IDelegationControllerOperatorDetails
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_IDelegationController *IDelegationControllerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*IDelegationControllerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerOperatorRegisteredIterator{contract: _IDelegationController.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_IDelegationController *IDelegationControllerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *IDelegationControllerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerOperatorRegistered)
				if err := _IDelegationController.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseOperatorRegistered(log types.Log) (*IDelegationControllerOperatorRegistered, error) {
	event := new(IDelegationControllerOperatorRegistered)
	if err := _IDelegationController.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the IDelegationController contract.
type IDelegationControllerOperatorSharesDecreasedIterator struct {
	Event *IDelegationControllerOperatorSharesDecreased // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerOperatorSharesDecreased)
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
		it.Event = new(IDelegationControllerOperatorSharesDecreased)
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
func (it *IDelegationControllerOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the IDelegationController contract.
type IDelegationControllerOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Pool     common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_IDelegationController *IDelegationControllerFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*IDelegationControllerOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerOperatorSharesDecreasedIterator{contract: _IDelegationController.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_IDelegationController *IDelegationControllerFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *IDelegationControllerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerOperatorSharesDecreased)
				if err := _IDelegationController.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseOperatorSharesDecreased(log types.Log) (*IDelegationControllerOperatorSharesDecreased, error) {
	event := new(IDelegationControllerOperatorSharesDecreased)
	if err := _IDelegationController.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the IDelegationController contract.
type IDelegationControllerOperatorSharesIncreasedIterator struct {
	Event *IDelegationControllerOperatorSharesIncreased // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerOperatorSharesIncreased)
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
		it.Event = new(IDelegationControllerOperatorSharesIncreased)
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
func (it *IDelegationControllerOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the IDelegationController contract.
type IDelegationControllerOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Pool     common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_IDelegationController *IDelegationControllerFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*IDelegationControllerOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerOperatorSharesIncreasedIterator{contract: _IDelegationController.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_IDelegationController *IDelegationControllerFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *IDelegationControllerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerOperatorSharesIncreased)
				if err := _IDelegationController.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseOperatorSharesIncreased(log types.Log) (*IDelegationControllerOperatorSharesIncreased, error) {
	event := new(IDelegationControllerOperatorSharesIncreased)
	if err := _IDelegationController.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerPoolWithdrawalDelaySetIterator is returned from FilterPoolWithdrawalDelaySet and is used to iterate over the raw logs and unpacked data for PoolWithdrawalDelaySet events raised by the IDelegationController contract.
type IDelegationControllerPoolWithdrawalDelaySetIterator struct {
	Event *IDelegationControllerPoolWithdrawalDelaySet // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerPoolWithdrawalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerPoolWithdrawalDelaySet)
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
		it.Event = new(IDelegationControllerPoolWithdrawalDelaySet)
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
func (it *IDelegationControllerPoolWithdrawalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerPoolWithdrawalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerPoolWithdrawalDelaySet represents a PoolWithdrawalDelaySet event raised by the IDelegationController contract.
type IDelegationControllerPoolWithdrawalDelaySet struct {
	Pool          common.Address
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPoolWithdrawalDelaySet is a free log retrieval operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_IDelegationController *IDelegationControllerFilterer) FilterPoolWithdrawalDelaySet(opts *bind.FilterOpts) (*IDelegationControllerPoolWithdrawalDelaySetIterator, error) {

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "PoolWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerPoolWithdrawalDelaySetIterator{contract: _IDelegationController.contract, event: "PoolWithdrawalDelaySet", logs: logs, sub: sub}, nil
}

// WatchPoolWithdrawalDelaySet is a free log subscription operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_IDelegationController *IDelegationControllerFilterer) WatchPoolWithdrawalDelaySet(opts *bind.WatchOpts, sink chan<- *IDelegationControllerPoolWithdrawalDelaySet) (event.Subscription, error) {

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "PoolWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerPoolWithdrawalDelaySet)
				if err := _IDelegationController.contract.UnpackLog(event, "PoolWithdrawalDelaySet", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParsePoolWithdrawalDelaySet(log types.Log) (*IDelegationControllerPoolWithdrawalDelaySet, error) {
	event := new(IDelegationControllerPoolWithdrawalDelaySet)
	if err := _IDelegationController.contract.UnpackLog(event, "PoolWithdrawalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the IDelegationController contract.
type IDelegationControllerStakerDelegatedIterator struct {
	Event *IDelegationControllerStakerDelegated // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerStakerDelegated)
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
		it.Event = new(IDelegationControllerStakerDelegated)
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
func (it *IDelegationControllerStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerStakerDelegated represents a StakerDelegated event raised by the IDelegationController contract.
type IDelegationControllerStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_IDelegationController *IDelegationControllerFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*IDelegationControllerStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerStakerDelegatedIterator{contract: _IDelegationController.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_IDelegationController *IDelegationControllerFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *IDelegationControllerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerStakerDelegated)
				if err := _IDelegationController.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseStakerDelegated(log types.Log) (*IDelegationControllerStakerDelegated, error) {
	event := new(IDelegationControllerStakerDelegated)
	if err := _IDelegationController.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the IDelegationController contract.
type IDelegationControllerStakerForceUndelegatedIterator struct {
	Event *IDelegationControllerStakerForceUndelegated // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerStakerForceUndelegated)
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
		it.Event = new(IDelegationControllerStakerForceUndelegated)
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
func (it *IDelegationControllerStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerStakerForceUndelegated represents a StakerForceUndelegated event raised by the IDelegationController contract.
type IDelegationControllerStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_IDelegationController *IDelegationControllerFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*IDelegationControllerStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerStakerForceUndelegatedIterator{contract: _IDelegationController.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_IDelegationController *IDelegationControllerFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *IDelegationControllerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerStakerForceUndelegated)
				if err := _IDelegationController.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseStakerForceUndelegated(log types.Log) (*IDelegationControllerStakerForceUndelegated, error) {
	event := new(IDelegationControllerStakerForceUndelegated)
	if err := _IDelegationController.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the IDelegationController contract.
type IDelegationControllerStakerUndelegatedIterator struct {
	Event *IDelegationControllerStakerUndelegated // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerStakerUndelegated)
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
		it.Event = new(IDelegationControllerStakerUndelegated)
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
func (it *IDelegationControllerStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerStakerUndelegated represents a StakerUndelegated event raised by the IDelegationController contract.
type IDelegationControllerStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_IDelegationController *IDelegationControllerFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*IDelegationControllerStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerStakerUndelegatedIterator{contract: _IDelegationController.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_IDelegationController *IDelegationControllerFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *IDelegationControllerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerStakerUndelegated)
				if err := _IDelegationController.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseStakerUndelegated(log types.Log) (*IDelegationControllerStakerUndelegated, error) {
	event := new(IDelegationControllerStakerUndelegated)
	if err := _IDelegationController.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerUpdateWrappedTokenGatewayIterator is returned from FilterUpdateWrappedTokenGateway and is used to iterate over the raw logs and unpacked data for UpdateWrappedTokenGateway events raised by the IDelegationController contract.
type IDelegationControllerUpdateWrappedTokenGatewayIterator struct {
	Event *IDelegationControllerUpdateWrappedTokenGateway // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerUpdateWrappedTokenGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerUpdateWrappedTokenGateway)
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
		it.Event = new(IDelegationControllerUpdateWrappedTokenGateway)
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
func (it *IDelegationControllerUpdateWrappedTokenGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerUpdateWrappedTokenGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerUpdateWrappedTokenGateway represents a UpdateWrappedTokenGateway event raised by the IDelegationController contract.
type IDelegationControllerUpdateWrappedTokenGateway struct {
	PreviousGateway common.Address
	CurrentGateway  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateWrappedTokenGateway is a free log retrieval operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_IDelegationController *IDelegationControllerFilterer) FilterUpdateWrappedTokenGateway(opts *bind.FilterOpts) (*IDelegationControllerUpdateWrappedTokenGatewayIterator, error) {

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "UpdateWrappedTokenGateway")
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerUpdateWrappedTokenGatewayIterator{contract: _IDelegationController.contract, event: "UpdateWrappedTokenGateway", logs: logs, sub: sub}, nil
}

// WatchUpdateWrappedTokenGateway is a free log subscription operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_IDelegationController *IDelegationControllerFilterer) WatchUpdateWrappedTokenGateway(opts *bind.WatchOpts, sink chan<- *IDelegationControllerUpdateWrappedTokenGateway) (event.Subscription, error) {

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "UpdateWrappedTokenGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerUpdateWrappedTokenGateway)
				if err := _IDelegationController.contract.UnpackLog(event, "UpdateWrappedTokenGateway", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseUpdateWrappedTokenGateway(log types.Log) (*IDelegationControllerUpdateWrappedTokenGateway, error) {
	event := new(IDelegationControllerUpdateWrappedTokenGateway)
	if err := _IDelegationController.contract.UnpackLog(event, "UpdateWrappedTokenGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the IDelegationController contract.
type IDelegationControllerWithdrawalCompletedIterator struct {
	Event *IDelegationControllerWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerWithdrawalCompleted)
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
		it.Event = new(IDelegationControllerWithdrawalCompleted)
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
func (it *IDelegationControllerWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerWithdrawalCompleted represents a WithdrawalCompleted event raised by the IDelegationController contract.
type IDelegationControllerWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_IDelegationController *IDelegationControllerFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*IDelegationControllerWithdrawalCompletedIterator, error) {

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerWithdrawalCompletedIterator{contract: _IDelegationController.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_IDelegationController *IDelegationControllerFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *IDelegationControllerWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerWithdrawalCompleted)
				if err := _IDelegationController.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseWithdrawalCompleted(log types.Log) (*IDelegationControllerWithdrawalCompleted, error) {
	event := new(IDelegationControllerWithdrawalCompleted)
	if err := _IDelegationController.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelegationControllerWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the IDelegationController contract.
type IDelegationControllerWithdrawalQueuedIterator struct {
	Event *IDelegationControllerWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *IDelegationControllerWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelegationControllerWithdrawalQueued)
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
		it.Event = new(IDelegationControllerWithdrawalQueued)
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
func (it *IDelegationControllerWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelegationControllerWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelegationControllerWithdrawalQueued represents a WithdrawalQueued event raised by the IDelegationController contract.
type IDelegationControllerWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     IDelegationControllerWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_IDelegationController *IDelegationControllerFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*IDelegationControllerWithdrawalQueuedIterator, error) {

	logs, sub, err := _IDelegationController.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &IDelegationControllerWithdrawalQueuedIterator{contract: _IDelegationController.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_IDelegationController *IDelegationControllerFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *IDelegationControllerWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _IDelegationController.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelegationControllerWithdrawalQueued)
				if err := _IDelegationController.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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
func (_IDelegationController *IDelegationControllerFilterer) ParseWithdrawalQueued(log types.Log) (*IDelegationControllerWithdrawalQueued, error) {
	event := new(IDelegationControllerWithdrawalQueued)
	if err := _IDelegationController.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
