// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ListaGateway

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

// ListaGatewayMetaData contains all meta data concerning the ListaGateway contract.
var ListaGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slisBNB\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_listaStakeManager\",\"type\":\"address\",\"internalType\":\"contractIListaStakeManager\"},{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"_poolController\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositNativeToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"emergencyNativeTransfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyTokenTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getListaContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"_listaStakeManager\",\"type\":\"address\",\"internalType\":\"contractIListaStakeManager\"},{\"name\":\"_slisBNB\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620010d5380380620010d583398101604081905262000035916200015a565b6200004033620000f1565b6001600160a01b0380861660805283811660a05282811660c052811660e0526200006a84620000f1565b60e05160405163095ea7b360e01b81526001600160a01b03918216600482015260001960248201529086169063095ea7b3906044016020604051808303816000875af1158015620000bf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000e59190620001da565b50505050505062000205565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146200015757600080fd5b50565b600080600080600060a086880312156200017357600080fd5b8551620001808162000141565b6020870151909550620001938162000141565b6040870151909450620001a68162000141565b6060870151909350620001b98162000141565b6080870151909250620001cc8162000141565b809150509295509295909350565b600060208284031215620001ed57600080fd5b81518015158114620001fe57600080fd5b9392505050565b60805160a05160c05160e051610e6d6200026860003960006105bf0152600081816102e10152610565015260008181609c0152818161026201526103e40152600081816102870152818161036001528181610492015261058d0152610e6d6000f3fe60806040526004361061007f5760003560e01c8063a3d5b2551161004e578063a3d5b2551461021a578063bafe6e871461023a578063f2fde38b146102b2578063f586c6d9146102d25761012a565b80631653a0c71461018c578063715018a6146101ac57806379433d8b146101c15780638da5cb5b146101c95761012a565b3661012a573373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610128576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f52656365697665206e6f7420616c6c6f7765640000000000000000000000000060448201526064015b60405180910390fd5b005b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f46616c6c6261636b206e6f7420616c6c6f776564000000000000000000000000604482015260640161011f565b34801561019857600080fd5b506101286101a7366004610ca3565b610305565b3480156101b857600080fd5b5061012861031b565b61012861032f565b3480156101d557600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561022657600080fd5b50610128610235366004610ccd565b610634565b34801561024657600080fd5b506040805173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f000000000000000000000000000000000000000000000000000000000000000016602082015201610211565b3480156102be57600080fd5b506101286102cd366004610d09565b610662565b3480156102de57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101f0565b61030d610719565b610317828261079a565b5050565b610323610719565b61032d600061087e565b565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156103bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e09190610d2b565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561044a57600080fd5b505af115801561045e573d6000803e3d6000fd5b50506040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600093507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1692506370a082319150602401602060405180830381865afa1580156104f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105149190610d2b565b905060006105228383610d44565b6040517f856abb2900000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f000000000000000000000000000000000000000000000000000000000000000081166044830152606482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063856abb29906084016020604051808303816000875af115801561060a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062e9190610d2b565b50505050565b61063c610719565b61065d73ffffffffffffffffffffffffffffffffffffffff841683836108f3565b505050565b61066a610719565b73ffffffffffffffffffffffffffffffffffffffff811661070d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161011f565b6107168161087e565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461032d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161011f565b6040805160008082526020820190925273ffffffffffffffffffffffffffffffffffffffff84169083906040516107d19190610da8565b60006040518083038185875af1925050503d806000811461080e576040519150601f19603f3d011682016040523d82523d6000602084013e610813565b606091505b505090508061065d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e41544956455f5452414e534645525f4641494c454400000000000000000000604482015260640161011f565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040805173ffffffffffffffffffffffffffffffffffffffff848116602483015260448083018590528351808403909101815260649092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649084015261065d928692916000916109be918516908490610a6b565b90508051600014806109df5750808060200190518101906109df9190610dc4565b61065d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161011f565b6060610a7a8484600085610a82565b949350505050565b606082471015610b14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161011f565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610b3d9190610da8565b60006040518083038185875af1925050503d8060008114610b7a576040519150601f19603f3d011682016040523d82523d6000602084013e610b7f565b606091505b5091509150610b9087838387610b9b565b979650505050505050565b60608315610c31578251600003610c2a5773ffffffffffffffffffffffffffffffffffffffff85163b610c2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161011f565b5081610a7a565b610a7a8383815115610c465781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011f9190610de6565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c9e57600080fd5b919050565b60008060408385031215610cb657600080fd5b610cbf83610c7a565b946020939093013593505050565b600080600060608486031215610ce257600080fd5b610ceb84610c7a565b9250610cf960208501610c7a565b9150604084013590509250925092565b600060208284031215610d1b57600080fd5b610d2482610c7a565b9392505050565b600060208284031215610d3d57600080fd5b5051919050565b81810381811115610d7e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60005b83811015610d9f578181015183820152602001610d87565b50506000910152565b60008251610dba818460208701610d84565b9190910192915050565b600060208284031215610dd657600080fd5b81518015158114610d2457600080fd5b6020815260008251806020840152610e05816040850160208701610d84565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea26469706673582212206cbfe77b5f1eba6142b8a54ef69e49167d25ce6561336939f96083ffaf22dc6464736f6c63430008140033",
}

// ListaGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use ListaGatewayMetaData.ABI instead.
var ListaGatewayABI = ListaGatewayMetaData.ABI

// ListaGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ListaGatewayMetaData.Bin instead.
var ListaGatewayBin = ListaGatewayMetaData.Bin

// DeployListaGateway deploys a new Ethereum contract, binding an instance of ListaGateway to it.
func DeployListaGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _slisBNB common.Address, _owner common.Address, _listaStakeManager common.Address, _pool common.Address, _poolController common.Address) (common.Address, *types.Transaction, *ListaGateway, error) {
	parsed, err := ListaGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ListaGatewayBin), backend, _slisBNB, _owner, _listaStakeManager, _pool, _poolController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ListaGateway{ListaGatewayCaller: ListaGatewayCaller{contract: contract}, ListaGatewayTransactor: ListaGatewayTransactor{contract: contract}, ListaGatewayFilterer: ListaGatewayFilterer{contract: contract}}, nil
}

// ListaGateway is an auto generated Go binding around an Ethereum contract.
type ListaGateway struct {
	ListaGatewayCaller     // Read-only binding to the contract
	ListaGatewayTransactor // Write-only binding to the contract
	ListaGatewayFilterer   // Log filterer for contract events
}

// ListaGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type ListaGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListaGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ListaGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListaGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ListaGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListaGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ListaGatewaySession struct {
	Contract     *ListaGateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ListaGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ListaGatewayCallerSession struct {
	Contract *ListaGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ListaGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ListaGatewayTransactorSession struct {
	Contract     *ListaGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ListaGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type ListaGatewayRaw struct {
	Contract *ListaGateway // Generic contract binding to access the raw methods on
}

// ListaGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ListaGatewayCallerRaw struct {
	Contract *ListaGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// ListaGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ListaGatewayTransactorRaw struct {
	Contract *ListaGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewListaGateway creates a new instance of ListaGateway, bound to a specific deployed contract.
func NewListaGateway(address common.Address, backend bind.ContractBackend) (*ListaGateway, error) {
	contract, err := bindListaGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ListaGateway{ListaGatewayCaller: ListaGatewayCaller{contract: contract}, ListaGatewayTransactor: ListaGatewayTransactor{contract: contract}, ListaGatewayFilterer: ListaGatewayFilterer{contract: contract}}, nil
}

// NewListaGatewayCaller creates a new read-only instance of ListaGateway, bound to a specific deployed contract.
func NewListaGatewayCaller(address common.Address, caller bind.ContractCaller) (*ListaGatewayCaller, error) {
	contract, err := bindListaGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ListaGatewayCaller{contract: contract}, nil
}

// NewListaGatewayTransactor creates a new write-only instance of ListaGateway, bound to a specific deployed contract.
func NewListaGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*ListaGatewayTransactor, error) {
	contract, err := bindListaGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ListaGatewayTransactor{contract: contract}, nil
}

// NewListaGatewayFilterer creates a new log filterer instance of ListaGateway, bound to a specific deployed contract.
func NewListaGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*ListaGatewayFilterer, error) {
	contract, err := bindListaGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ListaGatewayFilterer{contract: contract}, nil
}

// bindListaGateway binds a generic wrapper to an already deployed contract.
func bindListaGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ListaGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ListaGateway *ListaGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ListaGateway.Contract.ListaGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ListaGateway *ListaGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListaGateway.Contract.ListaGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ListaGateway *ListaGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ListaGateway.Contract.ListaGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ListaGateway *ListaGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ListaGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ListaGateway *ListaGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListaGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ListaGateway *ListaGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ListaGateway.Contract.contract.Transact(opts, method, params...)
}

// GetListaContracts is a free data retrieval call binding the contract method 0xbafe6e87.
//
// Solidity: function getListaContracts() view returns(address _listaStakeManager, address _slisBNB)
func (_ListaGateway *ListaGatewayCaller) GetListaContracts(opts *bind.CallOpts) (struct {
	ListaStakeManager common.Address
	SlisBNB           common.Address
}, error) {
	var out []interface{}
	err := _ListaGateway.contract.Call(opts, &out, "getListaContracts")

	outstruct := new(struct {
		ListaStakeManager common.Address
		SlisBNB           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListaStakeManager = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SlisBNB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetListaContracts is a free data retrieval call binding the contract method 0xbafe6e87.
//
// Solidity: function getListaContracts() view returns(address _listaStakeManager, address _slisBNB)
func (_ListaGateway *ListaGatewaySession) GetListaContracts() (struct {
	ListaStakeManager common.Address
	SlisBNB           common.Address
}, error) {
	return _ListaGateway.Contract.GetListaContracts(&_ListaGateway.CallOpts)
}

// GetListaContracts is a free data retrieval call binding the contract method 0xbafe6e87.
//
// Solidity: function getListaContracts() view returns(address _listaStakeManager, address _slisBNB)
func (_ListaGateway *ListaGatewayCallerSession) GetListaContracts() (struct {
	ListaStakeManager common.Address
	SlisBNB           common.Address
}, error) {
	return _ListaGateway.Contract.GetListaContracts(&_ListaGateway.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_ListaGateway *ListaGatewayCaller) GetPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ListaGateway.contract.Call(opts, &out, "getPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_ListaGateway *ListaGatewaySession) GetPoolAddress() (common.Address, error) {
	return _ListaGateway.Contract.GetPoolAddress(&_ListaGateway.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_ListaGateway *ListaGatewayCallerSession) GetPoolAddress() (common.Address, error) {
	return _ListaGateway.Contract.GetPoolAddress(&_ListaGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ListaGateway *ListaGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ListaGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ListaGateway *ListaGatewaySession) Owner() (common.Address, error) {
	return _ListaGateway.Contract.Owner(&_ListaGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ListaGateway *ListaGatewayCallerSession) Owner() (common.Address, error) {
	return _ListaGateway.Contract.Owner(&_ListaGateway.CallOpts)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_ListaGateway *ListaGatewayTransactor) DepositNativeToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListaGateway.contract.Transact(opts, "depositNativeToken")
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_ListaGateway *ListaGatewaySession) DepositNativeToken() (*types.Transaction, error) {
	return _ListaGateway.Contract.DepositNativeToken(&_ListaGateway.TransactOpts)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x79433d8b.
//
// Solidity: function depositNativeToken() payable returns()
func (_ListaGateway *ListaGatewayTransactorSession) DepositNativeToken() (*types.Transaction, error) {
	return _ListaGateway.Contract.DepositNativeToken(&_ListaGateway.TransactOpts)
}

// EmergencyNativeTransfer is a paid mutator transaction binding the contract method 0x1653a0c7.
//
// Solidity: function emergencyNativeTransfer(address to, uint256 amount) returns()
func (_ListaGateway *ListaGatewayTransactor) EmergencyNativeTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ListaGateway.contract.Transact(opts, "emergencyNativeTransfer", to, amount)
}

// EmergencyNativeTransfer is a paid mutator transaction binding the contract method 0x1653a0c7.
//
// Solidity: function emergencyNativeTransfer(address to, uint256 amount) returns()
func (_ListaGateway *ListaGatewaySession) EmergencyNativeTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ListaGateway.Contract.EmergencyNativeTransfer(&_ListaGateway.TransactOpts, to, amount)
}

// EmergencyNativeTransfer is a paid mutator transaction binding the contract method 0x1653a0c7.
//
// Solidity: function emergencyNativeTransfer(address to, uint256 amount) returns()
func (_ListaGateway *ListaGatewayTransactorSession) EmergencyNativeTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ListaGateway.Contract.EmergencyNativeTransfer(&_ListaGateway.TransactOpts, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_ListaGateway *ListaGatewayTransactor) EmergencyTokenTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ListaGateway.contract.Transact(opts, "emergencyTokenTransfer", token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_ListaGateway *ListaGatewaySession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ListaGateway.Contract.EmergencyTokenTransfer(&_ListaGateway.TransactOpts, token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_ListaGateway *ListaGatewayTransactorSession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ListaGateway.Contract.EmergencyTokenTransfer(&_ListaGateway.TransactOpts, token, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ListaGateway *ListaGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListaGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ListaGateway *ListaGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _ListaGateway.Contract.RenounceOwnership(&_ListaGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ListaGateway *ListaGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ListaGateway.Contract.RenounceOwnership(&_ListaGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ListaGateway *ListaGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ListaGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ListaGateway *ListaGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ListaGateway.Contract.TransferOwnership(&_ListaGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ListaGateway *ListaGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ListaGateway.Contract.TransferOwnership(&_ListaGateway.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ListaGateway *ListaGatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ListaGateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ListaGateway *ListaGatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ListaGateway.Contract.Fallback(&_ListaGateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ListaGateway *ListaGatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ListaGateway.Contract.Fallback(&_ListaGateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ListaGateway *ListaGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListaGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ListaGateway *ListaGatewaySession) Receive() (*types.Transaction, error) {
	return _ListaGateway.Contract.Receive(&_ListaGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ListaGateway *ListaGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _ListaGateway.Contract.Receive(&_ListaGateway.TransactOpts)
}

// ListaGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ListaGateway contract.
type ListaGatewayOwnershipTransferredIterator struct {
	Event *ListaGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ListaGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListaGatewayOwnershipTransferred)
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
		it.Event = new(ListaGatewayOwnershipTransferred)
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
func (it *ListaGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListaGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListaGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the ListaGateway contract.
type ListaGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ListaGateway *ListaGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ListaGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ListaGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ListaGatewayOwnershipTransferredIterator{contract: _ListaGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ListaGateway *ListaGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ListaGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ListaGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListaGatewayOwnershipTransferred)
				if err := _ListaGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ListaGateway *ListaGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*ListaGatewayOwnershipTransferred, error) {
	event := new(ListaGatewayOwnershipTransferred)
	if err := _ListaGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
