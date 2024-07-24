// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EmptyContract

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

// EmptyContractMetaData contains all meta data concerning the EmptyContract contract.
var EmptyContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"foo\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50606580601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c298557814602d575b600080fd5b00fea26469706673582212203604b928547c9faaf6b94a75e3cf1ecca4e93934110f4f8e59e795e3a8dffcde64736f6c63430008140033",
}

// EmptyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use EmptyContractMetaData.ABI instead.
var EmptyContractABI = EmptyContractMetaData.ABI

// EmptyContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmptyContractMetaData.Bin instead.
var EmptyContractBin = EmptyContractMetaData.Bin

// DeployEmptyContract deploys a new Ethereum contract, binding an instance of EmptyContract to it.
func DeployEmptyContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmptyContract, error) {
	parsed, err := EmptyContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmptyContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmptyContract{EmptyContractCaller: EmptyContractCaller{contract: contract}, EmptyContractTransactor: EmptyContractTransactor{contract: contract}, EmptyContractFilterer: EmptyContractFilterer{contract: contract}}, nil
}

// EmptyContract is an auto generated Go binding around an Ethereum contract.
type EmptyContract struct {
	EmptyContractCaller     // Read-only binding to the contract
	EmptyContractTransactor // Write-only binding to the contract
	EmptyContractFilterer   // Log filterer for contract events
}

// EmptyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmptyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmptyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmptyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmptyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmptyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmptyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmptyContractSession struct {
	Contract     *EmptyContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmptyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmptyContractCallerSession struct {
	Contract *EmptyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EmptyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmptyContractTransactorSession struct {
	Contract     *EmptyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EmptyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmptyContractRaw struct {
	Contract *EmptyContract // Generic contract binding to access the raw methods on
}

// EmptyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmptyContractCallerRaw struct {
	Contract *EmptyContractCaller // Generic read-only contract binding to access the raw methods on
}

// EmptyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmptyContractTransactorRaw struct {
	Contract *EmptyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmptyContract creates a new instance of EmptyContract, bound to a specific deployed contract.
func NewEmptyContract(address common.Address, backend bind.ContractBackend) (*EmptyContract, error) {
	contract, err := bindEmptyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmptyContract{EmptyContractCaller: EmptyContractCaller{contract: contract}, EmptyContractTransactor: EmptyContractTransactor{contract: contract}, EmptyContractFilterer: EmptyContractFilterer{contract: contract}}, nil
}

// NewEmptyContractCaller creates a new read-only instance of EmptyContract, bound to a specific deployed contract.
func NewEmptyContractCaller(address common.Address, caller bind.ContractCaller) (*EmptyContractCaller, error) {
	contract, err := bindEmptyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmptyContractCaller{contract: contract}, nil
}

// NewEmptyContractTransactor creates a new write-only instance of EmptyContract, bound to a specific deployed contract.
func NewEmptyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*EmptyContractTransactor, error) {
	contract, err := bindEmptyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmptyContractTransactor{contract: contract}, nil
}

// NewEmptyContractFilterer creates a new log filterer instance of EmptyContract, bound to a specific deployed contract.
func NewEmptyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*EmptyContractFilterer, error) {
	contract, err := bindEmptyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmptyContractFilterer{contract: contract}, nil
}

// bindEmptyContract binds a generic wrapper to an already deployed contract.
func bindEmptyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EmptyContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmptyContract *EmptyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmptyContract.Contract.EmptyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmptyContract *EmptyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmptyContract.Contract.EmptyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmptyContract *EmptyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmptyContract.Contract.EmptyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmptyContract *EmptyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmptyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmptyContract *EmptyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmptyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmptyContract *EmptyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmptyContract.Contract.contract.Transact(opts, method, params...)
}

// Foo is a paid mutator transaction binding the contract method 0xc2985578.
//
// Solidity: function foo() returns()
func (_EmptyContract *EmptyContractTransactor) Foo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmptyContract.contract.Transact(opts, "foo")
}

// Foo is a paid mutator transaction binding the contract method 0xc2985578.
//
// Solidity: function foo() returns()
func (_EmptyContract *EmptyContractSession) Foo() (*types.Transaction, error) {
	return _EmptyContract.Contract.Foo(&_EmptyContract.TransactOpts)
}

// Foo is a paid mutator transaction binding the contract method 0xc2985578.
//
// Solidity: function foo() returns()
func (_EmptyContract *EmptyContractTransactorSession) Foo() (*types.Transaction, error) {
	return _EmptyContract.Contract.Foo(&_EmptyContract.TransactOpts)
}
