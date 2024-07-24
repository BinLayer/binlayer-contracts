// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPoolController

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

// IPoolControllerMetaData contains all meta data concerning the IPoolController contract.
var IPoolControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addPoolsToDepositWhitelist\",\"inputs\":[{\"name\":\"poolsToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoPool\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoPoolWithSignature\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoPoolWithStaker\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePoolsFromDepositWhitelist\",\"inputs\":[{\"name\":\"poolsToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerPoolListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerPoolShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// IPoolControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolControllerMetaData.ABI instead.
var IPoolControllerABI = IPoolControllerMetaData.ABI

// IPoolController is an auto generated Go binding around an Ethereum contract.
type IPoolController struct {
	IPoolControllerCaller     // Read-only binding to the contract
	IPoolControllerTransactor // Write-only binding to the contract
	IPoolControllerFilterer   // Log filterer for contract events
}

// IPoolControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolControllerSession struct {
	Contract     *IPoolController  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolControllerCallerSession struct {
	Contract *IPoolControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPoolControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolControllerTransactorSession struct {
	Contract     *IPoolControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPoolControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolControllerRaw struct {
	Contract *IPoolController // Generic contract binding to access the raw methods on
}

// IPoolControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolControllerCallerRaw struct {
	Contract *IPoolControllerCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolControllerTransactorRaw struct {
	Contract *IPoolControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolController creates a new instance of IPoolController, bound to a specific deployed contract.
func NewIPoolController(address common.Address, backend bind.ContractBackend) (*IPoolController, error) {
	contract, err := bindIPoolController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolController{IPoolControllerCaller: IPoolControllerCaller{contract: contract}, IPoolControllerTransactor: IPoolControllerTransactor{contract: contract}, IPoolControllerFilterer: IPoolControllerFilterer{contract: contract}}, nil
}

// NewIPoolControllerCaller creates a new read-only instance of IPoolController, bound to a specific deployed contract.
func NewIPoolControllerCaller(address common.Address, caller bind.ContractCaller) (*IPoolControllerCaller, error) {
	contract, err := bindIPoolController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolControllerCaller{contract: contract}, nil
}

// NewIPoolControllerTransactor creates a new write-only instance of IPoolController, bound to a specific deployed contract.
func NewIPoolControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolControllerTransactor, error) {
	contract, err := bindIPoolController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolControllerTransactor{contract: contract}, nil
}

// NewIPoolControllerFilterer creates a new log filterer instance of IPoolController, bound to a specific deployed contract.
func NewIPoolControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolControllerFilterer, error) {
	contract, err := bindIPoolController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolControllerFilterer{contract: contract}, nil
}

// bindIPoolController binds a generic wrapper to an already deployed contract.
func bindIPoolController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoolControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolController *IPoolControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolController.Contract.IPoolControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolController *IPoolControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolController.Contract.IPoolControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolController *IPoolControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolController.Contract.IPoolControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolController *IPoolControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolController *IPoolControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolController *IPoolControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolController.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_IPoolController *IPoolControllerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_IPoolController *IPoolControllerSession) Delegation() (common.Address, error) {
	return _IPoolController.Contract.Delegation(&_IPoolController.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_IPoolController *IPoolControllerCallerSession) Delegation() (common.Address, error) {
	return _IPoolController.Contract.Delegation(&_IPoolController.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_IPoolController *IPoolControllerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_IPoolController *IPoolControllerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _IPoolController.Contract.GetDeposits(&_IPoolController.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_IPoolController *IPoolControllerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _IPoolController.Contract.GetDeposits(&_IPoolController.CallOpts, staker)
}

// PoolIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0xd8f41f74.
//
// Solidity: function poolIsWhitelistedForDeposit(address pool) view returns(bool)
func (_IPoolController *IPoolControllerCaller) PoolIsWhitelistedForDeposit(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "poolIsWhitelistedForDeposit", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0xd8f41f74.
//
// Solidity: function poolIsWhitelistedForDeposit(address pool) view returns(bool)
func (_IPoolController *IPoolControllerSession) PoolIsWhitelistedForDeposit(pool common.Address) (bool, error) {
	return _IPoolController.Contract.PoolIsWhitelistedForDeposit(&_IPoolController.CallOpts, pool)
}

// PoolIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0xd8f41f74.
//
// Solidity: function poolIsWhitelistedForDeposit(address pool) view returns(bool)
func (_IPoolController *IPoolControllerCallerSession) PoolIsWhitelistedForDeposit(pool common.Address) (bool, error) {
	return _IPoolController.Contract.PoolIsWhitelistedForDeposit(&_IPoolController.CallOpts, pool)
}

// PoolWhitelister is a free data retrieval call binding the contract method 0xd02eaed6.
//
// Solidity: function poolWhitelister() view returns(address)
func (_IPoolController *IPoolControllerCaller) PoolWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "poolWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolWhitelister is a free data retrieval call binding the contract method 0xd02eaed6.
//
// Solidity: function poolWhitelister() view returns(address)
func (_IPoolController *IPoolControllerSession) PoolWhitelister() (common.Address, error) {
	return _IPoolController.Contract.PoolWhitelister(&_IPoolController.CallOpts)
}

// PoolWhitelister is a free data retrieval call binding the contract method 0xd02eaed6.
//
// Solidity: function poolWhitelister() view returns(address)
func (_IPoolController *IPoolControllerCallerSession) PoolWhitelister() (common.Address, error) {
	return _IPoolController.Contract.PoolWhitelister(&_IPoolController.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IPoolController *IPoolControllerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IPoolController *IPoolControllerSession) Slasher() (common.Address, error) {
	return _IPoolController.Contract.Slasher(&_IPoolController.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_IPoolController *IPoolControllerCallerSession) Slasher() (common.Address, error) {
	return _IPoolController.Contract.Slasher(&_IPoolController.CallOpts)
}

// StakerPoolListLength is a free data retrieval call binding the contract method 0xdf8c804c.
//
// Solidity: function stakerPoolListLength(address staker) view returns(uint256)
func (_IPoolController *IPoolControllerCaller) StakerPoolListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "stakerPoolListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerPoolListLength is a free data retrieval call binding the contract method 0xdf8c804c.
//
// Solidity: function stakerPoolListLength(address staker) view returns(uint256)
func (_IPoolController *IPoolControllerSession) StakerPoolListLength(staker common.Address) (*big.Int, error) {
	return _IPoolController.Contract.StakerPoolListLength(&_IPoolController.CallOpts, staker)
}

// StakerPoolListLength is a free data retrieval call binding the contract method 0xdf8c804c.
//
// Solidity: function stakerPoolListLength(address staker) view returns(uint256)
func (_IPoolController *IPoolControllerCallerSession) StakerPoolListLength(staker common.Address) (*big.Int, error) {
	return _IPoolController.Contract.StakerPoolListLength(&_IPoolController.CallOpts, staker)
}

// StakerPoolShares is a free data retrieval call binding the contract method 0xb6230d5f.
//
// Solidity: function stakerPoolShares(address user, address pool) view returns(uint256 shares)
func (_IPoolController *IPoolControllerCaller) StakerPoolShares(opts *bind.CallOpts, user common.Address, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "stakerPoolShares", user, pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerPoolShares is a free data retrieval call binding the contract method 0xb6230d5f.
//
// Solidity: function stakerPoolShares(address user, address pool) view returns(uint256 shares)
func (_IPoolController *IPoolControllerSession) StakerPoolShares(user common.Address, pool common.Address) (*big.Int, error) {
	return _IPoolController.Contract.StakerPoolShares(&_IPoolController.CallOpts, user, pool)
}

// StakerPoolShares is a free data retrieval call binding the contract method 0xb6230d5f.
//
// Solidity: function stakerPoolShares(address user, address pool) view returns(uint256 shares)
func (_IPoolController *IPoolControllerCallerSession) StakerPoolShares(user common.Address, pool common.Address) (*big.Int, error) {
	return _IPoolController.Contract.StakerPoolShares(&_IPoolController.CallOpts, user, pool)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address pool) view returns(bool)
func (_IPoolController *IPoolControllerCaller) ThirdPartyTransfersForbidden(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _IPoolController.contract.Call(opts, &out, "thirdPartyTransfersForbidden", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address pool) view returns(bool)
func (_IPoolController *IPoolControllerSession) ThirdPartyTransfersForbidden(pool common.Address) (bool, error) {
	return _IPoolController.Contract.ThirdPartyTransfersForbidden(&_IPoolController.CallOpts, pool)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address pool) view returns(bool)
func (_IPoolController *IPoolControllerCallerSession) ThirdPartyTransfersForbidden(pool common.Address) (bool, error) {
	return _IPoolController.Contract.ThirdPartyTransfersForbidden(&_IPoolController.CallOpts, pool)
}

// AddPoolsToDepositWhitelist is a paid mutator transaction binding the contract method 0xa185a29a.
//
// Solidity: function addPoolsToDepositWhitelist(address[] poolsToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_IPoolController *IPoolControllerTransactor) AddPoolsToDepositWhitelist(opts *bind.TransactOpts, poolsToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "addPoolsToDepositWhitelist", poolsToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddPoolsToDepositWhitelist is a paid mutator transaction binding the contract method 0xa185a29a.
//
// Solidity: function addPoolsToDepositWhitelist(address[] poolsToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_IPoolController *IPoolControllerSession) AddPoolsToDepositWhitelist(poolsToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _IPoolController.Contract.AddPoolsToDepositWhitelist(&_IPoolController.TransactOpts, poolsToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddPoolsToDepositWhitelist is a paid mutator transaction binding the contract method 0xa185a29a.
//
// Solidity: function addPoolsToDepositWhitelist(address[] poolsToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_IPoolController *IPoolControllerTransactorSession) AddPoolsToDepositWhitelist(poolsToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _IPoolController.Contract.AddPoolsToDepositWhitelist(&_IPoolController.TransactOpts, poolsToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address pool, uint256 shares) returns()
func (_IPoolController *IPoolControllerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, token common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "addShares", staker, token, pool, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address pool, uint256 shares) returns()
func (_IPoolController *IPoolControllerSession) AddShares(staker common.Address, token common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.AddShares(&_IPoolController.TransactOpts, staker, token, pool, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address pool, uint256 shares) returns()
func (_IPoolController *IPoolControllerTransactorSession) AddShares(staker common.Address, token common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.AddShares(&_IPoolController.TransactOpts, staker, token, pool, shares)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x6fcafe5a.
//
// Solidity: function depositIntoPool(address pool, address token, uint256 amount) returns(uint256 shares)
func (_IPoolController *IPoolControllerTransactor) DepositIntoPool(opts *bind.TransactOpts, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "depositIntoPool", pool, token, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x6fcafe5a.
//
// Solidity: function depositIntoPool(address pool, address token, uint256 amount) returns(uint256 shares)
func (_IPoolController *IPoolControllerSession) DepositIntoPool(pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.DepositIntoPool(&_IPoolController.TransactOpts, pool, token, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x6fcafe5a.
//
// Solidity: function depositIntoPool(address pool, address token, uint256 amount) returns(uint256 shares)
func (_IPoolController *IPoolControllerTransactorSession) DepositIntoPool(pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.DepositIntoPool(&_IPoolController.TransactOpts, pool, token, amount)
}

// DepositIntoPoolWithSignature is a paid mutator transaction binding the contract method 0x0e4ed53d.
//
// Solidity: function depositIntoPoolWithSignature(address pool, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_IPoolController *IPoolControllerTransactor) DepositIntoPoolWithSignature(opts *bind.TransactOpts, pool common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "depositIntoPoolWithSignature", pool, token, amount, staker, expiry, signature)
}

// DepositIntoPoolWithSignature is a paid mutator transaction binding the contract method 0x0e4ed53d.
//
// Solidity: function depositIntoPoolWithSignature(address pool, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_IPoolController *IPoolControllerSession) DepositIntoPoolWithSignature(pool common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IPoolController.Contract.DepositIntoPoolWithSignature(&_IPoolController.TransactOpts, pool, token, amount, staker, expiry, signature)
}

// DepositIntoPoolWithSignature is a paid mutator transaction binding the contract method 0x0e4ed53d.
//
// Solidity: function depositIntoPoolWithSignature(address pool, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_IPoolController *IPoolControllerTransactorSession) DepositIntoPoolWithSignature(pool common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _IPoolController.Contract.DepositIntoPoolWithSignature(&_IPoolController.TransactOpts, pool, token, amount, staker, expiry, signature)
}

// DepositIntoPoolWithStaker is a paid mutator transaction binding the contract method 0x856abb29.
//
// Solidity: function depositIntoPoolWithStaker(address staker, address pool, address token, uint256 amount) returns(uint256 shares)
func (_IPoolController *IPoolControllerTransactor) DepositIntoPoolWithStaker(opts *bind.TransactOpts, staker common.Address, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "depositIntoPoolWithStaker", staker, pool, token, amount)
}

// DepositIntoPoolWithStaker is a paid mutator transaction binding the contract method 0x856abb29.
//
// Solidity: function depositIntoPoolWithStaker(address staker, address pool, address token, uint256 amount) returns(uint256 shares)
func (_IPoolController *IPoolControllerSession) DepositIntoPoolWithStaker(staker common.Address, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.DepositIntoPoolWithStaker(&_IPoolController.TransactOpts, staker, pool, token, amount)
}

// DepositIntoPoolWithStaker is a paid mutator transaction binding the contract method 0x856abb29.
//
// Solidity: function depositIntoPoolWithStaker(address staker, address pool, address token, uint256 amount) returns(uint256 shares)
func (_IPoolController *IPoolControllerTransactorSession) DepositIntoPoolWithStaker(staker common.Address, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.DepositIntoPoolWithStaker(&_IPoolController.TransactOpts, staker, pool, token, amount)
}

// RemovePoolsFromDepositWhitelist is a paid mutator transaction binding the contract method 0x53721791.
//
// Solidity: function removePoolsFromDepositWhitelist(address[] poolsToRemoveFromWhitelist) returns()
func (_IPoolController *IPoolControllerTransactor) RemovePoolsFromDepositWhitelist(opts *bind.TransactOpts, poolsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "removePoolsFromDepositWhitelist", poolsToRemoveFromWhitelist)
}

// RemovePoolsFromDepositWhitelist is a paid mutator transaction binding the contract method 0x53721791.
//
// Solidity: function removePoolsFromDepositWhitelist(address[] poolsToRemoveFromWhitelist) returns()
func (_IPoolController *IPoolControllerSession) RemovePoolsFromDepositWhitelist(poolsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IPoolController.Contract.RemovePoolsFromDepositWhitelist(&_IPoolController.TransactOpts, poolsToRemoveFromWhitelist)
}

// RemovePoolsFromDepositWhitelist is a paid mutator transaction binding the contract method 0x53721791.
//
// Solidity: function removePoolsFromDepositWhitelist(address[] poolsToRemoveFromWhitelist) returns()
func (_IPoolController *IPoolControllerTransactorSession) RemovePoolsFromDepositWhitelist(poolsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _IPoolController.Contract.RemovePoolsFromDepositWhitelist(&_IPoolController.TransactOpts, poolsToRemoveFromWhitelist)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address pool, uint256 shares) returns()
func (_IPoolController *IPoolControllerTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "removeShares", staker, pool, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address pool, uint256 shares) returns()
func (_IPoolController *IPoolControllerSession) RemoveShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.RemoveShares(&_IPoolController.TransactOpts, staker, pool, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address pool, uint256 shares) returns()
func (_IPoolController *IPoolControllerTransactorSession) RemoveShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IPoolController.Contract.RemoveShares(&_IPoolController.TransactOpts, staker, pool, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address pool, uint256 shares, address token) returns()
func (_IPoolController *IPoolControllerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, recipient common.Address, pool common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _IPoolController.contract.Transact(opts, "withdrawSharesAsTokens", recipient, pool, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address pool, uint256 shares, address token) returns()
func (_IPoolController *IPoolControllerSession) WithdrawSharesAsTokens(recipient common.Address, pool common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _IPoolController.Contract.WithdrawSharesAsTokens(&_IPoolController.TransactOpts, recipient, pool, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address pool, uint256 shares, address token) returns()
func (_IPoolController *IPoolControllerTransactorSession) WithdrawSharesAsTokens(recipient common.Address, pool common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _IPoolController.Contract.WithdrawSharesAsTokens(&_IPoolController.TransactOpts, recipient, pool, shares, token)
}

// IPoolControllerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IPoolController contract.
type IPoolControllerDepositIterator struct {
	Event *IPoolControllerDeposit // Event containing the contract specifics and raw log

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
func (it *IPoolControllerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolControllerDeposit)
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
		it.Event = new(IPoolControllerDeposit)
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
func (it *IPoolControllerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolControllerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolControllerDeposit represents a Deposit event raised by the IPoolController contract.
type IPoolControllerDeposit struct {
	Staker common.Address
	Token  common.Address
	Pool   common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address pool, uint256 shares)
func (_IPoolController *IPoolControllerFilterer) FilterDeposit(opts *bind.FilterOpts) (*IPoolControllerDepositIterator, error) {

	logs, sub, err := _IPoolController.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &IPoolControllerDepositIterator{contract: _IPoolController.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address pool, uint256 shares)
func (_IPoolController *IPoolControllerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IPoolControllerDeposit) (event.Subscription, error) {

	logs, sub, err := _IPoolController.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolControllerDeposit)
				if err := _IPoolController.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address pool, uint256 shares)
func (_IPoolController *IPoolControllerFilterer) ParseDeposit(log types.Log) (*IPoolControllerDeposit, error) {
	event := new(IPoolControllerDeposit)
	if err := _IPoolController.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolControllerPoolAddedToDepositWhitelistIterator is returned from FilterPoolAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for PoolAddedToDepositWhitelist events raised by the IPoolController contract.
type IPoolControllerPoolAddedToDepositWhitelistIterator struct {
	Event *IPoolControllerPoolAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *IPoolControllerPoolAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolControllerPoolAddedToDepositWhitelist)
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
		it.Event = new(IPoolControllerPoolAddedToDepositWhitelist)
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
func (it *IPoolControllerPoolAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolControllerPoolAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolControllerPoolAddedToDepositWhitelist represents a PoolAddedToDepositWhitelist event raised by the IPoolController contract.
type IPoolControllerPoolAddedToDepositWhitelist struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x39daccb429428846b0171894bd4200a7051a1c25c7b82ee455a39f89d72e634c.
//
// Solidity: event PoolAddedToDepositWhitelist(address pool)
func (_IPoolController *IPoolControllerFilterer) FilterPoolAddedToDepositWhitelist(opts *bind.FilterOpts) (*IPoolControllerPoolAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _IPoolController.contract.FilterLogs(opts, "PoolAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &IPoolControllerPoolAddedToDepositWhitelistIterator{contract: _IPoolController.contract, event: "PoolAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchPoolAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x39daccb429428846b0171894bd4200a7051a1c25c7b82ee455a39f89d72e634c.
//
// Solidity: event PoolAddedToDepositWhitelist(address pool)
func (_IPoolController *IPoolControllerFilterer) WatchPoolAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *IPoolControllerPoolAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _IPoolController.contract.WatchLogs(opts, "PoolAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolControllerPoolAddedToDepositWhitelist)
				if err := _IPoolController.contract.UnpackLog(event, "PoolAddedToDepositWhitelist", log); err != nil {
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

// ParsePoolAddedToDepositWhitelist is a log parse operation binding the contract event 0x39daccb429428846b0171894bd4200a7051a1c25c7b82ee455a39f89d72e634c.
//
// Solidity: event PoolAddedToDepositWhitelist(address pool)
func (_IPoolController *IPoolControllerFilterer) ParsePoolAddedToDepositWhitelist(log types.Log) (*IPoolControllerPoolAddedToDepositWhitelist, error) {
	event := new(IPoolControllerPoolAddedToDepositWhitelist)
	if err := _IPoolController.contract.UnpackLog(event, "PoolAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolControllerPoolRemovedFromDepositWhitelistIterator is returned from FilterPoolRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for PoolRemovedFromDepositWhitelist events raised by the IPoolController contract.
type IPoolControllerPoolRemovedFromDepositWhitelistIterator struct {
	Event *IPoolControllerPoolRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *IPoolControllerPoolRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolControllerPoolRemovedFromDepositWhitelist)
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
		it.Event = new(IPoolControllerPoolRemovedFromDepositWhitelist)
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
func (it *IPoolControllerPoolRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolControllerPoolRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolControllerPoolRemovedFromDepositWhitelist represents a PoolRemovedFromDepositWhitelist event raised by the IPoolController contract.
type IPoolControllerPoolRemovedFromDepositWhitelist struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x8ac71de5f0eed990ceee2d31cbfe4280db9f27a8f0bf5567c507309de33ed723.
//
// Solidity: event PoolRemovedFromDepositWhitelist(address pool)
func (_IPoolController *IPoolControllerFilterer) FilterPoolRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*IPoolControllerPoolRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _IPoolController.contract.FilterLogs(opts, "PoolRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &IPoolControllerPoolRemovedFromDepositWhitelistIterator{contract: _IPoolController.contract, event: "PoolRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchPoolRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x8ac71de5f0eed990ceee2d31cbfe4280db9f27a8f0bf5567c507309de33ed723.
//
// Solidity: event PoolRemovedFromDepositWhitelist(address pool)
func (_IPoolController *IPoolControllerFilterer) WatchPoolRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *IPoolControllerPoolRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _IPoolController.contract.WatchLogs(opts, "PoolRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolControllerPoolRemovedFromDepositWhitelist)
				if err := _IPoolController.contract.UnpackLog(event, "PoolRemovedFromDepositWhitelist", log); err != nil {
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

// ParsePoolRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x8ac71de5f0eed990ceee2d31cbfe4280db9f27a8f0bf5567c507309de33ed723.
//
// Solidity: event PoolRemovedFromDepositWhitelist(address pool)
func (_IPoolController *IPoolControllerFilterer) ParsePoolRemovedFromDepositWhitelist(log types.Log) (*IPoolControllerPoolRemovedFromDepositWhitelist, error) {
	event := new(IPoolControllerPoolRemovedFromDepositWhitelist)
	if err := _IPoolController.contract.UnpackLog(event, "PoolRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolControllerPoolWhitelisterChangedIterator is returned from FilterPoolWhitelisterChanged and is used to iterate over the raw logs and unpacked data for PoolWhitelisterChanged events raised by the IPoolController contract.
type IPoolControllerPoolWhitelisterChangedIterator struct {
	Event *IPoolControllerPoolWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *IPoolControllerPoolWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolControllerPoolWhitelisterChanged)
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
		it.Event = new(IPoolControllerPoolWhitelisterChanged)
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
func (it *IPoolControllerPoolWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolControllerPoolWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolControllerPoolWhitelisterChanged represents a PoolWhitelisterChanged event raised by the IPoolController contract.
type IPoolControllerPoolWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPoolWhitelisterChanged is a free log retrieval operation binding the contract event 0xb10a4c4f1bee1e88074ca37b6c3ce3990495d771d4f7735a5fed00a773af424f.
//
// Solidity: event PoolWhitelisterChanged(address previousAddress, address newAddress)
func (_IPoolController *IPoolControllerFilterer) FilterPoolWhitelisterChanged(opts *bind.FilterOpts) (*IPoolControllerPoolWhitelisterChangedIterator, error) {

	logs, sub, err := _IPoolController.contract.FilterLogs(opts, "PoolWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &IPoolControllerPoolWhitelisterChangedIterator{contract: _IPoolController.contract, event: "PoolWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchPoolWhitelisterChanged is a free log subscription operation binding the contract event 0xb10a4c4f1bee1e88074ca37b6c3ce3990495d771d4f7735a5fed00a773af424f.
//
// Solidity: event PoolWhitelisterChanged(address previousAddress, address newAddress)
func (_IPoolController *IPoolControllerFilterer) WatchPoolWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *IPoolControllerPoolWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _IPoolController.contract.WatchLogs(opts, "PoolWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolControllerPoolWhitelisterChanged)
				if err := _IPoolController.contract.UnpackLog(event, "PoolWhitelisterChanged", log); err != nil {
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

// ParsePoolWhitelisterChanged is a log parse operation binding the contract event 0xb10a4c4f1bee1e88074ca37b6c3ce3990495d771d4f7735a5fed00a773af424f.
//
// Solidity: event PoolWhitelisterChanged(address previousAddress, address newAddress)
func (_IPoolController *IPoolControllerFilterer) ParsePoolWhitelisterChanged(log types.Log) (*IPoolControllerPoolWhitelisterChanged, error) {
	event := new(IPoolControllerPoolWhitelisterChanged)
	if err := _IPoolController.contract.UnpackLog(event, "PoolWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator is returned from FilterUpdatedThirdPartyTransfersForbidden and is used to iterate over the raw logs and unpacked data for UpdatedThirdPartyTransfersForbidden events raised by the IPoolController contract.
type IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator struct {
	Event *IPoolControllerUpdatedThirdPartyTransfersForbidden // Event containing the contract specifics and raw log

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
func (it *IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolControllerUpdatedThirdPartyTransfersForbidden)
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
		it.Event = new(IPoolControllerUpdatedThirdPartyTransfersForbidden)
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
func (it *IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolControllerUpdatedThirdPartyTransfersForbidden represents a UpdatedThirdPartyTransfersForbidden event raised by the IPoolController contract.
type IPoolControllerUpdatedThirdPartyTransfersForbidden struct {
	Pool  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedThirdPartyTransfersForbidden is a free log retrieval operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address pool, bool value)
func (_IPoolController *IPoolControllerFilterer) FilterUpdatedThirdPartyTransfersForbidden(opts *bind.FilterOpts) (*IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator, error) {

	logs, sub, err := _IPoolController.contract.FilterLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return &IPoolControllerUpdatedThirdPartyTransfersForbiddenIterator{contract: _IPoolController.contract, event: "UpdatedThirdPartyTransfersForbidden", logs: logs, sub: sub}, nil
}

// WatchUpdatedThirdPartyTransfersForbidden is a free log subscription operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address pool, bool value)
func (_IPoolController *IPoolControllerFilterer) WatchUpdatedThirdPartyTransfersForbidden(opts *bind.WatchOpts, sink chan<- *IPoolControllerUpdatedThirdPartyTransfersForbidden) (event.Subscription, error) {

	logs, sub, err := _IPoolController.contract.WatchLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolControllerUpdatedThirdPartyTransfersForbidden)
				if err := _IPoolController.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
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

// ParseUpdatedThirdPartyTransfersForbidden is a log parse operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address pool, bool value)
func (_IPoolController *IPoolControllerFilterer) ParseUpdatedThirdPartyTransfersForbidden(log types.Log) (*IPoolControllerUpdatedThirdPartyTransfersForbidden, error) {
	event := new(IPoolControllerUpdatedThirdPartyTransfersForbidden)
	if err := _IPoolController.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
