// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import '../interfaces/IPoolController.sol';
import '../interfaces/IPool.sol';
import '../interfaces/IDelegationController.sol';
import '../interfaces/ISlasher.sol';

/**
 * @title Storage variables for the `PoolController.sol` contract.
 * @notice This storage contract is separate from the logic to simplify the upgrade process.
 */
abstract contract PoolControllerStorage is IPoolController {
  /// @notice The EIP-712 typehash for the contract's domain
  bytes32 public constant DOMAIN_TYPEHASH = keccak256('EIP712Domain(string name,uint256 chainId,address verifyingContract)');
  /// @notice The EIP-712 typehash for the deposit struct used by the contract
  bytes32 public constant DEPOSIT_TYPEHASH =
    keccak256('Deposit(address staker,address pool,address token,uint256 amount,uint256 nonce,uint256 expiry)');
  // maximum length of dynamic arrays in `stakerPoolList` mapping, for sanity's sake
  uint8 internal constant MAX_STAKER_POOL_LIST_LENGTH = 32;

  // system contracts
  IDelegationController public immutable delegation;
  ISlasher public immutable slasher;

  /**
   * @notice Original EIP-712 Domain separator for this contract.
   * @dev The domain separator may change in the event of a fork that modifies the ChainID.
   * Use the getter function `domainSeparator` to get the current domain separator for this contract.
   */
  bytes32 internal _DOMAIN_SEPARATOR;
  // staker => number of signed deposit nonce (used in depositIntoPoolWithSignature)
  mapping(address => uint256) public nonces;
  /// @notice Permissioned role, which can be changed by the contract owner. Has the ability to edit the pool whitelist
  address public poolWhitelister;
  /// @notice Mapping: staker => Pool => number of shares which they currently hold
  mapping(address => mapping(IPool => uint256)) public stakerPoolShares;
  /// @notice Mapping: staker => array of strategies in which they have nonzero shares
  mapping(address => IPool[]) public stakerPoolList;
  /// @notice Mapping: pool => whether or not stakers are allowed to deposit into it
  mapping(IPool => bool) public poolIsWhitelistedForDeposit;

  /**
   * @notice Mapping: pool => whether or not stakers are allowed to transfer pool shares to another address
   * if true for a pool, a user cannot depositIntoPoolWithSignature into that pool for another staker
   * and also when performing queueWithdrawals, a staker can only withdraw to themselves
   */
  mapping(IPool => bool) public thirdPartyTransfersForbidden;

  constructor(IDelegationController _delegation, ISlasher _slasher) {
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
