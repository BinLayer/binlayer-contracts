// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Errors

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

// ErrorsMetaData contains all meta data concerning the Errors contract.
var ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ACTION_NOT_IN_QUEUE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ALREADY_DELEGATED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"AMOUNT_CANNOT_BE_ZERO\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"AMOUNT_TOO_LARGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ARRAY_LENGTH_MISMATCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CALLER_CANNOT_UNDELEGATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CALLER_NOT_OPERATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANNOT_BE_IN_THE_FEATURE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANNOT_CANCEL_SPENT_SALT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANNOT_UNDELEGATE_OPERATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONTRACT_PAUSED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CUMULATIVE_EARNINGS_MUST_BE_GREATER_THAN_CUMULATIVE_LAIMED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DECREASE_OPT_OUT_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_EXCEEDS_MAX_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_ONLY_UNDERLYING_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DURATION_EXCEEDS_MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DURATION_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INDEX_PAUSED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIALIZE_ONCE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INPUT_LENGTH_MISMATCH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_EARNER_CLAIM_PROOF\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_EARNER_LEAF_INDEX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_PAUSE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_POOL_CONSIDERED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_ROOT_INDEX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_TOKEN_CLAIM_PROOF\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_TOKEN_LEAF_INDEX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INVALID_UNPAUSE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_DEPOSITS_EXCEEDED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PER_DEPOSIT_EXCEEDS_MAX_TOTAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PER_DEPOSIT_LIMIT_EXCEEDED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_WITHDRAWAL_DELAY_EXCEEDS_MAX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_WITHDRAWAL_DELAY_NOT_PASSED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NEW_ROOT_MUST_BE_NEWER_CALCULATED_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_CLAIMER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_CREATE_REWARDS_FOR_ALL_SUBMITTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_DELEGATION_CONTROLLER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_PAUSER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_POOL_WHITELISTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_REGISTERED_IN_BINLAYER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_REWARDS_UPDATER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOT_UNPAUSER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NO_POOLS_SET\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ONLY_POOL_CONTROLLER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ONLY_SUPPORT_WRAPPED_TOKEN_POOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ONLY_WITHDRAWER_CAN_COMPLETE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ALREADY_REGISTERED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPT_OUT_WINDOW_EXCEEDS_MAX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOLS_CANNOT_BE_EMPTY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOLS_MUST_BE_ASCENDING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL_NOT_FOUND\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL_NOT_WHITELISTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL_WITHDRAWAL_DELAY_EXCEEDS_MAX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROOT_ALREADY_ACTIVATED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROOT_ALREADY_DISABLED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROOT_IS_DISABLED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROOT_NOT_ACTIVATED_YET\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROOT_NOT_FOUND\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SALT_ALREADY_SPENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SHARE_AMOUNT_TOO_HIGH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIGNATURE_EXPIRED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_MUST_BE_DELEGATED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STARTTIMESTAMP_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STARTTIMESTAMP_TOO_FAR_IN_THE_FUTURE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STARTTIMESTAMP_TOO_FAR_IN_THE_PAST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THIRD_PARTY_TRANSFERS_DISABLED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THIRD_PARTY_TRANSFERS_FORBIDDEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWAL_DELAY_NOT_PASSED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_MUST_BE_STAKER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_NOT_STAKER_OR_GATEWAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_AMOUNT_SHARES_TOO_HIGH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_ONLY_UNDERLYING_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ZERO_ADDRESS_NOT_VALID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ZERO_SHARES_NOT_VALID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6116fd61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106104e85760003560e01c806388adebd211610290578063ca2633cd11610175578063d98555a1116100ed578063f059972f116100a1578063f10c942111610086578063f10c9421146115a7578063fb765715146115e3578063ff9a0df71461161f57600080fd5b8063f059972f1461152f578063f086b2911461156b57600080fd5b8063eb565197116100d2578063eb5651971461147b578063ee6d21aa146114b7578063efa2ab10146114f357600080fd5b8063d98555a114611403578063eb4d08461461143f57600080fd5b8063d1b25b3111610144578063d2b3b17611610129578063d2b3b1761461134f578063d577ae141461138b578063d6de421d146113c757600080fd5b8063d1b25b31146112d7578063d28c87f21461131357600080fd5b8063ca2633cd146111e7578063ca5768c214611223578063cdaff8141461125f578063d14bb17a1461129b57600080fd5b8063a41cf69f11610208578063b6f7ae43116101d7578063bf71cbd2116101bc578063bf71cbd214611133578063c3fcd8531461116f578063c619c91f146111ab57600080fd5b8063b6f7ae43146110bb578063b92066b2146110f757600080fd5b8063a41cf69f14610fcb578063a57751fe14611007578063a7651b9814611043578063ae7a3f6d1461107f57600080fd5b806395714f7b1161025f5780639c06be3e116102445780639c06be3e14610f175780639f840bfb14610f53578063a0afa0cc14610f8f57600080fd5b806395714f7b14610e9f57806399adb1c714610edb57600080fd5b806388adebd214610daf5780638a30162414610deb5780638c7fc1b214610e27578063941abc7714610e6357600080fd5b80634b5f220a116103d1578063702025981161034957806378833b66116102fd5780638520044e116102e25780638520044e14610cfb578063872f365214610d37578063884b210c14610d7357600080fd5b806378833b6614610c835780637965672514610cbf57600080fd5b806371925c671161032e57806371925c6714610bcf578063733a577d14610c0b57806375bfb89414610c4757600080fd5b80637020259814610b575780637054e88714610b9357600080fd5b80635b33eb0e116103a057806367fd5c911161038557806367fd5c9114610aa357806369bba99814610adf5780636b2e1e6114610b1b57600080fd5b80635b33eb0e14610a2b578063606980c014610a6757600080fd5b80634b5f220a1461093b5780634c4c015b146109775780634e729aee146109b357806355b533a6146109ef57600080fd5b8063193198271161046457806331a6fb971161043357806342170b851161041857806342170b8514610887578063490ff719146108c35780634ae9ddfa146108ff57600080fd5b806331a6fb971461080f57806332635eba1461084b57600080fd5b8063193198271461071f57806319ff7c781461075b5780632ec313d3146107975780632ec82ef4146107d357600080fd5b80630b8cd592116104bb5780630ee30d10116104a05780630ee30d101461066b57806315fbec3b146106a75780631863ca35146106e357600080fd5b80630b8cd592146105f35780630cc6ae8a1461062f57600080fd5b806301de8c6d146104ed578063059bef991461053f57806307ae55ba1461057b5780630b8c0aa1146105b7575b600080fd5b6105296040518060400160405280600281526020017f313800000000000000000000000000000000000000000000000000000000000081525081565b604051610536919061165b565b60405180910390f35b6105296040518060400160405280600281526020017f363000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313600000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363800000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353700000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363500000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343700000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353500000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f373100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343500000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313900000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f373000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f380000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f373200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313500000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323800000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f340000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343600000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363600000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323600000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333600000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333800000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353900000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343800000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343200000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333100000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363700000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f360000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f320000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f390000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353600000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f330000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f313000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f363900000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343900000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323900000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333500000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333900000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f343000000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f373300000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f333700000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323700000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f373400000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f323500000000000000000000000000000000000000000000000000000000000081525081565b6105296040518060400160405280600281526020017f353800000000000000000000000000000000000000000000000000000000000081525081565b600060208083528351808285015260005b818110156116885785810183015185820160400152820161166c565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea2646970667358221220e13df18507562d0a3d6f12e6245aaf6c80427bda0978a4626def96bf1d0393a764736f6c63430008140033",
}

// ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorsMetaData.ABI instead.
var ErrorsABI = ErrorsMetaData.ABI

// ErrorsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ErrorsMetaData.Bin instead.
var ErrorsBin = ErrorsMetaData.Bin

// DeployErrors deploys a new Ethereum contract, binding an instance of Errors to it.
func DeployErrors(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Errors, error) {
	parsed, err := ErrorsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ErrorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// Errors is an auto generated Go binding around an Ethereum contract.
type Errors struct {
	ErrorsCaller     // Read-only binding to the contract
	ErrorsTransactor // Write-only binding to the contract
	ErrorsFilterer   // Log filterer for contract events
}

// ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorsSession struct {
	Contract     *Errors           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorsCallerSession struct {
	Contract *ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorsTransactorSession struct {
	Contract     *ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorsRaw struct {
	Contract *Errors // Generic contract binding to access the raw methods on
}

// ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorsCallerRaw struct {
	Contract *ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorsTransactorRaw struct {
	Contract *ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrors creates a new instance of Errors, bound to a specific deployed contract.
func NewErrors(address common.Address, backend bind.ContractBackend) (*Errors, error) {
	contract, err := bindErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// NewErrorsCaller creates a new read-only instance of Errors, bound to a specific deployed contract.
func NewErrorsCaller(address common.Address, caller bind.ContractCaller) (*ErrorsCaller, error) {
	contract, err := bindErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsCaller{contract: contract}, nil
}

// NewErrorsTransactor creates a new write-only instance of Errors, bound to a specific deployed contract.
func NewErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorsTransactor, error) {
	contract, err := bindErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsTransactor{contract: contract}, nil
}

// NewErrorsFilterer creates a new log filterer instance of Errors, bound to a specific deployed contract.
func NewErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorsFilterer, error) {
	contract, err := bindErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorsFilterer{contract: contract}, nil
}

// bindErrors binds a generic wrapper to an already deployed contract.
func bindErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transact(opts, method, params...)
}

// ACTIONNOTINQUEUE is a free data retrieval call binding the contract method 0x872f3652.
//
// Solidity: function ACTION_NOT_IN_QUEUE() view returns(string)
func (_Errors *ErrorsCaller) ACTIONNOTINQUEUE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ACTION_NOT_IN_QUEUE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ACTIONNOTINQUEUE is a free data retrieval call binding the contract method 0x872f3652.
//
// Solidity: function ACTION_NOT_IN_QUEUE() view returns(string)
func (_Errors *ErrorsSession) ACTIONNOTINQUEUE() (string, error) {
	return _Errors.Contract.ACTIONNOTINQUEUE(&_Errors.CallOpts)
}

// ACTIONNOTINQUEUE is a free data retrieval call binding the contract method 0x872f3652.
//
// Solidity: function ACTION_NOT_IN_QUEUE() view returns(string)
func (_Errors *ErrorsCallerSession) ACTIONNOTINQUEUE() (string, error) {
	return _Errors.Contract.ACTIONNOTINQUEUE(&_Errors.CallOpts)
}

// ALREADYDELEGATED is a free data retrieval call binding the contract method 0xd6de421d.
//
// Solidity: function ALREADY_DELEGATED() view returns(string)
func (_Errors *ErrorsCaller) ALREADYDELEGATED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ALREADY_DELEGATED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ALREADYDELEGATED is a free data retrieval call binding the contract method 0xd6de421d.
//
// Solidity: function ALREADY_DELEGATED() view returns(string)
func (_Errors *ErrorsSession) ALREADYDELEGATED() (string, error) {
	return _Errors.Contract.ALREADYDELEGATED(&_Errors.CallOpts)
}

// ALREADYDELEGATED is a free data retrieval call binding the contract method 0xd6de421d.
//
// Solidity: function ALREADY_DELEGATED() view returns(string)
func (_Errors *ErrorsCallerSession) ALREADYDELEGATED() (string, error) {
	return _Errors.Contract.ALREADYDELEGATED(&_Errors.CallOpts)
}

// AMOUNTCANNOTBEZERO is a free data retrieval call binding the contract method 0x9f840bfb.
//
// Solidity: function AMOUNT_CANNOT_BE_ZERO() view returns(string)
func (_Errors *ErrorsCaller) AMOUNTCANNOTBEZERO(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AMOUNT_CANNOT_BE_ZERO")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AMOUNTCANNOTBEZERO is a free data retrieval call binding the contract method 0x9f840bfb.
//
// Solidity: function AMOUNT_CANNOT_BE_ZERO() view returns(string)
func (_Errors *ErrorsSession) AMOUNTCANNOTBEZERO() (string, error) {
	return _Errors.Contract.AMOUNTCANNOTBEZERO(&_Errors.CallOpts)
}

// AMOUNTCANNOTBEZERO is a free data retrieval call binding the contract method 0x9f840bfb.
//
// Solidity: function AMOUNT_CANNOT_BE_ZERO() view returns(string)
func (_Errors *ErrorsCallerSession) AMOUNTCANNOTBEZERO() (string, error) {
	return _Errors.Contract.AMOUNTCANNOTBEZERO(&_Errors.CallOpts)
}

// AMOUNTTOOLARGE is a free data retrieval call binding the contract method 0x059bef99.
//
// Solidity: function AMOUNT_TOO_LARGE() view returns(string)
func (_Errors *ErrorsCaller) AMOUNTTOOLARGE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "AMOUNT_TOO_LARGE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AMOUNTTOOLARGE is a free data retrieval call binding the contract method 0x059bef99.
//
// Solidity: function AMOUNT_TOO_LARGE() view returns(string)
func (_Errors *ErrorsSession) AMOUNTTOOLARGE() (string, error) {
	return _Errors.Contract.AMOUNTTOOLARGE(&_Errors.CallOpts)
}

// AMOUNTTOOLARGE is a free data retrieval call binding the contract method 0x059bef99.
//
// Solidity: function AMOUNT_TOO_LARGE() view returns(string)
func (_Errors *ErrorsCallerSession) AMOUNTTOOLARGE() (string, error) {
	return _Errors.Contract.AMOUNTTOOLARGE(&_Errors.CallOpts)
}

// ARRAYLENGTHMISMATCH is a free data retrieval call binding the contract method 0x88adebd2.
//
// Solidity: function ARRAY_LENGTH_MISMATCH() view returns(string)
func (_Errors *ErrorsCaller) ARRAYLENGTHMISMATCH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ARRAY_LENGTH_MISMATCH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ARRAYLENGTHMISMATCH is a free data retrieval call binding the contract method 0x88adebd2.
//
// Solidity: function ARRAY_LENGTH_MISMATCH() view returns(string)
func (_Errors *ErrorsSession) ARRAYLENGTHMISMATCH() (string, error) {
	return _Errors.Contract.ARRAYLENGTHMISMATCH(&_Errors.CallOpts)
}

// ARRAYLENGTHMISMATCH is a free data retrieval call binding the contract method 0x88adebd2.
//
// Solidity: function ARRAY_LENGTH_MISMATCH() view returns(string)
func (_Errors *ErrorsCallerSession) ARRAYLENGTHMISMATCH() (string, error) {
	return _Errors.Contract.ARRAYLENGTHMISMATCH(&_Errors.CallOpts)
}

// CALLERCANNOTUNDELEGATE is a free data retrieval call binding the contract method 0x19ff7c78.
//
// Solidity: function CALLER_CANNOT_UNDELEGATE() view returns(string)
func (_Errors *ErrorsCaller) CALLERCANNOTUNDELEGATE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CALLER_CANNOT_UNDELEGATE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CALLERCANNOTUNDELEGATE is a free data retrieval call binding the contract method 0x19ff7c78.
//
// Solidity: function CALLER_CANNOT_UNDELEGATE() view returns(string)
func (_Errors *ErrorsSession) CALLERCANNOTUNDELEGATE() (string, error) {
	return _Errors.Contract.CALLERCANNOTUNDELEGATE(&_Errors.CallOpts)
}

// CALLERCANNOTUNDELEGATE is a free data retrieval call binding the contract method 0x19ff7c78.
//
// Solidity: function CALLER_CANNOT_UNDELEGATE() view returns(string)
func (_Errors *ErrorsCallerSession) CALLERCANNOTUNDELEGATE() (string, error) {
	return _Errors.Contract.CALLERCANNOTUNDELEGATE(&_Errors.CallOpts)
}

// CALLERNOTOPERATOR is a free data retrieval call binding the contract method 0x42170b85.
//
// Solidity: function CALLER_NOT_OPERATOR() view returns(string)
func (_Errors *ErrorsCaller) CALLERNOTOPERATOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CALLER_NOT_OPERATOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CALLERNOTOPERATOR is a free data retrieval call binding the contract method 0x42170b85.
//
// Solidity: function CALLER_NOT_OPERATOR() view returns(string)
func (_Errors *ErrorsSession) CALLERNOTOPERATOR() (string, error) {
	return _Errors.Contract.CALLERNOTOPERATOR(&_Errors.CallOpts)
}

// CALLERNOTOPERATOR is a free data retrieval call binding the contract method 0x42170b85.
//
// Solidity: function CALLER_NOT_OPERATOR() view returns(string)
func (_Errors *ErrorsCallerSession) CALLERNOTOPERATOR() (string, error) {
	return _Errors.Contract.CALLERNOTOPERATOR(&_Errors.CallOpts)
}

// CANNOTBEINTHEFEATURE is a free data retrieval call binding the contract method 0x4c4c015b.
//
// Solidity: function CANNOT_BE_IN_THE_FEATURE() view returns(string)
func (_Errors *ErrorsCaller) CANNOTBEINTHEFEATURE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CANNOT_BE_IN_THE_FEATURE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CANNOTBEINTHEFEATURE is a free data retrieval call binding the contract method 0x4c4c015b.
//
// Solidity: function CANNOT_BE_IN_THE_FEATURE() view returns(string)
func (_Errors *ErrorsSession) CANNOTBEINTHEFEATURE() (string, error) {
	return _Errors.Contract.CANNOTBEINTHEFEATURE(&_Errors.CallOpts)
}

// CANNOTBEINTHEFEATURE is a free data retrieval call binding the contract method 0x4c4c015b.
//
// Solidity: function CANNOT_BE_IN_THE_FEATURE() view returns(string)
func (_Errors *ErrorsCallerSession) CANNOTBEINTHEFEATURE() (string, error) {
	return _Errors.Contract.CANNOTBEINTHEFEATURE(&_Errors.CallOpts)
}

// CANNOTCANCELSPENTSALT is a free data retrieval call binding the contract method 0xa41cf69f.
//
// Solidity: function CANNOT_CANCEL_SPENT_SALT() view returns(string)
func (_Errors *ErrorsCaller) CANNOTCANCELSPENTSALT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CANNOT_CANCEL_SPENT_SALT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CANNOTCANCELSPENTSALT is a free data retrieval call binding the contract method 0xa41cf69f.
//
// Solidity: function CANNOT_CANCEL_SPENT_SALT() view returns(string)
func (_Errors *ErrorsSession) CANNOTCANCELSPENTSALT() (string, error) {
	return _Errors.Contract.CANNOTCANCELSPENTSALT(&_Errors.CallOpts)
}

// CANNOTCANCELSPENTSALT is a free data retrieval call binding the contract method 0xa41cf69f.
//
// Solidity: function CANNOT_CANCEL_SPENT_SALT() view returns(string)
func (_Errors *ErrorsCallerSession) CANNOTCANCELSPENTSALT() (string, error) {
	return _Errors.Contract.CANNOTCANCELSPENTSALT(&_Errors.CallOpts)
}

// CANNOTUNDELEGATEOPERATOR is a free data retrieval call binding the contract method 0x75bfb894.
//
// Solidity: function CANNOT_UNDELEGATE_OPERATOR() view returns(string)
func (_Errors *ErrorsCaller) CANNOTUNDELEGATEOPERATOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CANNOT_UNDELEGATE_OPERATOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CANNOTUNDELEGATEOPERATOR is a free data retrieval call binding the contract method 0x75bfb894.
//
// Solidity: function CANNOT_UNDELEGATE_OPERATOR() view returns(string)
func (_Errors *ErrorsSession) CANNOTUNDELEGATEOPERATOR() (string, error) {
	return _Errors.Contract.CANNOTUNDELEGATEOPERATOR(&_Errors.CallOpts)
}

// CANNOTUNDELEGATEOPERATOR is a free data retrieval call binding the contract method 0x75bfb894.
//
// Solidity: function CANNOT_UNDELEGATE_OPERATOR() view returns(string)
func (_Errors *ErrorsCallerSession) CANNOTUNDELEGATEOPERATOR() (string, error) {
	return _Errors.Contract.CANNOTUNDELEGATEOPERATOR(&_Errors.CallOpts)
}

// CONTRACTPAUSED is a free data retrieval call binding the contract method 0x884b210c.
//
// Solidity: function CONTRACT_PAUSED() view returns(string)
func (_Errors *ErrorsCaller) CONTRACTPAUSED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CONTRACT_PAUSED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CONTRACTPAUSED is a free data retrieval call binding the contract method 0x884b210c.
//
// Solidity: function CONTRACT_PAUSED() view returns(string)
func (_Errors *ErrorsSession) CONTRACTPAUSED() (string, error) {
	return _Errors.Contract.CONTRACTPAUSED(&_Errors.CallOpts)
}

// CONTRACTPAUSED is a free data retrieval call binding the contract method 0x884b210c.
//
// Solidity: function CONTRACT_PAUSED() view returns(string)
func (_Errors *ErrorsCallerSession) CONTRACTPAUSED() (string, error) {
	return _Errors.Contract.CONTRACTPAUSED(&_Errors.CallOpts)
}

// CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED is a free data retrieval call binding the contract method 0x8520044e.
//
// Solidity: function CUMULATIVE_EARNINGS_MUST_BE_GREATER_THAN_CUMULATIVE_LAIMED() view returns(string)
func (_Errors *ErrorsCaller) CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CUMULATIVE_EARNINGS_MUST_BE_GREATER_THAN_CUMULATIVE_LAIMED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED is a free data retrieval call binding the contract method 0x8520044e.
//
// Solidity: function CUMULATIVE_EARNINGS_MUST_BE_GREATER_THAN_CUMULATIVE_LAIMED() view returns(string)
func (_Errors *ErrorsSession) CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED() (string, error) {
	return _Errors.Contract.CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED(&_Errors.CallOpts)
}

// CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED is a free data retrieval call binding the contract method 0x8520044e.
//
// Solidity: function CUMULATIVE_EARNINGS_MUST_BE_GREATER_THAN_CUMULATIVE_LAIMED() view returns(string)
func (_Errors *ErrorsCallerSession) CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED() (string, error) {
	return _Errors.Contract.CUMULATIVEEARNINGSMUSTBEGREATERTHANCUMULATIVELAIMED(&_Errors.CallOpts)
}

// DECREASEOPTOUTWINDOW is a free data retrieval call binding the contract method 0x79656725.
//
// Solidity: function DECREASE_OPT_OUT_WINDOW() view returns(string)
func (_Errors *ErrorsCaller) DECREASEOPTOUTWINDOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DECREASE_OPT_OUT_WINDOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DECREASEOPTOUTWINDOW is a free data retrieval call binding the contract method 0x79656725.
//
// Solidity: function DECREASE_OPT_OUT_WINDOW() view returns(string)
func (_Errors *ErrorsSession) DECREASEOPTOUTWINDOW() (string, error) {
	return _Errors.Contract.DECREASEOPTOUTWINDOW(&_Errors.CallOpts)
}

// DECREASEOPTOUTWINDOW is a free data retrieval call binding the contract method 0x79656725.
//
// Solidity: function DECREASE_OPT_OUT_WINDOW() view returns(string)
func (_Errors *ErrorsCallerSession) DECREASEOPTOUTWINDOW() (string, error) {
	return _Errors.Contract.DECREASEOPTOUTWINDOW(&_Errors.CallOpts)
}

// DEPOSITEXCEEDSMAXLENGTH is a free data retrieval call binding the contract method 0x0b8cd592.
//
// Solidity: function DEPOSIT_EXCEEDS_MAX_LENGTH() view returns(string)
func (_Errors *ErrorsCaller) DEPOSITEXCEEDSMAXLENGTH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DEPOSIT_EXCEEDS_MAX_LENGTH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DEPOSITEXCEEDSMAXLENGTH is a free data retrieval call binding the contract method 0x0b8cd592.
//
// Solidity: function DEPOSIT_EXCEEDS_MAX_LENGTH() view returns(string)
func (_Errors *ErrorsSession) DEPOSITEXCEEDSMAXLENGTH() (string, error) {
	return _Errors.Contract.DEPOSITEXCEEDSMAXLENGTH(&_Errors.CallOpts)
}

// DEPOSITEXCEEDSMAXLENGTH is a free data retrieval call binding the contract method 0x0b8cd592.
//
// Solidity: function DEPOSIT_EXCEEDS_MAX_LENGTH() view returns(string)
func (_Errors *ErrorsCallerSession) DEPOSITEXCEEDSMAXLENGTH() (string, error) {
	return _Errors.Contract.DEPOSITEXCEEDSMAXLENGTH(&_Errors.CallOpts)
}

// DEPOSITONLYUNDERLYINGTOKEN is a free data retrieval call binding the contract method 0x7054e887.
//
// Solidity: function DEPOSIT_ONLY_UNDERLYING_TOKEN() view returns(string)
func (_Errors *ErrorsCaller) DEPOSITONLYUNDERLYINGTOKEN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DEPOSIT_ONLY_UNDERLYING_TOKEN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DEPOSITONLYUNDERLYINGTOKEN is a free data retrieval call binding the contract method 0x7054e887.
//
// Solidity: function DEPOSIT_ONLY_UNDERLYING_TOKEN() view returns(string)
func (_Errors *ErrorsSession) DEPOSITONLYUNDERLYINGTOKEN() (string, error) {
	return _Errors.Contract.DEPOSITONLYUNDERLYINGTOKEN(&_Errors.CallOpts)
}

// DEPOSITONLYUNDERLYINGTOKEN is a free data retrieval call binding the contract method 0x7054e887.
//
// Solidity: function DEPOSIT_ONLY_UNDERLYING_TOKEN() view returns(string)
func (_Errors *ErrorsCallerSession) DEPOSITONLYUNDERLYINGTOKEN() (string, error) {
	return _Errors.Contract.DEPOSITONLYUNDERLYINGTOKEN(&_Errors.CallOpts)
}

// DURATIONEXCEEDSMAXREWARDSDURATION is a free data retrieval call binding the contract method 0x6b2e1e61.
//
// Solidity: function DURATION_EXCEEDS_MAX_REWARDS_DURATION() view returns(string)
func (_Errors *ErrorsCaller) DURATIONEXCEEDSMAXREWARDSDURATION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DURATION_EXCEEDS_MAX_REWARDS_DURATION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DURATIONEXCEEDSMAXREWARDSDURATION is a free data retrieval call binding the contract method 0x6b2e1e61.
//
// Solidity: function DURATION_EXCEEDS_MAX_REWARDS_DURATION() view returns(string)
func (_Errors *ErrorsSession) DURATIONEXCEEDSMAXREWARDSDURATION() (string, error) {
	return _Errors.Contract.DURATIONEXCEEDSMAXREWARDSDURATION(&_Errors.CallOpts)
}

// DURATIONEXCEEDSMAXREWARDSDURATION is a free data retrieval call binding the contract method 0x6b2e1e61.
//
// Solidity: function DURATION_EXCEEDS_MAX_REWARDS_DURATION() view returns(string)
func (_Errors *ErrorsCallerSession) DURATIONEXCEEDSMAXREWARDSDURATION() (string, error) {
	return _Errors.Contract.DURATIONEXCEEDSMAXREWARDSDURATION(&_Errors.CallOpts)
}

// DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x606980c0.
//
// Solidity: function DURATION_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS() view returns(string)
func (_Errors *ErrorsCaller) DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "DURATION_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x606980c0.
//
// Solidity: function DURATION_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS() view returns(string)
func (_Errors *ErrorsSession) DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS() (string, error) {
	return _Errors.Contract.DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS(&_Errors.CallOpts)
}

// DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x606980c0.
//
// Solidity: function DURATION_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS() view returns(string)
func (_Errors *ErrorsCallerSession) DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS() (string, error) {
	return _Errors.Contract.DURATIONMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS(&_Errors.CallOpts)
}

// INDEXPAUSED is a free data retrieval call binding the contract method 0x4ae9ddfa.
//
// Solidity: function INDEX_PAUSED() view returns(string)
func (_Errors *ErrorsCaller) INDEXPAUSED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INDEX_PAUSED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INDEXPAUSED is a free data retrieval call binding the contract method 0x4ae9ddfa.
//
// Solidity: function INDEX_PAUSED() view returns(string)
func (_Errors *ErrorsSession) INDEXPAUSED() (string, error) {
	return _Errors.Contract.INDEXPAUSED(&_Errors.CallOpts)
}

// INDEXPAUSED is a free data retrieval call binding the contract method 0x4ae9ddfa.
//
// Solidity: function INDEX_PAUSED() view returns(string)
func (_Errors *ErrorsCallerSession) INDEXPAUSED() (string, error) {
	return _Errors.Contract.INDEXPAUSED(&_Errors.CallOpts)
}

// INITIALIZEONCE is a free data retrieval call binding the contract method 0xb92066b2.
//
// Solidity: function INITIALIZE_ONCE() view returns(string)
func (_Errors *ErrorsCaller) INITIALIZEONCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INITIALIZE_ONCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INITIALIZEONCE is a free data retrieval call binding the contract method 0xb92066b2.
//
// Solidity: function INITIALIZE_ONCE() view returns(string)
func (_Errors *ErrorsSession) INITIALIZEONCE() (string, error) {
	return _Errors.Contract.INITIALIZEONCE(&_Errors.CallOpts)
}

// INITIALIZEONCE is a free data retrieval call binding the contract method 0xb92066b2.
//
// Solidity: function INITIALIZE_ONCE() view returns(string)
func (_Errors *ErrorsCallerSession) INITIALIZEONCE() (string, error) {
	return _Errors.Contract.INITIALIZEONCE(&_Errors.CallOpts)
}

// INPUTLENGTHMISMATCH is a free data retrieval call binding the contract method 0xfb765715.
//
// Solidity: function INPUT_LENGTH_MISMATCH() view returns(string)
func (_Errors *ErrorsCaller) INPUTLENGTHMISMATCH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INPUT_LENGTH_MISMATCH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INPUTLENGTHMISMATCH is a free data retrieval call binding the contract method 0xfb765715.
//
// Solidity: function INPUT_LENGTH_MISMATCH() view returns(string)
func (_Errors *ErrorsSession) INPUTLENGTHMISMATCH() (string, error) {
	return _Errors.Contract.INPUTLENGTHMISMATCH(&_Errors.CallOpts)
}

// INPUTLENGTHMISMATCH is a free data retrieval call binding the contract method 0xfb765715.
//
// Solidity: function INPUT_LENGTH_MISMATCH() view returns(string)
func (_Errors *ErrorsCallerSession) INPUTLENGTHMISMATCH() (string, error) {
	return _Errors.Contract.INPUTLENGTHMISMATCH(&_Errors.CallOpts)
}

// INVALIDEARNERCLAIMPROOF is a free data retrieval call binding the contract method 0xefa2ab10.
//
// Solidity: function INVALID_EARNER_CLAIM_PROOF() view returns(string)
func (_Errors *ErrorsCaller) INVALIDEARNERCLAIMPROOF(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_EARNER_CLAIM_PROOF")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDEARNERCLAIMPROOF is a free data retrieval call binding the contract method 0xefa2ab10.
//
// Solidity: function INVALID_EARNER_CLAIM_PROOF() view returns(string)
func (_Errors *ErrorsSession) INVALIDEARNERCLAIMPROOF() (string, error) {
	return _Errors.Contract.INVALIDEARNERCLAIMPROOF(&_Errors.CallOpts)
}

// INVALIDEARNERCLAIMPROOF is a free data retrieval call binding the contract method 0xefa2ab10.
//
// Solidity: function INVALID_EARNER_CLAIM_PROOF() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDEARNERCLAIMPROOF() (string, error) {
	return _Errors.Contract.INVALIDEARNERCLAIMPROOF(&_Errors.CallOpts)
}

// INVALIDEARNERLEAFINDEX is a free data retrieval call binding the contract method 0x733a577d.
//
// Solidity: function INVALID_EARNER_LEAF_INDEX() view returns(string)
func (_Errors *ErrorsCaller) INVALIDEARNERLEAFINDEX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_EARNER_LEAF_INDEX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDEARNERLEAFINDEX is a free data retrieval call binding the contract method 0x733a577d.
//
// Solidity: function INVALID_EARNER_LEAF_INDEX() view returns(string)
func (_Errors *ErrorsSession) INVALIDEARNERLEAFINDEX() (string, error) {
	return _Errors.Contract.INVALIDEARNERLEAFINDEX(&_Errors.CallOpts)
}

// INVALIDEARNERLEAFINDEX is a free data retrieval call binding the contract method 0x733a577d.
//
// Solidity: function INVALID_EARNER_LEAF_INDEX() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDEARNERLEAFINDEX() (string, error) {
	return _Errors.Contract.INVALIDEARNERLEAFINDEX(&_Errors.CallOpts)
}

// INVALIDPAUSE is a free data retrieval call binding the contract method 0x69bba998.
//
// Solidity: function INVALID_PAUSE() view returns(string)
func (_Errors *ErrorsCaller) INVALIDPAUSE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_PAUSE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDPAUSE is a free data retrieval call binding the contract method 0x69bba998.
//
// Solidity: function INVALID_PAUSE() view returns(string)
func (_Errors *ErrorsSession) INVALIDPAUSE() (string, error) {
	return _Errors.Contract.INVALIDPAUSE(&_Errors.CallOpts)
}

// INVALIDPAUSE is a free data retrieval call binding the contract method 0x69bba998.
//
// Solidity: function INVALID_PAUSE() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDPAUSE() (string, error) {
	return _Errors.Contract.INVALIDPAUSE(&_Errors.CallOpts)
}

// INVALIDPOOLCONSIDERED is a free data retrieval call binding the contract method 0x8c7fc1b2.
//
// Solidity: function INVALID_POOL_CONSIDERED() view returns(string)
func (_Errors *ErrorsCaller) INVALIDPOOLCONSIDERED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_POOL_CONSIDERED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDPOOLCONSIDERED is a free data retrieval call binding the contract method 0x8c7fc1b2.
//
// Solidity: function INVALID_POOL_CONSIDERED() view returns(string)
func (_Errors *ErrorsSession) INVALIDPOOLCONSIDERED() (string, error) {
	return _Errors.Contract.INVALIDPOOLCONSIDERED(&_Errors.CallOpts)
}

// INVALIDPOOLCONSIDERED is a free data retrieval call binding the contract method 0x8c7fc1b2.
//
// Solidity: function INVALID_POOL_CONSIDERED() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDPOOLCONSIDERED() (string, error) {
	return _Errors.Contract.INVALIDPOOLCONSIDERED(&_Errors.CallOpts)
}

// INVALIDROOTINDEX is a free data retrieval call binding the contract method 0x2ec313d3.
//
// Solidity: function INVALID_ROOT_INDEX() view returns(string)
func (_Errors *ErrorsCaller) INVALIDROOTINDEX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_ROOT_INDEX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDROOTINDEX is a free data retrieval call binding the contract method 0x2ec313d3.
//
// Solidity: function INVALID_ROOT_INDEX() view returns(string)
func (_Errors *ErrorsSession) INVALIDROOTINDEX() (string, error) {
	return _Errors.Contract.INVALIDROOTINDEX(&_Errors.CallOpts)
}

// INVALIDROOTINDEX is a free data retrieval call binding the contract method 0x2ec313d3.
//
// Solidity: function INVALID_ROOT_INDEX() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDROOTINDEX() (string, error) {
	return _Errors.Contract.INVALIDROOTINDEX(&_Errors.CallOpts)
}

// INVALIDTOKENCLAIMPROOF is a free data retrieval call binding the contract method 0x31a6fb97.
//
// Solidity: function INVALID_TOKEN_CLAIM_PROOF() view returns(string)
func (_Errors *ErrorsCaller) INVALIDTOKENCLAIMPROOF(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_TOKEN_CLAIM_PROOF")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDTOKENCLAIMPROOF is a free data retrieval call binding the contract method 0x31a6fb97.
//
// Solidity: function INVALID_TOKEN_CLAIM_PROOF() view returns(string)
func (_Errors *ErrorsSession) INVALIDTOKENCLAIMPROOF() (string, error) {
	return _Errors.Contract.INVALIDTOKENCLAIMPROOF(&_Errors.CallOpts)
}

// INVALIDTOKENCLAIMPROOF is a free data retrieval call binding the contract method 0x31a6fb97.
//
// Solidity: function INVALID_TOKEN_CLAIM_PROOF() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDTOKENCLAIMPROOF() (string, error) {
	return _Errors.Contract.INVALIDTOKENCLAIMPROOF(&_Errors.CallOpts)
}

// INVALIDTOKENLEAFINDEX is a free data retrieval call binding the contract method 0x67fd5c91.
//
// Solidity: function INVALID_TOKEN_LEAF_INDEX() view returns(string)
func (_Errors *ErrorsCaller) INVALIDTOKENLEAFINDEX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_TOKEN_LEAF_INDEX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDTOKENLEAFINDEX is a free data retrieval call binding the contract method 0x67fd5c91.
//
// Solidity: function INVALID_TOKEN_LEAF_INDEX() view returns(string)
func (_Errors *ErrorsSession) INVALIDTOKENLEAFINDEX() (string, error) {
	return _Errors.Contract.INVALIDTOKENLEAFINDEX(&_Errors.CallOpts)
}

// INVALIDTOKENLEAFINDEX is a free data retrieval call binding the contract method 0x67fd5c91.
//
// Solidity: function INVALID_TOKEN_LEAF_INDEX() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDTOKENLEAFINDEX() (string, error) {
	return _Errors.Contract.INVALIDTOKENLEAFINDEX(&_Errors.CallOpts)
}

// INVALIDUNPAUSE is a free data retrieval call binding the contract method 0xd577ae14.
//
// Solidity: function INVALID_UNPAUSE() view returns(string)
func (_Errors *ErrorsCaller) INVALIDUNPAUSE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "INVALID_UNPAUSE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// INVALIDUNPAUSE is a free data retrieval call binding the contract method 0xd577ae14.
//
// Solidity: function INVALID_UNPAUSE() view returns(string)
func (_Errors *ErrorsSession) INVALIDUNPAUSE() (string, error) {
	return _Errors.Contract.INVALIDUNPAUSE(&_Errors.CallOpts)
}

// INVALIDUNPAUSE is a free data retrieval call binding the contract method 0xd577ae14.
//
// Solidity: function INVALID_UNPAUSE() view returns(string)
func (_Errors *ErrorsCallerSession) INVALIDUNPAUSE() (string, error) {
	return _Errors.Contract.INVALIDUNPAUSE(&_Errors.CallOpts)
}

// MAXDEPOSITSEXCEEDED is a free data retrieval call binding the contract method 0x490ff719.
//
// Solidity: function MAX_DEPOSITS_EXCEEDED() view returns(string)
func (_Errors *ErrorsCaller) MAXDEPOSITSEXCEEDED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MAX_DEPOSITS_EXCEEDED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAXDEPOSITSEXCEEDED is a free data retrieval call binding the contract method 0x490ff719.
//
// Solidity: function MAX_DEPOSITS_EXCEEDED() view returns(string)
func (_Errors *ErrorsSession) MAXDEPOSITSEXCEEDED() (string, error) {
	return _Errors.Contract.MAXDEPOSITSEXCEEDED(&_Errors.CallOpts)
}

// MAXDEPOSITSEXCEEDED is a free data retrieval call binding the contract method 0x490ff719.
//
// Solidity: function MAX_DEPOSITS_EXCEEDED() view returns(string)
func (_Errors *ErrorsCallerSession) MAXDEPOSITSEXCEEDED() (string, error) {
	return _Errors.Contract.MAXDEPOSITSEXCEEDED(&_Errors.CallOpts)
}

// MAXPERDEPOSITEXCEEDSMAXTOTAL is a free data retrieval call binding the contract method 0x1863ca35.
//
// Solidity: function MAX_PER_DEPOSIT_EXCEEDS_MAX_TOTAL() view returns(string)
func (_Errors *ErrorsCaller) MAXPERDEPOSITEXCEEDSMAXTOTAL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MAX_PER_DEPOSIT_EXCEEDS_MAX_TOTAL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAXPERDEPOSITEXCEEDSMAXTOTAL is a free data retrieval call binding the contract method 0x1863ca35.
//
// Solidity: function MAX_PER_DEPOSIT_EXCEEDS_MAX_TOTAL() view returns(string)
func (_Errors *ErrorsSession) MAXPERDEPOSITEXCEEDSMAXTOTAL() (string, error) {
	return _Errors.Contract.MAXPERDEPOSITEXCEEDSMAXTOTAL(&_Errors.CallOpts)
}

// MAXPERDEPOSITEXCEEDSMAXTOTAL is a free data retrieval call binding the contract method 0x1863ca35.
//
// Solidity: function MAX_PER_DEPOSIT_EXCEEDS_MAX_TOTAL() view returns(string)
func (_Errors *ErrorsCallerSession) MAXPERDEPOSITEXCEEDSMAXTOTAL() (string, error) {
	return _Errors.Contract.MAXPERDEPOSITEXCEEDSMAXTOTAL(&_Errors.CallOpts)
}

// MAXPERDEPOSITLIMITEXCEEDED is a free data retrieval call binding the contract method 0x8a301624.
//
// Solidity: function MAX_PER_DEPOSIT_LIMIT_EXCEEDED() view returns(string)
func (_Errors *ErrorsCaller) MAXPERDEPOSITLIMITEXCEEDED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MAX_PER_DEPOSIT_LIMIT_EXCEEDED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MAXPERDEPOSITLIMITEXCEEDED is a free data retrieval call binding the contract method 0x8a301624.
//
// Solidity: function MAX_PER_DEPOSIT_LIMIT_EXCEEDED() view returns(string)
func (_Errors *ErrorsSession) MAXPERDEPOSITLIMITEXCEEDED() (string, error) {
	return _Errors.Contract.MAXPERDEPOSITLIMITEXCEEDED(&_Errors.CallOpts)
}

// MAXPERDEPOSITLIMITEXCEEDED is a free data retrieval call binding the contract method 0x8a301624.
//
// Solidity: function MAX_PER_DEPOSIT_LIMIT_EXCEEDED() view returns(string)
func (_Errors *ErrorsCallerSession) MAXPERDEPOSITLIMITEXCEEDED() (string, error) {
	return _Errors.Contract.MAXPERDEPOSITLIMITEXCEEDED(&_Errors.CallOpts)
}

// MINWITHDRAWALDELAYEXCEEDSMAX is a free data retrieval call binding the contract method 0x9c06be3e.
//
// Solidity: function MIN_WITHDRAWAL_DELAY_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsCaller) MINWITHDRAWALDELAYEXCEEDSMAX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MIN_WITHDRAWAL_DELAY_EXCEEDS_MAX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MINWITHDRAWALDELAYEXCEEDSMAX is a free data retrieval call binding the contract method 0x9c06be3e.
//
// Solidity: function MIN_WITHDRAWAL_DELAY_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsSession) MINWITHDRAWALDELAYEXCEEDSMAX() (string, error) {
	return _Errors.Contract.MINWITHDRAWALDELAYEXCEEDSMAX(&_Errors.CallOpts)
}

// MINWITHDRAWALDELAYEXCEEDSMAX is a free data retrieval call binding the contract method 0x9c06be3e.
//
// Solidity: function MIN_WITHDRAWAL_DELAY_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsCallerSession) MINWITHDRAWALDELAYEXCEEDSMAX() (string, error) {
	return _Errors.Contract.MINWITHDRAWALDELAYEXCEEDSMAX(&_Errors.CallOpts)
}

// MINWITHDRAWALDELAYNOTPASSED is a free data retrieval call binding the contract method 0x5b33eb0e.
//
// Solidity: function MIN_WITHDRAWAL_DELAY_NOT_PASSED() view returns(string)
func (_Errors *ErrorsCaller) MINWITHDRAWALDELAYNOTPASSED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MIN_WITHDRAWAL_DELAY_NOT_PASSED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MINWITHDRAWALDELAYNOTPASSED is a free data retrieval call binding the contract method 0x5b33eb0e.
//
// Solidity: function MIN_WITHDRAWAL_DELAY_NOT_PASSED() view returns(string)
func (_Errors *ErrorsSession) MINWITHDRAWALDELAYNOTPASSED() (string, error) {
	return _Errors.Contract.MINWITHDRAWALDELAYNOTPASSED(&_Errors.CallOpts)
}

// MINWITHDRAWALDELAYNOTPASSED is a free data retrieval call binding the contract method 0x5b33eb0e.
//
// Solidity: function MIN_WITHDRAWAL_DELAY_NOT_PASSED() view returns(string)
func (_Errors *ErrorsCallerSession) MINWITHDRAWALDELAYNOTPASSED() (string, error) {
	return _Errors.Contract.MINWITHDRAWALDELAYNOTPASSED(&_Errors.CallOpts)
}

// NEWROOTMUSTBENEWERCALCULATEDPERIOD is a free data retrieval call binding the contract method 0xcdaff814.
//
// Solidity: function NEW_ROOT_MUST_BE_NEWER_CALCULATED_PERIOD() view returns(string)
func (_Errors *ErrorsCaller) NEWROOTMUSTBENEWERCALCULATEDPERIOD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NEW_ROOT_MUST_BE_NEWER_CALCULATED_PERIOD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NEWROOTMUSTBENEWERCALCULATEDPERIOD is a free data retrieval call binding the contract method 0xcdaff814.
//
// Solidity: function NEW_ROOT_MUST_BE_NEWER_CALCULATED_PERIOD() view returns(string)
func (_Errors *ErrorsSession) NEWROOTMUSTBENEWERCALCULATEDPERIOD() (string, error) {
	return _Errors.Contract.NEWROOTMUSTBENEWERCALCULATEDPERIOD(&_Errors.CallOpts)
}

// NEWROOTMUSTBENEWERCALCULATEDPERIOD is a free data retrieval call binding the contract method 0xcdaff814.
//
// Solidity: function NEW_ROOT_MUST_BE_NEWER_CALCULATED_PERIOD() view returns(string)
func (_Errors *ErrorsCallerSession) NEWROOTMUSTBENEWERCALCULATEDPERIOD() (string, error) {
	return _Errors.Contract.NEWROOTMUSTBENEWERCALCULATEDPERIOD(&_Errors.CallOpts)
}

// NOTCLAIMER is a free data retrieval call binding the contract method 0x32635eba.
//
// Solidity: function NOT_CLAIMER() view returns(string)
func (_Errors *ErrorsCaller) NOTCLAIMER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_CLAIMER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTCLAIMER is a free data retrieval call binding the contract method 0x32635eba.
//
// Solidity: function NOT_CLAIMER() view returns(string)
func (_Errors *ErrorsSession) NOTCLAIMER() (string, error) {
	return _Errors.Contract.NOTCLAIMER(&_Errors.CallOpts)
}

// NOTCLAIMER is a free data retrieval call binding the contract method 0x32635eba.
//
// Solidity: function NOT_CLAIMER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTCLAIMER() (string, error) {
	return _Errors.Contract.NOTCLAIMER(&_Errors.CallOpts)
}

// NOTCREATEREWARDSFORALLSUBMITTER is a free data retrieval call binding the contract method 0xd98555a1.
//
// Solidity: function NOT_CREATE_REWARDS_FOR_ALL_SUBMITTER() view returns(string)
func (_Errors *ErrorsCaller) NOTCREATEREWARDSFORALLSUBMITTER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_CREATE_REWARDS_FOR_ALL_SUBMITTER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTCREATEREWARDSFORALLSUBMITTER is a free data retrieval call binding the contract method 0xd98555a1.
//
// Solidity: function NOT_CREATE_REWARDS_FOR_ALL_SUBMITTER() view returns(string)
func (_Errors *ErrorsSession) NOTCREATEREWARDSFORALLSUBMITTER() (string, error) {
	return _Errors.Contract.NOTCREATEREWARDSFORALLSUBMITTER(&_Errors.CallOpts)
}

// NOTCREATEREWARDSFORALLSUBMITTER is a free data retrieval call binding the contract method 0xd98555a1.
//
// Solidity: function NOT_CREATE_REWARDS_FOR_ALL_SUBMITTER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTCREATEREWARDSFORALLSUBMITTER() (string, error) {
	return _Errors.Contract.NOTCREATEREWARDSFORALLSUBMITTER(&_Errors.CallOpts)
}

// NOTDELEGATIONCONTROLLER is a free data retrieval call binding the contract method 0x70202598.
//
// Solidity: function NOT_DELEGATION_CONTROLLER() view returns(string)
func (_Errors *ErrorsCaller) NOTDELEGATIONCONTROLLER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_DELEGATION_CONTROLLER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTDELEGATIONCONTROLLER is a free data retrieval call binding the contract method 0x70202598.
//
// Solidity: function NOT_DELEGATION_CONTROLLER() view returns(string)
func (_Errors *ErrorsSession) NOTDELEGATIONCONTROLLER() (string, error) {
	return _Errors.Contract.NOTDELEGATIONCONTROLLER(&_Errors.CallOpts)
}

// NOTDELEGATIONCONTROLLER is a free data retrieval call binding the contract method 0x70202598.
//
// Solidity: function NOT_DELEGATION_CONTROLLER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTDELEGATIONCONTROLLER() (string, error) {
	return _Errors.Contract.NOTDELEGATIONCONTROLLER(&_Errors.CallOpts)
}

// NOTPAUSER is a free data retrieval call binding the contract method 0x941abc77.
//
// Solidity: function NOT_PAUSER() view returns(string)
func (_Errors *ErrorsCaller) NOTPAUSER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_PAUSER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTPAUSER is a free data retrieval call binding the contract method 0x941abc77.
//
// Solidity: function NOT_PAUSER() view returns(string)
func (_Errors *ErrorsSession) NOTPAUSER() (string, error) {
	return _Errors.Contract.NOTPAUSER(&_Errors.CallOpts)
}

// NOTPAUSER is a free data retrieval call binding the contract method 0x941abc77.
//
// Solidity: function NOT_PAUSER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTPAUSER() (string, error) {
	return _Errors.Contract.NOTPAUSER(&_Errors.CallOpts)
}

// NOTPOOLWHITELISTER is a free data retrieval call binding the contract method 0xc619c91f.
//
// Solidity: function NOT_POOL_WHITELISTER() view returns(string)
func (_Errors *ErrorsCaller) NOTPOOLWHITELISTER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_POOL_WHITELISTER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTPOOLWHITELISTER is a free data retrieval call binding the contract method 0xc619c91f.
//
// Solidity: function NOT_POOL_WHITELISTER() view returns(string)
func (_Errors *ErrorsSession) NOTPOOLWHITELISTER() (string, error) {
	return _Errors.Contract.NOTPOOLWHITELISTER(&_Errors.CallOpts)
}

// NOTPOOLWHITELISTER is a free data retrieval call binding the contract method 0xc619c91f.
//
// Solidity: function NOT_POOL_WHITELISTER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTPOOLWHITELISTER() (string, error) {
	return _Errors.Contract.NOTPOOLWHITELISTER(&_Errors.CallOpts)
}

// NOTREGISTEREDINBINLAYER is a free data retrieval call binding the contract method 0xca5768c2.
//
// Solidity: function NOT_REGISTERED_IN_BINLAYER() view returns(string)
func (_Errors *ErrorsCaller) NOTREGISTEREDINBINLAYER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_REGISTERED_IN_BINLAYER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTREGISTEREDINBINLAYER is a free data retrieval call binding the contract method 0xca5768c2.
//
// Solidity: function NOT_REGISTERED_IN_BINLAYER() view returns(string)
func (_Errors *ErrorsSession) NOTREGISTEREDINBINLAYER() (string, error) {
	return _Errors.Contract.NOTREGISTEREDINBINLAYER(&_Errors.CallOpts)
}

// NOTREGISTEREDINBINLAYER is a free data retrieval call binding the contract method 0xca5768c2.
//
// Solidity: function NOT_REGISTERED_IN_BINLAYER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTREGISTEREDINBINLAYER() (string, error) {
	return _Errors.Contract.NOTREGISTEREDINBINLAYER(&_Errors.CallOpts)
}

// NOTREWARDSUPDATER is a free data retrieval call binding the contract method 0xd2b3b176.
//
// Solidity: function NOT_REWARDS_UPDATER() view returns(string)
func (_Errors *ErrorsCaller) NOTREWARDSUPDATER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_REWARDS_UPDATER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTREWARDSUPDATER is a free data retrieval call binding the contract method 0xd2b3b176.
//
// Solidity: function NOT_REWARDS_UPDATER() view returns(string)
func (_Errors *ErrorsSession) NOTREWARDSUPDATER() (string, error) {
	return _Errors.Contract.NOTREWARDSUPDATER(&_Errors.CallOpts)
}

// NOTREWARDSUPDATER is a free data retrieval call binding the contract method 0xd2b3b176.
//
// Solidity: function NOT_REWARDS_UPDATER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTREWARDSUPDATER() (string, error) {
	return _Errors.Contract.NOTREWARDSUPDATER(&_Errors.CallOpts)
}

// NOTUNPAUSER is a free data retrieval call binding the contract method 0xc3fcd853.
//
// Solidity: function NOT_UNPAUSER() view returns(string)
func (_Errors *ErrorsCaller) NOTUNPAUSER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NOT_UNPAUSER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOTUNPAUSER is a free data retrieval call binding the contract method 0xc3fcd853.
//
// Solidity: function NOT_UNPAUSER() view returns(string)
func (_Errors *ErrorsSession) NOTUNPAUSER() (string, error) {
	return _Errors.Contract.NOTUNPAUSER(&_Errors.CallOpts)
}

// NOTUNPAUSER is a free data retrieval call binding the contract method 0xc3fcd853.
//
// Solidity: function NOT_UNPAUSER() view returns(string)
func (_Errors *ErrorsCallerSession) NOTUNPAUSER() (string, error) {
	return _Errors.Contract.NOTUNPAUSER(&_Errors.CallOpts)
}

// NOPOOLSSET is a free data retrieval call binding the contract method 0xff9a0df7.
//
// Solidity: function NO_POOLS_SET() view returns(string)
func (_Errors *ErrorsCaller) NOPOOLSSET(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "NO_POOLS_SET")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NOPOOLSSET is a free data retrieval call binding the contract method 0xff9a0df7.
//
// Solidity: function NO_POOLS_SET() view returns(string)
func (_Errors *ErrorsSession) NOPOOLSSET() (string, error) {
	return _Errors.Contract.NOPOOLSSET(&_Errors.CallOpts)
}

// NOPOOLSSET is a free data retrieval call binding the contract method 0xff9a0df7.
//
// Solidity: function NO_POOLS_SET() view returns(string)
func (_Errors *ErrorsCallerSession) NOPOOLSSET() (string, error) {
	return _Errors.Contract.NOPOOLSSET(&_Errors.CallOpts)
}

// ONLYPOOLCONTROLLER is a free data retrieval call binding the contract method 0x4b5f220a.
//
// Solidity: function ONLY_POOL_CONTROLLER() view returns(string)
func (_Errors *ErrorsCaller) ONLYPOOLCONTROLLER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ONLY_POOL_CONTROLLER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ONLYPOOLCONTROLLER is a free data retrieval call binding the contract method 0x4b5f220a.
//
// Solidity: function ONLY_POOL_CONTROLLER() view returns(string)
func (_Errors *ErrorsSession) ONLYPOOLCONTROLLER() (string, error) {
	return _Errors.Contract.ONLYPOOLCONTROLLER(&_Errors.CallOpts)
}

// ONLYPOOLCONTROLLER is a free data retrieval call binding the contract method 0x4b5f220a.
//
// Solidity: function ONLY_POOL_CONTROLLER() view returns(string)
func (_Errors *ErrorsCallerSession) ONLYPOOLCONTROLLER() (string, error) {
	return _Errors.Contract.ONLYPOOLCONTROLLER(&_Errors.CallOpts)
}

// ONLYSUPPORTWRAPPEDTOKENPOOL is a free data retrieval call binding the contract method 0xa0afa0cc.
//
// Solidity: function ONLY_SUPPORT_WRAPPED_TOKEN_POOL() view returns(string)
func (_Errors *ErrorsCaller) ONLYSUPPORTWRAPPEDTOKENPOOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ONLY_SUPPORT_WRAPPED_TOKEN_POOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ONLYSUPPORTWRAPPEDTOKENPOOL is a free data retrieval call binding the contract method 0xa0afa0cc.
//
// Solidity: function ONLY_SUPPORT_WRAPPED_TOKEN_POOL() view returns(string)
func (_Errors *ErrorsSession) ONLYSUPPORTWRAPPEDTOKENPOOL() (string, error) {
	return _Errors.Contract.ONLYSUPPORTWRAPPEDTOKENPOOL(&_Errors.CallOpts)
}

// ONLYSUPPORTWRAPPEDTOKENPOOL is a free data retrieval call binding the contract method 0xa0afa0cc.
//
// Solidity: function ONLY_SUPPORT_WRAPPED_TOKEN_POOL() view returns(string)
func (_Errors *ErrorsCallerSession) ONLYSUPPORTWRAPPEDTOKENPOOL() (string, error) {
	return _Errors.Contract.ONLYSUPPORTWRAPPEDTOKENPOOL(&_Errors.CallOpts)
}

// ONLYWITHDRAWERCANCOMPLETE is a free data retrieval call binding the contract method 0x55b533a6.
//
// Solidity: function ONLY_WITHDRAWER_CAN_COMPLETE() view returns(string)
func (_Errors *ErrorsCaller) ONLYWITHDRAWERCANCOMPLETE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ONLY_WITHDRAWER_CAN_COMPLETE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ONLYWITHDRAWERCANCOMPLETE is a free data retrieval call binding the contract method 0x55b533a6.
//
// Solidity: function ONLY_WITHDRAWER_CAN_COMPLETE() view returns(string)
func (_Errors *ErrorsSession) ONLYWITHDRAWERCANCOMPLETE() (string, error) {
	return _Errors.Contract.ONLYWITHDRAWERCANCOMPLETE(&_Errors.CallOpts)
}

// ONLYWITHDRAWERCANCOMPLETE is a free data retrieval call binding the contract method 0x55b533a6.
//
// Solidity: function ONLY_WITHDRAWER_CAN_COMPLETE() view returns(string)
func (_Errors *ErrorsCallerSession) ONLYWITHDRAWERCANCOMPLETE() (string, error) {
	return _Errors.Contract.ONLYWITHDRAWERCANCOMPLETE(&_Errors.CallOpts)
}

// OPERATORALREADYREGISTERED is a free data retrieval call binding the contract method 0x0b8c0aa1.
//
// Solidity: function OPERATOR_ALREADY_REGISTERED() view returns(string)
func (_Errors *ErrorsCaller) OPERATORALREADYREGISTERED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "OPERATOR_ALREADY_REGISTERED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OPERATORALREADYREGISTERED is a free data retrieval call binding the contract method 0x0b8c0aa1.
//
// Solidity: function OPERATOR_ALREADY_REGISTERED() view returns(string)
func (_Errors *ErrorsSession) OPERATORALREADYREGISTERED() (string, error) {
	return _Errors.Contract.OPERATORALREADYREGISTERED(&_Errors.CallOpts)
}

// OPERATORALREADYREGISTERED is a free data retrieval call binding the contract method 0x0b8c0aa1.
//
// Solidity: function OPERATOR_ALREADY_REGISTERED() view returns(string)
func (_Errors *ErrorsCallerSession) OPERATORALREADYREGISTERED() (string, error) {
	return _Errors.Contract.OPERATORALREADYREGISTERED(&_Errors.CallOpts)
}

// OPTOUTWINDOWEXCEEDSMAX is a free data retrieval call binding the contract method 0xf086b291.
//
// Solidity: function OPT_OUT_WINDOW_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsCaller) OPTOUTWINDOWEXCEEDSMAX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "OPT_OUT_WINDOW_EXCEEDS_MAX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OPTOUTWINDOWEXCEEDSMAX is a free data retrieval call binding the contract method 0xf086b291.
//
// Solidity: function OPT_OUT_WINDOW_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsSession) OPTOUTWINDOWEXCEEDSMAX() (string, error) {
	return _Errors.Contract.OPTOUTWINDOWEXCEEDSMAX(&_Errors.CallOpts)
}

// OPTOUTWINDOWEXCEEDSMAX is a free data retrieval call binding the contract method 0xf086b291.
//
// Solidity: function OPT_OUT_WINDOW_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsCallerSession) OPTOUTWINDOWEXCEEDSMAX() (string, error) {
	return _Errors.Contract.OPTOUTWINDOWEXCEEDSMAX(&_Errors.CallOpts)
}

// POOLSCANNOTBEEMPTY is a free data retrieval call binding the contract method 0x99adb1c7.
//
// Solidity: function POOLS_CANNOT_BE_EMPTY() view returns(string)
func (_Errors *ErrorsCaller) POOLSCANNOTBEEMPTY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOLS_CANNOT_BE_EMPTY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLSCANNOTBEEMPTY is a free data retrieval call binding the contract method 0x99adb1c7.
//
// Solidity: function POOLS_CANNOT_BE_EMPTY() view returns(string)
func (_Errors *ErrorsSession) POOLSCANNOTBEEMPTY() (string, error) {
	return _Errors.Contract.POOLSCANNOTBEEMPTY(&_Errors.CallOpts)
}

// POOLSCANNOTBEEMPTY is a free data retrieval call binding the contract method 0x99adb1c7.
//
// Solidity: function POOLS_CANNOT_BE_EMPTY() view returns(string)
func (_Errors *ErrorsCallerSession) POOLSCANNOTBEEMPTY() (string, error) {
	return _Errors.Contract.POOLSCANNOTBEEMPTY(&_Errors.CallOpts)
}

// POOLSMUSTBEASCENDING is a free data retrieval call binding the contract method 0xae7a3f6d.
//
// Solidity: function POOLS_MUST_BE_ASCENDING() view returns(string)
func (_Errors *ErrorsCaller) POOLSMUSTBEASCENDING(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOLS_MUST_BE_ASCENDING")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLSMUSTBEASCENDING is a free data retrieval call binding the contract method 0xae7a3f6d.
//
// Solidity: function POOLS_MUST_BE_ASCENDING() view returns(string)
func (_Errors *ErrorsSession) POOLSMUSTBEASCENDING() (string, error) {
	return _Errors.Contract.POOLSMUSTBEASCENDING(&_Errors.CallOpts)
}

// POOLSMUSTBEASCENDING is a free data retrieval call binding the contract method 0xae7a3f6d.
//
// Solidity: function POOLS_MUST_BE_ASCENDING() view returns(string)
func (_Errors *ErrorsCallerSession) POOLSMUSTBEASCENDING() (string, error) {
	return _Errors.Contract.POOLSMUSTBEASCENDING(&_Errors.CallOpts)
}

// POOLNOTFOUND is a free data retrieval call binding the contract method 0x01de8c6d.
//
// Solidity: function POOL_NOT_FOUND() view returns(string)
func (_Errors *ErrorsCaller) POOLNOTFOUND(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_NOT_FOUND")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLNOTFOUND is a free data retrieval call binding the contract method 0x01de8c6d.
//
// Solidity: function POOL_NOT_FOUND() view returns(string)
func (_Errors *ErrorsSession) POOLNOTFOUND() (string, error) {
	return _Errors.Contract.POOLNOTFOUND(&_Errors.CallOpts)
}

// POOLNOTFOUND is a free data retrieval call binding the contract method 0x01de8c6d.
//
// Solidity: function POOL_NOT_FOUND() view returns(string)
func (_Errors *ErrorsCallerSession) POOLNOTFOUND() (string, error) {
	return _Errors.Contract.POOLNOTFOUND(&_Errors.CallOpts)
}

// POOLNOTWHITELISTED is a free data retrieval call binding the contract method 0xd1b25b31.
//
// Solidity: function POOL_NOT_WHITELISTED() view returns(string)
func (_Errors *ErrorsCaller) POOLNOTWHITELISTED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_NOT_WHITELISTED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLNOTWHITELISTED is a free data retrieval call binding the contract method 0xd1b25b31.
//
// Solidity: function POOL_NOT_WHITELISTED() view returns(string)
func (_Errors *ErrorsSession) POOLNOTWHITELISTED() (string, error) {
	return _Errors.Contract.POOLNOTWHITELISTED(&_Errors.CallOpts)
}

// POOLNOTWHITELISTED is a free data retrieval call binding the contract method 0xd1b25b31.
//
// Solidity: function POOL_NOT_WHITELISTED() view returns(string)
func (_Errors *ErrorsCallerSession) POOLNOTWHITELISTED() (string, error) {
	return _Errors.Contract.POOLNOTWHITELISTED(&_Errors.CallOpts)
}

// POOLWITHDRAWALDELAYEXCEEDSMAX is a free data retrieval call binding the contract method 0xeb565197.
//
// Solidity: function POOL_WITHDRAWAL_DELAY_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsCaller) POOLWITHDRAWALDELAYEXCEEDSMAX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "POOL_WITHDRAWAL_DELAY_EXCEEDS_MAX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POOLWITHDRAWALDELAYEXCEEDSMAX is a free data retrieval call binding the contract method 0xeb565197.
//
// Solidity: function POOL_WITHDRAWAL_DELAY_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsSession) POOLWITHDRAWALDELAYEXCEEDSMAX() (string, error) {
	return _Errors.Contract.POOLWITHDRAWALDELAYEXCEEDSMAX(&_Errors.CallOpts)
}

// POOLWITHDRAWALDELAYEXCEEDSMAX is a free data retrieval call binding the contract method 0xeb565197.
//
// Solidity: function POOL_WITHDRAWAL_DELAY_EXCEEDS_MAX() view returns(string)
func (_Errors *ErrorsCallerSession) POOLWITHDRAWALDELAYEXCEEDSMAX() (string, error) {
	return _Errors.Contract.POOLWITHDRAWALDELAYEXCEEDSMAX(&_Errors.CallOpts)
}

// ROOTALREADYACTIVATED is a free data retrieval call binding the contract method 0x0ee30d10.
//
// Solidity: function ROOT_ALREADY_ACTIVATED() view returns(string)
func (_Errors *ErrorsCaller) ROOTALREADYACTIVATED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ROOT_ALREADY_ACTIVATED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROOTALREADYACTIVATED is a free data retrieval call binding the contract method 0x0ee30d10.
//
// Solidity: function ROOT_ALREADY_ACTIVATED() view returns(string)
func (_Errors *ErrorsSession) ROOTALREADYACTIVATED() (string, error) {
	return _Errors.Contract.ROOTALREADYACTIVATED(&_Errors.CallOpts)
}

// ROOTALREADYACTIVATED is a free data retrieval call binding the contract method 0x0ee30d10.
//
// Solidity: function ROOT_ALREADY_ACTIVATED() view returns(string)
func (_Errors *ErrorsCallerSession) ROOTALREADYACTIVATED() (string, error) {
	return _Errors.Contract.ROOTALREADYACTIVATED(&_Errors.CallOpts)
}

// ROOTALREADYDISABLED is a free data retrieval call binding the contract method 0xca2633cd.
//
// Solidity: function ROOT_ALREADY_DISABLED() view returns(string)
func (_Errors *ErrorsCaller) ROOTALREADYDISABLED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ROOT_ALREADY_DISABLED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROOTALREADYDISABLED is a free data retrieval call binding the contract method 0xca2633cd.
//
// Solidity: function ROOT_ALREADY_DISABLED() view returns(string)
func (_Errors *ErrorsSession) ROOTALREADYDISABLED() (string, error) {
	return _Errors.Contract.ROOTALREADYDISABLED(&_Errors.CallOpts)
}

// ROOTALREADYDISABLED is a free data retrieval call binding the contract method 0xca2633cd.
//
// Solidity: function ROOT_ALREADY_DISABLED() view returns(string)
func (_Errors *ErrorsCallerSession) ROOTALREADYDISABLED() (string, error) {
	return _Errors.Contract.ROOTALREADYDISABLED(&_Errors.CallOpts)
}

// ROOTISDISABLED is a free data retrieval call binding the contract method 0x0cc6ae8a.
//
// Solidity: function ROOT_IS_DISABLED() view returns(string)
func (_Errors *ErrorsCaller) ROOTISDISABLED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ROOT_IS_DISABLED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROOTISDISABLED is a free data retrieval call binding the contract method 0x0cc6ae8a.
//
// Solidity: function ROOT_IS_DISABLED() view returns(string)
func (_Errors *ErrorsSession) ROOTISDISABLED() (string, error) {
	return _Errors.Contract.ROOTISDISABLED(&_Errors.CallOpts)
}

// ROOTISDISABLED is a free data retrieval call binding the contract method 0x0cc6ae8a.
//
// Solidity: function ROOT_IS_DISABLED() view returns(string)
func (_Errors *ErrorsCallerSession) ROOTISDISABLED() (string, error) {
	return _Errors.Contract.ROOTISDISABLED(&_Errors.CallOpts)
}

// ROOTNOTACTIVATEDYET is a free data retrieval call binding the contract method 0xd28c87f2.
//
// Solidity: function ROOT_NOT_ACTIVATED_YET() view returns(string)
func (_Errors *ErrorsCaller) ROOTNOTACTIVATEDYET(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ROOT_NOT_ACTIVATED_YET")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROOTNOTACTIVATEDYET is a free data retrieval call binding the contract method 0xd28c87f2.
//
// Solidity: function ROOT_NOT_ACTIVATED_YET() view returns(string)
func (_Errors *ErrorsSession) ROOTNOTACTIVATEDYET() (string, error) {
	return _Errors.Contract.ROOTNOTACTIVATEDYET(&_Errors.CallOpts)
}

// ROOTNOTACTIVATEDYET is a free data retrieval call binding the contract method 0xd28c87f2.
//
// Solidity: function ROOT_NOT_ACTIVATED_YET() view returns(string)
func (_Errors *ErrorsCallerSession) ROOTNOTACTIVATEDYET() (string, error) {
	return _Errors.Contract.ROOTNOTACTIVATEDYET(&_Errors.CallOpts)
}

// ROOTNOTFOUND is a free data retrieval call binding the contract method 0xf10c9421.
//
// Solidity: function ROOT_NOT_FOUND() view returns(string)
func (_Errors *ErrorsCaller) ROOTNOTFOUND(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ROOT_NOT_FOUND")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROOTNOTFOUND is a free data retrieval call binding the contract method 0xf10c9421.
//
// Solidity: function ROOT_NOT_FOUND() view returns(string)
func (_Errors *ErrorsSession) ROOTNOTFOUND() (string, error) {
	return _Errors.Contract.ROOTNOTFOUND(&_Errors.CallOpts)
}

// ROOTNOTFOUND is a free data retrieval call binding the contract method 0xf10c9421.
//
// Solidity: function ROOT_NOT_FOUND() view returns(string)
func (_Errors *ErrorsCallerSession) ROOTNOTFOUND() (string, error) {
	return _Errors.Contract.ROOTNOTFOUND(&_Errors.CallOpts)
}

// SALTALREADYSPENT is a free data retrieval call binding the contract method 0xa7651b98.
//
// Solidity: function SALT_ALREADY_SPENT() view returns(string)
func (_Errors *ErrorsCaller) SALTALREADYSPENT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "SALT_ALREADY_SPENT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SALTALREADYSPENT is a free data retrieval call binding the contract method 0xa7651b98.
//
// Solidity: function SALT_ALREADY_SPENT() view returns(string)
func (_Errors *ErrorsSession) SALTALREADYSPENT() (string, error) {
	return _Errors.Contract.SALTALREADYSPENT(&_Errors.CallOpts)
}

// SALTALREADYSPENT is a free data retrieval call binding the contract method 0xa7651b98.
//
// Solidity: function SALT_ALREADY_SPENT() view returns(string)
func (_Errors *ErrorsCallerSession) SALTALREADYSPENT() (string, error) {
	return _Errors.Contract.SALTALREADYSPENT(&_Errors.CallOpts)
}

// SHAREAMOUNTTOOHIGH is a free data retrieval call binding the contract method 0x2ec82ef4.
//
// Solidity: function SHARE_AMOUNT_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCaller) SHAREAMOUNTTOOHIGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "SHARE_AMOUNT_TOO_HIGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SHAREAMOUNTTOOHIGH is a free data retrieval call binding the contract method 0x2ec82ef4.
//
// Solidity: function SHARE_AMOUNT_TOO_HIGH() view returns(string)
func (_Errors *ErrorsSession) SHAREAMOUNTTOOHIGH() (string, error) {
	return _Errors.Contract.SHAREAMOUNTTOOHIGH(&_Errors.CallOpts)
}

// SHAREAMOUNTTOOHIGH is a free data retrieval call binding the contract method 0x2ec82ef4.
//
// Solidity: function SHARE_AMOUNT_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCallerSession) SHAREAMOUNTTOOHIGH() (string, error) {
	return _Errors.Contract.SHAREAMOUNTTOOHIGH(&_Errors.CallOpts)
}

// SIGNATUREEXPIRED is a free data retrieval call binding the contract method 0x19319827.
//
// Solidity: function SIGNATURE_EXPIRED() view returns(string)
func (_Errors *ErrorsCaller) SIGNATUREEXPIRED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "SIGNATURE_EXPIRED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SIGNATUREEXPIRED is a free data retrieval call binding the contract method 0x19319827.
//
// Solidity: function SIGNATURE_EXPIRED() view returns(string)
func (_Errors *ErrorsSession) SIGNATUREEXPIRED() (string, error) {
	return _Errors.Contract.SIGNATUREEXPIRED(&_Errors.CallOpts)
}

// SIGNATUREEXPIRED is a free data retrieval call binding the contract method 0x19319827.
//
// Solidity: function SIGNATURE_EXPIRED() view returns(string)
func (_Errors *ErrorsCallerSession) SIGNATUREEXPIRED() (string, error) {
	return _Errors.Contract.SIGNATUREEXPIRED(&_Errors.CallOpts)
}

// STAKERMUSTBEDELEGATED is a free data retrieval call binding the contract method 0x07ae55ba.
//
// Solidity: function STAKER_MUST_BE_DELEGATED() view returns(string)
func (_Errors *ErrorsCaller) STAKERMUSTBEDELEGATED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "STAKER_MUST_BE_DELEGATED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKERMUSTBEDELEGATED is a free data retrieval call binding the contract method 0x07ae55ba.
//
// Solidity: function STAKER_MUST_BE_DELEGATED() view returns(string)
func (_Errors *ErrorsSession) STAKERMUSTBEDELEGATED() (string, error) {
	return _Errors.Contract.STAKERMUSTBEDELEGATED(&_Errors.CallOpts)
}

// STAKERMUSTBEDELEGATED is a free data retrieval call binding the contract method 0x07ae55ba.
//
// Solidity: function STAKER_MUST_BE_DELEGATED() view returns(string)
func (_Errors *ErrorsCallerSession) STAKERMUSTBEDELEGATED() (string, error) {
	return _Errors.Contract.STAKERMUSTBEDELEGATED(&_Errors.CallOpts)
}

// STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0xb6f7ae43.
//
// Solidity: function STARTTIMESTAMP_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS() view returns(string)
func (_Errors *ErrorsCaller) STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "STARTTIMESTAMP_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0xb6f7ae43.
//
// Solidity: function STARTTIMESTAMP_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS() view returns(string)
func (_Errors *ErrorsSession) STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS() (string, error) {
	return _Errors.Contract.STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS(&_Errors.CallOpts)
}

// STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0xb6f7ae43.
//
// Solidity: function STARTTIMESTAMP_MUST_BE_MULTIPLE_OF_CALCULATION_INTERVAL_SECONDS() view returns(string)
func (_Errors *ErrorsCallerSession) STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS() (string, error) {
	return _Errors.Contract.STARTTIMESTAMPMUSTBEMULTIPLEOFCALCULATIONINTERVALSECONDS(&_Errors.CallOpts)
}

// STARTTIMESTAMPTOOFARINTHEFUTURE is a free data retrieval call binding the contract method 0x15fbec3b.
//
// Solidity: function STARTTIMESTAMP_TOO_FAR_IN_THE_FUTURE() view returns(string)
func (_Errors *ErrorsCaller) STARTTIMESTAMPTOOFARINTHEFUTURE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "STARTTIMESTAMP_TOO_FAR_IN_THE_FUTURE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STARTTIMESTAMPTOOFARINTHEFUTURE is a free data retrieval call binding the contract method 0x15fbec3b.
//
// Solidity: function STARTTIMESTAMP_TOO_FAR_IN_THE_FUTURE() view returns(string)
func (_Errors *ErrorsSession) STARTTIMESTAMPTOOFARINTHEFUTURE() (string, error) {
	return _Errors.Contract.STARTTIMESTAMPTOOFARINTHEFUTURE(&_Errors.CallOpts)
}

// STARTTIMESTAMPTOOFARINTHEFUTURE is a free data retrieval call binding the contract method 0x15fbec3b.
//
// Solidity: function STARTTIMESTAMP_TOO_FAR_IN_THE_FUTURE() view returns(string)
func (_Errors *ErrorsCallerSession) STARTTIMESTAMPTOOFARINTHEFUTURE() (string, error) {
	return _Errors.Contract.STARTTIMESTAMPTOOFARINTHEFUTURE(&_Errors.CallOpts)
}

// STARTTIMESTAMPTOOFARINTHEPAST is a free data retrieval call binding the contract method 0x4e729aee.
//
// Solidity: function STARTTIMESTAMP_TOO_FAR_IN_THE_PAST() view returns(string)
func (_Errors *ErrorsCaller) STARTTIMESTAMPTOOFARINTHEPAST(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "STARTTIMESTAMP_TOO_FAR_IN_THE_PAST")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STARTTIMESTAMPTOOFARINTHEPAST is a free data retrieval call binding the contract method 0x4e729aee.
//
// Solidity: function STARTTIMESTAMP_TOO_FAR_IN_THE_PAST() view returns(string)
func (_Errors *ErrorsSession) STARTTIMESTAMPTOOFARINTHEPAST() (string, error) {
	return _Errors.Contract.STARTTIMESTAMPTOOFARINTHEPAST(&_Errors.CallOpts)
}

// STARTTIMESTAMPTOOFARINTHEPAST is a free data retrieval call binding the contract method 0x4e729aee.
//
// Solidity: function STARTTIMESTAMP_TOO_FAR_IN_THE_PAST() view returns(string)
func (_Errors *ErrorsCallerSession) STARTTIMESTAMPTOOFARINTHEPAST() (string, error) {
	return _Errors.Contract.STARTTIMESTAMPTOOFARINTHEPAST(&_Errors.CallOpts)
}

// THIRDPARTYTRANSFERSDISABLED is a free data retrieval call binding the contract method 0x71925c67.
//
// Solidity: function THIRD_PARTY_TRANSFERS_DISABLED() view returns(string)
func (_Errors *ErrorsCaller) THIRDPARTYTRANSFERSDISABLED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "THIRD_PARTY_TRANSFERS_DISABLED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// THIRDPARTYTRANSFERSDISABLED is a free data retrieval call binding the contract method 0x71925c67.
//
// Solidity: function THIRD_PARTY_TRANSFERS_DISABLED() view returns(string)
func (_Errors *ErrorsSession) THIRDPARTYTRANSFERSDISABLED() (string, error) {
	return _Errors.Contract.THIRDPARTYTRANSFERSDISABLED(&_Errors.CallOpts)
}

// THIRDPARTYTRANSFERSDISABLED is a free data retrieval call binding the contract method 0x71925c67.
//
// Solidity: function THIRD_PARTY_TRANSFERS_DISABLED() view returns(string)
func (_Errors *ErrorsCallerSession) THIRDPARTYTRANSFERSDISABLED() (string, error) {
	return _Errors.Contract.THIRDPARTYTRANSFERSDISABLED(&_Errors.CallOpts)
}

// THIRDPARTYTRANSFERSFORBIDDEN is a free data retrieval call binding the contract method 0xf059972f.
//
// Solidity: function THIRD_PARTY_TRANSFERS_FORBIDDEN() view returns(string)
func (_Errors *ErrorsCaller) THIRDPARTYTRANSFERSFORBIDDEN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "THIRD_PARTY_TRANSFERS_FORBIDDEN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// THIRDPARTYTRANSFERSFORBIDDEN is a free data retrieval call binding the contract method 0xf059972f.
//
// Solidity: function THIRD_PARTY_TRANSFERS_FORBIDDEN() view returns(string)
func (_Errors *ErrorsSession) THIRDPARTYTRANSFERSFORBIDDEN() (string, error) {
	return _Errors.Contract.THIRDPARTYTRANSFERSFORBIDDEN(&_Errors.CallOpts)
}

// THIRDPARTYTRANSFERSFORBIDDEN is a free data retrieval call binding the contract method 0xf059972f.
//
// Solidity: function THIRD_PARTY_TRANSFERS_FORBIDDEN() view returns(string)
func (_Errors *ErrorsCallerSession) THIRDPARTYTRANSFERSFORBIDDEN() (string, error) {
	return _Errors.Contract.THIRDPARTYTRANSFERSFORBIDDEN(&_Errors.CallOpts)
}

// WITHDRAWALDELAYNOTPASSED is a free data retrieval call binding the contract method 0xeb4d0846.
//
// Solidity: function WITHDRAWAL_DELAY_NOT_PASSED() view returns(string)
func (_Errors *ErrorsCaller) WITHDRAWALDELAYNOTPASSED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WITHDRAWAL_DELAY_NOT_PASSED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WITHDRAWALDELAYNOTPASSED is a free data retrieval call binding the contract method 0xeb4d0846.
//
// Solidity: function WITHDRAWAL_DELAY_NOT_PASSED() view returns(string)
func (_Errors *ErrorsSession) WITHDRAWALDELAYNOTPASSED() (string, error) {
	return _Errors.Contract.WITHDRAWALDELAYNOTPASSED(&_Errors.CallOpts)
}

// WITHDRAWALDELAYNOTPASSED is a free data retrieval call binding the contract method 0xeb4d0846.
//
// Solidity: function WITHDRAWAL_DELAY_NOT_PASSED() view returns(string)
func (_Errors *ErrorsCallerSession) WITHDRAWALDELAYNOTPASSED() (string, error) {
	return _Errors.Contract.WITHDRAWALDELAYNOTPASSED(&_Errors.CallOpts)
}

// WITHDRAWERMUSTBESTAKER is a free data retrieval call binding the contract method 0xee6d21aa.
//
// Solidity: function WITHDRAWER_MUST_BE_STAKER() view returns(string)
func (_Errors *ErrorsCaller) WITHDRAWERMUSTBESTAKER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WITHDRAWER_MUST_BE_STAKER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WITHDRAWERMUSTBESTAKER is a free data retrieval call binding the contract method 0xee6d21aa.
//
// Solidity: function WITHDRAWER_MUST_BE_STAKER() view returns(string)
func (_Errors *ErrorsSession) WITHDRAWERMUSTBESTAKER() (string, error) {
	return _Errors.Contract.WITHDRAWERMUSTBESTAKER(&_Errors.CallOpts)
}

// WITHDRAWERMUSTBESTAKER is a free data retrieval call binding the contract method 0xee6d21aa.
//
// Solidity: function WITHDRAWER_MUST_BE_STAKER() view returns(string)
func (_Errors *ErrorsCallerSession) WITHDRAWERMUSTBESTAKER() (string, error) {
	return _Errors.Contract.WITHDRAWERMUSTBESTAKER(&_Errors.CallOpts)
}

// WITHDRAWERNOTSTAKERORGATEWAY is a free data retrieval call binding the contract method 0x95714f7b.
//
// Solidity: function WITHDRAWER_NOT_STAKER_OR_GATEWAY() view returns(string)
func (_Errors *ErrorsCaller) WITHDRAWERNOTSTAKERORGATEWAY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WITHDRAWER_NOT_STAKER_OR_GATEWAY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WITHDRAWERNOTSTAKERORGATEWAY is a free data retrieval call binding the contract method 0x95714f7b.
//
// Solidity: function WITHDRAWER_NOT_STAKER_OR_GATEWAY() view returns(string)
func (_Errors *ErrorsSession) WITHDRAWERNOTSTAKERORGATEWAY() (string, error) {
	return _Errors.Contract.WITHDRAWERNOTSTAKERORGATEWAY(&_Errors.CallOpts)
}

// WITHDRAWERNOTSTAKERORGATEWAY is a free data retrieval call binding the contract method 0x95714f7b.
//
// Solidity: function WITHDRAWER_NOT_STAKER_OR_GATEWAY() view returns(string)
func (_Errors *ErrorsCallerSession) WITHDRAWERNOTSTAKERORGATEWAY() (string, error) {
	return _Errors.Contract.WITHDRAWERNOTSTAKERORGATEWAY(&_Errors.CallOpts)
}

// WITHDRAWAMOUNTSHARESTOOHIGH is a free data retrieval call binding the contract method 0xa57751fe.
//
// Solidity: function WITHDRAW_AMOUNT_SHARES_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCaller) WITHDRAWAMOUNTSHARESTOOHIGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WITHDRAW_AMOUNT_SHARES_TOO_HIGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WITHDRAWAMOUNTSHARESTOOHIGH is a free data retrieval call binding the contract method 0xa57751fe.
//
// Solidity: function WITHDRAW_AMOUNT_SHARES_TOO_HIGH() view returns(string)
func (_Errors *ErrorsSession) WITHDRAWAMOUNTSHARESTOOHIGH() (string, error) {
	return _Errors.Contract.WITHDRAWAMOUNTSHARESTOOHIGH(&_Errors.CallOpts)
}

// WITHDRAWAMOUNTSHARESTOOHIGH is a free data retrieval call binding the contract method 0xa57751fe.
//
// Solidity: function WITHDRAW_AMOUNT_SHARES_TOO_HIGH() view returns(string)
func (_Errors *ErrorsCallerSession) WITHDRAWAMOUNTSHARESTOOHIGH() (string, error) {
	return _Errors.Contract.WITHDRAWAMOUNTSHARESTOOHIGH(&_Errors.CallOpts)
}

// WITHDRAWONLYUNDERLYINGTOKEN is a free data retrieval call binding the contract method 0xbf71cbd2.
//
// Solidity: function WITHDRAW_ONLY_UNDERLYING_TOKEN() view returns(string)
func (_Errors *ErrorsCaller) WITHDRAWONLYUNDERLYINGTOKEN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "WITHDRAW_ONLY_UNDERLYING_TOKEN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WITHDRAWONLYUNDERLYINGTOKEN is a free data retrieval call binding the contract method 0xbf71cbd2.
//
// Solidity: function WITHDRAW_ONLY_UNDERLYING_TOKEN() view returns(string)
func (_Errors *ErrorsSession) WITHDRAWONLYUNDERLYINGTOKEN() (string, error) {
	return _Errors.Contract.WITHDRAWONLYUNDERLYINGTOKEN(&_Errors.CallOpts)
}

// WITHDRAWONLYUNDERLYINGTOKEN is a free data retrieval call binding the contract method 0xbf71cbd2.
//
// Solidity: function WITHDRAW_ONLY_UNDERLYING_TOKEN() view returns(string)
func (_Errors *ErrorsCallerSession) WITHDRAWONLYUNDERLYINGTOKEN() (string, error) {
	return _Errors.Contract.WITHDRAWONLYUNDERLYINGTOKEN(&_Errors.CallOpts)
}

// ZEROADDRESSNOTVALID is a free data retrieval call binding the contract method 0xd14bb17a.
//
// Solidity: function ZERO_ADDRESS_NOT_VALID() view returns(string)
func (_Errors *ErrorsCaller) ZEROADDRESSNOTVALID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ZERO_ADDRESS_NOT_VALID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ZEROADDRESSNOTVALID is a free data retrieval call binding the contract method 0xd14bb17a.
//
// Solidity: function ZERO_ADDRESS_NOT_VALID() view returns(string)
func (_Errors *ErrorsSession) ZEROADDRESSNOTVALID() (string, error) {
	return _Errors.Contract.ZEROADDRESSNOTVALID(&_Errors.CallOpts)
}

// ZEROADDRESSNOTVALID is a free data retrieval call binding the contract method 0xd14bb17a.
//
// Solidity: function ZERO_ADDRESS_NOT_VALID() view returns(string)
func (_Errors *ErrorsCallerSession) ZEROADDRESSNOTVALID() (string, error) {
	return _Errors.Contract.ZEROADDRESSNOTVALID(&_Errors.CallOpts)
}

// ZEROSHARESNOTVALID is a free data retrieval call binding the contract method 0x78833b66.
//
// Solidity: function ZERO_SHARES_NOT_VALID() view returns(string)
func (_Errors *ErrorsCaller) ZEROSHARESNOTVALID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "ZERO_SHARES_NOT_VALID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ZEROSHARESNOTVALID is a free data retrieval call binding the contract method 0x78833b66.
//
// Solidity: function ZERO_SHARES_NOT_VALID() view returns(string)
func (_Errors *ErrorsSession) ZEROSHARESNOTVALID() (string, error) {
	return _Errors.Contract.ZEROSHARESNOTVALID(&_Errors.CallOpts)
}

// ZEROSHARESNOTVALID is a free data retrieval call binding the contract method 0x78833b66.
//
// Solidity: function ZERO_SHARES_NOT_VALID() view returns(string)
func (_Errors *ErrorsCallerSession) ZEROSHARESNOTVALID() (string, error) {
	return _Errors.Contract.ZEROSHARESNOTVALID(&_Errors.CallOpts)
}
