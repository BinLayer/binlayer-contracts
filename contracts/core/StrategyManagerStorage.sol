// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import '../interfaces/IStrategyManager.sol';
import '../interfaces/IStrategy.sol';
import '../interfaces/IDelegationManager.sol';
import '../interfaces/ISlasher.sol';

/**
 * @title Storage variables for the `StrategyManager` contract.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract StrategyManagerStorage is IStrategyManager {
  /// @notice The EIP-712 typehash for the contract's domain
  bytes32 public constant DOMAIN_TYPEHASH = keccak256('EIP712Domain(string name,uint256 chainId,address verifyingContract)');
  /// @notice The EIP-712 typehash for the deposit struct used by the contract
  bytes32 public constant DEPOSIT_TYPEHASH =
    keccak256('Deposit(address staker,address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)');
  // maximum length of dynamic arrays in `stakerStrategyList` mapping, for sanity's sake
  uint8 internal constant MAX_STAKER_STRATEGY_LIST_LENGTH = 32;

  // system contracts
  IDelegationManager public immutable delegation;
  ISlasher public immutable slasher;

  /**
   * @notice Original EIP-712 Domain separator for this contract.
   * @dev The domain separator may change in the event of a fork that modifies the ChainID.
   * Use the getter function `domainSeparator` to get the current domain separator for this contract.
   */
  bytes32 internal _DOMAIN_SEPARATOR;
  // staker => number of signed deposit nonce (used in depositIntoStrategyWithSignature)
  mapping(address => uint256) public nonces;
  /// @notice Permissioned role, which can be changed by the contract owner. Has the ability to edit the strategy whitelist
  address public strategyWhitelister;
  /// @notice Mapping: staker => Strategy => number of shares which they currently hold
  mapping(address => mapping(IStrategy => uint256)) public stakerStrategyShares;
  /// @notice Mapping: staker => array of strategies in which they have nonzero shares
  mapping(address => IStrategy[]) public stakerStrategyList;
  /// @notice Mapping: strategy => whether or not stakers are allowed to deposit into it
  mapping(IStrategy => bool) public strategyIsWhitelistedForDeposit;

  /**
   * @notice Mapping: strategy => whether or not stakers are allowed to transfer strategy shares to another address
   * if true for a strategy, a user cannot depositIntoStrategyWithSignature into that strategy for another staker
   * and also when performing queueWithdrawals, a staker can only withdraw to themselves
   */
  mapping(IStrategy => bool) public thirdPartyTransfersForbidden;

  constructor(IDelegationManager _delegation, ISlasher _slasher) {
    delegation = _delegation;
    slasher = _slasher;
  }

  /**
   * @dev This empty reserved space is put in place to allow future versions to add new
   * variables without shifting down storage in the inheritance chain.
   * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
   */
  uint256[49] private __gap;
}
