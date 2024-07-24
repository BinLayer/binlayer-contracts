// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IListaGateway

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

// IListaGatewayMetaData contains all meta data concerning the IListaGateway contract.
var IListaGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"depositNativeToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// IListaGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IListaGatewayMetaData.ABI instead.
var IListaGatewayABI = IListaGatewayMetaData.ABI

// IListaGateway is an auto generated Go binding around an Ethereum contract.
type IListaGateway struct {
	IListaGatewayCaller     // Read-only binding to the contract
	IListaGatewayTransactor // Write-only binding to the contract
	IListaGatewayFilterer   // Log filterer for contract events
}

// IListaGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IListaGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IListaGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IListaGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IListaGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IListaGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IListaGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IListaGatewaySession struct {
	Contract     *IListaGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IListaGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IListaGatewayCallerSession struct {
	Contract *IListaGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IListaGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IListaGatewayTransactorSession struct {
	Contract     *IListaGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IListaGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IListaGatewayRaw struct {
	Contract *IListaGateway // Generic contract binding to access the raw methods on
}

// IListaGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IListaGatewayCallerRaw struct {
	Contract *IListaGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IListaGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IListaGatewayTransactorRaw struct {
	Contract *IListaGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIListaGateway creates a new instance of IListaGateway, bound to a specific deployed contract.
func NewIListaGateway(address common.Address, backend bind.ContractBackend) (*IListaGateway, error) {
	contract, err := bindIListaGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IListaGateway{IListaGatewayCaller: IListaGatewayCaller{contract: contract}, IListaGatewayTransactor: IListaGatewayTransactor{contract: contract}, IListaGatewayFilterer: IListaGatewayFilterer{contract: contract}}, nil
}

// NewIListaGatewayCaller creates a new read-only instance of IListaGateway, bound to a specific deployed contract.
func NewIListaGatewayCaller(address common.Address, caller bind.ContractCaller) (*IListaGatewayCaller, error) {
	contract, err := bindIListaGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IListaGatewayCaller{contract: contract}, nil
}

// NewIListaGatewayTransactor creates a new write-only instance of IListaGateway, bound to a specific deployed contract.
func NewIListaGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IListaGatewayTransactor, error) {
	contract, err := bindIListaGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IListaGatewayTransactor{contract: contract}, nil
}

// NewIListaGatewayFilterer creates a new log filterer instance of IListaGateway, bound to a specific deployed contract.
func NewIListaGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IListaGatewayFilterer, error) {
	contract, err := bindIListaGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IListaGatewayFilterer{contract: contract}, nil
}

// bindIListaGateway binds a generic wrapper to an already deployed contract.
func bindIListaGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IListaGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IListaGateway *IListaGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IListaGateway.Contract.IListaGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IListaGateway *IListaGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IListaGateway.Contract.IListaGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IListaGateway *IListaGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IListaGateway.Contract.IListaGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IListaGateway *IListaGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IListaGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IListaGateway *IListaGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IListaGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IListaGateway *IListaGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IListaGateway.Contract.contract.Transact(opts, method, params...)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_IListaGateway *IListaGatewayTransactor) DepositNativeToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IListaGateway.contract.Transact(opts, "depositNativeToken")
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_IListaGateway *IListaGatewaySession) DepositNativeToken() (*types.Transaction, error) {
	return _IListaGateway.Contract.DepositNativeToken(&_IListaGateway.TransactOpts)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_IListaGateway *IListaGatewayTransactorSession) DepositNativeToken() (*types.Transaction, error) {
	return _IListaGateway.Contract.DepositNativeToken(&_IListaGateway.TransactOpts)
}
