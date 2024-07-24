// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC677

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

// IERC677MetaData contains all meta data concerning the IERC677 contract.
var IERC677MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"transferAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// IERC677ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC677MetaData.ABI instead.
var IERC677ABI = IERC677MetaData.ABI

// IERC677 is an auto generated Go binding around an Ethereum contract.
type IERC677 struct {
	IERC677Caller     // Read-only binding to the contract
	IERC677Transactor // Write-only binding to the contract
	IERC677Filterer   // Log filterer for contract events
}

// IERC677Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC677Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC677Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC677Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC677Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC677Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC677Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC677Session struct {
	Contract     *IERC677          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC677CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC677CallerSession struct {
	Contract *IERC677Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC677TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC677TransactorSession struct {
	Contract     *IERC677Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC677Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC677Raw struct {
	Contract *IERC677 // Generic contract binding to access the raw methods on
}

// IERC677CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC677CallerRaw struct {
	Contract *IERC677Caller // Generic read-only contract binding to access the raw methods on
}

// IERC677TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC677TransactorRaw struct {
	Contract *IERC677Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC677 creates a new instance of IERC677, bound to a specific deployed contract.
func NewIERC677(address common.Address, backend bind.ContractBackend) (*IERC677, error) {
	contract, err := bindIERC677(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC677{IERC677Caller: IERC677Caller{contract: contract}, IERC677Transactor: IERC677Transactor{contract: contract}, IERC677Filterer: IERC677Filterer{contract: contract}}, nil
}

// NewIERC677Caller creates a new read-only instance of IERC677, bound to a specific deployed contract.
func NewIERC677Caller(address common.Address, caller bind.ContractCaller) (*IERC677Caller, error) {
	contract, err := bindIERC677(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC677Caller{contract: contract}, nil
}

// NewIERC677Transactor creates a new write-only instance of IERC677, bound to a specific deployed contract.
func NewIERC677Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC677Transactor, error) {
	contract, err := bindIERC677(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC677Transactor{contract: contract}, nil
}

// NewIERC677Filterer creates a new log filterer instance of IERC677, bound to a specific deployed contract.
func NewIERC677Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC677Filterer, error) {
	contract, err := bindIERC677(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC677Filterer{contract: contract}, nil
}

// bindIERC677 binds a generic wrapper to an already deployed contract.
func bindIERC677(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC677MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC677 *IERC677Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC677.Contract.IERC677Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC677 *IERC677Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC677.Contract.IERC677Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC677 *IERC677Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC677.Contract.IERC677Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC677 *IERC677CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC677.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC677 *IERC677TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC677.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC677 *IERC677TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC677.Contract.contract.Transact(opts, method, params...)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_IERC677 *IERC677Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_IERC677 *IERC677Session) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677.Contract.TransferAndCall(&_IERC677.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_IERC677 *IERC677TransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677.Contract.TransferAndCall(&_IERC677.TransactOpts, to, value, data)
}

// IERC677TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC677 contract.
type IERC677TransferIterator struct {
	Event *IERC677Transfer // Event containing the contract specifics and raw log

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
func (it *IERC677TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC677Transfer)
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
		it.Event = new(IERC677Transfer)
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
func (it *IERC677TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC677TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC677Transfer represents a Transfer event raised by the IERC677 contract.
type IERC677Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_IERC677 *IERC677Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC677TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC677.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC677TransferIterator{contract: _IERC677.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_IERC677 *IERC677Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC677Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC677.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC677Transfer)
				if err := _IERC677.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_IERC677 *IERC677Filterer) ParseTransfer(log types.Log) (*IERC677Transfer, error) {
	event := new(IERC677Transfer)
	if err := _IERC677.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
