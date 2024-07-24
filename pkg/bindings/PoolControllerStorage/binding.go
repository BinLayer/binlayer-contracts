// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PoolControllerStorage

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

// PoolControllerStorageMetaData contains all meta data concerning the PoolControllerStorage contract.
var PoolControllerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addPoolsToDepositWhitelist\",\"inputs\":[{\"name\":\"poolsToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoPool\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoPoolWithSignature\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoPoolWithStaker\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePoolsFromDepositWhitelist\",\"inputs\":[{\"name\":\"poolsToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerPoolList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerPoolListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerPoolShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// PoolControllerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolControllerStorageMetaData.ABI instead.
var PoolControllerStorageABI = PoolControllerStorageMetaData.ABI

// PoolControllerStorage is an auto generated Go binding around an Ethereum contract.
type PoolControllerStorage struct {
	PoolControllerStorageCaller     // Read-only binding to the contract
	PoolControllerStorageTransactor // Write-only binding to the contract
	PoolControllerStorageFilterer   // Log filterer for contract events
}

// PoolControllerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolControllerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolControllerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolControllerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolControllerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolControllerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolControllerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolControllerStorageSession struct {
	Contract     *PoolControllerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolControllerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolControllerStorageCallerSession struct {
	Contract *PoolControllerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PoolControllerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolControllerStorageTransactorSession struct {
	Contract     *PoolControllerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PoolControllerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolControllerStorageRaw struct {
	Contract *PoolControllerStorage // Generic contract binding to access the raw methods on
}

// PoolControllerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolControllerStorageCallerRaw struct {
	Contract *PoolControllerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PoolControllerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolControllerStorageTransactorRaw struct {
	Contract *PoolControllerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolControllerStorage creates a new instance of PoolControllerStorage, bound to a specific deployed contract.
func NewPoolControllerStorage(address common.Address, backend bind.ContractBackend) (*PoolControllerStorage, error) {
	contract, err := bindPoolControllerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolControllerStorage{PoolControllerStorageCaller: PoolControllerStorageCaller{contract: contract}, PoolControllerStorageTransactor: PoolControllerStorageTransactor{contract: contract}, PoolControllerStorageFilterer: PoolControllerStorageFilterer{contract: contract}}, nil
}

// NewPoolControllerStorageCaller creates a new read-only instance of PoolControllerStorage, bound to a specific deployed contract.
func NewPoolControllerStorageCaller(address common.Address, caller bind.ContractCaller) (*PoolControllerStorageCaller, error) {
	contract, err := bindPoolControllerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolControllerStorageCaller{contract: contract}, nil
}

// NewPoolControllerStorageTransactor creates a new write-only instance of PoolControllerStorage, bound to a specific deployed contract.
func NewPoolControllerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolControllerStorageTransactor, error) {
	contract, err := bindPoolControllerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolControllerStorageTransactor{contract: contract}, nil
}

// NewPoolControllerStorageFilterer creates a new log filterer instance of PoolControllerStorage, bound to a specific deployed contract.
func NewPoolControllerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolControllerStorageFilterer, error) {
	contract, err := bindPoolControllerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolControllerStorageFilterer{contract: contract}, nil
}

// bindPoolControllerStorage binds a generic wrapper to an already deployed contract.
func bindPoolControllerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolControllerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolControllerStorage *PoolControllerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolControllerStorage.Contract.PoolControllerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolControllerStorage *PoolControllerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.PoolControllerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolControllerStorage *PoolControllerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.PoolControllerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolControllerStorage *PoolControllerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolControllerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolControllerStorage *PoolControllerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolControllerStorage *PoolControllerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_PoolControllerStorage *PoolControllerStorageCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_PoolControllerStorage *PoolControllerStorageSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _PoolControllerStorage.Contract.DEPOSITTYPEHASH(&_PoolControllerStorage.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _PoolControllerStorage.Contract.DEPOSITTYPEHASH(&_PoolControllerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_PoolControllerStorage *PoolControllerStorageCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_PoolControllerStorage *PoolControllerStorageSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _PoolControllerStorage.Contract.DOMAINTYPEHASH(&_PoolControllerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _PoolControllerStorage.Contract.DOMAINTYPEHASH(&_PoolControllerStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageSession) Delegation() (common.Address, error) {
	return _PoolControllerStorage.Contract.Delegation(&_PoolControllerStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) Delegation() (common.Address, error) {
	return _PoolControllerStorage.Contract.Delegation(&_PoolControllerStorage.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_PoolControllerStorage *PoolControllerStorageCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "getDeposits", staker)

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
func (_PoolControllerStorage *PoolControllerStorageSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _PoolControllerStorage.Contract.GetDeposits(&_PoolControllerStorage.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_PoolControllerStorage *PoolControllerStorageCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _PoolControllerStorage.Contract.GetDeposits(&_PoolControllerStorage.CallOpts, staker)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PoolControllerStorage.Contract.Nonces(&_PoolControllerStorage.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PoolControllerStorage.Contract.Nonces(&_PoolControllerStorage.CallOpts, arg0)
}

// PoolIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0xd8f41f74.
//
// Solidity: function poolIsWhitelistedForDeposit(address ) view returns(bool)
func (_PoolControllerStorage *PoolControllerStorageCaller) PoolIsWhitelistedForDeposit(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "poolIsWhitelistedForDeposit", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0xd8f41f74.
//
// Solidity: function poolIsWhitelistedForDeposit(address ) view returns(bool)
func (_PoolControllerStorage *PoolControllerStorageSession) PoolIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _PoolControllerStorage.Contract.PoolIsWhitelistedForDeposit(&_PoolControllerStorage.CallOpts, arg0)
}

// PoolIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0xd8f41f74.
//
// Solidity: function poolIsWhitelistedForDeposit(address ) view returns(bool)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) PoolIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _PoolControllerStorage.Contract.PoolIsWhitelistedForDeposit(&_PoolControllerStorage.CallOpts, arg0)
}

// PoolWhitelister is a free data retrieval call binding the contract method 0xd02eaed6.
//
// Solidity: function poolWhitelister() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCaller) PoolWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "poolWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolWhitelister is a free data retrieval call binding the contract method 0xd02eaed6.
//
// Solidity: function poolWhitelister() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageSession) PoolWhitelister() (common.Address, error) {
	return _PoolControllerStorage.Contract.PoolWhitelister(&_PoolControllerStorage.CallOpts)
}

// PoolWhitelister is a free data retrieval call binding the contract method 0xd02eaed6.
//
// Solidity: function poolWhitelister() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) PoolWhitelister() (common.Address, error) {
	return _PoolControllerStorage.Contract.PoolWhitelister(&_PoolControllerStorage.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageSession) Slasher() (common.Address, error) {
	return _PoolControllerStorage.Contract.Slasher(&_PoolControllerStorage.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) Slasher() (common.Address, error) {
	return _PoolControllerStorage.Contract.Slasher(&_PoolControllerStorage.CallOpts)
}

// StakerPoolList is a free data retrieval call binding the contract method 0x062860d8.
//
// Solidity: function stakerPoolList(address , uint256 ) view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCaller) StakerPoolList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "stakerPoolList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerPoolList is a free data retrieval call binding the contract method 0x062860d8.
//
// Solidity: function stakerPoolList(address , uint256 ) view returns(address)
func (_PoolControllerStorage *PoolControllerStorageSession) StakerPoolList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _PoolControllerStorage.Contract.StakerPoolList(&_PoolControllerStorage.CallOpts, arg0, arg1)
}

// StakerPoolList is a free data retrieval call binding the contract method 0x062860d8.
//
// Solidity: function stakerPoolList(address , uint256 ) view returns(address)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) StakerPoolList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _PoolControllerStorage.Contract.StakerPoolList(&_PoolControllerStorage.CallOpts, arg0, arg1)
}

// StakerPoolListLength is a free data retrieval call binding the contract method 0xdf8c804c.
//
// Solidity: function stakerPoolListLength(address staker) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageCaller) StakerPoolListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "stakerPoolListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerPoolListLength is a free data retrieval call binding the contract method 0xdf8c804c.
//
// Solidity: function stakerPoolListLength(address staker) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageSession) StakerPoolListLength(staker common.Address) (*big.Int, error) {
	return _PoolControllerStorage.Contract.StakerPoolListLength(&_PoolControllerStorage.CallOpts, staker)
}

// StakerPoolListLength is a free data retrieval call binding the contract method 0xdf8c804c.
//
// Solidity: function stakerPoolListLength(address staker) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) StakerPoolListLength(staker common.Address) (*big.Int, error) {
	return _PoolControllerStorage.Contract.StakerPoolListLength(&_PoolControllerStorage.CallOpts, staker)
}

// StakerPoolShares is a free data retrieval call binding the contract method 0xb6230d5f.
//
// Solidity: function stakerPoolShares(address , address ) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageCaller) StakerPoolShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "stakerPoolShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerPoolShares is a free data retrieval call binding the contract method 0xb6230d5f.
//
// Solidity: function stakerPoolShares(address , address ) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageSession) StakerPoolShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PoolControllerStorage.Contract.StakerPoolShares(&_PoolControllerStorage.CallOpts, arg0, arg1)
}

// StakerPoolShares is a free data retrieval call binding the contract method 0xb6230d5f.
//
// Solidity: function stakerPoolShares(address , address ) view returns(uint256)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) StakerPoolShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PoolControllerStorage.Contract.StakerPoolShares(&_PoolControllerStorage.CallOpts, arg0, arg1)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_PoolControllerStorage *PoolControllerStorageCaller) ThirdPartyTransfersForbidden(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PoolControllerStorage.contract.Call(opts, &out, "thirdPartyTransfersForbidden", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_PoolControllerStorage *PoolControllerStorageSession) ThirdPartyTransfersForbidden(arg0 common.Address) (bool, error) {
	return _PoolControllerStorage.Contract.ThirdPartyTransfersForbidden(&_PoolControllerStorage.CallOpts, arg0)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_PoolControllerStorage *PoolControllerStorageCallerSession) ThirdPartyTransfersForbidden(arg0 common.Address) (bool, error) {
	return _PoolControllerStorage.Contract.ThirdPartyTransfersForbidden(&_PoolControllerStorage.CallOpts, arg0)
}

// AddPoolsToDepositWhitelist is a paid mutator transaction binding the contract method 0xa185a29a.
//
// Solidity: function addPoolsToDepositWhitelist(address[] poolsToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactor) AddPoolsToDepositWhitelist(opts *bind.TransactOpts, poolsToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "addPoolsToDepositWhitelist", poolsToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddPoolsToDepositWhitelist is a paid mutator transaction binding the contract method 0xa185a29a.
//
// Solidity: function addPoolsToDepositWhitelist(address[] poolsToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_PoolControllerStorage *PoolControllerStorageSession) AddPoolsToDepositWhitelist(poolsToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.AddPoolsToDepositWhitelist(&_PoolControllerStorage.TransactOpts, poolsToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddPoolsToDepositWhitelist is a paid mutator transaction binding the contract method 0xa185a29a.
//
// Solidity: function addPoolsToDepositWhitelist(address[] poolsToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) AddPoolsToDepositWhitelist(poolsToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.AddPoolsToDepositWhitelist(&_PoolControllerStorage.TransactOpts, poolsToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address pool, uint256 shares) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, token common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "addShares", staker, token, pool, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address pool, uint256 shares) returns()
func (_PoolControllerStorage *PoolControllerStorageSession) AddShares(staker common.Address, token common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.AddShares(&_PoolControllerStorage.TransactOpts, staker, token, pool, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address pool, uint256 shares) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) AddShares(staker common.Address, token common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.AddShares(&_PoolControllerStorage.TransactOpts, staker, token, pool, shares)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x6fcafe5a.
//
// Solidity: function depositIntoPool(address pool, address token, uint256 amount) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageTransactor) DepositIntoPool(opts *bind.TransactOpts, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "depositIntoPool", pool, token, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x6fcafe5a.
//
// Solidity: function depositIntoPool(address pool, address token, uint256 amount) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageSession) DepositIntoPool(pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.DepositIntoPool(&_PoolControllerStorage.TransactOpts, pool, token, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x6fcafe5a.
//
// Solidity: function depositIntoPool(address pool, address token, uint256 amount) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) DepositIntoPool(pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.DepositIntoPool(&_PoolControllerStorage.TransactOpts, pool, token, amount)
}

// DepositIntoPoolWithSignature is a paid mutator transaction binding the contract method 0x0e4ed53d.
//
// Solidity: function depositIntoPoolWithSignature(address pool, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageTransactor) DepositIntoPoolWithSignature(opts *bind.TransactOpts, pool common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "depositIntoPoolWithSignature", pool, token, amount, staker, expiry, signature)
}

// DepositIntoPoolWithSignature is a paid mutator transaction binding the contract method 0x0e4ed53d.
//
// Solidity: function depositIntoPoolWithSignature(address pool, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageSession) DepositIntoPoolWithSignature(pool common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.DepositIntoPoolWithSignature(&_PoolControllerStorage.TransactOpts, pool, token, amount, staker, expiry, signature)
}

// DepositIntoPoolWithSignature is a paid mutator transaction binding the contract method 0x0e4ed53d.
//
// Solidity: function depositIntoPoolWithSignature(address pool, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) DepositIntoPoolWithSignature(pool common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.DepositIntoPoolWithSignature(&_PoolControllerStorage.TransactOpts, pool, token, amount, staker, expiry, signature)
}

// DepositIntoPoolWithStaker is a paid mutator transaction binding the contract method 0x856abb29.
//
// Solidity: function depositIntoPoolWithStaker(address staker, address pool, address token, uint256 amount) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageTransactor) DepositIntoPoolWithStaker(opts *bind.TransactOpts, staker common.Address, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "depositIntoPoolWithStaker", staker, pool, token, amount)
}

// DepositIntoPoolWithStaker is a paid mutator transaction binding the contract method 0x856abb29.
//
// Solidity: function depositIntoPoolWithStaker(address staker, address pool, address token, uint256 amount) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageSession) DepositIntoPoolWithStaker(staker common.Address, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.DepositIntoPoolWithStaker(&_PoolControllerStorage.TransactOpts, staker, pool, token, amount)
}

// DepositIntoPoolWithStaker is a paid mutator transaction binding the contract method 0x856abb29.
//
// Solidity: function depositIntoPoolWithStaker(address staker, address pool, address token, uint256 amount) returns(uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) DepositIntoPoolWithStaker(staker common.Address, pool common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.DepositIntoPoolWithStaker(&_PoolControllerStorage.TransactOpts, staker, pool, token, amount)
}

// RemovePoolsFromDepositWhitelist is a paid mutator transaction binding the contract method 0x53721791.
//
// Solidity: function removePoolsFromDepositWhitelist(address[] poolsToRemoveFromWhitelist) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactor) RemovePoolsFromDepositWhitelist(opts *bind.TransactOpts, poolsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "removePoolsFromDepositWhitelist", poolsToRemoveFromWhitelist)
}

// RemovePoolsFromDepositWhitelist is a paid mutator transaction binding the contract method 0x53721791.
//
// Solidity: function removePoolsFromDepositWhitelist(address[] poolsToRemoveFromWhitelist) returns()
func (_PoolControllerStorage *PoolControllerStorageSession) RemovePoolsFromDepositWhitelist(poolsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.RemovePoolsFromDepositWhitelist(&_PoolControllerStorage.TransactOpts, poolsToRemoveFromWhitelist)
}

// RemovePoolsFromDepositWhitelist is a paid mutator transaction binding the contract method 0x53721791.
//
// Solidity: function removePoolsFromDepositWhitelist(address[] poolsToRemoveFromWhitelist) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) RemovePoolsFromDepositWhitelist(poolsToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.RemovePoolsFromDepositWhitelist(&_PoolControllerStorage.TransactOpts, poolsToRemoveFromWhitelist)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address pool, uint256 shares) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "removeShares", staker, pool, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address pool, uint256 shares) returns()
func (_PoolControllerStorage *PoolControllerStorageSession) RemoveShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.RemoveShares(&_PoolControllerStorage.TransactOpts, staker, pool, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address pool, uint256 shares) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) RemoveShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.RemoveShares(&_PoolControllerStorage.TransactOpts, staker, pool, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address pool, uint256 shares, address token) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, recipient common.Address, pool common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _PoolControllerStorage.contract.Transact(opts, "withdrawSharesAsTokens", recipient, pool, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address pool, uint256 shares, address token) returns()
func (_PoolControllerStorage *PoolControllerStorageSession) WithdrawSharesAsTokens(recipient common.Address, pool common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.WithdrawSharesAsTokens(&_PoolControllerStorage.TransactOpts, recipient, pool, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address pool, uint256 shares, address token) returns()
func (_PoolControllerStorage *PoolControllerStorageTransactorSession) WithdrawSharesAsTokens(recipient common.Address, pool common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _PoolControllerStorage.Contract.WithdrawSharesAsTokens(&_PoolControllerStorage.TransactOpts, recipient, pool, shares, token)
}

// PoolControllerStorageDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PoolControllerStorage contract.
type PoolControllerStorageDepositIterator struct {
	Event *PoolControllerStorageDeposit // Event containing the contract specifics and raw log

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
func (it *PoolControllerStorageDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolControllerStorageDeposit)
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
		it.Event = new(PoolControllerStorageDeposit)
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
func (it *PoolControllerStorageDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolControllerStorageDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolControllerStorageDeposit represents a Deposit event raised by the PoolControllerStorage contract.
type PoolControllerStorageDeposit struct {
	Staker common.Address
	Token  common.Address
	Pool   common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address pool, uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageFilterer) FilterDeposit(opts *bind.FilterOpts) (*PoolControllerStorageDepositIterator, error) {

	logs, sub, err := _PoolControllerStorage.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &PoolControllerStorageDepositIterator{contract: _PoolControllerStorage.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address pool, uint256 shares)
func (_PoolControllerStorage *PoolControllerStorageFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PoolControllerStorageDeposit) (event.Subscription, error) {

	logs, sub, err := _PoolControllerStorage.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolControllerStorageDeposit)
				if err := _PoolControllerStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_PoolControllerStorage *PoolControllerStorageFilterer) ParseDeposit(log types.Log) (*PoolControllerStorageDeposit, error) {
	event := new(PoolControllerStorageDeposit)
	if err := _PoolControllerStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolControllerStoragePoolAddedToDepositWhitelistIterator is returned from FilterPoolAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for PoolAddedToDepositWhitelist events raised by the PoolControllerStorage contract.
type PoolControllerStoragePoolAddedToDepositWhitelistIterator struct {
	Event *PoolControllerStoragePoolAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *PoolControllerStoragePoolAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolControllerStoragePoolAddedToDepositWhitelist)
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
		it.Event = new(PoolControllerStoragePoolAddedToDepositWhitelist)
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
func (it *PoolControllerStoragePoolAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolControllerStoragePoolAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolControllerStoragePoolAddedToDepositWhitelist represents a PoolAddedToDepositWhitelist event raised by the PoolControllerStorage contract.
type PoolControllerStoragePoolAddedToDepositWhitelist struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x39daccb429428846b0171894bd4200a7051a1c25c7b82ee455a39f89d72e634c.
//
// Solidity: event PoolAddedToDepositWhitelist(address pool)
func (_PoolControllerStorage *PoolControllerStorageFilterer) FilterPoolAddedToDepositWhitelist(opts *bind.FilterOpts) (*PoolControllerStoragePoolAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _PoolControllerStorage.contract.FilterLogs(opts, "PoolAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &PoolControllerStoragePoolAddedToDepositWhitelistIterator{contract: _PoolControllerStorage.contract, event: "PoolAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchPoolAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x39daccb429428846b0171894bd4200a7051a1c25c7b82ee455a39f89d72e634c.
//
// Solidity: event PoolAddedToDepositWhitelist(address pool)
func (_PoolControllerStorage *PoolControllerStorageFilterer) WatchPoolAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *PoolControllerStoragePoolAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _PoolControllerStorage.contract.WatchLogs(opts, "PoolAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolControllerStoragePoolAddedToDepositWhitelist)
				if err := _PoolControllerStorage.contract.UnpackLog(event, "PoolAddedToDepositWhitelist", log); err != nil {
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
func (_PoolControllerStorage *PoolControllerStorageFilterer) ParsePoolAddedToDepositWhitelist(log types.Log) (*PoolControllerStoragePoolAddedToDepositWhitelist, error) {
	event := new(PoolControllerStoragePoolAddedToDepositWhitelist)
	if err := _PoolControllerStorage.contract.UnpackLog(event, "PoolAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolControllerStoragePoolRemovedFromDepositWhitelistIterator is returned from FilterPoolRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for PoolRemovedFromDepositWhitelist events raised by the PoolControllerStorage contract.
type PoolControllerStoragePoolRemovedFromDepositWhitelistIterator struct {
	Event *PoolControllerStoragePoolRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *PoolControllerStoragePoolRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolControllerStoragePoolRemovedFromDepositWhitelist)
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
		it.Event = new(PoolControllerStoragePoolRemovedFromDepositWhitelist)
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
func (it *PoolControllerStoragePoolRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolControllerStoragePoolRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolControllerStoragePoolRemovedFromDepositWhitelist represents a PoolRemovedFromDepositWhitelist event raised by the PoolControllerStorage contract.
type PoolControllerStoragePoolRemovedFromDepositWhitelist struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x8ac71de5f0eed990ceee2d31cbfe4280db9f27a8f0bf5567c507309de33ed723.
//
// Solidity: event PoolRemovedFromDepositWhitelist(address pool)
func (_PoolControllerStorage *PoolControllerStorageFilterer) FilterPoolRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*PoolControllerStoragePoolRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _PoolControllerStorage.contract.FilterLogs(opts, "PoolRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &PoolControllerStoragePoolRemovedFromDepositWhitelistIterator{contract: _PoolControllerStorage.contract, event: "PoolRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchPoolRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x8ac71de5f0eed990ceee2d31cbfe4280db9f27a8f0bf5567c507309de33ed723.
//
// Solidity: event PoolRemovedFromDepositWhitelist(address pool)
func (_PoolControllerStorage *PoolControllerStorageFilterer) WatchPoolRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *PoolControllerStoragePoolRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _PoolControllerStorage.contract.WatchLogs(opts, "PoolRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolControllerStoragePoolRemovedFromDepositWhitelist)
				if err := _PoolControllerStorage.contract.UnpackLog(event, "PoolRemovedFromDepositWhitelist", log); err != nil {
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
func (_PoolControllerStorage *PoolControllerStorageFilterer) ParsePoolRemovedFromDepositWhitelist(log types.Log) (*PoolControllerStoragePoolRemovedFromDepositWhitelist, error) {
	event := new(PoolControllerStoragePoolRemovedFromDepositWhitelist)
	if err := _PoolControllerStorage.contract.UnpackLog(event, "PoolRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolControllerStoragePoolWhitelisterChangedIterator is returned from FilterPoolWhitelisterChanged and is used to iterate over the raw logs and unpacked data for PoolWhitelisterChanged events raised by the PoolControllerStorage contract.
type PoolControllerStoragePoolWhitelisterChangedIterator struct {
	Event *PoolControllerStoragePoolWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *PoolControllerStoragePoolWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolControllerStoragePoolWhitelisterChanged)
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
		it.Event = new(PoolControllerStoragePoolWhitelisterChanged)
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
func (it *PoolControllerStoragePoolWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolControllerStoragePoolWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolControllerStoragePoolWhitelisterChanged represents a PoolWhitelisterChanged event raised by the PoolControllerStorage contract.
type PoolControllerStoragePoolWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPoolWhitelisterChanged is a free log retrieval operation binding the contract event 0xb10a4c4f1bee1e88074ca37b6c3ce3990495d771d4f7735a5fed00a773af424f.
//
// Solidity: event PoolWhitelisterChanged(address previousAddress, address newAddress)
func (_PoolControllerStorage *PoolControllerStorageFilterer) FilterPoolWhitelisterChanged(opts *bind.FilterOpts) (*PoolControllerStoragePoolWhitelisterChangedIterator, error) {

	logs, sub, err := _PoolControllerStorage.contract.FilterLogs(opts, "PoolWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &PoolControllerStoragePoolWhitelisterChangedIterator{contract: _PoolControllerStorage.contract, event: "PoolWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchPoolWhitelisterChanged is a free log subscription operation binding the contract event 0xb10a4c4f1bee1e88074ca37b6c3ce3990495d771d4f7735a5fed00a773af424f.
//
// Solidity: event PoolWhitelisterChanged(address previousAddress, address newAddress)
func (_PoolControllerStorage *PoolControllerStorageFilterer) WatchPoolWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *PoolControllerStoragePoolWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _PoolControllerStorage.contract.WatchLogs(opts, "PoolWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolControllerStoragePoolWhitelisterChanged)
				if err := _PoolControllerStorage.contract.UnpackLog(event, "PoolWhitelisterChanged", log); err != nil {
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
func (_PoolControllerStorage *PoolControllerStorageFilterer) ParsePoolWhitelisterChanged(log types.Log) (*PoolControllerStoragePoolWhitelisterChanged, error) {
	event := new(PoolControllerStoragePoolWhitelisterChanged)
	if err := _PoolControllerStorage.contract.UnpackLog(event, "PoolWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator is returned from FilterUpdatedThirdPartyTransfersForbidden and is used to iterate over the raw logs and unpacked data for UpdatedThirdPartyTransfersForbidden events raised by the PoolControllerStorage contract.
type PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator struct {
	Event *PoolControllerStorageUpdatedThirdPartyTransfersForbidden // Event containing the contract specifics and raw log

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
func (it *PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolControllerStorageUpdatedThirdPartyTransfersForbidden)
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
		it.Event = new(PoolControllerStorageUpdatedThirdPartyTransfersForbidden)
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
func (it *PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolControllerStorageUpdatedThirdPartyTransfersForbidden represents a UpdatedThirdPartyTransfersForbidden event raised by the PoolControllerStorage contract.
type PoolControllerStorageUpdatedThirdPartyTransfersForbidden struct {
	Pool  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedThirdPartyTransfersForbidden is a free log retrieval operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address pool, bool value)
func (_PoolControllerStorage *PoolControllerStorageFilterer) FilterUpdatedThirdPartyTransfersForbidden(opts *bind.FilterOpts) (*PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator, error) {

	logs, sub, err := _PoolControllerStorage.contract.FilterLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return &PoolControllerStorageUpdatedThirdPartyTransfersForbiddenIterator{contract: _PoolControllerStorage.contract, event: "UpdatedThirdPartyTransfersForbidden", logs: logs, sub: sub}, nil
}

// WatchUpdatedThirdPartyTransfersForbidden is a free log subscription operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address pool, bool value)
func (_PoolControllerStorage *PoolControllerStorageFilterer) WatchUpdatedThirdPartyTransfersForbidden(opts *bind.WatchOpts, sink chan<- *PoolControllerStorageUpdatedThirdPartyTransfersForbidden) (event.Subscription, error) {

	logs, sub, err := _PoolControllerStorage.contract.WatchLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolControllerStorageUpdatedThirdPartyTransfersForbidden)
				if err := _PoolControllerStorage.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
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
func (_PoolControllerStorage *PoolControllerStorageFilterer) ParseUpdatedThirdPartyTransfersForbidden(log types.Log) (*PoolControllerStorageUpdatedThirdPartyTransfersForbidden, error) {
	event := new(PoolControllerStorageUpdatedThirdPartyTransfersForbidden)
	if err := _PoolControllerStorage.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
