// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC677Receiver

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

// IERC677ReceiverMetaData contains all meta data concerning the IERC677Receiver contract.
var IERC677ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onTokenTransfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IERC677ReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC677ReceiverMetaData.ABI instead.
var IERC677ReceiverABI = IERC677ReceiverMetaData.ABI

// IERC677Receiver is an auto generated Go binding around an Ethereum contract.
type IERC677Receiver struct {
	IERC677ReceiverCaller     // Read-only binding to the contract
	IERC677ReceiverTransactor // Write-only binding to the contract
	IERC677ReceiverFilterer   // Log filterer for contract events
}

// IERC677ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC677ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC677ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC677ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC677ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC677ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC677ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC677ReceiverSession struct {
	Contract     *IERC677Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC677ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC677ReceiverCallerSession struct {
	Contract *IERC677ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC677ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC677ReceiverTransactorSession struct {
	Contract     *IERC677ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC677ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC677ReceiverRaw struct {
	Contract *IERC677Receiver // Generic contract binding to access the raw methods on
}

// IERC677ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC677ReceiverCallerRaw struct {
	Contract *IERC677ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC677ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC677ReceiverTransactorRaw struct {
	Contract *IERC677ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC677Receiver creates a new instance of IERC677Receiver, bound to a specific deployed contract.
func NewIERC677Receiver(address common.Address, backend bind.ContractBackend) (*IERC677Receiver, error) {
	contract, err := bindIERC677Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC677Receiver{IERC677ReceiverCaller: IERC677ReceiverCaller{contract: contract}, IERC677ReceiverTransactor: IERC677ReceiverTransactor{contract: contract}, IERC677ReceiverFilterer: IERC677ReceiverFilterer{contract: contract}}, nil
}

// NewIERC677ReceiverCaller creates a new read-only instance of IERC677Receiver, bound to a specific deployed contract.
func NewIERC677ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC677ReceiverCaller, error) {
	contract, err := bindIERC677Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC677ReceiverCaller{contract: contract}, nil
}

// NewIERC677ReceiverTransactor creates a new write-only instance of IERC677Receiver, bound to a specific deployed contract.
func NewIERC677ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC677ReceiverTransactor, error) {
	contract, err := bindIERC677Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC677ReceiverTransactor{contract: contract}, nil
}

// NewIERC677ReceiverFilterer creates a new log filterer instance of IERC677Receiver, bound to a specific deployed contract.
func NewIERC677ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC677ReceiverFilterer, error) {
	contract, err := bindIERC677Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC677ReceiverFilterer{contract: contract}, nil
}

// bindIERC677Receiver binds a generic wrapper to an already deployed contract.
func bindIERC677Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC677ReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC677Receiver *IERC677ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC677Receiver.Contract.IERC677ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC677Receiver *IERC677ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.IERC677ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC677Receiver *IERC677ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.IERC677ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC677Receiver *IERC677ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC677Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC677Receiver *IERC677ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC677Receiver *IERC677ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_IERC677Receiver *IERC677ReceiverTransactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677Receiver.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_IERC677Receiver *IERC677ReceiverSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.OnTokenTransfer(&_IERC677Receiver.TransactOpts, sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_IERC677Receiver *IERC677ReceiverTransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.OnTokenTransfer(&_IERC677Receiver.TransactOpts, sender, amount, data)
}
