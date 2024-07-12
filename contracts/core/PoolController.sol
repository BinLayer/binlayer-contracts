// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import '@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol';
import '@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol';
import '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import '../permissions/Pausable.sol';
import './PoolControllerStorage.sol';
import '../libraries/EIP1271SignatureUtils.sol';
import '../helpers/Errors.sol';

/**
 * @title The primary entry- and exit-point for funds into and out of BinLayer.
 *
 * @notice This contract is for managing deposits in different pools. The main
 * functionalities are:
 * - adding and removing pools that any delegator can deposit into
 * - enabling deposit of assets into specified pool(s)
 */
contract PoolController is Initializable, OwnableUpgradeable, ReentrancyGuardUpgradeable, Pausable, PoolControllerStorage {
  using SafeERC20 for IERC20;

  // index for flag that pauses deposits when set
  uint8 internal constant PAUSED_DEPOSITS = 0;

  // chain id at the time of contract deployment
  uint256 internal immutable ORIGINAL_CHAIN_ID;

  modifier onlyPoolWhitelister() {
    require(msg.sender == poolWhitelister, Errors.NOT_POOL_WHITELISTER);
    _;
  }

  modifier onlyPoolsWhitelistedForDeposit(IPool pool) {
    require(poolIsWhitelistedForDeposit[pool], Errors.POOL_NOT_WHITELISTED);
    _;
  }

  modifier onlyDelegationController() {
    require(msg.sender == address(delegation), Errors.NOT_DELEGATION_CONTROLLER);
    _;
  }

  /**
   * @param _delegation The delegation contract of BinLayer.
   * @param _slasher The primary slashing contract of BinLayer.
   */
  constructor(IDelegationController _delegation, ISlasher _slasher) PoolControllerStorage(_delegation, _slasher) {
    _disableInitializers();
    ORIGINAL_CHAIN_ID = block.chainid;
  }

  // EXTERNAL FUNCTIONS

  /**
   * @notice Initializes the pool controller contract. Sets the `pauserRegistry` (currently **not** modifiable after being set),
   * and transfers contract ownership to the specified `initialOwner`.
   * @param _pauserRegistry Used for access control of pausing.
   * @param initialOwner Ownership of this contract is transferred to this address.
   * @param initialPoolWhitelister The initial value of `poolWhitelister` to set.
   * @param  initialPausedStatus The initial value of `_paused` to set.
   */
  function initialize(
    address initialOwner,
    address initialPoolWhitelister,
    IPauserRegistry _pauserRegistry,
    uint256 initialPausedStatus
  ) external initializer {
    _DOMAIN_SEPARATOR = _calculateDomainSeparator();
    __ReentrancyGuard_init();
    _initializePauser(_pauserRegistry, initialPausedStatus);
    _transferOwnership(initialOwner);
    _setPoolWhitelister(initialPoolWhitelister);
  }

  /**
   * @notice Deposits `amount` of `token` into the specified `pool`, with the resultant shares credited to `msg.sender`
   * @param pool is the specified pool where deposit is to be made,
   * @param token is the denomination in which the deposit is to be made,
   * @param amount is the amount of token to be deposited in the pool by the staker
   * @return shares The amount of new shares in the `pool` created as part of the action.
   * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
   *
   * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a pool is not recommended.  This can lead to attack vectors
   *          where the token balance and corresponding pool shares are not in sync upon reentrancy.
   */
  function depositIntoPool(
    IPool pool,
    IERC20 token,
    uint256 amount
  ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 shares) {
    shares = _depositIntoPool(msg.sender, pool, token, amount);
  }

  /**
   * @notice Deposits `amount` of `token` into the specified `pool`, with the resultant shares credited to `staker`
   * @param staker Staker address
   * @param pool is the specified pool where deposit is to be made,
   * @param token is the denomination in which the deposit is to be made,
   * @param amount is the amount of token to be deposited in the pool by the staker
   * @return shares The amount of new shares in the `pool` created as part of the action.
   * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
   *
   * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a pool is not recommended.  This can lead to attack vectors
   *          where the token balance and corresponding pool shares are not in sync upon reentrancy.
   */
  function depositIntoPoolWithStaker(
    address staker,
    IPool pool,
    IERC20 token,
    uint256 amount
  ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 shares) {
    if (staker != msg.sender) {
      require(!thirdPartyTransfersForbidden[pool], Errors.THIRD_PARTY_TRANSFERS_DISABLED);
    }
    shares = _depositIntoPool(staker, pool, token, amount);
  }

  /**
   * @notice Used for depositing an asset into the specified pool with the resultant shares credited to `staker`,
   * who must sign off on the action.
   * Note that the assets are transferred out/from the `msg.sender`, not from the `staker`; this function is explicitly designed
   * purely to help one address deposit 'for' another.
   * @param pool is the specified pool where deposit is to be made,
   * @param token is the denomination in which the deposit is to be made,
   * @param amount is the amount of token to be deposited in the pool by the staker
   * @param staker the staker that the deposited assets will be credited to
   * @param expiry the timestamp at which the signature expires
   * @param signature is a valid signature from the `staker`. either an ECDSA signature if the `staker` is an EOA, or data to forward
   * following EIP-1271 if the `staker` is a contract
   * @return shares The amount of new shares in the `pool` created as part of the action.
   * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
   * @dev A signature is required for this function to eliminate the possibility of griefing attacks, specifically those
   * targeting stakers who may be attempting to undelegate.
   * @dev Cannot be called if thirdPartyTransfersForbidden is set to true for this pool
   *
   *  WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a pool is not recommended.  This can lead to attack vectors
   *          where the token balance and corresponding pool shares are not in sync upon reentrancy
   */
  function depositIntoPoolWithSignature(
    IPool pool,
    IERC20 token,
    uint256 amount,
    address staker,
    uint256 expiry,
    bytes memory signature
  ) external onlyWhenNotPaused(PAUSED_DEPOSITS) nonReentrant returns (uint256 shares) {
    require(!thirdPartyTransfersForbidden[pool], Errors.THIRD_PARTY_TRANSFERS_DISABLED);
    require(expiry >= block.timestamp, Errors.SIGNATURE_EXPIRED);
    // calculate struct hash, then increment `staker`'s nonce
    uint256 nonce = nonces[staker];
    bytes32 structHash = keccak256(abi.encode(DEPOSIT_TYPEHASH, staker, pool, token, amount, nonce, expiry));
    unchecked {
      nonces[staker] = nonce + 1;
    }

    // calculate the digest hash
    bytes32 digestHash = keccak256(abi.encodePacked('\x19\x01', domainSeparator(), structHash));

    /**
     * check validity of signature:
     * 1) if `staker` is an EOA, then `signature` must be a valid ECDSA signature from `staker`,
     * indicating their intention for this action
     * 2) if `staker` is a contract, then `signature` will be checked according to EIP-1271
     */
    EIP1271SignatureUtils.checkSignature_EIP1271(staker, digestHash, signature);

    // deposit the tokens (from the `msg.sender`) and credit the new shares to the `staker`
    shares = _depositIntoPool(staker, pool, token, amount);
  }

  /// @notice Used by the DelegationController.sol to remove a Staker's shares from a particular pool when entering the withdrawal queue
  function removeShares(address staker, IPool pool, uint256 shares) external onlyDelegationController {
    _removeShares(staker, pool, shares);
  }

  /// @notice Used by the DelegationController.sol to award a Staker some shares that have passed through the withdrawal queue
  function addShares(address staker, IERC20 token, IPool pool, uint256 shares) external onlyDelegationController {
    _addShares(staker, token, pool, shares);
  }

  /// @notice Used by the DelegationController.sol to convert withdrawn shares to tokens and send them to a recipient
  function withdrawSharesAsTokens(address recipient, IPool pool, uint256 shares, IERC20 token) external onlyDelegationController {
    pool.withdraw(recipient, token, shares);
  }

  /**
   * If true for a pool, a user cannot depositIntoPoolWithSignature into that pool for another staker
   * and also when performing DelegationController.sol.queueWithdrawals, a staker can only withdraw to themselves.
   * Defaulted to false for all existing pools.
   * @param pool The pool to set `thirdPartyTransfersForbidden` value to
   * @param value bool value to set `thirdPartyTransfersForbidden` to
   */
  function setThirdPartyTransfersForbidden(IPool pool, bool value) external onlyPoolWhitelister {
    _setThirdPartyTransfersForbidden(pool, value);
  }

  /**
   * @notice Owner-only function to change the `poolWhitelister` address.
   * @param newPoolWhitelister new address for the `poolWhitelister`.
   */
  function setPoolWhitelister(address newPoolWhitelister) external onlyOwner {
    _setPoolWhitelister(newPoolWhitelister);
  }

  /**
   * @notice Owner-only function that adds the provided Pools to the 'whitelist' of pools that stakers can deposit into
   * @param poolsToWhitelist Pools that will be added to the `poolIsWhitelistedForDeposit` mapping (if they aren't in it already)
   * @param thirdPartyTransfersForbiddenValues bool values to set `thirdPartyTransfersForbidden` to for each pool
   */
  function addPoolsToDepositWhitelist(
    IPool[] calldata poolsToWhitelist,
    bool[] calldata thirdPartyTransfersForbiddenValues
  ) external onlyPoolWhitelister {
    require(poolsToWhitelist.length == thirdPartyTransfersForbiddenValues.length, Errors.ARRAY_LENGTH_MISMATCH);
    uint256 poolsToWhitelistLength = poolsToWhitelist.length;
    for (uint256 i = 0; i < poolsToWhitelistLength; ) {
      // change storage and emit event only if pool is not already in whitelist
      if (!poolIsWhitelistedForDeposit[poolsToWhitelist[i]]) {
        poolIsWhitelistedForDeposit[poolsToWhitelist[i]] = true;
        emit PoolAddedToDepositWhitelist(poolsToWhitelist[i]);
        _setThirdPartyTransfersForbidden(poolsToWhitelist[i], thirdPartyTransfersForbiddenValues[i]);
      }
      unchecked {
        ++i;
      }
    }
  }

  /**
   * @notice Owner-only function that removes the provided Pools from the 'whitelist' of pools that stakers can deposit into
   * @param poolsToRemoveFromWhitelist Pools that will be removed to the `poolIsWhitelistedForDeposit` mapping (if they are in it)
   */
  function removePoolsFromDepositWhitelist(IPool[] calldata poolsToRemoveFromWhitelist) external onlyPoolWhitelister {
    uint256 poolsToRemoveFromWhitelistLength = poolsToRemoveFromWhitelist.length;
    for (uint256 i = 0; i < poolsToRemoveFromWhitelistLength; ) {
      // change storage and emit event only if pool is already in whitelist
      if (poolIsWhitelistedForDeposit[poolsToRemoveFromWhitelist[i]]) {
        poolIsWhitelistedForDeposit[poolsToRemoveFromWhitelist[i]] = false;
        emit PoolRemovedFromDepositWhitelist(poolsToRemoveFromWhitelist[i]);
        // Set mapping value to default false value
        _setThirdPartyTransfersForbidden(poolsToRemoveFromWhitelist[i], false);
      }
      unchecked {
        ++i;
      }
    }
  }

  // INTERNAL FUNCTIONS

  /**
   * @notice This function adds `shares` for a given `pool` to the `staker` and runs through the necessary update logic.
   * @param staker The address to add shares to
   * @param token The token that is being deposited (used for indexing)
   * @param pool The Pool in which the `staker` is receiving shares
   * @param shares The amount of shares to grant to the `staker`
   * @dev In particular, this function calls `delegation.increaseDelegatedShares(staker, pool, shares)` to ensure that all
   * delegated shares are tracked, increases the stored share amount in `stakerPoolShares[staker][pool]`, and adds `pool`
   * to the `staker`'s list of pools, if it is not in the list already.
   */
  function _addShares(address staker, IERC20 token, IPool pool, uint256 shares) internal {
    // sanity checks on inputs
    require(staker != address(0), Errors.ZERO_ADDRESS_NOT_VALID);
    require(shares != 0, Errors.ZERO_SHARES_NOT_VALID);

    // if they dont have existing shares of this pool, add it to their strats
    if (stakerPoolShares[staker][pool] == 0) {
      require(stakerPoolList[staker].length < MAX_STAKER_STRATEGY_LIST_LENGTH, Errors.DEPOSIT_EXCEEDS_MAX_LENGTH);
      stakerPoolList[staker].push(pool);
    }

    // add the returned shares to their existing shares for this pool
    stakerPoolShares[staker][pool] += shares;

    emit Deposit(staker, token, pool, shares);
  }

  /**
   * @notice Internal function in which `amount` of ERC20 `token` is transferred from `msg.sender` to the Pool-type contract
   * `pool`, with the resulting shares credited to `staker`.
   * @param staker The address that will be credited with the new shares.
   * @param pool The Pool contract to deposit into.
   * @param token The ERC20 token to deposit.
   * @param amount The amount of `token` to deposit.
   * @return shares The amount of *new* shares in `pool` that have been credited to the `staker`.
   */
  function _depositIntoPool(
    address staker,
    IPool pool,
    IERC20 token,
    uint256 amount
  ) internal onlyPoolsWhitelistedForDeposit(pool) returns (uint256 shares) {
    // transfer tokens from the sender to the pool
    token.safeTransferFrom(msg.sender, address(pool), amount);

    // deposit the assets into the specified pool and get the equivalent amount of shares in that pool
    shares = pool.deposit(token, amount);

    // add the returned shares to the staker's existing shares for this pool
    _addShares(staker, token, pool, shares);

    // Increase shares delegated to operator, if needed
    delegation.increaseDelegatedShares(staker, pool, shares);

    return shares;
  }

  /**
   * @notice Decreases the shares that `staker` holds in `pool` by `shareAmount`.
   * @param staker The address to decrement shares from
   * @param pool The pool for which the `staker`'s shares are being decremented
   * @param shareAmount The amount of shares to decrement
   * @dev If the amount of shares represents all of the staker`s shares in said pool,
   * then the pool is removed from stakerPoolList[staker] and 'true' is returned. Otherwise 'false' is returned.
   */
  function _removeShares(address staker, IPool pool, uint256 shareAmount) internal returns (bool) {
    // sanity checks on inputs
    require(shareAmount != 0, Errors.ZERO_SHARES_NOT_VALID);

    //check that the user has sufficient shares
    uint256 userShares = stakerPoolShares[staker][pool];

    require(shareAmount <= userShares, Errors.SHARE_AMOUNT_TOO_HIGH);
    //unchecked arithmetic since we just checked this above
    unchecked {
      userShares = userShares - shareAmount;
    }

    // subtract the shares from the staker's existing shares for this pool
    stakerPoolShares[staker][pool] = userShares;

    // if no existing shares, remove the pool from the staker's dynamic array of pools
    if (userShares == 0) {
      _removePoolFromStakerPoolList(staker, pool);

      // return true in the event that the pool was removed from stakerPoolList[staker]
      return true;
    }
    // return false in the event that the pool was *not* removed from stakerPoolList[staker]
    return false;
  }

  /**
   * @notice Removes `pool` from `staker`'s dynamic array of pools, i.e. from `stakerPoolList[staker]`
   * @param staker The user whose array will have an entry removed
   * @param pool The Pool to remove from `stakerPoolList[staker]`
   */
  function _removePoolFromStakerPoolList(address staker, IPool pool) internal {
    //loop through all of the pools, find the right one, then replace
    uint256 stratsLength = stakerPoolList[staker].length;
    uint256 j = 0;
    for (; j < stratsLength; ) {
      if (stakerPoolList[staker][j] == pool) {
        //replace the pool with the last pool in the list
        stakerPoolList[staker][j] = stakerPoolList[staker][stakerPoolList[staker].length - 1];
        break;
      }
      unchecked {
        ++j;
      }
    }
    // if we didn't find the pool, revert
    require(j != stratsLength, Errors.POOL_NOT_FOUND);
    // pop off the last entry in the list of pools
    stakerPoolList[staker].pop();
  }

  /**
   * @notice Internal function for modifying `thirdPartyTransfersForbidden`.
   * Used inside of the `setThirdPartyTransfersForbidden` and `addPoolsToDepositWhitelist` functions.
   * @param pool The pool to set `thirdPartyTransfersForbidden` value to
   * @param value bool value to set `thirdPartyTransfersForbidden` to
   */
  function _setThirdPartyTransfersForbidden(IPool pool, bool value) internal {
    emit UpdatedThirdPartyTransfersForbidden(pool, value);
    thirdPartyTransfersForbidden[pool] = value;
  }

  /**
   * @notice Internal function for modifying the `poolWhitelister`. Used inside of the `setPoolWhitelister` and `initialize` functions.
   * @param newPoolWhitelister The new address for the `poolWhitelister` to take.
   */
  function _setPoolWhitelister(address newPoolWhitelister) internal {
    emit PoolWhitelisterChanged(poolWhitelister, newPoolWhitelister);
    poolWhitelister = newPoolWhitelister;
  }

  // VIEW FUNCTIONS
  /**
   * @notice Get all details on the staker's deposits and corresponding shares
   * @param staker The staker of interest, whose deposits this function will fetch
   * @return (staker's pools, shares in these pools)
   */
  function getDeposits(address staker) external view returns (IPool[] memory, uint256[] memory) {
    uint256 poolsLength = stakerPoolList[staker].length;
    uint256[] memory shares = new uint256[](poolsLength);

    for (uint256 i = 0; i < poolsLength; ) {
      shares[i] = stakerPoolShares[staker][stakerPoolList[staker][i]];
      unchecked {
        ++i;
      }
    }
    return (stakerPoolList[staker], shares);
  }

  /// @notice Simple getter function that returns `stakerPoolList[staker].length`.
  function stakerPoolListLength(address staker) external view returns (uint256) {
    return stakerPoolList[staker].length;
  }

  /**
   * @notice Getter function for the current EIP-712 domain separator for this contract.
   * @dev The domain separator will change in the event of a fork that changes the ChainID.
   */
  function domainSeparator() public view returns (bytes32) {
    if (block.chainid == ORIGINAL_CHAIN_ID) {
      return _DOMAIN_SEPARATOR;
    } else {
      return _calculateDomainSeparator();
    }
  }

  // @notice Internal function for calculating the current domain separator of this contract
  function _calculateDomainSeparator() internal view returns (bytes32) {
    return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes('BinLayer')), block.chainid, address(this)));
  }
}
