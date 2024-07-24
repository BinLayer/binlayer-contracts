// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DelegationController

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

// IDelegationControllerOperatorDetails is an auto generated low-level Go binding around an user-defined struct.
type IDelegationControllerOperatorDetails struct {
	EarningsReceiver   common.Address
	DelegationApprover common.Address
	StakerOptOutWindow uint32
}

// IDelegationControllerUnstakeParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationControllerUnstakeParams struct {
	Pools      []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

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

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// DelegationControllerMetaData contains all meta data concerning the DelegationController contract.
var DelegationControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_poolController\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_STAKER_OPT_OUT_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_WITHDRAWAL_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakerNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateToBySignature\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"earningsReceiver\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalDelay\",\"inputs\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minWithdrawalDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"_withdrawalDelay\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolWithdrawalDelay\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawalDelay\",\"inputs\":[{\"name\":\"newMinWithdrawalDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolWithdrawalDelay\",\"inputs\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"withdrawalDelay\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerOptOutWindow\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakes\",\"inputs\":[{\"name\":\"unstakeParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.UnstakeParams[]\",\"components\":[{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateWrappedTokenGateway\",\"inputs\":[{\"name\":\"_newWrappedTokenGateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraws\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationController.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wrappedTokenGateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinWithdrawalDelaySet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDetailsModified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolWithdrawalDelaySet\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPool\"},{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateWrappedTokenGateway\",\"inputs\":[{\"name\":\"previousGateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currentGateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationController.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"contractIPool[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200563138038062005631833981016040819052620000349162000137565b6001600160a01b03808316608052811660a052620000516200005d565b50504660c05262000176565b600054610100900460ff1615620000ca5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146200011c576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013457600080fd5b50565b600080604083850312156200014b57600080fd5b825162000158816200011e565b60208401519092506200016b816200011e565b809150509250929050565b60805160a05160c05161545a620001d76000396000612437015260006107be0152600081816105ca01528181610ec10152818161137e01528181611dd601528181612b6f015281816139d401528181613ae00152613f93015261545a6000f3fe608060405234801561001057600080fd5b506004361061036d5760003560e01c806365da1264116101d3578063b6d65ea311610104578063da8be864116100a2578063f16172b01161007c578063f16172b01461099b578063f2fde38b146109ae578063f698da25146109c1578063fabc1cbc146109c957600080fd5b8063da8be86414610962578063eafe9f3114610975578063eea9064b1461098857600080fd5b8063be0525d7116100de578063be0525d714610864578063c5e480db14610877578063c94b51111461092e578063cf80873e1461094157600080fd5b8063b6d65ea314610800578063b7f06ebe14610813578063bb45fef21461083657600080fd5b80639004134711610171578063a17884841161014b578063a17884841461078f578063a238f9df146107af578063b1344271146107b9578063b6805c54146107e057600080fd5b80639004134714610752578063973a67961461077257806399be81c81461077c57600080fd5b8063778e55f3116101ad578063778e55f3146106f05780637f5480711461071b578063886f11951461072e5780638da5cb5b1461074157600080fd5b806365da1264146106915780636d70f7ae146106ba578063715018a6146106e857600080fd5b806322bf40e4116102ad578063595c6a671161024b5780635c975abb116102255780635c975abb1461062a5780635f966f141461063257806363152b131461065e57806365108e921461067e57600080fd5b8063595c6a67146105ec578063597b36da146105f45780635ac86ab71461060757600080fd5b80633cdeb5e0116102875780633cdeb5e0146105315780633e28391d14610560578063433773821461059e5780634aa9d585146105c557600080fd5b806322bf40e4146104eb57806328a573ae146104fe57806329c77d4f1461051157600080fd5b806310d67a2f1161031a5780631bbce091116102f45780631bbce0911461043c5780631d1bf7f21461044f5780631e1198381461049957806320606b70146104c457600080fd5b806310d67a2f14610403578063132d496714610416578063136439dd1461042957600080fd5b80630b9f487a1161034b5780630b9f487a146103d45780630d5b0067146103e75780630f589e59146103f057600080fd5b80630449ca391461037257806304a4f979146103985780630af71b77146103bf575b600080fd5b610385610380366004614337565b6109dc565b6040519081526020015b60405180910390f35b6103857f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6103d26103cd36600461436d565b610a63565b005b6103856103e2366004614456565b610b93565b610385609e5481565b6103d26103fe36600461450b565b610c70565b6103d261041136600461455f565b610d96565b6103d2610424366004614583565b610e82565b6103d26104373660046145c4565b610f4e565b61038561044a366004614583565b6110c0565b61038561045d36600461455f565b6001600160a01b03166000908152609a602052604090206001015474010000000000000000000000000000000000000000900463ffffffff1690565b60a2546104ac906001600160a01b031681565b6040516001600160a01b03909116815260200161038f565b6103857f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6103d26104f93660046145dd565b6110ee565b6103d261050c366004614583565b61133f565b61038561051f36600461455f565b609c6020526000908152604090205481565b6104ac61053f36600461455f565b6001600160a01b039081166000908152609a60205260409020600101541690565b61058e61056e36600461455f565b6001600160a01b039081166000908152609b602052604090205416151590565b604051901515815260200161038f565b6103857f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b81565b6104ac7f000000000000000000000000000000000000000000000000000000000000000081565b6103d2611404565b6103856106023660046148dd565b61153b565b61058e61061536600461491a565b606654600160ff9092169190911b9081161490565b606654610385565b6104ac61064036600461455f565b6001600160a01b039081166000908152609a60205260409020541690565b61067161066c366004614337565b61156b565b60405161038f919061493d565b6103d261068c36600461455f565b61191d565b6104ac61069f36600461455f565b609b602052600090815260409020546001600160a01b031681565b61058e6106c836600461455f565b6001600160a01b039081166000908152609a602052604090205416151590565b6103d26119a6565b6103856106fe366004614981565b609960209081526000928352604080842090915290825290205481565b6103d2610729366004614a63565b6119ba565b6065546104ac906001600160a01b031681565b6033546001600160a01b03166104ac565b610765610760366004614af4565b611a98565b60405161038f9190614b7f565b61038562ed4e0081565b6103d261078a366004614b92565b611b73565b61038561079d36600461455f565b60a06020526000908152604090205481565b61038562278d0081565b6104ac7f000000000000000000000000000000000000000000000000000000000000000081565b6103856107ee36600461455f565b60a16020526000908152604090205481565b6103d261080e366004614bd6565b611c2a565b61058e6108213660046145c4565b609f6020526000908152604090205460ff1681565b61058e610844366004614c66565b609d60209081526000928352604080842090915290825290205460ff1681565b6103d26108723660046145c4565b611caf565b6108f861088536600461455f565b6040805160608082018352600080835260208084018290529284018190526001600160a01b039485168152609a83528390208351918201845280548516825260010154938416918101919091527401000000000000000000000000000000000000000090920463ffffffff169082015290565b6040805182516001600160a01b039081168252602080850151909116908201529181015163ffffffff169082015260600161038f565b61038561093c366004614c92565b611cc0565b61095461094f36600461455f565b611d94565b60405161038f929190614d13565b61067161097036600461455f565b611e51565b6103d2610983366004614d38565b61230c565b6103d2610996366004614da4565b612320565b6103d26109a9366004614dfd565b61232c565b6103d26109bc36600461455f565b6123a6565b610385612433565b6103d26109d73660046145c4565b61250e565b609e54600090815b83811015610a5957600060a16000878785818110610a0457610a04614e19565b9050602002016020810190610a19919061455f565b6001600160a01b03166001600160a01b0316815260200190815260200160002054905082811115610a48578092505b50610a5281614e77565b90506109e4565b5090505b92915050565b60665460408051808201909152600181527f35000000000000000000000000000000000000000000000000000000000000006020820152600291600490811603610ac95760405162461bcd60e51b8152600401610ac09190614ef5565b60405180910390fd5b50610ad2612686565b60005b88811015610b7d57610b6d8a8a83818110610af257610af2614e19565b9050602002810190610b049190614f08565b898984818110610b1657610b16614e19565b9050602002810190610b289190614f46565b898986818110610b3a57610b3a614e19565b90506020020135888887818110610b5357610b53614e19565b9050602002016020810190610b689190614fae565b6126df565b610b7681614e77565b9050610ad5565b50610b88600160d455565b505050505050505050565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad6020808301919091526001600160a01b038681168385015288811660608401528716608083015260a0820185905260c08083018590528351808403909101815260e0909201909252805191012060009081610c11612433565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101919091526042810183905260620160408051808303601f19018152919052805160209091012098975050505050505050565b336000908152609a6020908152604091829020548251808401909352600283527f3230000000000000000000000000000000000000000000000000000000000000918301919091526001600160a01b031615610cdf5760405162461bcd60e51b8152600401610ac09190614ef5565b50610cea3384612d5f565b604080518082019091526060815260006020820152610d0c3380836000612f41565b336001600160a01b03167f8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e285604051610d459190614fcb565b60405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051610d8892919061501d565b60405180910390a250505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610de9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0d919061504c565b6001600160a01b0316336001600160a01b0316146040518060400160405280600181526020017f320000000000000000000000000000000000000000000000000000000000000081525090610e755760405162461bcd60e51b8152600401610ac09190614ef5565b50610e7f816132e9565b50565b60408051808201909152600281527f31390000000000000000000000000000000000000000000000000000000000006020820152336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610eff5760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038084166000908152609b60205260409020541615610f49576001600160a01b038084166000908152609b602052604090205416610f47818585856133c6565b505b505050565b6065546040517f46fbf68e0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610faf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd39190615069565b6040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250906110275760405162461bcd60e51b8152600401610ac09190614ef5565b5060665460408051808201909152600181527f3700000000000000000000000000000000000000000000000000000000000000602082015290828116146110815760405162461bcd60e51b8152600401610ac09190614ef5565b50606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6001600160a01b0383166000908152609c60205260408120546110e585828686611cc0565b95945050505050565b600054610100900460ff161580801561110e5750600054600160ff909116105b806111285750303b158015611128575060005460ff166001145b61119a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610ac0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156111f857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b611202888861344f565b604080518082018252600881527f42696e4c6179657200000000000000000000000000000000000000000000000060209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527fc3e1ab4858383d79d3a0e54f9365ec1da08f828fcebd1952f7ea695e26718a6081840152466060820152306080808301919091528351808303909101815260a090910190925281519101206098556112b461350c565b6112bd89613591565b6112c6866135fb565b6112d285858585613695565b8015610b8857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b60408051808201909152600281527f31390000000000000000000000000000000000000000000000000000000000006020820152336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146113bc5760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038084166000908152609b60205260409020541615610f49576001600160a01b038084166000908152609b602052604090205416610f4781858585613821565b6065546040517f46fbf68e0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611465573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114899190615069565b6040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250906114dd5760405162461bcd60e51b8152600401610ac09190614ef5565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60008160405160200161154e91906150fb565b604051602081830303815290604052805190602001209050919050565b606654606090600190600290811614156040518060400160405280600181526020017f3500000000000000000000000000000000000000000000000000000000000000815250906115cf5760405162461bcd60e51b8152600401610ac09190614ef5565b5060008367ffffffffffffffff8111156115eb576115eb614664565b604051908082528060200260200182016040528015611614578160200160208202803683370190505b50336000908152609b60205260408120549192506001600160a01b03909116905b858110156119125786868281811061164f5761164f614e19565b9050602002810190611661919061510e565b61166f906020810190614f46565b905087878381811061168357611683614e19565b9050602002810190611695919061510e565b61169f9080614f46565b9050146040518060400160405280600281526020017f3235000000000000000000000000000000000000000000000000000000000000815250906116f65760405162461bcd60e51b8152600401610ac09190614ef5565b503387878381811061170a5761170a614e19565b905060200281019061171c919061510e565b61172d90606081019060400161455f565b6001600160a01b03161480611789575060a2546001600160a01b031687878381811061175b5761175b614e19565b905060200281019061176d919061510e565b61177e90606081019060400161455f565b6001600160a01b0316145b6040518060400160405280600281526020017f3236000000000000000000000000000000000000000000000000000000000000815250906117dd5760405162461bcd60e51b8152600401610ac09190614ef5565b506118e333838989858181106117f5576117f5614e19565b9050602002810190611807919061510e565b61181890606081019060400161455f565b8a8a8681811061182a5761182a614e19565b905060200281019061183c919061510e565b6118469080614f46565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508d915088905081811061188c5761188c614e19565b905060200281019061189e919061510e565b6118ac906020810190614f46565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506138aa92505050565b8382815181106118f5576118f5614e19565b60209081029190910101528061190a81614e77565b915050611635565b509095945050505050565b611925613cdf565b60a254604080516001600160a01b03928316815291831660208301527f6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b910160405180910390a160a280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6119ae613cdf565b6119b86000613591565b565b42836020015110156040518060400160405280600281526020017f313300000000000000000000000000000000000000000000000000000000000081525090611a165760405162461bcd60e51b8152600401610ac09190614ef5565b506000609c6000876001600160a01b03166001600160a01b031681526020019081526020016000205490506000611a538783888860200151611cc0565b6001600160a01b0388166000908152609c602052604090206001840190558551909150611a839088908390613d39565b611a8f87878686612f41565b50505050505050565b60606000825167ffffffffffffffff811115611ab657611ab6614664565b604051908082528060200260200182016040528015611adf578160200160208202803683370190505b50905060005b8351811015610a59576001600160a01b03851660009081526099602052604081208551909190869084908110611b1d57611b1d614e19565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054828281518110611b5857611b58614e19565b6020908102919091010152611b6c81614e77565b9050611ae5565b336000908152609a60205260409020546001600160a01b031615156040518060400160405280600281526020017f323100000000000000000000000000000000000000000000000000000000000081525090611be25760405162461bcd60e51b8152600401610ac09190614ef5565b50336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051611c1e92919061501d565b60405180910390a25050565b60665460408051808201909152600181527f35000000000000000000000000000000000000000000000000000000000000006020820152600291600490811603611c875760405162461bcd60e51b8152600401610ac09190614ef5565b50611c90612686565b611c9d86868686866126df565b611ca7600160d455565b505050505050565b611cb7613cdf565b610e7f816135fb565b604080517f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b6020808301919091526001600160a01b0387811683850152851660608301526080820186905260a08083018590528351808403909101815260c0909201909252805191012060009081611d36612433565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101919091526042810183905260620160408051808303601f190181529190528051602090910120979650505050505050565b6040517f94f649dd0000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152606091829160009182917f000000000000000000000000000000000000000000000000000000000000000016906394f649dd90602401600060405180830381865afa158015611e1d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611e45919081019061519d565b90969095509350505050565b606654606090600190600290811614156040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525090611eb55760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038084166000908152609b60205260409020541615156040518060400160405280600281526020017f323200000000000000000000000000000000000000000000000000000000000081525090611f275760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038084166000908152609a6020526040902054161515156040518060400160405280600281526020017f323300000000000000000000000000000000000000000000000000000000000081525090611f9a5760405162461bcd60e51b8152600401610ac09190614ef5565b5060408051808201909152600181527f330000000000000000000000000000000000000000000000000000000000000060208201526001600160a01b038416611ff65760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038084166000818152609b60205260409020549091169033148061202a5750336001600160a01b038216145b8061205157506001600160a01b038181166000908152609a60205260409020600101541633145b6040518060400160405280600281526020017f3234000000000000000000000000000000000000000000000000000000000000815250906120a55760405162461bcd60e51b8152600401610ac09190614ef5565b506000806120b286611d94565b9092509050336001600160a01b0387161461210857826001600160a01b0316866001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a35b826001600160a01b0316866001600160a01b03167ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467660405160405180910390a36001600160a01b0386166000908152609b6020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055825190036121a4576040805160008152602081019091529450612303565b815167ffffffffffffffff8111156121be576121be614664565b6040519080825280602002602001820160405280156121e7578160200160208202803683370190505b50945060005b82518110156123015760408051600180825281830190925260009160208083019080368337505060408051600180825281830190925292935060009291506020808301908036833701905050905084838151811061224d5761224d614e19565b60200260200101518260008151811061226857612268614e19565b60200260200101906001600160a01b031690816001600160a01b03168152505083838151811061229a5761229a614e19565b6020026020010151816000815181106122b5576122b5614e19565b6020026020010181815250506122ce89878b85856138aa565b8884815181106122e0576122e0614e19565b602002602001018181525050505080806122f990614e77565b9150506121ed565b505b50505050919050565b612314613cdf565b610f4784848484613695565b610f4933848484612f41565b336000908152609a60205260409020546001600160a01b031615156040518060400160405280600281526020017f32310000000000000000000000000000000000000000000000000000000000008152509061239b5760405162461bcd60e51b8152600401610ac09190614ef5565b50610e7f3382612d5f565b6123ae613cdf565b6001600160a01b03811661242a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610ac0565b610e7f81613591565b60007f00000000000000000000000000000000000000000000000000000000000000004603612463575060985490565b50604080518082018252600881527f42696e4c6179657200000000000000000000000000000000000000000000000060209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527fc3e1ab4858383d79d3a0e54f9365ec1da08f828fcebd1952f7ea695e26718a6081840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612561573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612585919061504c565b6001600160a01b0316336001600160a01b0316146040518060400160405280600181526020017f3200000000000000000000000000000000000000000000000000000000000000815250906125ed5760405162461bcd60e51b8152600401610ac09190614ef5565b506066541981196066541916146040518060400160405280600181526020017f38000000000000000000000000000000000000000000000000000000000000008152509061264e5760405162461bcd60e51b8152600401610ac09190614ef5565b50606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016110b5565b600260d454036126d85760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610ac0565b600260d455565b60006126ed61060287615258565b6000818152609f6020908152604091829020548251808401909352600283527f3332000000000000000000000000000000000000000000000000000000000000918301919091529192509060ff166127585760405162461bcd60e51b8152600401610ac09190614ef5565b50609e54429061276e60a0890160808a01615264565b63ffffffff1661277e9190615281565b11156040518060400160405280600281526020017f3333000000000000000000000000000000000000000000000000000000000000815250906127d45760405162461bcd60e51b8152600401610ac09190614ef5565b506127e5606087016040880161455f565b6001600160a01b0316336001600160a01b0316146040518060400160405280600281526020017f33340000000000000000000000000000000000000000000000000000000000008152509061284d5760405162461bcd60e51b8152600401610ac09190614ef5565b5081156128b85761286160a0870187614f46565b60408051808201909152600281527f32350000000000000000000000000000000000000000000000000000000000006020820152915085146128b65760405162461bcd60e51b8152600401610ac09190614ef5565b505b6000818152609f6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558115612a745760005b61290260a0880188614f46565b9050811015612a6e574260a1600061291d60a08b018b614f46565b8581811061292d5761292d614e19565b9050602002016020810190612942919061455f565b6001600160a01b0316815260208101919091526040016000205461296c60a08a0160808b01615264565b63ffffffff1661297c9190615281565b11156040518060400160405280600281526020017f3335000000000000000000000000000000000000000000000000000000000000815250906129d25760405162461bcd60e51b8152600401610ac09190614ef5565b50612a666129e3602089018961455f565b336129f160a08b018b614f46565b85818110612a0157612a01614e19565b9050602002016020810190612a16919061455f565b612a2360c08c018c614f46565b86818110612a3357612a33614e19565b905060200201358a8a87818110612a4c57612a4c614e19565b9050602002016020810190612a61919061455f565b613f44565b6001016128f5565b50612d1d565b336000908152609b60205260408120546001600160a01b0316905b612a9c60a0890189614f46565b9050811015612d1a574260a16000612ab760a08c018c614f46565b85818110612ac757612ac7614e19565b9050602002016020810190612adc919061455f565b6001600160a01b03168152602081019190915260400160002054612b0660a08b0160808c01615264565b63ffffffff16612b169190615281565b11156040518060400160405280600281526020017f333500000000000000000000000000000000000000000000000000000000000081525090612b6c5760405162461bcd60e51b8152600401610ac09190614ef5565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c4623ea133898985818110612baf57612baf614e19565b9050602002016020810190612bc4919061455f565b612bd160a08d018d614f46565b86818110612be157612be1614e19565b9050602002016020810190612bf6919061455f565b612c0360c08e018e614f46565b87818110612c1357612c13614e19565b60405160e088901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b03968716600482015294861660248601529290941660448401526020909102013560648201526084019050600060405180830381600087803b158015612c8b57600080fd5b505af1158015612c9f573d6000803e3d6000fd5b505050506001600160a01b03821615612d1257612d128233612cc460a08c018c614f46565b85818110612cd457612cd4614e19565b9050602002016020810190612ce9919061455f565b612cf660c08d018d614f46565b86818110612d0657612d06614e19565b90506020020135613821565b600101612a8f565b50505b6040518181527fc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d9060200160405180910390a1505050505050565b600160d455565b6000612d6e602083018361455f565b6001600160a01b031614156040518060400160405280600181526020017f330000000000000000000000000000000000000000000000000000000000000081525090612dcd5760405162461bcd60e51b8152600401610ac09190614ef5565b5062ed4e00612de26060830160408401615264565b63ffffffff1611156040518060400160405280600281526020017f323700000000000000000000000000000000000000000000000000000000000081525090612e3e5760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b0382166000908152609a6020526040908190206001015474010000000000000000000000000000000000000000900463ffffffff1690612e8c9060608401908401615264565b63ffffffff1610156040518060400160405280600281526020017f323800000000000000000000000000000000000000000000000000000000000081525090612ee85760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b0382166000908152609a602052604090208190612f0d8282615294565b505060405133907ffebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac90611c1e908490614fcb565b6066546040805180820190915260018082527f35000000000000000000000000000000000000000000000000000000000000006020830152600092811603612f9c5760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038086166000908152609b6020526040902054161515156040518060400160405280600281526020017f32390000000000000000000000000000000000000000000000000000000000008152509061300f5760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038085166000908152609a60205260409020541615156040518060400160405280600281526020017f3330000000000000000000000000000000000000000000000000000000000000815250906130815760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038085166000908152609a60205260409020600101541680158015906130b85750336001600160a01b03821614155b80156130cd5750336001600160a01b03861614155b1561321c5742846020015110156040518060400160405280600281526020017f31330000000000000000000000000000000000000000000000000000000000008152509061312e5760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b0381166000908152609d60209081526040808320868452825291829020548251808401909352600283527f33310000000000000000000000000000000000000000000000000000000000009183019190915260ff16156131a95760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b0381166000908152609d602090815260408083208684528252822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055850151613209908890889085908890610b93565b905061321a82828760000151613d39565b505b6001600160a01b038681166000818152609b602052604080822080547fffffffffffffffffffffffff000000000000000000000000000000000000000016948a169485179055517fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d87433049190a360008061329388611d94565b9150915060005b8251811015610b88576132e1888a8584815181106132ba576132ba614e19565b60200260200101518585815181106132d4576132d4614e19565b6020026020010151613821565b60010161329a565b60408051808201909152600181527f330000000000000000000000000000000000000000000000000000000000000060208201526001600160a01b0382166133445760405162461bcd60e51b8152600401610ac09190614ef5565b50606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b038085166000908152609960209081526040808320938616835292905290812080548392906133fd90849061536e565b9091555050604080516001600160a01b0385811682528481166020830152918101839052908516907f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd90606001610d88565b6065546001600160a01b031615801561347057506001600160a01b03821615155b6040518060400160405280600181526020017f3600000000000000000000000000000000000000000000000000000000000000815250906134c45760405162461bcd60e51b8152600401610ac09190614ef5565b50606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613508826132e9565b5050565b600054610100900460ff166135895760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610ac0565b6119b8613feb565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60408051808201909152600281527f3338000000000000000000000000000000000000000000000000000000000000602082015262278d008211156136535760405162461bcd60e51b8152600401610ac09190614ef5565b50609e5460408051918252602082018390527f338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3910160405180910390a1609e55565b60408051808201909152600281527f323500000000000000000000000000000000000000000000000000000000000060208201528382146136e95760405162461bcd60e51b8152600401610ac09190614ef5565b508260005b81811015611ca757600086868381811061370a5761370a614e19565b905060200201602081019061371f919061455f565b6001600160a01b038116600090815260a1602052604081205491925086868581811061374d5761374d614e19565b90506020020135905062278d008111156040518060400160405280600281526020017f3339000000000000000000000000000000000000000000000000000000000000815250906137b15760405162461bcd60e51b8152600401610ac09190614ef5565b506001600160a01b038316600081815260a160209081526040918290208490558151928352820184905281018290527f108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc29060600160405180910390a15050508061381a90614e77565b90506136ee565b6001600160a01b03808516600090815260996020908152604080832093861683529290529081208054839290613858908490615281565b9091555050604080516001600160a01b0385811682528481166020830152918101839052908516907f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c90606001610d88565b60408051808201909152600181527f330000000000000000000000000000000000000000000000000000000000000060208201526000906001600160a01b0387166139085760405162461bcd60e51b8152600401610ac09190614ef5565b50825160408051808201909152600281527f333600000000000000000000000000000000000000000000000000000000000060208201529061395d5760405162461bcd60e51b8152600401610ac09190614ef5565b5060005b8351811015613bcf576001600160a01b038616156139b7576139b7868886848151811061399057613990614e19565b60200260200101518685815181106139aa576139aa614e19565b60200260200101516133c6565b846001600160a01b0316876001600160a01b03161480613a8957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639b4da03d858381518110613a1357613a13614e19565b60200260200101516040518263ffffffff1660e01b8152600401613a4691906001600160a01b0391909116815260200190565b602060405180830381865afa158015613a63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a879190615069565b155b6040518060400160405280600281526020017f333700000000000000000000000000000000000000000000000000000000000081525090613add5760405162461bcd60e51b8152600401610ac09190614ef5565b507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638c80d4e588868481518110613b2057613b20614e19565b6020026020010151868581518110613b3a57613b3a614e19565b60209081029190910101516040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606401600060405180830381600087803b158015613bac57600080fd5b505af1158015613bc0573d6000803e3d6000fd5b50505050806001019050613961565b506001600160a01b038616600090815260a060205260408120805491829190613bf783614e77565b919050555060006040518060e00160405280896001600160a01b03168152602001886001600160a01b03168152602001876001600160a01b031681526020018381526020014263ffffffff1681526020018681526020018581525090506000613c5f8261153b565b6000818152609f60205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055519091507f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f990613ccb9083908590615381565b60405180910390a198975050505050505050565b6033546001600160a01b031633146119b85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610ac0565b6001600160a01b0383163b15613e8e576040517f1626ba7e00000000000000000000000000000000000000000000000000000000808252906001600160a01b03851690631626ba7e90613d92908690869060040161539a565b602060405180830381865afa158015613daf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dd391906153b3565b7fffffffff000000000000000000000000000000000000000000000000000000001614610f495760405162461bcd60e51b815260206004820152605360248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a2045524331323731207369676e61747572652060648201527f766572696669636174696f6e206661696c656400000000000000000000000000608482015260a401610ac0565b826001600160a01b0316613ea28383614068565b6001600160a01b031614610f495760405162461bcd60e51b815260206004820152604760248201527f454950313237315369676e61747572655574696c732e636865636b5369676e6160448201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d60648201527f207369676e657200000000000000000000000000000000000000000000000000608482015260a401610ac0565b6040517fc608c7f30000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015284811660248301526044820184905282811660648301527f0000000000000000000000000000000000000000000000000000000000000000169063c608c7f390608401600060405180830381600087803b158015613fd757600080fd5b505af1158015610b88573d6000803e3d6000fd5b600054610100900460ff16612d585760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610ac0565b60008060006140778585614084565b91509150610a59816140c9565b60008082516041036140ba5760208301516040840151606085015160001a6140ae8782858561422e565b945094505050506140c2565b506000905060025b9250929050565b60008160048111156140dd576140dd6153f5565b036140e55750565b60018160048111156140f9576140f96153f5565b036141465760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610ac0565b600281600481111561415a5761415a6153f5565b036141a75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610ac0565b60038160048111156141bb576141bb6153f5565b03610e7f5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610ac0565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561426557506000905060036142e9565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156142b9573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166142e2576000600192509250506142e9565b9150600090505b94509492505050565b60008083601f84011261430457600080fd5b50813567ffffffffffffffff81111561431c57600080fd5b6020830191508360208260051b85010111156140c257600080fd5b6000806020838503121561434a57600080fd5b823567ffffffffffffffff81111561436157600080fd5b611e45858286016142f2565b6000806000806000806000806080898b03121561438957600080fd5b883567ffffffffffffffff808211156143a157600080fd5b6143ad8c838d016142f2565b909a50985060208b01359150808211156143c657600080fd5b6143d28c838d016142f2565b909850965060408b01359150808211156143eb57600080fd5b6143f78c838d016142f2565b909650945060608b013591508082111561441057600080fd5b5061441d8b828c016142f2565b999c989b5096995094979396929594505050565b6001600160a01b0381168114610e7f57600080fd5b803561445181614431565b919050565b600080600080600060a0868803121561446e57600080fd5b853561447981614431565b9450602086013561448981614431565b9350604086013561449981614431565b94979396509394606081013594506080013592915050565b6000606082840312156144c357600080fd5b50919050565b60008083601f8401126144db57600080fd5b50813567ffffffffffffffff8111156144f357600080fd5b6020830191508360208285010111156140c257600080fd5b60008060006080848603121561452057600080fd5b61452a85856144b1565b9250606084013567ffffffffffffffff81111561454657600080fd5b614552868287016144c9565b9497909650939450505050565b60006020828403121561457157600080fd5b813561457c81614431565b9392505050565b60008060006060848603121561459857600080fd5b83356145a381614431565b925060208401356145b381614431565b929592945050506040919091013590565b6000602082840312156145d657600080fd5b5035919050565b60008060008060008060008060c0898b0312156145f957600080fd5b883561460481614431565b9750602089013561461481614431565b96506040890135955060608901359450608089013567ffffffffffffffff8082111561463f57600080fd5b61464b8c838d016142f2565b909650945060a08b013591508082111561441057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff811182821017156146b6576146b6614664565b60405290565b6040805190810167ffffffffffffffff811182821017156146b6576146b6614664565b604051601f8201601f1916810167ffffffffffffffff8111828210171561470857614708614664565b604052919050565b63ffffffff81168114610e7f57600080fd5b803561445181614710565b600067ffffffffffffffff82111561474757614747614664565b5060051b60200190565b600082601f83011261476257600080fd5b813560206147776147728361472d565b6146df565b82815260059290921b8401810191818101908684111561479657600080fd5b8286015b848110156147ba5780356147ad81614431565b835291830191830161479a565b509695505050505050565b600082601f8301126147d657600080fd5b813560206147e66147728361472d565b82815260059290921b8401810191818101908684111561480557600080fd5b8286015b848110156147ba5780358352918301918301614809565b600060e0828403121561483257600080fd5b61483a614693565b905061484582614446565b815261485360208301614446565b602082015261486460408301614446565b60408201526060820135606082015261487f60808301614722565b608082015260a082013567ffffffffffffffff8082111561489f57600080fd5b6148ab85838601614751565b60a084015260c08401359150808211156148c457600080fd5b506148d1848285016147c5565b60c08301525092915050565b6000602082840312156148ef57600080fd5b813567ffffffffffffffff81111561490657600080fd5b61491284828501614820565b949350505050565b60006020828403121561492c57600080fd5b813560ff8116811461457c57600080fd5b6020808252825182820181905260009190848201906040850190845b8181101561497557835183529284019291840191600101614959565b50909695505050505050565b6000806040838503121561499457600080fd5b823561499f81614431565b915060208301356149af81614431565b809150509250929050565b6000604082840312156149cc57600080fd5b6149d46146bc565b9050813567ffffffffffffffff808211156149ee57600080fd5b818401915084601f830112614a0257600080fd5b8135602082821115614a1657614a16614664565b614a2881601f19601f850116016146df565b92508183528681838601011115614a3e57600080fd5b8181850182850137600081838501015282855280860135818601525050505092915050565b600080600080600060a08688031215614a7b57600080fd5b8535614a8681614431565b94506020860135614a9681614431565b9350604086013567ffffffffffffffff80821115614ab357600080fd5b614abf89838a016149ba565b94506060880135915080821115614ad557600080fd5b50614ae2888289016149ba565b95989497509295608001359392505050565b60008060408385031215614b0757600080fd5b8235614b1281614431565b9150602083013567ffffffffffffffff811115614b2e57600080fd5b614b3a85828601614751565b9150509250929050565b600081518084526020808501945080840160005b83811015614b7457815187529582019590820190600101614b58565b509495945050505050565b60208152600061457c6020830184614b44565b60008060208385031215614ba557600080fd5b823567ffffffffffffffff811115614bbc57600080fd5b611e45858286016144c9565b8015158114610e7f57600080fd5b600080600080600060808688031215614bee57600080fd5b853567ffffffffffffffff80821115614c0657600080fd5b9087019060e0828a031215614c1a57600080fd5b90955060208701359080821115614c3057600080fd5b50614c3d888289016142f2565b909550935050604086013591506060860135614c5881614bc8565b809150509295509295909350565b60008060408385031215614c7957600080fd5b8235614c8481614431565b946020939093013593505050565b60008060008060808587031215614ca857600080fd5b8435614cb381614431565b9350602085013592506040850135614cca81614431565b9396929550929360600135925050565b600081518084526020808501945080840160005b83811015614b745781516001600160a01b031687529582019590820190600101614cee565b604081526000614d266040830185614cda565b82810360208401526110e58185614b44565b60008060008060408587031215614d4e57600080fd5b843567ffffffffffffffff80821115614d6657600080fd5b614d72888389016142f2565b90965094506020870135915080821115614d8b57600080fd5b50614d98878288016142f2565b95989497509550505050565b600080600060608486031215614db957600080fd5b8335614dc481614431565b9250602084013567ffffffffffffffff811115614de057600080fd5b614dec868287016149ba565b925050604084013590509250925092565b600060608284031215614e0f57600080fd5b61457c83836144b1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614ea857614ea8614e48565b5060010190565b6000815180845260005b81811015614ed557602081850181015186830182015201614eb9565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061457c6020830184614eaf565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21833603018112614f3c57600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614f7b57600080fd5b83018035915067ffffffffffffffff821115614f9657600080fd5b6020019150600581901b36038213156140c257600080fd5b600060208284031215614fc057600080fd5b813561457c81614bc8565b606081018235614fda81614431565b6001600160a01b039081168352602084013590614ff682614431565b166020830152604083013561500a81614710565b63ffffffff811660408401525092915050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561505e57600080fd5b815161457c81614431565b60006020828403121561507b57600080fd5b815161457c81614bc8565b60006001600160a01b03808351168452806020840151166020850152806040840151166040850152506060820151606084015263ffffffff608083015116608084015260a082015160e060a08501526150e260e0850182614cda565b905060c083015184820360c08601526110e58282614b44565b60208152600061457c6020830184615086565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112614f3c57600080fd5b600082601f83011261515357600080fd5b815160206151636147728361472d565b82815260059290921b8401810191818101908684111561518257600080fd5b8286015b848110156147ba5780518352918301918301615186565b600080604083850312156151b057600080fd5b825167ffffffffffffffff808211156151c857600080fd5b818501915085601f8301126151dc57600080fd5b815160206151ec6147728361472d565b82815260059290921b8401810191818101908984111561520b57600080fd5b948201945b8386101561523257855161522381614431565b82529482019490820190615210565b9188015191965090935050508082111561524b57600080fd5b50614b3a85828601615142565b6000610a5d3683614820565b60006020828403121561527657600080fd5b813561457c81614710565b80820180821115610a5d57610a5d614e48565b813561529f81614431565b81547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038216178255506001810160208301356152e381614431565b81547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03821617825550604083013561532381614710565b81547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff1660a09190911b77ffffffff0000000000000000000000000000000000000000161790555050565b81810381811115610a5d57610a5d614e48565b8281526040602082015260006149126040830184615086565b8281526040602082015260006149126040830184614eaf565b6000602082840312156153c557600080fd5b81517fffffffff000000000000000000000000000000000000000000000000000000008116811461457c57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea2646970667358221220c701d4cf114dbd358849c4b73b92cef9f082bde423e572d2437b40bdb27e90c264736f6c63430008140033",
}

// DelegationControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegationControllerMetaData.ABI instead.
var DelegationControllerABI = DelegationControllerMetaData.ABI

// DelegationControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegationControllerMetaData.Bin instead.
var DelegationControllerBin = DelegationControllerMetaData.Bin

// DeployDelegationController deploys a new Ethereum contract, binding an instance of DelegationController to it.
func DeployDelegationController(auth *bind.TransactOpts, backend bind.ContractBackend, _poolController common.Address, _slasher common.Address) (common.Address, *types.Transaction, *DelegationController, error) {
	parsed, err := DelegationControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegationControllerBin), backend, _poolController, _slasher)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelegationController{DelegationControllerCaller: DelegationControllerCaller{contract: contract}, DelegationControllerTransactor: DelegationControllerTransactor{contract: contract}, DelegationControllerFilterer: DelegationControllerFilterer{contract: contract}}, nil
}

// DelegationController is an auto generated Go binding around an Ethereum contract.
type DelegationController struct {
	DelegationControllerCaller     // Read-only binding to the contract
	DelegationControllerTransactor // Write-only binding to the contract
	DelegationControllerFilterer   // Log filterer for contract events
}

// DelegationControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegationControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegationControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegationControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegationControllerSession struct {
	Contract     *DelegationController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DelegationControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegationControllerCallerSession struct {
	Contract *DelegationControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DelegationControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegationControllerTransactorSession struct {
	Contract     *DelegationControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DelegationControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegationControllerRaw struct {
	Contract *DelegationController // Generic contract binding to access the raw methods on
}

// DelegationControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegationControllerCallerRaw struct {
	Contract *DelegationControllerCaller // Generic read-only contract binding to access the raw methods on
}

// DelegationControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegationControllerTransactorRaw struct {
	Contract *DelegationControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegationController creates a new instance of DelegationController, bound to a specific deployed contract.
func NewDelegationController(address common.Address, backend bind.ContractBackend) (*DelegationController, error) {
	contract, err := bindDelegationController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegationController{DelegationControllerCaller: DelegationControllerCaller{contract: contract}, DelegationControllerTransactor: DelegationControllerTransactor{contract: contract}, DelegationControllerFilterer: DelegationControllerFilterer{contract: contract}}, nil
}

// NewDelegationControllerCaller creates a new read-only instance of DelegationController, bound to a specific deployed contract.
func NewDelegationControllerCaller(address common.Address, caller bind.ContractCaller) (*DelegationControllerCaller, error) {
	contract, err := bindDelegationController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerCaller{contract: contract}, nil
}

// NewDelegationControllerTransactor creates a new write-only instance of DelegationController, bound to a specific deployed contract.
func NewDelegationControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegationControllerTransactor, error) {
	contract, err := bindDelegationController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerTransactor{contract: contract}, nil
}

// NewDelegationControllerFilterer creates a new log filterer instance of DelegationController, bound to a specific deployed contract.
func NewDelegationControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegationControllerFilterer, error) {
	contract, err := bindDelegationController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerFilterer{contract: contract}, nil
}

// bindDelegationController binds a generic wrapper to an already deployed contract.
func bindDelegationController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegationControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationController *DelegationControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegationController.Contract.DelegationControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationController *DelegationControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationController.Contract.DelegationControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationController *DelegationControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationController.Contract.DelegationControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationController *DelegationControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegationController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationController *DelegationControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationController *DelegationControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationController.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _DelegationController.Contract.DELEGATIONAPPROVALTYPEHASH(&_DelegationController.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _DelegationController.Contract.DELEGATIONAPPROVALTYPEHASH(&_DelegationController.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DelegationController.Contract.DOMAINTYPEHASH(&_DelegationController.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DelegationController.Contract.DOMAINTYPEHASH(&_DelegationController.CallOpts)
}

// MAXSTAKEROPTOUTWINDOW is a free data retrieval call binding the contract method 0x973a6796.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW() view returns(uint256)
func (_DelegationController *DelegationControllerCaller) MAXSTAKEROPTOUTWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "MAX_STAKER_OPT_OUT_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKEROPTOUTWINDOW is a free data retrieval call binding the contract method 0x973a6796.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW() view returns(uint256)
func (_DelegationController *DelegationControllerSession) MAXSTAKEROPTOUTWINDOW() (*big.Int, error) {
	return _DelegationController.Contract.MAXSTAKEROPTOUTWINDOW(&_DelegationController.CallOpts)
}

// MAXSTAKEROPTOUTWINDOW is a free data retrieval call binding the contract method 0x973a6796.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW() view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) MAXSTAKEROPTOUTWINDOW() (*big.Int, error) {
	return _DelegationController.Contract.MAXSTAKEROPTOUTWINDOW(&_DelegationController.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_DelegationController *DelegationControllerCaller) MAXWITHDRAWALDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_DelegationController *DelegationControllerSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _DelegationController.Contract.MAXWITHDRAWALDELAY(&_DelegationController.CallOpts)
}

// MAXWITHDRAWALDELAY is a free data retrieval call binding the contract method 0xa238f9df.
//
// Solidity: function MAX_WITHDRAWAL_DELAY() view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) MAXWITHDRAWALDELAY() (*big.Int, error) {
	return _DelegationController.Contract.MAXWITHDRAWALDELAY(&_DelegationController.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) STAKERDELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "STAKER_DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _DelegationController.Contract.STAKERDELEGATIONTYPEHASH(&_DelegationController.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _DelegationController.Contract.STAKERDELEGATIONTYPEHASH(&_DelegationController.CallOpts)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) CalculateCurrentStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "calculateCurrentStakerDelegationDigestHash", staker, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationController.Contract.CalculateCurrentStakerDelegationDigestHash(&_DelegationController.CallOpts, staker, operator, expiry)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationController.Contract.CalculateCurrentStakerDelegationDigestHash(&_DelegationController.CallOpts, staker, operator, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, _delegationApprover, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationController.Contract.CalculateDelegationApprovalDigestHash(&_DelegationController.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationController.Contract.CalculateDelegationApprovalDigestHash(&_DelegationController.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) CalculateStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "calculateStakerDelegationDigestHash", staker, _stakerNonce, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationController.Contract.CalculateStakerDelegationDigestHash(&_DelegationController.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationController.Contract.CalculateStakerDelegationDigestHash(&_DelegationController.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationController *DelegationControllerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationController *DelegationControllerSession) CalculateWithdrawalRoot(withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	return _DelegationController.Contract.CalculateWithdrawalRoot(&_DelegationController.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationControllerWithdrawal) ([32]byte, error) {
	return _DelegationController.Contract.CalculateWithdrawalRoot(&_DelegationController.CallOpts, withdrawal)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationController *DelegationControllerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationController *DelegationControllerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.CumulativeWithdrawalsQueued(&_DelegationController.CallOpts, arg0)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.CumulativeWithdrawalsQueued(&_DelegationController.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationController *DelegationControllerCaller) DelegatedTo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "delegatedTo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationController *DelegationControllerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _DelegationController.Contract.DelegatedTo(&_DelegationController.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationController *DelegationControllerCallerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _DelegationController.Contract.DelegatedTo(&_DelegationController.CallOpts, arg0)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationController *DelegationControllerCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationController *DelegationControllerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _DelegationController.Contract.DelegationApprover(&_DelegationController.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationController *DelegationControllerCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _DelegationController.Contract.DelegationApprover(&_DelegationController.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationController *DelegationControllerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "delegationApproverSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationController *DelegationControllerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _DelegationController.Contract.DelegationApproverSaltIsSpent(&_DelegationController.CallOpts, arg0, arg1)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationController *DelegationControllerCallerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _DelegationController.Contract.DelegationApproverSaltIsSpent(&_DelegationController.CallOpts, arg0, arg1)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationController *DelegationControllerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationController *DelegationControllerSession) DomainSeparator() ([32]byte, error) {
	return _DelegationController.Contract.DomainSeparator(&_DelegationController.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationController *DelegationControllerCallerSession) DomainSeparator() ([32]byte, error) {
	return _DelegationController.Contract.DomainSeparator(&_DelegationController.CallOpts)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_DelegationController *DelegationControllerCaller) EarningsReceiver(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "earningsReceiver", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_DelegationController *DelegationControllerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _DelegationController.Contract.EarningsReceiver(&_DelegationController.CallOpts, operator)
}

// EarningsReceiver is a free data retrieval call binding the contract method 0x5f966f14.
//
// Solidity: function earningsReceiver(address operator) view returns(address)
func (_DelegationController *DelegationControllerCallerSession) EarningsReceiver(operator common.Address) (common.Address, error) {
	return _DelegationController.Contract.EarningsReceiver(&_DelegationController.CallOpts, operator)
}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_DelegationController *DelegationControllerCaller) GetDelegatableShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "getDelegatableShares", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_DelegationController *DelegationControllerSession) GetDelegatableShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _DelegationController.Contract.GetDelegatableShares(&_DelegationController.CallOpts, staker)
}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_DelegationController *DelegationControllerCallerSession) GetDelegatableShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _DelegationController.Contract.GetDelegatableShares(&_DelegationController.CallOpts, staker)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_DelegationController *DelegationControllerCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address, pools []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "getOperatorShares", operator, pools)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_DelegationController *DelegationControllerSession) GetOperatorShares(operator common.Address, pools []common.Address) ([]*big.Int, error) {
	return _DelegationController.Contract.GetOperatorShares(&_DelegationController.CallOpts, operator, pools)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] pools) view returns(uint256[])
func (_DelegationController *DelegationControllerCallerSession) GetOperatorShares(operator common.Address, pools []common.Address) ([]*big.Int, error) {
	return _DelegationController.Contract.GetOperatorShares(&_DelegationController.CallOpts, operator, pools)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_DelegationController *DelegationControllerCaller) GetWithdrawalDelay(opts *bind.CallOpts, pools []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "getWithdrawalDelay", pools)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_DelegationController *DelegationControllerSession) GetWithdrawalDelay(pools []common.Address) (*big.Int, error) {
	return _DelegationController.Contract.GetWithdrawalDelay(&_DelegationController.CallOpts, pools)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] pools) view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) GetWithdrawalDelay(pools []common.Address) (*big.Int, error) {
	return _DelegationController.Contract.GetWithdrawalDelay(&_DelegationController.CallOpts, pools)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationController *DelegationControllerCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationController *DelegationControllerSession) IsDelegated(staker common.Address) (bool, error) {
	return _DelegationController.Contract.IsDelegated(&_DelegationController.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationController *DelegationControllerCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _DelegationController.Contract.IsDelegated(&_DelegationController.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationController *DelegationControllerCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationController *DelegationControllerSession) IsOperator(operator common.Address) (bool, error) {
	return _DelegationController.Contract.IsOperator(&_DelegationController.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationController *DelegationControllerCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _DelegationController.Contract.IsOperator(&_DelegationController.CallOpts, operator)
}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_DelegationController *DelegationControllerCaller) MinWithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "minWithdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_DelegationController *DelegationControllerSession) MinWithdrawalDelay() (*big.Int, error) {
	return _DelegationController.Contract.MinWithdrawalDelay(&_DelegationController.CallOpts)
}

// MinWithdrawalDelay is a free data retrieval call binding the contract method 0x0d5b0067.
//
// Solidity: function minWithdrawalDelay() view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) MinWithdrawalDelay() (*big.Int, error) {
	return _DelegationController.Contract.MinWithdrawalDelay(&_DelegationController.CallOpts)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationController *DelegationControllerCaller) OperatorDetails(opts *bind.CallOpts, operator common.Address) (IDelegationControllerOperatorDetails, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "operatorDetails", operator)

	if err != nil {
		return *new(IDelegationControllerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationControllerOperatorDetails)).(*IDelegationControllerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationController *DelegationControllerSession) OperatorDetails(operator common.Address) (IDelegationControllerOperatorDetails, error) {
	return _DelegationController.Contract.OperatorDetails(&_DelegationController.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationController *DelegationControllerCallerSession) OperatorDetails(operator common.Address) (IDelegationControllerOperatorDetails, error) {
	return _DelegationController.Contract.OperatorDetails(&_DelegationController.CallOpts, operator)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationController *DelegationControllerCaller) OperatorShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "operatorShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationController *DelegationControllerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.OperatorShares(&_DelegationController.CallOpts, arg0, arg1)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.OperatorShares(&_DelegationController.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegationController *DelegationControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegationController *DelegationControllerSession) Owner() (common.Address, error) {
	return _DelegationController.Contract.Owner(&_DelegationController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegationController *DelegationControllerCallerSession) Owner() (common.Address, error) {
	return _DelegationController.Contract.Owner(&_DelegationController.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelegationController *DelegationControllerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelegationController *DelegationControllerSession) Paused(index uint8) (bool, error) {
	return _DelegationController.Contract.Paused(&_DelegationController.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelegationController *DelegationControllerCallerSession) Paused(index uint8) (bool, error) {
	return _DelegationController.Contract.Paused(&_DelegationController.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelegationController *DelegationControllerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelegationController *DelegationControllerSession) Paused0() (*big.Int, error) {
	return _DelegationController.Contract.Paused0(&_DelegationController.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) Paused0() (*big.Int, error) {
	return _DelegationController.Contract.Paused0(&_DelegationController.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelegationController *DelegationControllerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelegationController *DelegationControllerSession) PauserRegistry() (common.Address, error) {
	return _DelegationController.Contract.PauserRegistry(&_DelegationController.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelegationController *DelegationControllerCallerSession) PauserRegistry() (common.Address, error) {
	return _DelegationController.Contract.PauserRegistry(&_DelegationController.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationController *DelegationControllerCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationController *DelegationControllerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _DelegationController.Contract.PendingWithdrawals(&_DelegationController.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationController *DelegationControllerCallerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _DelegationController.Contract.PendingWithdrawals(&_DelegationController.CallOpts, arg0)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_DelegationController *DelegationControllerCaller) PoolController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "poolController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_DelegationController *DelegationControllerSession) PoolController() (common.Address, error) {
	return _DelegationController.Contract.PoolController(&_DelegationController.CallOpts)
}

// PoolController is a free data retrieval call binding the contract method 0x4aa9d585.
//
// Solidity: function poolController() view returns(address)
func (_DelegationController *DelegationControllerCallerSession) PoolController() (common.Address, error) {
	return _DelegationController.Contract.PoolController(&_DelegationController.CallOpts)
}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address ) view returns(uint256)
func (_DelegationController *DelegationControllerCaller) PoolWithdrawalDelay(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "poolWithdrawalDelay", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address ) view returns(uint256)
func (_DelegationController *DelegationControllerSession) PoolWithdrawalDelay(arg0 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.PoolWithdrawalDelay(&_DelegationController.CallOpts, arg0)
}

// PoolWithdrawalDelay is a free data retrieval call binding the contract method 0xb6805c54.
//
// Solidity: function poolWithdrawalDelay(address ) view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) PoolWithdrawalDelay(arg0 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.PoolWithdrawalDelay(&_DelegationController.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationController *DelegationControllerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationController *DelegationControllerSession) Slasher() (common.Address, error) {
	return _DelegationController.Contract.Slasher(&_DelegationController.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationController *DelegationControllerCallerSession) Slasher() (common.Address, error) {
	return _DelegationController.Contract.Slasher(&_DelegationController.CallOpts)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationController *DelegationControllerCaller) StakerNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "stakerNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationController *DelegationControllerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.StakerNonce(&_DelegationController.CallOpts, arg0)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _DelegationController.Contract.StakerNonce(&_DelegationController.CallOpts, arg0)
}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_DelegationController *DelegationControllerCaller) StakerOptOutWindow(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "stakerOptOutWindow", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_DelegationController *DelegationControllerSession) StakerOptOutWindow(operator common.Address) (*big.Int, error) {
	return _DelegationController.Contract.StakerOptOutWindow(&_DelegationController.CallOpts, operator)
}

// StakerOptOutWindow is a free data retrieval call binding the contract method 0x1d1bf7f2.
//
// Solidity: function stakerOptOutWindow(address operator) view returns(uint256)
func (_DelegationController *DelegationControllerCallerSession) StakerOptOutWindow(operator common.Address) (*big.Int, error) {
	return _DelegationController.Contract.StakerOptOutWindow(&_DelegationController.CallOpts, operator)
}

// WrappedTokenGateway is a free data retrieval call binding the contract method 0x1e119838.
//
// Solidity: function wrappedTokenGateway() view returns(address)
func (_DelegationController *DelegationControllerCaller) WrappedTokenGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationController.contract.Call(opts, &out, "wrappedTokenGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenGateway is a free data retrieval call binding the contract method 0x1e119838.
//
// Solidity: function wrappedTokenGateway() view returns(address)
func (_DelegationController *DelegationControllerSession) WrappedTokenGateway() (common.Address, error) {
	return _DelegationController.Contract.WrappedTokenGateway(&_DelegationController.CallOpts)
}

// WrappedTokenGateway is a free data retrieval call binding the contract method 0x1e119838.
//
// Solidity: function wrappedTokenGateway() view returns(address)
func (_DelegationController *DelegationControllerCallerSession) WrappedTokenGateway() (common.Address, error) {
	return _DelegationController.Contract.WrappedTokenGateway(&_DelegationController.CallOpts)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationController *DelegationControllerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "decreaseDelegatedShares", staker, pool, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationController *DelegationControllerSession) DecreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.DecreaseDelegatedShares(&_DelegationController.TransactOpts, staker, pool, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationController *DelegationControllerTransactorSession) DecreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.DecreaseDelegatedShares(&_DelegationController.TransactOpts, staker, pool, shares)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationController *DelegationControllerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationController *DelegationControllerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationController.Contract.DelegateTo(&_DelegationController.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationController *DelegationControllerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationController.Contract.DelegateTo(&_DelegationController.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationController *DelegationControllerTransactor) DelegateToBySignature(opts *bind.TransactOpts, staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "delegateToBySignature", staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationController *DelegationControllerSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationController.Contract.DelegateToBySignature(&_DelegationController.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationController *DelegationControllerTransactorSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationController.Contract.DelegateToBySignature(&_DelegationController.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationController *DelegationControllerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "increaseDelegatedShares", staker, pool, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationController *DelegationControllerSession) IncreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.IncreaseDelegatedShares(&_DelegationController.TransactOpts, staker, pool, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address pool, uint256 shares) returns()
func (_DelegationController *DelegationControllerTransactorSession) IncreaseDelegatedShares(staker common.Address, pool common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.IncreaseDelegatedShares(&_DelegationController.TransactOpts, staker, pool, shares)
}

// Initialize is a paid mutator transaction binding the contract method 0x22bf40e4.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, uint256 _minWithdrawalDelay, address[] _pools, uint256[] _withdrawalDelay) returns()
func (_DelegationController *DelegationControllerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _minWithdrawalDelay *big.Int, _pools []common.Address, _withdrawalDelay []*big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelay, _pools, _withdrawalDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x22bf40e4.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, uint256 _minWithdrawalDelay, address[] _pools, uint256[] _withdrawalDelay) returns()
func (_DelegationController *DelegationControllerSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _minWithdrawalDelay *big.Int, _pools []common.Address, _withdrawalDelay []*big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.Initialize(&_DelegationController.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelay, _pools, _withdrawalDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x22bf40e4.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, uint256 _minWithdrawalDelay, address[] _pools, uint256[] _withdrawalDelay) returns()
func (_DelegationController *DelegationControllerTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _minWithdrawalDelay *big.Int, _pools []common.Address, _withdrawalDelay []*big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.Initialize(&_DelegationController.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelay, _pools, _withdrawalDelay)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationController *DelegationControllerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "modifyOperatorDetails", newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationController *DelegationControllerSession) ModifyOperatorDetails(newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _DelegationController.Contract.ModifyOperatorDetails(&_DelegationController.TransactOpts, newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationController *DelegationControllerTransactorSession) ModifyOperatorDetails(newOperatorDetails IDelegationControllerOperatorDetails) (*types.Transaction, error) {
	return _DelegationController.Contract.ModifyOperatorDetails(&_DelegationController.TransactOpts, newOperatorDetails)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelegationController *DelegationControllerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelegationController *DelegationControllerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.Pause(&_DelegationController.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelegationController *DelegationControllerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.Pause(&_DelegationController.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelegationController *DelegationControllerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelegationController *DelegationControllerSession) PauseAll() (*types.Transaction, error) {
	return _DelegationController.Contract.PauseAll(&_DelegationController.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelegationController *DelegationControllerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _DelegationController.Contract.PauseAll(&_DelegationController.TransactOpts)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationController *DelegationControllerTransactor) RegisterAsOperator(opts *bind.TransactOpts, registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "registerAsOperator", registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationController *DelegationControllerSession) RegisterAsOperator(registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationController.Contract.RegisterAsOperator(&_DelegationController.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationController *DelegationControllerTransactorSession) RegisterAsOperator(registeringOperatorDetails IDelegationControllerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationController.Contract.RegisterAsOperator(&_DelegationController.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegationController *DelegationControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegationController *DelegationControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelegationController.Contract.RenounceOwnership(&_DelegationController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegationController *DelegationControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelegationController.Contract.RenounceOwnership(&_DelegationController.TransactOpts)
}

// SetMinWithdrawalDelay is a paid mutator transaction binding the contract method 0xbe0525d7.
//
// Solidity: function setMinWithdrawalDelay(uint256 newMinWithdrawalDelay) returns()
func (_DelegationController *DelegationControllerTransactor) SetMinWithdrawalDelay(opts *bind.TransactOpts, newMinWithdrawalDelay *big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "setMinWithdrawalDelay", newMinWithdrawalDelay)
}

// SetMinWithdrawalDelay is a paid mutator transaction binding the contract method 0xbe0525d7.
//
// Solidity: function setMinWithdrawalDelay(uint256 newMinWithdrawalDelay) returns()
func (_DelegationController *DelegationControllerSession) SetMinWithdrawalDelay(newMinWithdrawalDelay *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.SetMinWithdrawalDelay(&_DelegationController.TransactOpts, newMinWithdrawalDelay)
}

// SetMinWithdrawalDelay is a paid mutator transaction binding the contract method 0xbe0525d7.
//
// Solidity: function setMinWithdrawalDelay(uint256 newMinWithdrawalDelay) returns()
func (_DelegationController *DelegationControllerTransactorSession) SetMinWithdrawalDelay(newMinWithdrawalDelay *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.SetMinWithdrawalDelay(&_DelegationController.TransactOpts, newMinWithdrawalDelay)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelegationController *DelegationControllerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelegationController *DelegationControllerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.SetPauserRegistry(&_DelegationController.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelegationController *DelegationControllerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.SetPauserRegistry(&_DelegationController.TransactOpts, newPauserRegistry)
}

// SetPoolWithdrawalDelay is a paid mutator transaction binding the contract method 0xeafe9f31.
//
// Solidity: function setPoolWithdrawalDelay(address[] pools, uint256[] withdrawalDelay) returns()
func (_DelegationController *DelegationControllerTransactor) SetPoolWithdrawalDelay(opts *bind.TransactOpts, pools []common.Address, withdrawalDelay []*big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "setPoolWithdrawalDelay", pools, withdrawalDelay)
}

// SetPoolWithdrawalDelay is a paid mutator transaction binding the contract method 0xeafe9f31.
//
// Solidity: function setPoolWithdrawalDelay(address[] pools, uint256[] withdrawalDelay) returns()
func (_DelegationController *DelegationControllerSession) SetPoolWithdrawalDelay(pools []common.Address, withdrawalDelay []*big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.SetPoolWithdrawalDelay(&_DelegationController.TransactOpts, pools, withdrawalDelay)
}

// SetPoolWithdrawalDelay is a paid mutator transaction binding the contract method 0xeafe9f31.
//
// Solidity: function setPoolWithdrawalDelay(address[] pools, uint256[] withdrawalDelay) returns()
func (_DelegationController *DelegationControllerTransactorSession) SetPoolWithdrawalDelay(pools []common.Address, withdrawalDelay []*big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.SetPoolWithdrawalDelay(&_DelegationController.TransactOpts, pools, withdrawalDelay)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegationController *DelegationControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegationController *DelegationControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.TransferOwnership(&_DelegationController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegationController *DelegationControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.TransferOwnership(&_DelegationController.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_DelegationController *DelegationControllerTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_DelegationController *DelegationControllerSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.Undelegate(&_DelegationController.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_DelegationController *DelegationControllerTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.Undelegate(&_DelegationController.TransactOpts, staker)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelegationController *DelegationControllerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelegationController *DelegationControllerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.Unpause(&_DelegationController.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelegationController *DelegationControllerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationController.Contract.Unpause(&_DelegationController.TransactOpts, newPausedStatus)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_DelegationController *DelegationControllerTransactor) Unstakes(opts *bind.TransactOpts, unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "unstakes", unstakeParams)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_DelegationController *DelegationControllerSession) Unstakes(unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _DelegationController.Contract.Unstakes(&_DelegationController.TransactOpts, unstakeParams)
}

// Unstakes is a paid mutator transaction binding the contract method 0x63152b13.
//
// Solidity: function unstakes((address[],uint256[],address)[] unstakeParams) returns(bytes32[])
func (_DelegationController *DelegationControllerTransactorSession) Unstakes(unstakeParams []IDelegationControllerUnstakeParams) (*types.Transaction, error) {
	return _DelegationController.Contract.Unstakes(&_DelegationController.TransactOpts, unstakeParams)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationController *DelegationControllerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "updateOperatorMetadataURI", metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationController *DelegationControllerSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _DelegationController.Contract.UpdateOperatorMetadataURI(&_DelegationController.TransactOpts, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationController *DelegationControllerTransactorSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _DelegationController.Contract.UpdateOperatorMetadataURI(&_DelegationController.TransactOpts, metadataURI)
}

// UpdateWrappedTokenGateway is a paid mutator transaction binding the contract method 0x65108e92.
//
// Solidity: function updateWrappedTokenGateway(address _newWrappedTokenGateway) returns()
func (_DelegationController *DelegationControllerTransactor) UpdateWrappedTokenGateway(opts *bind.TransactOpts, _newWrappedTokenGateway common.Address) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "updateWrappedTokenGateway", _newWrappedTokenGateway)
}

// UpdateWrappedTokenGateway is a paid mutator transaction binding the contract method 0x65108e92.
//
// Solidity: function updateWrappedTokenGateway(address _newWrappedTokenGateway) returns()
func (_DelegationController *DelegationControllerSession) UpdateWrappedTokenGateway(_newWrappedTokenGateway common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.UpdateWrappedTokenGateway(&_DelegationController.TransactOpts, _newWrappedTokenGateway)
}

// UpdateWrappedTokenGateway is a paid mutator transaction binding the contract method 0x65108e92.
//
// Solidity: function updateWrappedTokenGateway(address _newWrappedTokenGateway) returns()
func (_DelegationController *DelegationControllerTransactorSession) UpdateWrappedTokenGateway(_newWrappedTokenGateway common.Address) (*types.Transaction, error) {
	return _DelegationController.Contract.UpdateWrappedTokenGateway(&_DelegationController.TransactOpts, _newWrappedTokenGateway)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationController *DelegationControllerTransactor) Withdraw(opts *bind.TransactOpts, withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "withdraw", withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationController *DelegationControllerSession) Withdraw(withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationController.Contract.Withdraw(&_DelegationController.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb6d65ea3.
//
// Solidity: function withdraw((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationController *DelegationControllerTransactorSession) Withdraw(withdrawal IDelegationControllerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationController.Contract.Withdraw(&_DelegationController.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationController *DelegationControllerTransactor) Withdraws(opts *bind.TransactOpts, withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationController.contract.Transact(opts, "withdraws", withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationController *DelegationControllerSession) Withdraws(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationController.Contract.Withdraws(&_DelegationController.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// Withdraws is a paid mutator transaction binding the contract method 0x0af71b77.
//
// Solidity: function withdraws((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationController *DelegationControllerTransactorSession) Withdraws(withdrawals []IDelegationControllerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationController.Contract.Withdraws(&_DelegationController.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// DelegationControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DelegationController contract.
type DelegationControllerInitializedIterator struct {
	Event *DelegationControllerInitialized // Event containing the contract specifics and raw log

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
func (it *DelegationControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerInitialized)
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
		it.Event = new(DelegationControllerInitialized)
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
func (it *DelegationControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerInitialized represents a Initialized event raised by the DelegationController contract.
type DelegationControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelegationController *DelegationControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelegationControllerInitializedIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerInitializedIterator{contract: _DelegationController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelegationController *DelegationControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelegationControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerInitialized)
				if err := _DelegationController.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DelegationController *DelegationControllerFilterer) ParseInitialized(log types.Log) (*DelegationControllerInitialized, error) {
	event := new(DelegationControllerInitialized)
	if err := _DelegationController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerMinWithdrawalDelaySetIterator is returned from FilterMinWithdrawalDelaySet and is used to iterate over the raw logs and unpacked data for MinWithdrawalDelaySet events raised by the DelegationController contract.
type DelegationControllerMinWithdrawalDelaySetIterator struct {
	Event *DelegationControllerMinWithdrawalDelaySet // Event containing the contract specifics and raw log

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
func (it *DelegationControllerMinWithdrawalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerMinWithdrawalDelaySet)
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
		it.Event = new(DelegationControllerMinWithdrawalDelaySet)
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
func (it *DelegationControllerMinWithdrawalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerMinWithdrawalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerMinWithdrawalDelaySet represents a MinWithdrawalDelaySet event raised by the DelegationController contract.
type DelegationControllerMinWithdrawalDelaySet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinWithdrawalDelaySet is a free log retrieval operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_DelegationController *DelegationControllerFilterer) FilterMinWithdrawalDelaySet(opts *bind.FilterOpts) (*DelegationControllerMinWithdrawalDelaySetIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "MinWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerMinWithdrawalDelaySetIterator{contract: _DelegationController.contract, event: "MinWithdrawalDelaySet", logs: logs, sub: sub}, nil
}

// WatchMinWithdrawalDelaySet is a free log subscription operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_DelegationController *DelegationControllerFilterer) WatchMinWithdrawalDelaySet(opts *bind.WatchOpts, sink chan<- *DelegationControllerMinWithdrawalDelaySet) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "MinWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerMinWithdrawalDelaySet)
				if err := _DelegationController.contract.UnpackLog(event, "MinWithdrawalDelaySet", log); err != nil {
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

// ParseMinWithdrawalDelaySet is a log parse operation binding the contract event 0x338caf1431dddfb34caa16bfc51573f97922fa2f8eb6d70d27476d8b4a89d5c3.
//
// Solidity: event MinWithdrawalDelaySet(uint256 previousValue, uint256 newValue)
func (_DelegationController *DelegationControllerFilterer) ParseMinWithdrawalDelaySet(log types.Log) (*DelegationControllerMinWithdrawalDelaySet, error) {
	event := new(DelegationControllerMinWithdrawalDelaySet)
	if err := _DelegationController.contract.UnpackLog(event, "MinWithdrawalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerOperatorDetailsModifiedIterator is returned from FilterOperatorDetailsModified and is used to iterate over the raw logs and unpacked data for OperatorDetailsModified events raised by the DelegationController contract.
type DelegationControllerOperatorDetailsModifiedIterator struct {
	Event *DelegationControllerOperatorDetailsModified // Event containing the contract specifics and raw log

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
func (it *DelegationControllerOperatorDetailsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerOperatorDetailsModified)
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
		it.Event = new(DelegationControllerOperatorDetailsModified)
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
func (it *DelegationControllerOperatorDetailsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerOperatorDetailsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerOperatorDetailsModified represents a OperatorDetailsModified event raised by the DelegationController contract.
type DelegationControllerOperatorDetailsModified struct {
	Operator           common.Address
	NewOperatorDetails IDelegationControllerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorDetailsModified is a free log retrieval operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationController *DelegationControllerFilterer) FilterOperatorDetailsModified(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerOperatorDetailsModifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerOperatorDetailsModifiedIterator{contract: _DelegationController.contract, event: "OperatorDetailsModified", logs: logs, sub: sub}, nil
}

// WatchOperatorDetailsModified is a free log subscription operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationController *DelegationControllerFilterer) WatchOperatorDetailsModified(opts *bind.WatchOpts, sink chan<- *DelegationControllerOperatorDetailsModified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerOperatorDetailsModified)
				if err := _DelegationController.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
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

// ParseOperatorDetailsModified is a log parse operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationController *DelegationControllerFilterer) ParseOperatorDetailsModified(log types.Log) (*DelegationControllerOperatorDetailsModified, error) {
	event := new(DelegationControllerOperatorDetailsModified)
	if err := _DelegationController.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the DelegationController contract.
type DelegationControllerOperatorMetadataURIUpdatedIterator struct {
	Event *DelegationControllerOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *DelegationControllerOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerOperatorMetadataURIUpdated)
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
		it.Event = new(DelegationControllerOperatorMetadataURIUpdated)
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
func (it *DelegationControllerOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the DelegationController contract.
type DelegationControllerOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationController *DelegationControllerFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerOperatorMetadataURIUpdatedIterator{contract: _DelegationController.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationController *DelegationControllerFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *DelegationControllerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerOperatorMetadataURIUpdated)
				if err := _DelegationController.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
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

// ParseOperatorMetadataURIUpdated is a log parse operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationController *DelegationControllerFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*DelegationControllerOperatorMetadataURIUpdated, error) {
	event := new(DelegationControllerOperatorMetadataURIUpdated)
	if err := _DelegationController.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the DelegationController contract.
type DelegationControllerOperatorRegisteredIterator struct {
	Event *DelegationControllerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *DelegationControllerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerOperatorRegistered)
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
		it.Event = new(DelegationControllerOperatorRegistered)
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
func (it *DelegationControllerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerOperatorRegistered represents a OperatorRegistered event raised by the DelegationController contract.
type DelegationControllerOperatorRegistered struct {
	Operator        common.Address
	OperatorDetails IDelegationControllerOperatorDetails
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationController *DelegationControllerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerOperatorRegisteredIterator{contract: _DelegationController.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationController *DelegationControllerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *DelegationControllerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerOperatorRegistered)
				if err := _DelegationController.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationController *DelegationControllerFilterer) ParseOperatorRegistered(log types.Log) (*DelegationControllerOperatorRegistered, error) {
	event := new(DelegationControllerOperatorRegistered)
	if err := _DelegationController.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the DelegationController contract.
type DelegationControllerOperatorSharesDecreasedIterator struct {
	Event *DelegationControllerOperatorSharesDecreased // Event containing the contract specifics and raw log

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
func (it *DelegationControllerOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerOperatorSharesDecreased)
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
		it.Event = new(DelegationControllerOperatorSharesDecreased)
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
func (it *DelegationControllerOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the DelegationController contract.
type DelegationControllerOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Pool     common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationController *DelegationControllerFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerOperatorSharesDecreasedIterator{contract: _DelegationController.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationController *DelegationControllerFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *DelegationControllerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerOperatorSharesDecreased)
				if err := _DelegationController.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
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

// ParseOperatorSharesDecreased is a log parse operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationController *DelegationControllerFilterer) ParseOperatorSharesDecreased(log types.Log) (*DelegationControllerOperatorSharesDecreased, error) {
	event := new(DelegationControllerOperatorSharesDecreased)
	if err := _DelegationController.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the DelegationController contract.
type DelegationControllerOperatorSharesIncreasedIterator struct {
	Event *DelegationControllerOperatorSharesIncreased // Event containing the contract specifics and raw log

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
func (it *DelegationControllerOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerOperatorSharesIncreased)
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
		it.Event = new(DelegationControllerOperatorSharesIncreased)
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
func (it *DelegationControllerOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the DelegationController contract.
type DelegationControllerOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Pool     common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationController *DelegationControllerFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*DelegationControllerOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerOperatorSharesIncreasedIterator{contract: _DelegationController.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationController *DelegationControllerFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *DelegationControllerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerOperatorSharesIncreased)
				if err := _DelegationController.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
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

// ParseOperatorSharesIncreased is a log parse operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address pool, uint256 shares)
func (_DelegationController *DelegationControllerFilterer) ParseOperatorSharesIncreased(log types.Log) (*DelegationControllerOperatorSharesIncreased, error) {
	event := new(DelegationControllerOperatorSharesIncreased)
	if err := _DelegationController.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelegationController contract.
type DelegationControllerOwnershipTransferredIterator struct {
	Event *DelegationControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelegationControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerOwnershipTransferred)
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
		it.Event = new(DelegationControllerOwnershipTransferred)
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
func (it *DelegationControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerOwnershipTransferred represents a OwnershipTransferred event raised by the DelegationController contract.
type DelegationControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegationController *DelegationControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelegationControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerOwnershipTransferredIterator{contract: _DelegationController.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegationController *DelegationControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelegationControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerOwnershipTransferred)
				if err := _DelegationController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DelegationController *DelegationControllerFilterer) ParseOwnershipTransferred(log types.Log) (*DelegationControllerOwnershipTransferred, error) {
	event := new(DelegationControllerOwnershipTransferred)
	if err := _DelegationController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DelegationController contract.
type DelegationControllerPausedIterator struct {
	Event *DelegationControllerPaused // Event containing the contract specifics and raw log

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
func (it *DelegationControllerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerPaused)
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
		it.Event = new(DelegationControllerPaused)
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
func (it *DelegationControllerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerPaused represents a Paused event raised by the DelegationController contract.
type DelegationControllerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelegationController *DelegationControllerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*DelegationControllerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerPausedIterator{contract: _DelegationController.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelegationController *DelegationControllerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DelegationControllerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerPaused)
				if err := _DelegationController.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DelegationController *DelegationControllerFilterer) ParsePaused(log types.Log) (*DelegationControllerPaused, error) {
	event := new(DelegationControllerPaused)
	if err := _DelegationController.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the DelegationController contract.
type DelegationControllerPauserRegistrySetIterator struct {
	Event *DelegationControllerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *DelegationControllerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerPauserRegistrySet)
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
		it.Event = new(DelegationControllerPauserRegistrySet)
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
func (it *DelegationControllerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerPauserRegistrySet represents a PauserRegistrySet event raised by the DelegationController contract.
type DelegationControllerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelegationController *DelegationControllerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*DelegationControllerPauserRegistrySetIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerPauserRegistrySetIterator{contract: _DelegationController.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelegationController *DelegationControllerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *DelegationControllerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerPauserRegistrySet)
				if err := _DelegationController.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_DelegationController *DelegationControllerFilterer) ParsePauserRegistrySet(log types.Log) (*DelegationControllerPauserRegistrySet, error) {
	event := new(DelegationControllerPauserRegistrySet)
	if err := _DelegationController.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerPoolWithdrawalDelaySetIterator is returned from FilterPoolWithdrawalDelaySet and is used to iterate over the raw logs and unpacked data for PoolWithdrawalDelaySet events raised by the DelegationController contract.
type DelegationControllerPoolWithdrawalDelaySetIterator struct {
	Event *DelegationControllerPoolWithdrawalDelaySet // Event containing the contract specifics and raw log

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
func (it *DelegationControllerPoolWithdrawalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerPoolWithdrawalDelaySet)
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
		it.Event = new(DelegationControllerPoolWithdrawalDelaySet)
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
func (it *DelegationControllerPoolWithdrawalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerPoolWithdrawalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerPoolWithdrawalDelaySet represents a PoolWithdrawalDelaySet event raised by the DelegationController contract.
type DelegationControllerPoolWithdrawalDelaySet struct {
	Pool          common.Address
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPoolWithdrawalDelaySet is a free log retrieval operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_DelegationController *DelegationControllerFilterer) FilterPoolWithdrawalDelaySet(opts *bind.FilterOpts) (*DelegationControllerPoolWithdrawalDelaySetIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "PoolWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerPoolWithdrawalDelaySetIterator{contract: _DelegationController.contract, event: "PoolWithdrawalDelaySet", logs: logs, sub: sub}, nil
}

// WatchPoolWithdrawalDelaySet is a free log subscription operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_DelegationController *DelegationControllerFilterer) WatchPoolWithdrawalDelaySet(opts *bind.WatchOpts, sink chan<- *DelegationControllerPoolWithdrawalDelaySet) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "PoolWithdrawalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerPoolWithdrawalDelaySet)
				if err := _DelegationController.contract.UnpackLog(event, "PoolWithdrawalDelaySet", log); err != nil {
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

// ParsePoolWithdrawalDelaySet is a log parse operation binding the contract event 0x108376441269231310d8eb7d2c786779cb19c9b7f967e2e95ab366f0fde46dc2.
//
// Solidity: event PoolWithdrawalDelaySet(address pool, uint256 previousValue, uint256 newValue)
func (_DelegationController *DelegationControllerFilterer) ParsePoolWithdrawalDelaySet(log types.Log) (*DelegationControllerPoolWithdrawalDelaySet, error) {
	event := new(DelegationControllerPoolWithdrawalDelaySet)
	if err := _DelegationController.contract.UnpackLog(event, "PoolWithdrawalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the DelegationController contract.
type DelegationControllerStakerDelegatedIterator struct {
	Event *DelegationControllerStakerDelegated // Event containing the contract specifics and raw log

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
func (it *DelegationControllerStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStakerDelegated)
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
		it.Event = new(DelegationControllerStakerDelegated)
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
func (it *DelegationControllerStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStakerDelegated represents a StakerDelegated event raised by the DelegationController contract.
type DelegationControllerStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationControllerStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStakerDelegatedIterator{contract: _DelegationController.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStakerDelegated)
				if err := _DelegationController.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
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

// ParseStakerDelegated is a log parse operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) ParseStakerDelegated(log types.Log) (*DelegationControllerStakerDelegated, error) {
	event := new(DelegationControllerStakerDelegated)
	if err := _DelegationController.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the DelegationController contract.
type DelegationControllerStakerForceUndelegatedIterator struct {
	Event *DelegationControllerStakerForceUndelegated // Event containing the contract specifics and raw log

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
func (it *DelegationControllerStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStakerForceUndelegated)
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
		it.Event = new(DelegationControllerStakerForceUndelegated)
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
func (it *DelegationControllerStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStakerForceUndelegated represents a StakerForceUndelegated event raised by the DelegationController contract.
type DelegationControllerStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationControllerStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStakerForceUndelegatedIterator{contract: _DelegationController.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStakerForceUndelegated)
				if err := _DelegationController.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
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

// ParseStakerForceUndelegated is a log parse operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) ParseStakerForceUndelegated(log types.Log) (*DelegationControllerStakerForceUndelegated, error) {
	event := new(DelegationControllerStakerForceUndelegated)
	if err := _DelegationController.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the DelegationController contract.
type DelegationControllerStakerUndelegatedIterator struct {
	Event *DelegationControllerStakerUndelegated // Event containing the contract specifics and raw log

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
func (it *DelegationControllerStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerStakerUndelegated)
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
		it.Event = new(DelegationControllerStakerUndelegated)
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
func (it *DelegationControllerStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerStakerUndelegated represents a StakerUndelegated event raised by the DelegationController contract.
type DelegationControllerStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationControllerStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerStakerUndelegatedIterator{contract: _DelegationController.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *DelegationControllerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerStakerUndelegated)
				if err := _DelegationController.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
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

// ParseStakerUndelegated is a log parse operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationController *DelegationControllerFilterer) ParseStakerUndelegated(log types.Log) (*DelegationControllerStakerUndelegated, error) {
	event := new(DelegationControllerStakerUndelegated)
	if err := _DelegationController.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DelegationController contract.
type DelegationControllerUnpausedIterator struct {
	Event *DelegationControllerUnpaused // Event containing the contract specifics and raw log

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
func (it *DelegationControllerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerUnpaused)
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
		it.Event = new(DelegationControllerUnpaused)
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
func (it *DelegationControllerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerUnpaused represents a Unpaused event raised by the DelegationController contract.
type DelegationControllerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelegationController *DelegationControllerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*DelegationControllerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelegationControllerUnpausedIterator{contract: _DelegationController.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelegationController *DelegationControllerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DelegationControllerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerUnpaused)
				if err := _DelegationController.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DelegationController *DelegationControllerFilterer) ParseUnpaused(log types.Log) (*DelegationControllerUnpaused, error) {
	event := new(DelegationControllerUnpaused)
	if err := _DelegationController.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerUpdateWrappedTokenGatewayIterator is returned from FilterUpdateWrappedTokenGateway and is used to iterate over the raw logs and unpacked data for UpdateWrappedTokenGateway events raised by the DelegationController contract.
type DelegationControllerUpdateWrappedTokenGatewayIterator struct {
	Event *DelegationControllerUpdateWrappedTokenGateway // Event containing the contract specifics and raw log

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
func (it *DelegationControllerUpdateWrappedTokenGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerUpdateWrappedTokenGateway)
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
		it.Event = new(DelegationControllerUpdateWrappedTokenGateway)
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
func (it *DelegationControllerUpdateWrappedTokenGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerUpdateWrappedTokenGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerUpdateWrappedTokenGateway represents a UpdateWrappedTokenGateway event raised by the DelegationController contract.
type DelegationControllerUpdateWrappedTokenGateway struct {
	PreviousGateway common.Address
	CurrentGateway  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateWrappedTokenGateway is a free log retrieval operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_DelegationController *DelegationControllerFilterer) FilterUpdateWrappedTokenGateway(opts *bind.FilterOpts) (*DelegationControllerUpdateWrappedTokenGatewayIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "UpdateWrappedTokenGateway")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerUpdateWrappedTokenGatewayIterator{contract: _DelegationController.contract, event: "UpdateWrappedTokenGateway", logs: logs, sub: sub}, nil
}

// WatchUpdateWrappedTokenGateway is a free log subscription operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_DelegationController *DelegationControllerFilterer) WatchUpdateWrappedTokenGateway(opts *bind.WatchOpts, sink chan<- *DelegationControllerUpdateWrappedTokenGateway) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "UpdateWrappedTokenGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerUpdateWrappedTokenGateway)
				if err := _DelegationController.contract.UnpackLog(event, "UpdateWrappedTokenGateway", log); err != nil {
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

// ParseUpdateWrappedTokenGateway is a log parse operation binding the contract event 0x6ed22dc7330f7d5d7c2ceacb5a19323d459493561529441177421938a434815b.
//
// Solidity: event UpdateWrappedTokenGateway(address previousGateway, address currentGateway)
func (_DelegationController *DelegationControllerFilterer) ParseUpdateWrappedTokenGateway(log types.Log) (*DelegationControllerUpdateWrappedTokenGateway, error) {
	event := new(DelegationControllerUpdateWrappedTokenGateway)
	if err := _DelegationController.contract.UnpackLog(event, "UpdateWrappedTokenGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the DelegationController contract.
type DelegationControllerWithdrawalCompletedIterator struct {
	Event *DelegationControllerWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *DelegationControllerWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerWithdrawalCompleted)
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
		it.Event = new(DelegationControllerWithdrawalCompleted)
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
func (it *DelegationControllerWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerWithdrawalCompleted represents a WithdrawalCompleted event raised by the DelegationController contract.
type DelegationControllerWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationController *DelegationControllerFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*DelegationControllerWithdrawalCompletedIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerWithdrawalCompletedIterator{contract: _DelegationController.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationController *DelegationControllerFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *DelegationControllerWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerWithdrawalCompleted)
				if err := _DelegationController.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationController *DelegationControllerFilterer) ParseWithdrawalCompleted(log types.Log) (*DelegationControllerWithdrawalCompleted, error) {
	event := new(DelegationControllerWithdrawalCompleted)
	if err := _DelegationController.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationControllerWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the DelegationController contract.
type DelegationControllerWithdrawalQueuedIterator struct {
	Event *DelegationControllerWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *DelegationControllerWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationControllerWithdrawalQueued)
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
		it.Event = new(DelegationControllerWithdrawalQueued)
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
func (it *DelegationControllerWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationControllerWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationControllerWithdrawalQueued represents a WithdrawalQueued event raised by the DelegationController contract.
type DelegationControllerWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     IDelegationControllerWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationController *DelegationControllerFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*DelegationControllerWithdrawalQueuedIterator, error) {

	logs, sub, err := _DelegationController.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &DelegationControllerWithdrawalQueuedIterator{contract: _DelegationController.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationController *DelegationControllerFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *DelegationControllerWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _DelegationController.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationControllerWithdrawalQueued)
				if err := _DelegationController.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationController *DelegationControllerFilterer) ParseWithdrawalQueued(log types.Log) (*DelegationControllerWithdrawalQueued, error) {
	event := new(DelegationControllerWithdrawalQueued)
	if err := _DelegationController.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
