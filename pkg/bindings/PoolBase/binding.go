// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PoolBase

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

// PoolBaseMetaData contains all meta data concerning the PoolBase contract.
var PoolBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_poolController\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initializeBase\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001b2938038062001b29833981016040819052620000349162000113565b6001600160a01b0381166080526200004b62000052565b5062000145565b600054610100900460ff1615620000bf5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000111576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012657600080fd5b81516001600160a01b03811681146200013e57600080fd5b9392505050565b6080516119b3620001766000396000818161021d015281816106cc01528181610b7d0152610c9601526119b36000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80637a8b2637116100d8578063c4e01fa41161008c578063e3dae51c11610066578063e3dae51c1461033e578063f3e7387514610351578063fabc1cbc1461036457600080fd5b8063c4e01fa414610305578063ce7c2ac214610318578063d9caed121461032b57600080fd5b80638c871019116100bd5780638c871019146102ca5780638f6a6240146102dd578063ab5921e1146102f057600080fd5b80637a8b263714610291578063886f1195146102a457600080fd5b80634aa9d5851161012f578063595c6a6711610114578063595c6a67146102525780635ac86ab71461025a5780635c975abb1461028957600080fd5b80634aa9d58514610218578063553ca5f81461023f57600080fd5b80632495a599116101605780632495a599146101a45780633a98ef39146101ee57806347e7ef241461020557600080fd5b806310d67a2f1461017c578063136439dd14610191575b600080fd5b61018f61018a36600461169e565b610377565b005b61018f61019f3660046116c2565b6104a0565b6033546101c49073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101f760345481565b6040519081526020016101e5565b6101f76102133660046116db565b610623565b6101c47f000000000000000000000000000000000000000000000000000000000000000081565b6101f761024d36600461169e565b6107d2565b61018f6107e6565b610279610268366004611707565b6001805460ff9092161b9081161490565b60405190151581526020016101e5565b6001546101f7565b6101f761029f3660046116c2565b61092f565b6000546101c49062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6101f76102d83660046116c2565b61097a565b6101f76102eb36600461169e565b610985565b6102f8610993565b6040516101e5919061174e565b61018f61031336600461179f565b6109b3565b6101f761032636600461169e565b610b2f565b61018f6103393660046117d8565b610bea565b6101f761034c3660046116c2565b610d9d565b6101f761035f3660046116c2565b610dd6565b61018f6103723660046116c2565b610de1565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104089190611819565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250906104935760405162461bcd60e51b815260040161048a919061174e565b60405180910390fd5b5061049d81610f8d565b50565b6000546040517f46fbf68e0000000000000000000000000000000000000000000000000000000081523360048201526201000090910473ffffffffffffffffffffffffffffffffffffffff16906346fbf68e90602401602060405180830381865afa158015610513573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105379190611836565b6040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152509061058b5760405162461bcd60e51b815260040161048a919061174e565b5060018054604080518082019091529182527f37000000000000000000000000000000000000000000000000000000000000006020830152828116146105e45760405162461bcd60e51b815260040161048a919061174e565b50600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60018054604080518082019091528281527f35000000000000000000000000000000000000000000000000000000000000006020820152600092839281160361067f5760405162461bcd60e51b815260040161048a919061174e565b5060408051808201909152600281527f313900000000000000000000000000000000000000000000000000000000000060208201523373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461070a5760405162461bcd60e51b815260040161048a919061174e565b50610715848461109d565b60345460006107266103e883611887565b905060006103e861073561110d565b61073f9190611887565b9050600061074d878361189a565b90508061075a84896118ad565b61076491906118c4565b60408051808201909152600281527f31350000000000000000000000000000000000000000000000000000000000006020820152909650866107b95760405162461bcd60e51b815260040161048a919061174e565b506107c48685611887565b603455505050505092915050565b60006107e061029f83610b2f565b92915050565b6000546040517f46fbf68e0000000000000000000000000000000000000000000000000000000081523360048201526201000090910473ffffffffffffffffffffffffffffffffffffffff16906346fbf68e90602401602060405180830381865afa158015610859573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087d9190611836565b6040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250906108d15760405162461bcd60e51b815260040161048a919061174e565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000806103e86034546109429190611887565b905060006103e861095161110d565b61095b9190611887565b90508161096885836118ad565b61097291906118c4565b949350505050565b60006107e082610d9d565b60006107e061035f83610b2f565b606060405180608001604052806049815260200161193560499139905090565b600054610100900460ff16158080156109d35750600054600160ff909116105b806109ed5750303b1580156109ed575060005460ff166001145b610a5f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161048a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610abd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610ac783836111a5565b8015610b2a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6040517fb6230d5f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301523060248301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063b6230d5f90604401602060405180830381865afa158015610bc6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e091906118ff565b60018054600290811614156040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525090610c495760405162461bcd60e51b815260040161048a919061174e565b5060408051808201909152600281527f313900000000000000000000000000000000000000000000000000000000000060208201523373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610cd45760405162461bcd60e51b815260040161048a919061174e565b50610ce0848484611271565b60345460408051808201909152600281527f3432000000000000000000000000000000000000000000000000000000000000602082015281841115610d385760405162461bcd60e51b815260040161048a919061174e565b506000610d476103e883611887565b905060006103e8610d5661110d565b610d609190611887565b9050600082610d6f87846118ad565b610d7991906118c4565b9050610d85868561189a565b603455610d938888836112e7565b5050505050505050565b6000806103e8603454610db09190611887565b905060006103e8610dbf61110d565b610dc99190611887565b90508061096883866118ad565b60006107e08261092f565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e729190611819565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f320000000000000000000000000000000000000000000000000000000000000081525090610ef45760405162461bcd60e51b815260040161048a919061174e565b506001541981196001541916146040518060400160405280600181526020017f380000000000000000000000000000000000000000000000000000000000000081525090610f555760405162461bcd60e51b815260040161048a919061174e565b50600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610618565b60408051808201909152600181527f3300000000000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8216610ff55760405162461bcd60e51b815260040161048a919061174e565b506000546040805173ffffffffffffffffffffffffffffffffffffffff620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a16000805473ffffffffffffffffffffffffffffffffffffffff90921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b60335460408051808201909152600281527f343300000000000000000000000000000000000000000000000000000000000060208201529073ffffffffffffffffffffffffffffffffffffffff848116911614610b2a5760405162461bcd60e51b815260040161048a919061174e565b6033546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa15801561117c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a091906118ff565b905090565b600054610100900460ff166112225760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161048a565b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841617905561126d816000611308565b5050565b60335460408051808201909152600281527f343400000000000000000000000000000000000000000000000000000000000060208201529073ffffffffffffffffffffffffffffffffffffffff8481169116146112e15760405162461bcd60e51b815260040161048a919061174e565b50505050565b610b2a73ffffffffffffffffffffffffffffffffffffffff831684836113e1565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16158015611349575073ffffffffffffffffffffffffffffffffffffffff821615155b6040518060400160405280600181526020017f36000000000000000000000000000000000000000000000000000000000000008152509061139d5760405162461bcd60e51b815260040161048a919061174e565b50600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261126d82610f8d565b6040805173ffffffffffffffffffffffffffffffffffffffff848116602483015260448083018590528351808403909101815260649092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610b2a928692916000916114ac91851690849061153f565b90508051600014806114cd5750808060200190518101906114cd9190611836565b610b2a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161048a565b60606109728484600085856000808673ffffffffffffffffffffffffffffffffffffffff1685876040516115739190611918565b60006040518083038185875af1925050503d80600081146115b0576040519150601f19603f3d011682016040523d82523d6000602084013e6115b5565b606091505b50915091506115c6878383876115d1565b979650505050505050565b6060831561164d5782516000036116465773ffffffffffffffffffffffffffffffffffffffff85163b6116465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161048a565b5081610972565b61097283838151156116625781518083602001fd5b8060405162461bcd60e51b815260040161048a919061174e565b73ffffffffffffffffffffffffffffffffffffffff8116811461049d57600080fd5b6000602082840312156116b057600080fd5b81356116bb8161167c565b9392505050565b6000602082840312156116d457600080fd5b5035919050565b600080604083850312156116ee57600080fd5b82356116f98161167c565b946020939093013593505050565b60006020828403121561171957600080fd5b813560ff811681146116bb57600080fd5b60005b8381101561174557818101518382015260200161172d565b50506000910152565b602081526000825180602084015261176d81604085016020870161172a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600080604083850312156117b257600080fd5b82356117bd8161167c565b915060208301356117cd8161167c565b809150509250929050565b6000806000606084860312156117ed57600080fd5b83356117f88161167c565b925060208401356118088161167c565b929592945050506040919091013590565b60006020828403121561182b57600080fd5b81516116bb8161167c565b60006020828403121561184857600080fd5b815180151581146116bb57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156107e0576107e0611858565b818103818111156107e0576107e0611858565b80820281158282048414176107e0576107e0611858565b6000826118fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60006020828403121561191157600080fd5b5051919050565b6000825161192a81846020870161172a565b919091019291505056fe4261736520506f6f6c20696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a26469706673582212209133becf764fc98d710f414cb1799687bfa9fdda5aed9dfd0081b2db2e3722cd64736f6c63430008140033",
}

// PoolBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolBaseMetaData.ABI instead.
var PoolBaseABI = PoolBaseMetaData.ABI

// PoolBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolBaseMetaData.Bin instead.
var PoolBaseBin = PoolBaseMetaData.Bin

// DeployPoolBase deploys a new Ethereum contract, binding an instance of PoolBase to it.
func DeployPoolBase(auth *bind.TransactOpts, backend bind.ContractBackend, _poolController common.Address) (common.Address, *types.Transaction, *PoolBase, error) {
	parsed, err := PoolBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolBaseBin), backend, _poolController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolBase{PoolBaseCaller: PoolBaseCaller{contract: contract}, PoolBaseTransactor: PoolBaseTransactor{contract: contract}, PoolBaseFilterer: PoolBaseFilterer{contract: contract}}, nil
}

// PoolBase is an auto generated Go binding around an Ethereum contract.
type PoolBase struct {
	PoolBaseCaller     // Read-only binding to the contract
	PoolBaseTransactor // Write-only binding to the contract
	PoolBaseFilterer   // Log filterer for contract events
}

// PoolBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolBaseSession struct {
	Contract     *PoolBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolBaseCallerSession struct {
	Contract *PoolBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PoolBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolBaseTransactorSession struct {
	Contract     *PoolBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PoolBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolBaseRaw struct {
	Contract *PoolBase // Generic contract binding to access the raw methods on
}

// PoolBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolBaseCallerRaw struct {
	Contract *PoolBaseCaller // Generic read-only contract binding to access the raw methods on
}

// PoolBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolBaseTransactorRaw struct {
	Contract *PoolBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolBase creates a new instance of PoolBase, bound to a specific deployed contract.
func NewPoolBase(address common.Address, backend bind.ContractBackend) (*PoolBase, error) {
	contract, err := bindPoolBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolBase{PoolBaseCaller: PoolBaseCaller{contract: contract}, PoolBaseTransactor: PoolBaseTransactor{contract: contract}, PoolBaseFilterer: PoolBaseFilterer{contract: contract}}, nil
}

// NewPoolBaseCaller creates a new read-only instance of PoolBase, bound to a specific deployed contract.
func NewPoolBaseCaller(address common.Address, caller bind.ContractCaller) (*PoolBaseCaller, error) {
	contract, err := bindPoolBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolBaseCaller{contract: contract}, nil
}

// NewPoolBaseTransactor creates a new write-only instance of PoolBase, bound to a specific deployed contract.
func NewPoolBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolBaseTransactor, error) {
	contract, err := bindPoolBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTransactor{contract: contract}, nil
}

// NewPoolBaseFilterer creates a new log filterer instance of PoolBase, bound to a specific deployed contract.
func NewPoolBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolBaseFilterer, error) {
	contract, err := bindPoolBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolBaseFilterer{contract: contract}, nil
}

// bindPoolBase binds a generic wrapper to an already deployed contract.
func bindPoolBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolBase *PoolBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolBase.Contract.PoolBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolBase *PoolBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolBase.Contract.PoolBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolBase *PoolBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolBase.Contract.PoolBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolBase *PoolBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolBase *PoolBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolBase *PoolBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolBase.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_PoolBase *PoolBaseCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_PoolBase *PoolBaseSession) Explanation() (string, error) {
	return _PoolBase.Contract.Explanation(&_PoolBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_PoolBase *PoolBaseCallerSession) Explanation() (string, error) {
	return _PoolBase.Contract.Explanation(&_PoolBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PoolBase *PoolBaseCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PoolBase *PoolBaseSession) Paused(index uint8) (bool, error) {
	return _PoolBase.Contract.Paused(&_PoolBase.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PoolBase *PoolBaseCallerSession) Paused(index uint8) (bool, error) {
	return _PoolBase.Contract.Paused(&_PoolBase.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PoolBase *PoolBaseCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PoolBase *PoolBaseSession) Paused0() (*big.Int, error) {
	return _PoolBase.Contract.Paused0(&_PoolBase.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) Paused0() (*big.Int, error) {
	return _PoolBase.Contract.Paused0(&_PoolBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PoolBase *PoolBaseCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PoolBase *PoolBaseSession) PauserRegistry() (common.Address, error) {
	return _PoolBase.Contract.PauserRegistry(&_PoolBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PoolBase *PoolBaseCallerSession) PauserRegistry() (common.Address, error) {
	return _PoolBase.Contract.PauserRegistry(&_PoolBase.CallOpts)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_PoolBase *PoolBaseCaller) PoolController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "poolController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_PoolBase *PoolBaseSession) PoolController() (common.Address, error) {
	return _PoolBase.Contract.PoolController(&_PoolBase.CallOpts)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_PoolBase *PoolBaseCallerSession) PoolController() (common.Address, error) {
	return _PoolBase.Contract.PoolController(&_PoolBase.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_PoolBase *PoolBaseCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_PoolBase *PoolBaseSession) Shares(user common.Address) (*big.Int, error) {
	return _PoolBase.Contract.Shares(&_PoolBase.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _PoolBase.Contract.Shares(&_PoolBase.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_PoolBase *PoolBaseCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_PoolBase *PoolBaseSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.SharesToUnderlying(&_PoolBase.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.SharesToUnderlying(&_PoolBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_PoolBase *PoolBaseCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_PoolBase *PoolBaseSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.SharesToUnderlyingView(&_PoolBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.SharesToUnderlyingView(&_PoolBase.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PoolBase *PoolBaseCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PoolBase *PoolBaseSession) TotalShares() (*big.Int, error) {
	return _PoolBase.Contract.TotalShares(&_PoolBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) TotalShares() (*big.Int, error) {
	return _PoolBase.Contract.TotalShares(&_PoolBase.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_PoolBase *PoolBaseCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_PoolBase *PoolBaseSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.UnderlyingToShares(&_PoolBase.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.UnderlyingToShares(&_PoolBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_PoolBase *PoolBaseCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_PoolBase *PoolBaseSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.UnderlyingToSharesView(&_PoolBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBase.Contract.UnderlyingToSharesView(&_PoolBase.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolBase *PoolBaseCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolBase *PoolBaseSession) UnderlyingToken() (common.Address, error) {
	return _PoolBase.Contract.UnderlyingToken(&_PoolBase.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolBase *PoolBaseCallerSession) UnderlyingToken() (common.Address, error) {
	return _PoolBase.Contract.UnderlyingToken(&_PoolBase.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_PoolBase *PoolBaseCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolBase.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_PoolBase *PoolBaseSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _PoolBase.Contract.UserUnderlyingView(&_PoolBase.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_PoolBase *PoolBaseCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _PoolBase.Contract.UserUnderlyingView(&_PoolBase.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_PoolBase *PoolBaseTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_PoolBase *PoolBaseSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Deposit(&_PoolBase.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_PoolBase *PoolBaseTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Deposit(&_PoolBase.TransactOpts, token, amount)
}

// InitializeBase is a paid mutator transaction binding the contract method 0xc4e01fa4.
//
// Solidity: function initializeBase(address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBase *PoolBaseTransactor) InitializeBase(opts *bind.TransactOpts, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "initializeBase", _underlyingToken, _pauserRegistry)
}

// InitializeBase is a paid mutator transaction binding the contract method 0xc4e01fa4.
//
// Solidity: function initializeBase(address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBase *PoolBaseSession) InitializeBase(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBase.Contract.InitializeBase(&_PoolBase.TransactOpts, _underlyingToken, _pauserRegistry)
}

// InitializeBase is a paid mutator transaction binding the contract method 0xc4e01fa4.
//
// Solidity: function initializeBase(address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBase *PoolBaseTransactorSession) InitializeBase(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBase.Contract.InitializeBase(&_PoolBase.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PoolBase *PoolBaseTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PoolBase *PoolBaseSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Pause(&_PoolBase.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PoolBase *PoolBaseTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Pause(&_PoolBase.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PoolBase *PoolBaseTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PoolBase *PoolBaseSession) PauseAll() (*types.Transaction, error) {
	return _PoolBase.Contract.PauseAll(&_PoolBase.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PoolBase *PoolBaseTransactorSession) PauseAll() (*types.Transaction, error) {
	return _PoolBase.Contract.PauseAll(&_PoolBase.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PoolBase *PoolBaseTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PoolBase *PoolBaseSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBase.Contract.SetPauserRegistry(&_PoolBase.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PoolBase *PoolBaseTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBase.Contract.SetPauserRegistry(&_PoolBase.TransactOpts, newPauserRegistry)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PoolBase *PoolBaseTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PoolBase *PoolBaseSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Unpause(&_PoolBase.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PoolBase *PoolBaseTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Unpause(&_PoolBase.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_PoolBase *PoolBaseTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_PoolBase *PoolBaseSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _PoolBase.Contract.UserUnderlying(&_PoolBase.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_PoolBase *PoolBaseTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _PoolBase.Contract.UserUnderlying(&_PoolBase.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_PoolBase *PoolBaseTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _PoolBase.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_PoolBase *PoolBaseSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Withdraw(&_PoolBase.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_PoolBase *PoolBaseTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _PoolBase.Contract.Withdraw(&_PoolBase.TransactOpts, recipient, token, amountShares)
}

// PoolBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoolBase contract.
type PoolBaseInitializedIterator struct {
	Event *PoolBaseInitialized // Event containing the contract specifics and raw log

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
func (it *PoolBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseInitialized)
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
		it.Event = new(PoolBaseInitialized)
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
func (it *PoolBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseInitialized represents a Initialized event raised by the PoolBase contract.
type PoolBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolBase *PoolBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolBaseInitializedIterator, error) {

	logs, sub, err := _PoolBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolBaseInitializedIterator{contract: _PoolBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolBase *PoolBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _PoolBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseInitialized)
				if err := _PoolBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolBase *PoolBaseFilterer) ParseInitialized(log types.Log) (*PoolBaseInitialized, error) {
	event := new(PoolBaseInitialized)
	if err := _PoolBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PoolBase contract.
type PoolBasePausedIterator struct {
	Event *PoolBasePaused // Event containing the contract specifics and raw log

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
func (it *PoolBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBasePaused)
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
		it.Event = new(PoolBasePaused)
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
func (it *PoolBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBasePaused represents a Paused event raised by the PoolBase contract.
type PoolBasePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PoolBase *PoolBaseFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*PoolBasePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBase.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolBasePausedIterator{contract: _PoolBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PoolBase *PoolBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolBasePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBase.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBasePaused)
				if err := _PoolBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PoolBase *PoolBaseFilterer) ParsePaused(log types.Log) (*PoolBasePaused, error) {
	event := new(PoolBasePaused)
	if err := _PoolBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBasePauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the PoolBase contract.
type PoolBasePauserRegistrySetIterator struct {
	Event *PoolBasePauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *PoolBasePauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBasePauserRegistrySet)
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
		it.Event = new(PoolBasePauserRegistrySet)
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
func (it *PoolBasePauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBasePauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBasePauserRegistrySet represents a PauserRegistrySet event raised by the PoolBase contract.
type PoolBasePauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PoolBase *PoolBaseFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*PoolBasePauserRegistrySetIterator, error) {

	logs, sub, err := _PoolBase.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &PoolBasePauserRegistrySetIterator{contract: _PoolBase.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PoolBase *PoolBaseFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *PoolBasePauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _PoolBase.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBasePauserRegistrySet)
				if err := _PoolBase.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PoolBase *PoolBaseFilterer) ParsePauserRegistrySet(log types.Log) (*PoolBasePauserRegistrySet, error) {
	event := new(PoolBasePauserRegistrySet)
	if err := _PoolBase.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PoolBase contract.
type PoolBaseUnpausedIterator struct {
	Event *PoolBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseUnpaused)
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
		it.Event = new(PoolBaseUnpaused)
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
func (it *PoolBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseUnpaused represents a Unpaused event raised by the PoolBase contract.
type PoolBaseUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PoolBase *PoolBaseFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*PoolBaseUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBase.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolBaseUnpausedIterator{contract: _PoolBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PoolBase *PoolBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolBaseUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBase.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseUnpaused)
				if err := _PoolBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PoolBase *PoolBaseFilterer) ParseUnpaused(log types.Log) (*PoolBaseUnpaused, error) {
	event := new(PoolBaseUnpaused)
	if err := _PoolBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
