// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IListaStakeManager

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

// IListaStakeManagerMetaData contains all meta data concerning the IListaStakeManager contract.
var IListaStakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// IListaStakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IListaStakeManagerMetaData.ABI instead.
var IListaStakeManagerABI = IListaStakeManagerMetaData.ABI

// IListaStakeManager is an auto generated Go binding around an Ethereum contract.
type IListaStakeManager struct {
	IListaStakeManagerCaller     // Read-only binding to the contract
	IListaStakeManagerTransactor // Write-only binding to the contract
	IListaStakeManagerFilterer   // Log filterer for contract events
}

// IListaStakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IListaStakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IListaStakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IListaStakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IListaStakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IListaStakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IListaStakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IListaStakeManagerSession struct {
	Contract     *IListaStakeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IListaStakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IListaStakeManagerCallerSession struct {
	Contract *IListaStakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IListaStakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IListaStakeManagerTransactorSession struct {
	Contract     *IListaStakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IListaStakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IListaStakeManagerRaw struct {
	Contract *IListaStakeManager // Generic contract binding to access the raw methods on
}

// IListaStakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IListaStakeManagerCallerRaw struct {
	Contract *IListaStakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IListaStakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IListaStakeManagerTransactorRaw struct {
	Contract *IListaStakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIListaStakeManager creates a new instance of IListaStakeManager, bound to a specific deployed contract.
func NewIListaStakeManager(address common.Address, backend bind.ContractBackend) (*IListaStakeManager, error) {
	contract, err := bindIListaStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IListaStakeManager{IListaStakeManagerCaller: IListaStakeManagerCaller{contract: contract}, IListaStakeManagerTransactor: IListaStakeManagerTransactor{contract: contract}, IListaStakeManagerFilterer: IListaStakeManagerFilterer{contract: contract}}, nil
}

// NewIListaStakeManagerCaller creates a new read-only instance of IListaStakeManager, bound to a specific deployed contract.
func NewIListaStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*IListaStakeManagerCaller, error) {
	contract, err := bindIListaStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IListaStakeManagerCaller{contract: contract}, nil
}

// NewIListaStakeManagerTransactor creates a new write-only instance of IListaStakeManager, bound to a specific deployed contract.
func NewIListaStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IListaStakeManagerTransactor, error) {
	contract, err := bindIListaStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IListaStakeManagerTransactor{contract: contract}, nil
}

// NewIListaStakeManagerFilterer creates a new log filterer instance of IListaStakeManager, bound to a specific deployed contract.
func NewIListaStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IListaStakeManagerFilterer, error) {
	contract, err := bindIListaStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IListaStakeManagerFilterer{contract: contract}, nil
}

// bindIListaStakeManager binds a generic wrapper to an already deployed contract.
func bindIListaStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IListaStakeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IListaStakeManager *IListaStakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IListaStakeManager.Contract.IListaStakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IListaStakeManager *IListaStakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IListaStakeManager.Contract.IListaStakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IListaStakeManager *IListaStakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IListaStakeManager.Contract.IListaStakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IListaStakeManager *IListaStakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IListaStakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IListaStakeManager *IListaStakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IListaStakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IListaStakeManager *IListaStakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IListaStakeManager.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IListaStakeManager *IListaStakeManagerTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IListaStakeManager.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IListaStakeManager *IListaStakeManagerSession) Deposit() (*types.Transaction, error) {
	return _IListaStakeManager.Contract.Deposit(&_IListaStakeManager.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IListaStakeManager *IListaStakeManagerTransactorSession) Deposit() (*types.Transaction, error) {
	return _IListaStakeManager.Contract.Deposit(&_IListaStakeManager.TransactOpts)
}
