// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IWrappedToken

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

// IWrappedTokenMetaData contains all meta data concerning the IWrappedToken contract.
var IWrappedTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"guy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IWrappedTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IWrappedTokenMetaData.ABI instead.
var IWrappedTokenABI = IWrappedTokenMetaData.ABI

// IWrappedToken is an auto generated Go binding around an Ethereum contract.
type IWrappedToken struct {
	IWrappedTokenCaller     // Read-only binding to the contract
	IWrappedTokenTransactor // Write-only binding to the contract
	IWrappedTokenFilterer   // Log filterer for contract events
}

// IWrappedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWrappedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWrappedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWrappedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrappedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWrappedTokenSession struct {
	Contract     *IWrappedToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWrappedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWrappedTokenCallerSession struct {
	Contract *IWrappedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IWrappedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWrappedTokenTransactorSession struct {
	Contract     *IWrappedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IWrappedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWrappedTokenRaw struct {
	Contract *IWrappedToken // Generic contract binding to access the raw methods on
}

// IWrappedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWrappedTokenCallerRaw struct {
	Contract *IWrappedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IWrappedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWrappedTokenTransactorRaw struct {
	Contract *IWrappedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWrappedToken creates a new instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedToken(address common.Address, backend bind.ContractBackend) (*IWrappedToken, error) {
	contract, err := bindIWrappedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWrappedToken{IWrappedTokenCaller: IWrappedTokenCaller{contract: contract}, IWrappedTokenTransactor: IWrappedTokenTransactor{contract: contract}, IWrappedTokenFilterer: IWrappedTokenFilterer{contract: contract}}, nil
}

// NewIWrappedTokenCaller creates a new read-only instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedTokenCaller(address common.Address, caller bind.ContractCaller) (*IWrappedTokenCaller, error) {
	contract, err := bindIWrappedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenCaller{contract: contract}, nil
}

// NewIWrappedTokenTransactor creates a new write-only instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IWrappedTokenTransactor, error) {
	contract, err := bindIWrappedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenTransactor{contract: contract}, nil
}

// NewIWrappedTokenFilterer creates a new log filterer instance of IWrappedToken, bound to a specific deployed contract.
func NewIWrappedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IWrappedTokenFilterer, error) {
	contract, err := bindIWrappedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWrappedTokenFilterer{contract: contract}, nil
}

// bindIWrappedToken binds a generic wrapper to an already deployed contract.
func bindIWrappedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWrappedTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedToken *IWrappedTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedToken.Contract.IWrappedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedToken *IWrappedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedToken.Contract.IWrappedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedToken *IWrappedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedToken.Contract.IWrappedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrappedToken *IWrappedTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrappedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrappedToken *IWrappedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrappedToken *IWrappedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrappedToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWrappedToken *IWrappedTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWrappedToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWrappedToken *IWrappedTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWrappedToken.Contract.BalanceOf(&_IWrappedToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWrappedToken *IWrappedTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWrappedToken.Contract.BalanceOf(&_IWrappedToken.CallOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_IWrappedToken *IWrappedTokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_IWrappedToken *IWrappedTokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Approve(&_IWrappedToken.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_IWrappedToken *IWrappedTokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Approve(&_IWrappedToken.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWrappedToken *IWrappedTokenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWrappedToken *IWrappedTokenSession) Deposit() (*types.Transaction, error) {
	return _IWrappedToken.Contract.Deposit(&_IWrappedToken.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWrappedToken *IWrappedTokenTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWrappedToken.Contract.Deposit(&_IWrappedToken.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_IWrappedToken *IWrappedTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_IWrappedToken *IWrappedTokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.TransferFrom(&_IWrappedToken.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_IWrappedToken *IWrappedTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.TransferFrom(&_IWrappedToken.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWrappedToken *IWrappedTokenTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWrappedToken *IWrappedTokenSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Withdraw(&_IWrappedToken.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWrappedToken *IWrappedTokenTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWrappedToken.Contract.Withdraw(&_IWrappedToken.TransactOpts, arg0)
}
