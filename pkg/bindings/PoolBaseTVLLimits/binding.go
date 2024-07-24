// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PoolBaseTVLLimits

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

// PoolBaseTVLLimitsMetaData contains all meta data concerning the PoolBaseTVLLimits contract.
var PoolBaseTVLLimitsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_poolController\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTVLLimits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_maxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeBase\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxPerDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTotalDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMaxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPerDepositUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalDepositsUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200206338038062002063833981016040819052620000349162000115565b6001600160a01b038116608052806200004c62000054565b505062000147565b600054610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000113576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012857600080fd5b81516001600160a01b03811681146200014057600080fd5b9392505050565b608051611eeb6200017860003960008181610283015281816109fc01528181610ead0152610fc60152611eeb6000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c806361b01b5d116100ee578063c4e01fa411610097578063df6fadc111610071578063df6fadc1146103ad578063e3dae51c146103c8578063f3e73875146103db578063fabc1cbc146103ee57600080fd5b8063c4e01fa414610374578063ce7c2ac214610387578063d9caed121461039a57600080fd5b80638c871019116100c85780638c871019146103395780638f6a62401461034c578063ab5921e11461035f57600080fd5b806361b01b5d146102f75780637a8b263714610300578063886f11951461031357600080fd5b806343fe08b01161015b578063553ca5f811610135578063553ca5f8146102a5578063595c6a67146102b85780635ac86ab7146102c05780635c975abb146102ef57600080fd5b806343fe08b01461026257806347e7ef241461026b5780634aa9d5851461027e57600080fd5b8063136439dd1161018c578063136439dd146101ee5780632495a599146102015780633a98ef391461024b57600080fd5b8063019e2729146101b357806310d67a2f146101c857806311c70c9d146101db575b600080fd5b6101c66101c1366004611b6a565b610401565b005b6101c66101d6366004611bb4565b61058e565b6101c66101e9366004611bd8565b6106ae565b6101c66101fc366004611bfa565b6107d0565b6033546102219073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61025460345481565b604051908152602001610242565b61025460665481565b610254610279366004611c13565b610953565b6102217f000000000000000000000000000000000000000000000000000000000000000081565b6102546102b3366004611bb4565b610b02565b6101c6610b16565b6102df6102ce366004611c3f565b6001805460ff9092161b9081161490565b6040519015158152602001610242565b600154610254565b61025460675481565b61025461030e366004611bfa565b610c5f565b6000546102219062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b610254610347366004611bfa565b610caa565b61025461035a366004611bb4565b610cb5565b610367610cc3565b6040516102429190611c86565b6101c6610382366004611cd7565b610ce3565b610254610395366004611bb4565b610e5f565b6101c66103a8366004611d10565b610f1a565b60665460675460408051928352602083019190915201610242565b6102546103d6366004611bfa565b6110cd565b6102546103e9366004611bfa565b611106565b6101c66103fc366004611bfa565b611111565b600054610100900460ff16158080156104215750600054600160ff909116105b8061043b5750303b15801561043b575060005460ff166001145b6104b25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561051057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61051a85856112bd565b6105248383611396565b801561058757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061f9190611d51565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250906106a15760405162461bcd60e51b81526004016104a99190611c86565b506106ab8161145e565b50565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561071b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073f9190611d51565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250906107c15760405162461bcd60e51b81526004016104a99190611c86565b506107cc82826112bd565b5050565b6000546040517f46fbf68e0000000000000000000000000000000000000000000000000000000081523360048201526201000090910473ffffffffffffffffffffffffffffffffffffffff16906346fbf68e90602401602060405180830381865afa158015610843573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108679190611d6e565b6040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250906108bb5760405162461bcd60e51b81526004016104a99190611c86565b5060018054604080518082019091529182527f37000000000000000000000000000000000000000000000000000000000000006020830152828116146109145760405162461bcd60e51b81526004016104a99190611c86565b50600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60018054604080518082019091528281527f3500000000000000000000000000000000000000000000000000000000000000602082015260009283928116036109af5760405162461bcd60e51b81526004016104a99190611c86565b5060408051808201909152600281527f313900000000000000000000000000000000000000000000000000000000000060208201523373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a3a5760405162461bcd60e51b81526004016104a99190611c86565b50610a45848461156e565b6034546000610a566103e883611dbf565b905060006103e8610a65611635565b610a6f9190611dbf565b90506000610a7d8783611dd2565b905080610a8a8489611de5565b610a949190611dfc565b60408051808201909152600281527f3135000000000000000000000000000000000000000000000000000000000000602082015290965086610ae95760405162461bcd60e51b81526004016104a99190611c86565b50610af48685611dbf565b603455505050505092915050565b6000610b1061030e83610e5f565b92915050565b6000546040517f46fbf68e0000000000000000000000000000000000000000000000000000000081523360048201526201000090910473ffffffffffffffffffffffffffffffffffffffff16906346fbf68e90602401602060405180830381865afa158015610b89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bad9190611d6e565b6040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525090610c015760405162461bcd60e51b81526004016104a99190611c86565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000806103e8603454610c729190611dbf565b905060006103e8610c81611635565b610c8b9190611dbf565b905081610c988583611de5565b610ca29190611dfc565b949350505050565b6000610b10826110cd565b6000610b106103e983610e5f565b6060604051806080016040528060498152602001611e6d60499139905090565b600054610100900460ff1615808015610d035750600054600160ff909116105b80610d1d5750303b158015610d1d575060005460ff166001145b610d8f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104a9565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610ded57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610df78383611396565b8015610e5a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6040517fb6230d5f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301523060248301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063b6230d5f90604401602060405180830381865afa158015610ef6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b109190611e37565b60018054600290811614156040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525090610f795760405162461bcd60e51b81526004016104a99190611c86565b5060408051808201909152600281527f313900000000000000000000000000000000000000000000000000000000000060208201523373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146110045760405162461bcd60e51b81526004016104a99190611c86565b506110108484846116cd565b60345460408051808201909152600281527f34320000000000000000000000000000000000000000000000000000000000006020820152818411156110685760405162461bcd60e51b81526004016104a99190611c86565b5060006110776103e883611dbf565b905060006103e8611086611635565b6110909190611dbf565b905060008261109f8784611de5565b6110a99190611dfc565b90506110b58685611dd2565b6034556110c3888883611743565b5050505050505050565b6000806103e86034546110e09190611dbf565b905060006103e86110ef611635565b6110f99190611dbf565b905080610c988386611de5565b6000610b1082610c5f565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561117e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a29190611d51565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250906112245760405162461bcd60e51b81526004016104a99190611c86565b506001541981196001541916146040518060400160405280600181526020017f3800000000000000000000000000000000000000000000000000000000000000815250906112855760405162461bcd60e51b81526004016104a99190611c86565b50600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610948565b60665460408051918252602082018490527ff97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5910160405180910390a160675460408051918252602082018390527f6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452910160405180910390a160408051808201909152600281527f343700000000000000000000000000000000000000000000000000000000000060208201528183111561138a5760405162461bcd60e51b81526004016104a99190611c86565b50606691909155606755565b600054610100900460ff166114135760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a9565b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790556107cc816000611764565b60408051808201909152600181527f3300000000000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff82166114c65760405162461bcd60e51b81526004016104a99190611c86565b506000546040805173ffffffffffffffffffffffffffffffffffffffff620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a16000805473ffffffffffffffffffffffffffffffffffffffff90921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b6066548111156040518060400160405280600281526020017f3436000000000000000000000000000000000000000000000000000000000000815250906115c85760405162461bcd60e51b81526004016104a99190611c86565b506067546115d4611635565b11156040518060400160405280600281526020017f34350000000000000000000000000000000000000000000000000000000000008152509061162a5760405162461bcd60e51b81526004016104a99190611c86565b506107cc828261183d565b6033546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156116a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c89190611e37565b905090565b60335460408051808201909152600281527f343400000000000000000000000000000000000000000000000000000000000060208201529073ffffffffffffffffffffffffffffffffffffffff84811691161461173d5760405162461bcd60e51b81526004016104a99190611c86565b50505050565b610e5a73ffffffffffffffffffffffffffffffffffffffff831684836118ad565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff161580156117a5575073ffffffffffffffffffffffffffffffffffffffff821615155b6040518060400160405280600181526020017f3600000000000000000000000000000000000000000000000000000000000000815250906117f95760405162461bcd60e51b81526004016104a99190611c86565b50600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26107cc8261145e565b60335460408051808201909152600281527f343300000000000000000000000000000000000000000000000000000000000060208201529073ffffffffffffffffffffffffffffffffffffffff848116911614610e5a5760405162461bcd60e51b81526004016104a99190611c86565b6040805173ffffffffffffffffffffffffffffffffffffffff848116602483015260448083018590528351808403909101815260649092018352602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490840152610e5a92869291600091611978918516908490611a0b565b90508051600014806119995750808060200190518101906119999190611d6e565b610e5a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016104a9565b6060610ca28484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051611a3f9190611e50565b60006040518083038185875af1925050503d8060008114611a7c576040519150601f19603f3d011682016040523d82523d6000602084013e611a81565b606091505b5091509150611a9287838387611a9d565b979650505050505050565b60608315611b19578251600003611b125773ffffffffffffffffffffffffffffffffffffffff85163b611b125760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104a9565b5081610ca2565b610ca28383815115611b2e5781518083602001fd5b8060405162461bcd60e51b81526004016104a99190611c86565b73ffffffffffffffffffffffffffffffffffffffff811681146106ab57600080fd5b60008060008060808587031215611b8057600080fd5b84359350602085013592506040850135611b9981611b48565b91506060850135611ba981611b48565b939692955090935050565b600060208284031215611bc657600080fd5b8135611bd181611b48565b9392505050565b60008060408385031215611beb57600080fd5b50508035926020909101359150565b600060208284031215611c0c57600080fd5b5035919050565b60008060408385031215611c2657600080fd5b8235611c3181611b48565b946020939093013593505050565b600060208284031215611c5157600080fd5b813560ff81168114611bd157600080fd5b60005b83811015611c7d578181015183820152602001611c65565b50506000910152565b6020815260008251806020840152611ca5816040850160208701611c62565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60008060408385031215611cea57600080fd5b8235611cf581611b48565b91506020830135611d0581611b48565b809150509250929050565b600080600060608486031215611d2557600080fd5b8335611d3081611b48565b92506020840135611d4081611b48565b929592945050506040919091013590565b600060208284031215611d6357600080fd5b8151611bd181611b48565b600060208284031215611d8057600080fd5b81518015158114611bd157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610b1057610b10611d90565b81810381811115610b1057610b10611d90565b8082028115828204841417610b1057610b10611d90565b600082611e32577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600060208284031215611e4957600080fd5b5051919050565b60008251611e62818460208701611c62565b919091019291505056fe4261736520506f6f6c20696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a2646970667358221220783fa38a40028cabb1e0a464f52b0ede459b260af509a363dab7c45a3061267d64736f6c63430008140033",
}

// PoolBaseTVLLimitsABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolBaseTVLLimitsMetaData.ABI instead.
var PoolBaseTVLLimitsABI = PoolBaseTVLLimitsMetaData.ABI

// PoolBaseTVLLimitsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolBaseTVLLimitsMetaData.Bin instead.
var PoolBaseTVLLimitsBin = PoolBaseTVLLimitsMetaData.Bin

// DeployPoolBaseTVLLimits deploys a new Ethereum contract, binding an instance of PoolBaseTVLLimits to it.
func DeployPoolBaseTVLLimits(auth *bind.TransactOpts, backend bind.ContractBackend, _poolController common.Address) (common.Address, *types.Transaction, *PoolBaseTVLLimits, error) {
	parsed, err := PoolBaseTVLLimitsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolBaseTVLLimitsBin), backend, _poolController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolBaseTVLLimits{PoolBaseTVLLimitsCaller: PoolBaseTVLLimitsCaller{contract: contract}, PoolBaseTVLLimitsTransactor: PoolBaseTVLLimitsTransactor{contract: contract}, PoolBaseTVLLimitsFilterer: PoolBaseTVLLimitsFilterer{contract: contract}}, nil
}

// PoolBaseTVLLimits is an auto generated Go binding around an Ethereum contract.
type PoolBaseTVLLimits struct {
	PoolBaseTVLLimitsCaller     // Read-only binding to the contract
	PoolBaseTVLLimitsTransactor // Write-only binding to the contract
	PoolBaseTVLLimitsFilterer   // Log filterer for contract events
}

// PoolBaseTVLLimitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolBaseTVLLimitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolBaseTVLLimitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolBaseTVLLimitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolBaseTVLLimitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolBaseTVLLimitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolBaseTVLLimitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolBaseTVLLimitsSession struct {
	Contract     *PoolBaseTVLLimits // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PoolBaseTVLLimitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolBaseTVLLimitsCallerSession struct {
	Contract *PoolBaseTVLLimitsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PoolBaseTVLLimitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolBaseTVLLimitsTransactorSession struct {
	Contract     *PoolBaseTVLLimitsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PoolBaseTVLLimitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolBaseTVLLimitsRaw struct {
	Contract *PoolBaseTVLLimits // Generic contract binding to access the raw methods on
}

// PoolBaseTVLLimitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolBaseTVLLimitsCallerRaw struct {
	Contract *PoolBaseTVLLimitsCaller // Generic read-only contract binding to access the raw methods on
}

// PoolBaseTVLLimitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolBaseTVLLimitsTransactorRaw struct {
	Contract *PoolBaseTVLLimitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolBaseTVLLimits creates a new instance of PoolBaseTVLLimits, bound to a specific deployed contract.
func NewPoolBaseTVLLimits(address common.Address, backend bind.ContractBackend) (*PoolBaseTVLLimits, error) {
	contract, err := bindPoolBaseTVLLimits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimits{PoolBaseTVLLimitsCaller: PoolBaseTVLLimitsCaller{contract: contract}, PoolBaseTVLLimitsTransactor: PoolBaseTVLLimitsTransactor{contract: contract}, PoolBaseTVLLimitsFilterer: PoolBaseTVLLimitsFilterer{contract: contract}}, nil
}

// NewPoolBaseTVLLimitsCaller creates a new read-only instance of PoolBaseTVLLimits, bound to a specific deployed contract.
func NewPoolBaseTVLLimitsCaller(address common.Address, caller bind.ContractCaller) (*PoolBaseTVLLimitsCaller, error) {
	contract, err := bindPoolBaseTVLLimits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsCaller{contract: contract}, nil
}

// NewPoolBaseTVLLimitsTransactor creates a new write-only instance of PoolBaseTVLLimits, bound to a specific deployed contract.
func NewPoolBaseTVLLimitsTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolBaseTVLLimitsTransactor, error) {
	contract, err := bindPoolBaseTVLLimits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsTransactor{contract: contract}, nil
}

// NewPoolBaseTVLLimitsFilterer creates a new log filterer instance of PoolBaseTVLLimits, bound to a specific deployed contract.
func NewPoolBaseTVLLimitsFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolBaseTVLLimitsFilterer, error) {
	contract, err := bindPoolBaseTVLLimits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsFilterer{contract: contract}, nil
}

// bindPoolBaseTVLLimits binds a generic wrapper to an already deployed contract.
func bindPoolBaseTVLLimits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolBaseTVLLimitsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolBaseTVLLimits.Contract.PoolBaseTVLLimitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.PoolBaseTVLLimitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.PoolBaseTVLLimitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolBaseTVLLimits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Explanation() (string, error) {
	return _PoolBaseTVLLimits.Contract.Explanation(&_PoolBaseTVLLimits.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) Explanation() (string, error) {
	return _PoolBaseTVLLimits.Contract.Explanation(&_PoolBaseTVLLimits.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) GetTVLLimits(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "getTVLLimits")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _PoolBaseTVLLimits.Contract.GetTVLLimits(&_PoolBaseTVLLimits.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _PoolBaseTVLLimits.Contract.GetTVLLimits(&_PoolBaseTVLLimits.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) MaxPerDeposit() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.MaxPerDeposit(&_PoolBaseTVLLimits.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.MaxPerDeposit(&_PoolBaseTVLLimits.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) MaxTotalDeposits() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.MaxTotalDeposits(&_PoolBaseTVLLimits.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.MaxTotalDeposits(&_PoolBaseTVLLimits.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Paused(index uint8) (bool, error) {
	return _PoolBaseTVLLimits.Contract.Paused(&_PoolBaseTVLLimits.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) Paused(index uint8) (bool, error) {
	return _PoolBaseTVLLimits.Contract.Paused(&_PoolBaseTVLLimits.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Paused0() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.Paused0(&_PoolBaseTVLLimits.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) Paused0() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.Paused0(&_PoolBaseTVLLimits.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) PauserRegistry() (common.Address, error) {
	return _PoolBaseTVLLimits.Contract.PauserRegistry(&_PoolBaseTVLLimits.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) PauserRegistry() (common.Address, error) {
	return _PoolBaseTVLLimits.Contract.PauserRegistry(&_PoolBaseTVLLimits.CallOpts)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) PoolController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "poolController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) PoolController() (common.Address, error) {
	return _PoolBaseTVLLimits.Contract.PoolController(&_PoolBaseTVLLimits.CallOpts)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) PoolController() (common.Address, error) {
	return _PoolBaseTVLLimits.Contract.PoolController(&_PoolBaseTVLLimits.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Shares(user common.Address) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.Shares(&_PoolBaseTVLLimits.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.Shares(&_PoolBaseTVLLimits.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.SharesToUnderlying(&_PoolBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.SharesToUnderlying(&_PoolBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.SharesToUnderlyingView(&_PoolBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.SharesToUnderlyingView(&_PoolBaseTVLLimits.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) TotalShares() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.TotalShares(&_PoolBaseTVLLimits.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) TotalShares() (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.TotalShares(&_PoolBaseTVLLimits.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.UnderlyingToShares(&_PoolBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.UnderlyingToShares(&_PoolBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.UnderlyingToSharesView(&_PoolBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.UnderlyingToSharesView(&_PoolBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) UnderlyingToken() (common.Address, error) {
	return _PoolBaseTVLLimits.Contract.UnderlyingToken(&_PoolBaseTVLLimits.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) UnderlyingToken() (common.Address, error) {
	return _PoolBaseTVLLimits.Contract.UnderlyingToken(&_PoolBaseTVLLimits.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolBaseTVLLimits.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.UserUnderlyingView(&_PoolBaseTVLLimits.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _PoolBaseTVLLimits.Contract.UserUnderlyingView(&_PoolBaseTVLLimits.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Deposit(&_PoolBaseTVLLimits.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Deposit(&_PoolBaseTVLLimits.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) Initialize(opts *bind.TransactOpts, _maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "initialize", _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Initialize(&_PoolBaseTVLLimits.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Initialize(&_PoolBaseTVLLimits.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// InitializeBase is a paid mutator transaction binding the contract method 0xc4e01fa4.
//
// Solidity: function initializeBase(address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) InitializeBase(opts *bind.TransactOpts, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "initializeBase", _underlyingToken, _pauserRegistry)
}

// InitializeBase is a paid mutator transaction binding the contract method 0xc4e01fa4.
//
// Solidity: function initializeBase(address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) InitializeBase(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.InitializeBase(&_PoolBaseTVLLimits.TransactOpts, _underlyingToken, _pauserRegistry)
}

// InitializeBase is a paid mutator transaction binding the contract method 0xc4e01fa4.
//
// Solidity: function initializeBase(address _underlyingToken, address _pauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) InitializeBase(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.InitializeBase(&_PoolBaseTVLLimits.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Pause(&_PoolBaseTVLLimits.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Pause(&_PoolBaseTVLLimits.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) PauseAll() (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.PauseAll(&_PoolBaseTVLLimits.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) PauseAll() (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.PauseAll(&_PoolBaseTVLLimits.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.SetPauserRegistry(&_PoolBaseTVLLimits.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.SetPauserRegistry(&_PoolBaseTVLLimits.TransactOpts, newPauserRegistry)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) SetTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "setTVLLimits", newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.SetTVLLimits(&_PoolBaseTVLLimits.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.SetTVLLimits(&_PoolBaseTVLLimits.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Unpause(&_PoolBaseTVLLimits.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Unpause(&_PoolBaseTVLLimits.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.UserUnderlying(&_PoolBaseTVLLimits.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.UserUnderlying(&_PoolBaseTVLLimits.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Withdraw(&_PoolBaseTVLLimits.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _PoolBaseTVLLimits.Contract.Withdraw(&_PoolBaseTVLLimits.TransactOpts, recipient, token, amountShares)
}

// PoolBaseTVLLimitsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsInitializedIterator struct {
	Event *PoolBaseTVLLimitsInitialized // Event containing the contract specifics and raw log

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
func (it *PoolBaseTVLLimitsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseTVLLimitsInitialized)
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
		it.Event = new(PoolBaseTVLLimitsInitialized)
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
func (it *PoolBaseTVLLimitsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseTVLLimitsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseTVLLimitsInitialized represents a Initialized event raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolBaseTVLLimitsInitializedIterator, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsInitializedIterator{contract: _PoolBaseTVLLimits.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolBaseTVLLimitsInitialized) (event.Subscription, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseTVLLimitsInitialized)
				if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) ParseInitialized(log types.Log) (*PoolBaseTVLLimitsInitialized, error) {
	event := new(PoolBaseTVLLimitsInitialized)
	if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBaseTVLLimitsMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsMaxPerDepositUpdatedIterator struct {
	Event *PoolBaseTVLLimitsMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *PoolBaseTVLLimitsMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseTVLLimitsMaxPerDepositUpdated)
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
		it.Event = new(PoolBaseTVLLimitsMaxPerDepositUpdated)
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
func (it *PoolBaseTVLLimitsMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseTVLLimitsMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseTVLLimitsMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*PoolBaseTVLLimitsMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsMaxPerDepositUpdatedIterator{contract: _PoolBaseTVLLimits.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *PoolBaseTVLLimitsMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseTVLLimitsMaxPerDepositUpdated)
				if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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

// ParseMaxPerDepositUpdated is a log parse operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) ParseMaxPerDepositUpdated(log types.Log) (*PoolBaseTVLLimitsMaxPerDepositUpdated, error) {
	event := new(PoolBaseTVLLimitsMaxPerDepositUpdated)
	if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator struct {
	Event *PoolBaseTVLLimitsMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseTVLLimitsMaxTotalDepositsUpdated)
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
		it.Event = new(PoolBaseTVLLimitsMaxTotalDepositsUpdated)
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
func (it *PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseTVLLimitsMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsMaxTotalDepositsUpdatedIterator{contract: _PoolBaseTVLLimits.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *PoolBaseTVLLimitsMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseTVLLimitsMaxTotalDepositsUpdated)
				if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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

// ParseMaxTotalDepositsUpdated is a log parse operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*PoolBaseTVLLimitsMaxTotalDepositsUpdated, error) {
	event := new(PoolBaseTVLLimitsMaxTotalDepositsUpdated)
	if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBaseTVLLimitsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsPausedIterator struct {
	Event *PoolBaseTVLLimitsPaused // Event containing the contract specifics and raw log

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
func (it *PoolBaseTVLLimitsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseTVLLimitsPaused)
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
		it.Event = new(PoolBaseTVLLimitsPaused)
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
func (it *PoolBaseTVLLimitsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseTVLLimitsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseTVLLimitsPaused represents a Paused event raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*PoolBaseTVLLimitsPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBaseTVLLimits.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsPausedIterator{contract: _PoolBaseTVLLimits.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolBaseTVLLimitsPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBaseTVLLimits.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseTVLLimitsPaused)
				if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) ParsePaused(log types.Log) (*PoolBaseTVLLimitsPaused, error) {
	event := new(PoolBaseTVLLimitsPaused)
	if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBaseTVLLimitsPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsPauserRegistrySetIterator struct {
	Event *PoolBaseTVLLimitsPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *PoolBaseTVLLimitsPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseTVLLimitsPauserRegistrySet)
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
		it.Event = new(PoolBaseTVLLimitsPauserRegistrySet)
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
func (it *PoolBaseTVLLimitsPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseTVLLimitsPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseTVLLimitsPauserRegistrySet represents a PauserRegistrySet event raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*PoolBaseTVLLimitsPauserRegistrySetIterator, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsPauserRegistrySetIterator{contract: _PoolBaseTVLLimits.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *PoolBaseTVLLimitsPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _PoolBaseTVLLimits.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseTVLLimitsPauserRegistrySet)
				if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) ParsePauserRegistrySet(log types.Log) (*PoolBaseTVLLimitsPauserRegistrySet, error) {
	event := new(PoolBaseTVLLimitsPauserRegistrySet)
	if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBaseTVLLimitsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsUnpausedIterator struct {
	Event *PoolBaseTVLLimitsUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolBaseTVLLimitsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBaseTVLLimitsUnpaused)
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
		it.Event = new(PoolBaseTVLLimitsUnpaused)
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
func (it *PoolBaseTVLLimitsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBaseTVLLimitsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBaseTVLLimitsUnpaused represents a Unpaused event raised by the PoolBaseTVLLimits contract.
type PoolBaseTVLLimitsUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*PoolBaseTVLLimitsUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBaseTVLLimits.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolBaseTVLLimitsUnpausedIterator{contract: _PoolBaseTVLLimits.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolBaseTVLLimitsUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolBaseTVLLimits.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBaseTVLLimitsUnpaused)
				if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PoolBaseTVLLimits *PoolBaseTVLLimitsFilterer) ParseUnpaused(log types.Log) (*PoolBaseTVLLimitsUnpaused, error) {
	event := new(PoolBaseTVLLimitsUnpaused)
	if err := _PoolBaseTVLLimits.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
