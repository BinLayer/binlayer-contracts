// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PauserRegistry

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

// PauserRegistryMetaData contains all meta data concerning the PauserRegistry contract.
var PauserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pausers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_unpauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPauser\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setIsPauser\",\"inputs\":[{\"name\":\"newPauser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnpauser\",\"inputs\":[{\"name\":\"newUnpauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PauserStatusChanged\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"canPause\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnpauserChanged\",\"inputs\":[{\"name\":\"previousUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newUnpauser\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506040516108fa3803806108fa83398101604081905261002f9161020d565b60005b825181101561007757610065838281518110610050576100506102e3565b6020026020010151600161008860201b60201c565b8061006f816102f9565b915050610032565b5061008181610132565b505061036e565b6040805180820190915260018152603360f81b60208201526001600160a01b0383166100d05760405162461bcd60e51b81526004016100c79190610320565b60405180910390fd5b506001600160a01b03821660008181526020818152604091829020805460ff19168515159081179091558251938452908301527f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152910160405180910390a15050565b6040805180820190915260018152603360f81b60208201526001600160a01b0382166101715760405162461bcd60e51b81526004016100c79190610320565b50600154604080516001600160a01b03928316815291831660208301527f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892910160405180910390a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b038116811461020857600080fd5b919050565b6000806040838503121561022057600080fd5b82516001600160401b038082111561023757600080fd5b818501915085601f83011261024b57600080fd5b815160208282111561025f5761025f6101db565b8160051b604051601f19603f83011681018181108682111715610284576102846101db565b6040529283528183019350848101820192898411156102a257600080fd5b948201945b838610156102c7576102b8866101f1565b855294820194938201936102a7565b96506102d690508782016101f1565b9450505050509250929050565b634e487b7160e01b600052603260045260246000fd5b60006001820161031957634e487b7160e01b600052601160045260246000fd5b5060010190565b600060208083528351808285015260005b8181101561034d57858101830151858201604001528201610331565b506000604082860101526040601f19601f8301168501019250505092915050565b61057d8061037d6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806346fbf68e146100515780638568520614610089578063ce5484281461009e578063eab66d7a146100b1575b600080fd5b61007461005f36600461047d565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009c61009736600461049f565b6100f6565b005b61009c6100ac36600461047d565b610194565b6001546100d19073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610080565b60018054604080518082019091529182527f3200000000000000000000000000000000000000000000000000000000000000602083015273ffffffffffffffffffffffffffffffffffffffff163314610185576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c91906104db565b60405180910390fd5b506101908282610227565b5050565b60018054604080518082019091529182527f3200000000000000000000000000000000000000000000000000000000000000602083015273ffffffffffffffffffffffffffffffffffffffff16331461021a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c91906104db565b5061022481610336565b50565b60408051808201909152600181527f3300000000000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff83166102a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c91906104db565b5073ffffffffffffffffffffffffffffffffffffffff82166000818152602081815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168515159081179091558251938452908301527f65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152910160405180910390a15050565b60408051808201909152600181527f3300000000000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff82166103b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c91906104db565b506001546040805173ffffffffffffffffffffffffffffffffffffffff928316815291831660208301527f06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892910160405180910390a1600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b803573ffffffffffffffffffffffffffffffffffffffff8116811461047857600080fd5b919050565b60006020828403121561048f57600080fd5b61049882610454565b9392505050565b600080604083850312156104b257600080fd5b6104bb83610454565b9150602083013580151581146104d057600080fd5b809150509250929050565b600060208083528351808285015260005b81811015610508578581018301518582016040015282016104ec565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea264697066735822122046f987f1046733b84803122fe6b54e22cfe5e60f7620828bd392a912b1a66e9a64736f6c63430008140033",
}

// PauserRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PauserRegistryMetaData.ABI instead.
var PauserRegistryABI = PauserRegistryMetaData.ABI

// PauserRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PauserRegistryMetaData.Bin instead.
var PauserRegistryBin = PauserRegistryMetaData.Bin

// DeployPauserRegistry deploys a new Ethereum contract, binding an instance of PauserRegistry to it.
func DeployPauserRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _pausers []common.Address, _unpauser common.Address) (common.Address, *types.Transaction, *PauserRegistry, error) {
	parsed, err := PauserRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PauserRegistryBin), backend, _pausers, _unpauser)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PauserRegistry{PauserRegistryCaller: PauserRegistryCaller{contract: contract}, PauserRegistryTransactor: PauserRegistryTransactor{contract: contract}, PauserRegistryFilterer: PauserRegistryFilterer{contract: contract}}, nil
}

// PauserRegistry is an auto generated Go binding around an Ethereum contract.
type PauserRegistry struct {
	PauserRegistryCaller     // Read-only binding to the contract
	PauserRegistryTransactor // Write-only binding to the contract
	PauserRegistryFilterer   // Log filterer for contract events
}

// PauserRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserRegistrySession struct {
	Contract     *PauserRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserRegistryCallerSession struct {
	Contract *PauserRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PauserRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserRegistryTransactorSession struct {
	Contract     *PauserRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PauserRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRegistryRaw struct {
	Contract *PauserRegistry // Generic contract binding to access the raw methods on
}

// PauserRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserRegistryCallerRaw struct {
	Contract *PauserRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PauserRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserRegistryTransactorRaw struct {
	Contract *PauserRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauserRegistry creates a new instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistry(address common.Address, backend bind.ContractBackend) (*PauserRegistry, error) {
	contract, err := bindPauserRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauserRegistry{PauserRegistryCaller: PauserRegistryCaller{contract: contract}, PauserRegistryTransactor: PauserRegistryTransactor{contract: contract}, PauserRegistryFilterer: PauserRegistryFilterer{contract: contract}}, nil
}

// NewPauserRegistryCaller creates a new read-only instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryCaller(address common.Address, caller bind.ContractCaller) (*PauserRegistryCaller, error) {
	contract, err := bindPauserRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryCaller{contract: contract}, nil
}

// NewPauserRegistryTransactor creates a new write-only instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserRegistryTransactor, error) {
	contract, err := bindPauserRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryTransactor{contract: contract}, nil
}

// NewPauserRegistryFilterer creates a new log filterer instance of PauserRegistry, bound to a specific deployed contract.
func NewPauserRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserRegistryFilterer, error) {
	contract, err := bindPauserRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserRegistryFilterer{contract: contract}, nil
}

// bindPauserRegistry binds a generic wrapper to an already deployed contract.
func bindPauserRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PauserRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRegistry *PauserRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauserRegistry.Contract.PauserRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRegistry *PauserRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRegistry.Contract.PauserRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRegistry *PauserRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRegistry.Contract.PauserRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRegistry *PauserRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauserRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRegistry *PauserRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRegistry *PauserRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistryCaller) IsPauser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PauserRegistry.contract.Call(opts, &out, "isPauser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistrySession) IsPauser(arg0 common.Address) (bool, error) {
	return _PauserRegistry.Contract.IsPauser(&_PauserRegistry.CallOpts, arg0)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address ) view returns(bool)
func (_PauserRegistry *PauserRegistryCallerSession) IsPauser(arg0 common.Address) (bool, error) {
	return _PauserRegistry.Contract.IsPauser(&_PauserRegistry.CallOpts, arg0)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistryCaller) Unpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PauserRegistry.contract.Call(opts, &out, "unpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistrySession) Unpauser() (common.Address, error) {
	return _PauserRegistry.Contract.Unpauser(&_PauserRegistry.CallOpts)
}

// Unpauser is a free data retrieval call binding the contract method 0xeab66d7a.
//
// Solidity: function unpauser() view returns(address)
func (_PauserRegistry *PauserRegistryCallerSession) Unpauser() (common.Address, error) {
	return _PauserRegistry.Contract.Unpauser(&_PauserRegistry.CallOpts)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistryTransactor) SetIsPauser(opts *bind.TransactOpts, newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.contract.Transact(opts, "setIsPauser", newPauser, canPause)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistrySession) SetIsPauser(newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetIsPauser(&_PauserRegistry.TransactOpts, newPauser, canPause)
}

// SetIsPauser is a paid mutator transaction binding the contract method 0x85685206.
//
// Solidity: function setIsPauser(address newPauser, bool canPause) returns()
func (_PauserRegistry *PauserRegistryTransactorSession) SetIsPauser(newPauser common.Address, canPause bool) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetIsPauser(&_PauserRegistry.TransactOpts, newPauser, canPause)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistryTransactor) SetUnpauser(opts *bind.TransactOpts, newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.contract.Transact(opts, "setUnpauser", newUnpauser)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistrySession) SetUnpauser(newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetUnpauser(&_PauserRegistry.TransactOpts, newUnpauser)
}

// SetUnpauser is a paid mutator transaction binding the contract method 0xce548428.
//
// Solidity: function setUnpauser(address newUnpauser) returns()
func (_PauserRegistry *PauserRegistryTransactorSession) SetUnpauser(newUnpauser common.Address) (*types.Transaction, error) {
	return _PauserRegistry.Contract.SetUnpauser(&_PauserRegistry.TransactOpts, newUnpauser)
}

// PauserRegistryPauserStatusChangedIterator is returned from FilterPauserStatusChanged and is used to iterate over the raw logs and unpacked data for PauserStatusChanged events raised by the PauserRegistry contract.
type PauserRegistryPauserStatusChangedIterator struct {
	Event *PauserRegistryPauserStatusChanged // Event containing the contract specifics and raw log

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
func (it *PauserRegistryPauserStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRegistryPauserStatusChanged)
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
		it.Event = new(PauserRegistryPauserStatusChanged)
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
func (it *PauserRegistryPauserStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRegistryPauserStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRegistryPauserStatusChanged represents a PauserStatusChanged event raised by the PauserRegistry contract.
type PauserRegistryPauserStatusChanged struct {
	Pauser   common.Address
	CanPause bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauserStatusChanged is a free log retrieval operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) FilterPauserStatusChanged(opts *bind.FilterOpts) (*PauserRegistryPauserStatusChangedIterator, error) {

	logs, sub, err := _PauserRegistry.contract.FilterLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return &PauserRegistryPauserStatusChangedIterator{contract: _PauserRegistry.contract, event: "PauserStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPauserStatusChanged is a free log subscription operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) WatchPauserStatusChanged(opts *bind.WatchOpts, sink chan<- *PauserRegistryPauserStatusChanged) (event.Subscription, error) {

	logs, sub, err := _PauserRegistry.contract.WatchLogs(opts, "PauserStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRegistryPauserStatusChanged)
				if err := _PauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
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

// ParsePauserStatusChanged is a log parse operation binding the contract event 0x65d3a1fd4c13f05cba164f80d03ce90fb4b5e21946bfc3ab7dbd434c2d0b9152.
//
// Solidity: event PauserStatusChanged(address pauser, bool canPause)
func (_PauserRegistry *PauserRegistryFilterer) ParsePauserStatusChanged(log types.Log) (*PauserRegistryPauserStatusChanged, error) {
	event := new(PauserRegistryPauserStatusChanged)
	if err := _PauserRegistry.contract.UnpackLog(event, "PauserStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserRegistryUnpauserChangedIterator is returned from FilterUnpauserChanged and is used to iterate over the raw logs and unpacked data for UnpauserChanged events raised by the PauserRegistry contract.
type PauserRegistryUnpauserChangedIterator struct {
	Event *PauserRegistryUnpauserChanged // Event containing the contract specifics and raw log

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
func (it *PauserRegistryUnpauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRegistryUnpauserChanged)
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
		it.Event = new(PauserRegistryUnpauserChanged)
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
func (it *PauserRegistryUnpauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRegistryUnpauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRegistryUnpauserChanged represents a UnpauserChanged event raised by the PauserRegistry contract.
type PauserRegistryUnpauserChanged struct {
	PreviousUnpauser common.Address
	NewUnpauser      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnpauserChanged is a free log retrieval operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) FilterUnpauserChanged(opts *bind.FilterOpts) (*PauserRegistryUnpauserChangedIterator, error) {

	logs, sub, err := _PauserRegistry.contract.FilterLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return &PauserRegistryUnpauserChangedIterator{contract: _PauserRegistry.contract, event: "UnpauserChanged", logs: logs, sub: sub}, nil
}

// WatchUnpauserChanged is a free log subscription operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) WatchUnpauserChanged(opts *bind.WatchOpts, sink chan<- *PauserRegistryUnpauserChanged) (event.Subscription, error) {

	logs, sub, err := _PauserRegistry.contract.WatchLogs(opts, "UnpauserChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRegistryUnpauserChanged)
				if err := _PauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
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

// ParseUnpauserChanged is a log parse operation binding the contract event 0x06b4167a2528887a1e97a366eefe8549bfbf1ea3e6ac81cb2564a934d20e8892.
//
// Solidity: event UnpauserChanged(address previousUnpauser, address newUnpauser)
func (_PauserRegistry *PauserRegistryFilterer) ParseUnpauserChanged(log types.Log) (*PauserRegistryUnpauserChanged, error) {
	event := new(PauserRegistryUnpauserChanged)
	if err := _PauserRegistry.contract.UnpackLog(event, "UnpauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
