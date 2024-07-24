// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IWrappedTokenGateway

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

// IWrappedTokenGatewayMetaData contains all meta data concerning the IWrappedTokenGateway contract.
var IWrappedTokenGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"depositNativeToken\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawNativeTokens\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IWrappedTokenGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IWrappedTokenGatewayMetaData.ABI instead.
var IWrappedTokenGatewayABI = IWrappedTokenGatewayMetaData.ABI

// IWrappedTokenGateway is an auto generated Go binding around an Ethereum contract.
type IWrappedTokenGateway struct {
	IWrappedTokenGatewayCaller     // Read-only binding to the contract
	IWrappedTokenGatewayTransactor // Write-only binding to the contract
	IWrappedTokenGatewayFilterer   // Log filterer for contract events
}

// IWrappedTokenGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWrappedTokenGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWrappedTokenGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWrappedTokenGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWrappedTokenGatewaySession struct {
	Contract     *IWrappedTokenGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IWrappedTokenGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWrappedTokenGatewayCallerSession struct {
	Contract *IWrappedTokenGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IWrappedTokenGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWrappedTokenGatewayTransactorSession struct {
	Contract     *IWrappedTokenGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IWrappedTokenGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWrappedTokenGatewayRaw struct {
	Contract *IWrappedTokenGateway // Generic contract binding to access the raw methods on
}

// IWrappedTokenGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWrappedTokenGatewayCallerRaw struct {
	Contract *IWrappedTokenGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IWrappedTokenGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWrappedTokenGatewayTransactorRaw struct {
	Contract *IWrappedTokenGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWrappedTokenGateway creates a new instance of IWrappedTokenGateway, bound to a specific deployed contract.
func NewIWrappedTokenGateway(address common.Address, backend bind.ContractBackend) (*IWrappedTokenGateway, error) {
	contract, err := bindIWrappedTokenGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenGateway{IWrappedTokenGatewayCaller: IWrappedTokenGatewayCaller{contract: contract}, IWrappedTokenGatewayTransactor: IWrappedTokenGatewayTransactor{contract: contract}, IWrappedTokenGatewayFilterer: IWrappedTokenGatewayFilterer{contract: contract}}, nil
}

// NewIWrappedTokenGatewayCaller creates a new read-only instance of IWrappedTokenGateway, bound to a specific deployed contract.
func NewIWrappedTokenGatewayCaller(address common.Address, caller bind.ContractCaller) (*IWrappedTokenGatewayCaller, error) {
	contract, err := bindIWrappedTokenGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenGatewayCaller{contract: contract}, nil
}

// NewIWrappedTokenGatewayTransactor creates a new write-only instance of IWrappedTokenGateway, bound to a specific deployed contract.
func NewIWrappedTokenGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IWrappedTokenGatewayTransactor, error) {
	contract, err := bindIWrappedTokenGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenGatewayTransactor{contract: contract}, nil
}

// NewIWrappedTokenGatewayFilterer creates a new log filterer instance of IWrappedTokenGateway, bound to a specific deployed contract.
func NewIWrappedTokenGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IWrappedTokenGatewayFilterer, error) {
	contract, err := bindIWrappedTokenGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenGatewayFilterer{contract: contract}, nil
}

// bindIWrappedTokenGateway binds a generic wrapper to an already deployed contract.
func bindIWrappedTokenGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWrappedTokenGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedTokenGateway *IWrappedTokenGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedTokenGateway.Contract.IWrappedTokenGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedTokenGateway *IWrappedTokenGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.IWrappedTokenGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedTokenGateway *IWrappedTokenGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.IWrappedTokenGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedTokenGateway *IWrappedTokenGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedTokenGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedTokenGateway *IWrappedTokenGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedTokenGateway *IWrappedTokenGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.contract.Transact(opts, method, params...)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x20e2d818.
//
// Solidity: function depositNativeToken(address staker) payable returns()
func (_IWrappedTokenGateway *IWrappedTokenGatewayTransactor) DepositNativeToken(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _IWrappedTokenGateway.contract.Transact(opts, "depositNativeToken", staker)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x20e2d818.
//
// Solidity: function depositNativeToken(address staker) payable returns()
func (_IWrappedTokenGateway *IWrappedTokenGatewaySession) DepositNativeToken(staker common.Address) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.DepositNativeToken(&_IWrappedTokenGateway.TransactOpts, staker)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x20e2d818.
//
// Solidity: function depositNativeToken(address staker) payable returns()
func (_IWrappedTokenGateway *IWrappedTokenGatewayTransactorSession) DepositNativeToken(staker common.Address) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.DepositNativeToken(&_IWrappedTokenGateway.TransactOpts, staker)
}

// WithdrawNativeTokens is a paid mutator transaction binding the contract method 0xd9d46816.
//
// Solidity: function withdrawNativeTokens((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexs, bool[] receiveAsTokens) returns()
func (_IWrappedTokenGateway *IWrappedTokenGatewayTransactor) WithdrawNativeTokens(opts *bind.TransactOpts, withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexs []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _IWrappedTokenGateway.contract.Transact(opts, "withdrawNativeTokens", withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens)
}

// WithdrawNativeTokens is a paid mutator transaction binding the contract method 0xd9d46816.
//
// Solidity: function withdrawNativeTokens((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexs, bool[] receiveAsTokens) returns()
func (_IWrappedTokenGateway *IWrappedTokenGatewaySession) WithdrawNativeTokens(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexs []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.WithdrawNativeTokens(&_IWrappedTokenGateway.TransactOpts, withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens)
}

// WithdrawNativeTokens is a paid mutator transaction binding the contract method 0xd9d46816.
//
// Solidity: function withdrawNativeTokens((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexs, bool[] receiveAsTokens) returns()
func (_IWrappedTokenGateway *IWrappedTokenGatewayTransactorSession) WithdrawNativeTokens(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexs []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _IWrappedTokenGateway.Contract.WithdrawNativeTokens(&_IWrappedTokenGateway.TransactOpts, withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens)
}
