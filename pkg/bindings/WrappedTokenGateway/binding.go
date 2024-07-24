// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WrappedTokenGateway

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

// IDelegationControllerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationControllerWithdrawal struct {
	Staker         common.Address
	DelegatedTo    common.Address
	Withdrawer     common.Address
	Nonce          *big.Int
	StartTimestamp uint32
	Pools          []common.Address
	Shares         []*big.Int
}

// WrappedTokenGatewayMetaData contains all meta data concerning the WrappedTokenGateway contract.
var WrappedTokenGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_wrappedToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"_poolController\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"},{\"name\":\"_delegationController\",\"type\":\"address\",\"internalType\":\"contractIDelegationController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositNativeToken\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"emergencyNativeTransfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyTokenTransfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPoolAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWrappedTokenAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNativeTokens\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162001a8738038062001a8783398101604081905262000035916200015a565b6200004033620000f1565b6001600160a01b0380861660805283811660a05282811660c052811660e0526200006a84620000f1565b60c05160405163095ea7b360e01b81526001600160a01b03918216600482015260001960248201529086169063095ea7b3906044016020604051808303816000875af1158015620000bf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000e59190620001da565b50505050505062000205565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146200015757600080fd5b50565b600080600080600060a086880312156200017357600080fd5b8551620001808162000141565b6020870151909550620001938162000141565b6040870151909450620001a68162000141565b6060870151909350620001b98162000141565b6080870151909250620001cc8162000141565b809150509295509295909350565b600060208284031215620001ed57600080fd5b81518015158114620001fe57600080fd5b9392505050565b60805160a05160c05160e051611811620002766000396000610814015260006104340152600081816102e1015281816103de01526105eb01526000818160b7015281816101d60152818161031d0152818161040601528181610754015281816108b9015261097a01526118116000f3fe60806040526004361061009a5760003560e01c80638da5cb5b11610069578063d9d468161161004e578063d9d4681614610292578063f2fde38b146102b2578063f586c6d9146102d257610145565b80638da5cb5b14610247578063a3d5b2551461027257610145565b80631653a0c7146101a75780631dceb71f146101c757806320e2d8181461021f578063715018a61461023257610145565b36610145573373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610143576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f52656365697665206e6f7420616c6c6f7765640000000000000000000000000060448201526064015b60405180910390fd5b005b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f46616c6c6261636b206e6f7420616c6c6f776564000000000000000000000000604482015260640161013a565b3480156101b357600080fd5b506101436101c236600461104c565b610305565b3480156101d357600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b61014361022d366004611078565b61031b565b34801561023e57600080fd5b506101436104a5565b34801561025357600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff166101f6565b34801561027e57600080fd5b5061014361028d36600461109c565b6104b9565b34801561029e57600080fd5b506101436102ad366004611129565b6104e7565b3480156102be57600080fd5b506101436102cd366004611078565b610a02565b3480156102de57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006101f6565b61030d610ab9565b6103178282610b3a565b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561038357600080fd5b505af1158015610397573d6000803e3d6000fd5b50506040517f856abb2900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301527f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000811660448301523460648301527f000000000000000000000000000000000000000000000000000000000000000016935063856abb29925060840190506020604051808303816000875af1158015610481573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031791906111ed565b6104ad610ab9565b6104b76000610c1e565b565b6104c1610ab9565b6104e273ffffffffffffffffffffffffffffffffffffffff84168383610c93565b505050565b60005b87811015610722573389898381811061050557610505611206565b90506020028101906105179190611235565b610525906020810190611078565b73ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600281526020017f3430000000000000000000000000000000000000000000000000000000000000815250906105aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013a9190611297565b5060005b8989838181106105c0576105c0611206565b90506020028101906105d29190611235565b6105e09060a08101906112e8565b905081101561070f577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168a8a8481811061063257610632611206565b90506020028101906106449190611235565b6106529060a08101906112e8565b8381811061066257610662611206565b90506020020160208101906106779190611078565b73ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600281526020017f3431000000000000000000000000000000000000000000000000000000000000815250906106fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013a9190611297565b50806107078161137f565b9150506105ae565b508061071a8161137f565b9150506104ea565b506040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156107b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d491906111ed565b6040517f0af71b7700000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690630af71b7790610857908c908c908c908c908c908c908c908c906004016115d1565b600060405180830381600087803b15801561087157600080fd5b505af1158015610885573d6000803e3d6000fd5b50506040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600092507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1691506370a0823190602401602060405180830381865afa158015610916573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093a91906111ed565b905060006109488383611793565b6040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290529091507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b1580156109d357600080fd5b505af11580156109e7573d6000803e3d6000fd5b505050506109f53382610b3a565b5050505050505050505050565b610a0a610ab9565b73ffffffffffffffffffffffffffffffffffffffff8116610aad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161013a565b610ab681610c1e565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146104b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161013a565b6040805160008082526020820190925273ffffffffffffffffffffffffffffffffffffffff8416908390604051610b7191906117ac565b60006040518083038185875af1925050503d8060008114610bae576040519150601f19603f3d011682016040523d82523d6000602084013e610bb3565b606091505b50509050806104e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e41544956455f5452414e534645525f4641494c454400000000000000000000604482015260640161013a565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040805173ffffffffffffffffffffffffffffffffffffffff848116602483015260448083018590528351808403909101815260649092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526104e292869291600091610d5e918516908490610e0b565b9050805160001480610d7f575080806020019051810190610d7f91906117be565b6104e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161013a565b6060610e1a8484600085610e22565b949350505050565b606082471015610eb4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161013a565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610edd91906117ac565b60006040518083038185875af1925050503d8060008114610f1a576040519150601f19603f3d011682016040523d82523d6000602084013e610f1f565b606091505b5091509150610f3087838387610f3b565b979650505050505050565b60608315610fd1578251600003610fca5773ffffffffffffffffffffffffffffffffffffffff85163b610fca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161013a565b5081610e1a565b610e1a8383815115610fe65781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013a9190611297565b73ffffffffffffffffffffffffffffffffffffffff81168114610ab657600080fd5b80356110478161101a565b919050565b6000806040838503121561105f57600080fd5b823561106a8161101a565b946020939093013593505050565b60006020828403121561108a57600080fd5b81356110958161101a565b9392505050565b6000806000606084860312156110b157600080fd5b83356110bc8161101a565b925060208401356110cc8161101a565b929592945050506040919091013590565b60008083601f8401126110ef57600080fd5b50813567ffffffffffffffff81111561110757600080fd5b6020830191508360208260051b850101111561112257600080fd5b9250929050565b6000806000806000806000806080898b03121561114557600080fd5b883567ffffffffffffffff8082111561115d57600080fd5b6111698c838d016110dd565b909a50985060208b013591508082111561118257600080fd5b61118e8c838d016110dd565b909850965060408b01359150808211156111a757600080fd5b6111b38c838d016110dd565b909650945060608b01359150808211156111cc57600080fd5b506111d98b828c016110dd565b999c989b5096995094979396929594505050565b6000602082840312156111ff57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261126957600080fd5b9190910192915050565b60005b8381101561128e578181015183820152602001611276565b50506000910152565b60208152600082518060208401526112b6816040850160208701611273565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261131d57600080fd5b83018035915067ffffffffffffffff82111561133857600080fd5b6020019150600581901b360382131561112257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113b0576113b0611350565b5060010190565b803563ffffffff8116811461104757600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261140057600080fd5b830160208101925035905067ffffffffffffffff81111561142057600080fd5b8060051b360382131561112257600080fd5b8183526000602080850194508260005b8581101561147d5781356114558161101a565b73ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101611442565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156114ba57600080fd5b8260051b80836020870137939093016020019392505050565b60008383855260208086019550808560051b830101846000805b8881101561157e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868503018a5261152683896113cb565b808652868601845b828110156115695783356115418161101a565b73ffffffffffffffffffffffffffffffffffffffff168252928801929088019060010161152e565b509b87019b95505050918401916001016114ed565b509198975050505050505050565b8015158114610ab657600080fd5b8183526000602080850194508260005b8581101561147d5781356115bd8161158c565b1515875295820195908201906001016115aa565b60808082528101889052600060a060058a901b830181019083018b835b8c811015611744577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6086850301835281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218f360301811261164f57600080fd5b8e0160e0813561165e8161101a565b73ffffffffffffffffffffffffffffffffffffffff1686526020828101356116858161101a565b73ffffffffffffffffffffffffffffffffffffffff168188015260406116ac84820161103c565b73ffffffffffffffffffffffffffffffffffffffff1690880152606083810135908801526116dc608084016113b7565b63ffffffff1660808801526116f460a08401846113cb565b8360a08a0152611707848a018284611432565b9350505060c0611719818501856113cb565b9450888403828a015261172d848683611488565b9850505094850194939093019250506001016115ee565b505050828103602084015261175a81898b6114d3565b9050828103604084015261176f818789611488565b9050828103606084015261178481858761159a565b9b9a5050505050505050505050565b818103818111156117a6576117a6611350565b92915050565b60008251611269818460208701611273565b6000602082840312156117d057600080fd5b81516110958161158c56fea264697066735822122095394ea7e67f173411a0ab29d455e4baf54f10ce56b9d8ca5197c16590bc53fb64736f6c63430008140033",
}

// WrappedTokenGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedTokenGatewayMetaData.ABI instead.
var WrappedTokenGatewayABI = WrappedTokenGatewayMetaData.ABI

// WrappedTokenGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappedTokenGatewayMetaData.Bin instead.
var WrappedTokenGatewayBin = WrappedTokenGatewayMetaData.Bin

// DeployWrappedTokenGateway deploys a new Ethereum contract, binding an instance of WrappedTokenGateway to it.
func DeployWrappedTokenGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _wrappedToken common.Address, _owner common.Address, _pool common.Address, _poolController common.Address, _delegationController common.Address) (common.Address, *types.Transaction, *WrappedTokenGateway, error) {
	parsed, err := WrappedTokenGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedTokenGatewayBin), backend, _wrappedToken, _owner, _pool, _poolController, _delegationController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedTokenGateway{WrappedTokenGatewayCaller: WrappedTokenGatewayCaller{contract: contract}, WrappedTokenGatewayTransactor: WrappedTokenGatewayTransactor{contract: contract}, WrappedTokenGatewayFilterer: WrappedTokenGatewayFilterer{contract: contract}}, nil
}

// WrappedTokenGateway is an auto generated Go binding around an Ethereum contract.
type WrappedTokenGateway struct {
	WrappedTokenGatewayCaller     // Read-only binding to the contract
	WrappedTokenGatewayTransactor // Write-only binding to the contract
	WrappedTokenGatewayFilterer   // Log filterer for contract events
}

// WrappedTokenGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedTokenGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedTokenGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedTokenGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedTokenGatewaySession struct {
	Contract     *WrappedTokenGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WrappedTokenGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedTokenGatewayCallerSession struct {
	Contract *WrappedTokenGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// WrappedTokenGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedTokenGatewayTransactorSession struct {
	Contract     *WrappedTokenGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// WrappedTokenGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedTokenGatewayRaw struct {
	Contract *WrappedTokenGateway // Generic contract binding to access the raw methods on
}

// WrappedTokenGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedTokenGatewayCallerRaw struct {
	Contract *WrappedTokenGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedTokenGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedTokenGatewayTransactorRaw struct {
	Contract *WrappedTokenGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedTokenGateway creates a new instance of WrappedTokenGateway, bound to a specific deployed contract.
func NewWrappedTokenGateway(address common.Address, backend bind.ContractBackend) (*WrappedTokenGateway, error) {
	contract, err := bindWrappedTokenGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenGateway{WrappedTokenGatewayCaller: WrappedTokenGatewayCaller{contract: contract}, WrappedTokenGatewayTransactor: WrappedTokenGatewayTransactor{contract: contract}, WrappedTokenGatewayFilterer: WrappedTokenGatewayFilterer{contract: contract}}, nil
}

// NewWrappedTokenGatewayCaller creates a new read-only instance of WrappedTokenGateway, bound to a specific deployed contract.
func NewWrappedTokenGatewayCaller(address common.Address, caller bind.ContractCaller) (*WrappedTokenGatewayCaller, error) {
	contract, err := bindWrappedTokenGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenGatewayCaller{contract: contract}, nil
}

// NewWrappedTokenGatewayTransactor creates a new write-only instance of WrappedTokenGateway, bound to a specific deployed contract.
func NewWrappedTokenGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedTokenGatewayTransactor, error) {
	contract, err := bindWrappedTokenGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenGatewayTransactor{contract: contract}, nil
}

// NewWrappedTokenGatewayFilterer creates a new log filterer instance of WrappedTokenGateway, bound to a specific deployed contract.
func NewWrappedTokenGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedTokenGatewayFilterer, error) {
	contract, err := bindWrappedTokenGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenGatewayFilterer{contract: contract}, nil
}

// bindWrappedTokenGateway binds a generic wrapper to an already deployed contract.
func bindWrappedTokenGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedTokenGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedTokenGateway *WrappedTokenGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedTokenGateway.Contract.WrappedTokenGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedTokenGateway *WrappedTokenGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.WrappedTokenGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedTokenGateway *WrappedTokenGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.WrappedTokenGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedTokenGateway *WrappedTokenGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedTokenGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.contract.Transact(opts, method, params...)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewayCaller) GetPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedTokenGateway.contract.Call(opts, &out, "getPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewaySession) GetPoolAddress() (common.Address, error) {
	return _WrappedTokenGateway.Contract.GetPoolAddress(&_WrappedTokenGateway.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0xf586c6d9.
//
// Solidity: function getPoolAddress() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewayCallerSession) GetPoolAddress() (common.Address, error) {
	return _WrappedTokenGateway.Contract.GetPoolAddress(&_WrappedTokenGateway.CallOpts)
}

// GetWrappedTokenAddress is a free data retrieval call binding the contract method 0x1dceb71f.
//
// Solidity: function getWrappedTokenAddress() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewayCaller) GetWrappedTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedTokenGateway.contract.Call(opts, &out, "getWrappedTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenAddress is a free data retrieval call binding the contract method 0x1dceb71f.
//
// Solidity: function getWrappedTokenAddress() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewaySession) GetWrappedTokenAddress() (common.Address, error) {
	return _WrappedTokenGateway.Contract.GetWrappedTokenAddress(&_WrappedTokenGateway.CallOpts)
}

// GetWrappedTokenAddress is a free data retrieval call binding the contract method 0x1dceb71f.
//
// Solidity: function getWrappedTokenAddress() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewayCallerSession) GetWrappedTokenAddress() (common.Address, error) {
	return _WrappedTokenGateway.Contract.GetWrappedTokenAddress(&_WrappedTokenGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedTokenGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewaySession) Owner() (common.Address, error) {
	return _WrappedTokenGateway.Contract.Owner(&_WrappedTokenGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrappedTokenGateway *WrappedTokenGatewayCallerSession) Owner() (common.Address, error) {
	return _WrappedTokenGateway.Contract.Owner(&_WrappedTokenGateway.CallOpts)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x20e2d818.
//
// Solidity: function depositNativeToken(address staker) payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) DepositNativeToken(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.Transact(opts, "depositNativeToken", staker)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x20e2d818.
//
// Solidity: function depositNativeToken(address staker) payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) DepositNativeToken(staker common.Address) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.DepositNativeToken(&_WrappedTokenGateway.TransactOpts, staker)
}

// DepositNativeToken is a paid mutator transaction binding the contract method 0x20e2d818.
//
// Solidity: function depositNativeToken(address staker) payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) DepositNativeToken(staker common.Address) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.DepositNativeToken(&_WrappedTokenGateway.TransactOpts, staker)
}

// EmergencyNativeTransfer is a paid mutator transaction binding the contract method 0x1653a0c7.
//
// Solidity: function emergencyNativeTransfer(address to, uint256 amount) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) EmergencyNativeTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.Transact(opts, "emergencyNativeTransfer", to, amount)
}

// EmergencyNativeTransfer is a paid mutator transaction binding the contract method 0x1653a0c7.
//
// Solidity: function emergencyNativeTransfer(address to, uint256 amount) returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) EmergencyNativeTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.EmergencyNativeTransfer(&_WrappedTokenGateway.TransactOpts, to, amount)
}

// EmergencyNativeTransfer is a paid mutator transaction binding the contract method 0x1653a0c7.
//
// Solidity: function emergencyNativeTransfer(address to, uint256 amount) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) EmergencyNativeTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.EmergencyNativeTransfer(&_WrappedTokenGateway.TransactOpts, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) EmergencyTokenTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.Transact(opts, "emergencyTokenTransfer", token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.EmergencyTokenTransfer(&_WrappedTokenGateway.TransactOpts, token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.EmergencyTokenTransfer(&_WrappedTokenGateway.TransactOpts, token, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.RenounceOwnership(&_WrappedTokenGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.RenounceOwnership(&_WrappedTokenGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.TransferOwnership(&_WrappedTokenGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.TransferOwnership(&_WrappedTokenGateway.TransactOpts, newOwner)
}

// WithdrawNativeTokens is a paid mutator transaction binding the contract method 0xd9d46816.
//
// Solidity: function withdrawNativeTokens((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexs, bool[] receiveAsTokens) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) WithdrawNativeTokens(opts *bind.TransactOpts, withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexs []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.Transact(opts, "withdrawNativeTokens", withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens)
}

// WithdrawNativeTokens is a paid mutator transaction binding the contract method 0xd9d46816.
//
// Solidity: function withdrawNativeTokens((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexs, bool[] receiveAsTokens) returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) WithdrawNativeTokens(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexs []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.WithdrawNativeTokens(&_WrappedTokenGateway.TransactOpts, withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens)
}

// WithdrawNativeTokens is a paid mutator transaction binding the contract method 0xd9d46816.
//
// Solidity: function withdrawNativeTokens((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexs, bool[] receiveAsTokens) returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) WithdrawNativeTokens(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexs []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.WithdrawNativeTokens(&_WrappedTokenGateway.TransactOpts, withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.Fallback(&_WrappedTokenGateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.Fallback(&_WrappedTokenGateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewaySession) Receive() (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.Receive(&_WrappedTokenGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedTokenGateway *WrappedTokenGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _WrappedTokenGateway.Contract.Receive(&_WrappedTokenGateway.TransactOpts)
}

// WrappedTokenGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WrappedTokenGateway contract.
type WrappedTokenGatewayOwnershipTransferredIterator struct {
	Event *WrappedTokenGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WrappedTokenGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenGatewayOwnershipTransferred)
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
		it.Event = new(WrappedTokenGatewayOwnershipTransferred)
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
func (it *WrappedTokenGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the WrappedTokenGateway contract.
type WrappedTokenGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappedTokenGateway *WrappedTokenGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WrappedTokenGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WrappedTokenGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenGatewayOwnershipTransferredIterator{contract: _WrappedTokenGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WrappedTokenGateway *WrappedTokenGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrappedTokenGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WrappedTokenGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenGatewayOwnershipTransferred)
				if err := _WrappedTokenGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WrappedTokenGateway *WrappedTokenGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*WrappedTokenGatewayOwnershipTransferred, error) {
	event := new(WrappedTokenGatewayOwnershipTransferred)
	if err := _WrappedTokenGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
