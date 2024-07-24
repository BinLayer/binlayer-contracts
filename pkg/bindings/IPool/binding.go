// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPool

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

// IPoolMetaData contains all meta data concerning the IPool contract.
var IPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolMetaData.ABI instead.
var IPoolABI = IPoolMetaData.ABI

// IPool is an auto generated Go binding around an Ethereum contract.
type IPool struct {
	IPoolCaller     // Read-only binding to the contract
	IPoolTransactor // Write-only binding to the contract
	IPoolFilterer   // Log filterer for contract events
}

// IPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolSession struct {
	Contract     *IPool            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolCallerSession struct {
	Contract *IPoolCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolTransactorSession struct {
	Contract     *IPoolTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolRaw struct {
	Contract *IPool // Generic contract binding to access the raw methods on
}

// IPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolCallerRaw struct {
	Contract *IPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolTransactorRaw struct {
	Contract *IPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPool creates a new instance of IPool, bound to a specific deployed contract.
func NewIPool(address common.Address, backend bind.ContractBackend) (*IPool, error) {
	contract, err := bindIPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPool{IPoolCaller: IPoolCaller{contract: contract}, IPoolTransactor: IPoolTransactor{contract: contract}, IPoolFilterer: IPoolFilterer{contract: contract}}, nil
}

// NewIPoolCaller creates a new read-only instance of IPool, bound to a specific deployed contract.
func NewIPoolCaller(address common.Address, caller bind.ContractCaller) (*IPoolCaller, error) {
	contract, err := bindIPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolCaller{contract: contract}, nil
}

// NewIPoolTransactor creates a new write-only instance of IPool, bound to a specific deployed contract.
func NewIPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolTransactor, error) {
	contract, err := bindIPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolTransactor{contract: contract}, nil
}

// NewIPoolFilterer creates a new log filterer instance of IPool, bound to a specific deployed contract.
func NewIPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolFilterer, error) {
	contract, err := bindIPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolFilterer{contract: contract}, nil
}

// bindIPool binds a generic wrapper to an already deployed contract.
func bindIPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.IPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IPool *IPoolCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IPool *IPoolSession) Explanation() (string, error) {
	return _IPool.Contract.Explanation(&_IPool.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_IPool *IPoolCallerSession) Explanation() (string, error) {
	return _IPool.Contract.Explanation(&_IPool.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IPool *IPoolCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IPool *IPoolSession) Shares(user common.Address) (*big.Int, error) {
	return _IPool.Contract.Shares(&_IPool.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_IPool *IPoolCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _IPool.Contract.Shares(&_IPool.CallOpts, user)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IPool *IPoolCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IPool *IPoolSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _IPool.Contract.SharesToUnderlyingView(&_IPool.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_IPool *IPoolCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _IPool.Contract.SharesToUnderlyingView(&_IPool.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IPool *IPoolCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IPool *IPoolSession) TotalShares() (*big.Int, error) {
	return _IPool.Contract.TotalShares(&_IPool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_IPool *IPoolCallerSession) TotalShares() (*big.Int, error) {
	return _IPool.Contract.TotalShares(&_IPool.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IPool *IPoolCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IPool *IPoolSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _IPool.Contract.UnderlyingToSharesView(&_IPool.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_IPool *IPoolCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _IPool.Contract.UnderlyingToSharesView(&_IPool.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IPool *IPoolCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IPool *IPoolSession) UnderlyingToken() (common.Address, error) {
	return _IPool.Contract.UnderlyingToken(&_IPool.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_IPool *IPoolCallerSession) UnderlyingToken() (common.Address, error) {
	return _IPool.Contract.UnderlyingToken(&_IPool.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IPool *IPoolCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPool.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IPool *IPoolSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _IPool.Contract.UserUnderlyingView(&_IPool.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_IPool *IPoolCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _IPool.Contract.UserUnderlyingView(&_IPool.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IPool *IPoolTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IPool *IPoolSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Deposit(&_IPool.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_IPool *IPoolTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Deposit(&_IPool.TransactOpts, token, amount)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IPool *IPoolTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IPool *IPoolSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.SharesToUnderlying(&_IPool.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_IPool *IPoolTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.SharesToUnderlying(&_IPool.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IPool *IPoolTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IPool *IPoolSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.UnderlyingToShares(&_IPool.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_IPool *IPoolTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.UnderlyingToShares(&_IPool.TransactOpts, amountUnderlying)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IPool *IPoolTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IPool *IPoolSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _IPool.Contract.UserUnderlying(&_IPool.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_IPool *IPoolTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _IPool.Contract.UserUnderlying(&_IPool.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_IPool *IPoolTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_IPool *IPoolSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Withdraw(&_IPool.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_IPool *IPoolTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Withdraw(&_IPool.TransactOpts, recipient, token, amountShares)
}
