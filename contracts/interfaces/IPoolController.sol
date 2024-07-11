// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import './IPool.sol';
import './ISlasher.sol';
import './IDelegationController.sol';

/**
 * @title Interface for the primary entrypoint for funds into BinLayer.
 * @notice See the `PoolController` contract itself for implementation details.
 */
interface IPoolController {
  /**
   * @notice Emitted when a new deposit occurs on behalf of `staker`.
   * @param staker Is the staker who is depositing funds into BinLayer.
   * @param pool Is the pool that `staker` has deposited into.
   * @param token Is the token that `staker` deposited.
   * @param shares Is the number of new shares `staker` has been granted in `pool`.
   */
  event Deposit(address staker, IERC20 token, IPool pool, uint256 shares);

  /// @notice Emitted when `thirdPartyTransfersForbidden` is updated for a pool and value by the owner
  event UpdatedThirdPartyTransfersForbidden(IPool pool, bool value);

  /// @notice Emitted when the `poolWhitelister` is changed
  event PoolWhitelisterChanged(address previousAddress, address newAddress);

  /// @notice Emitted when a pool is added to the approved list of pools for deposit
  event PoolAddedToDepositWhitelist(IPool pool);

  /// @notice Emitted when a pool is removed from the approved list of pools for deposit
  event PoolRemovedFromDepositWhitelist(IPool pool);

  /**
   * @notice Deposits `amount` of `token` into the specified `pool`, with the resultant shares credited to `msg.sender`
   * @param pool is the specified pool where deposit is to be made,
   * @param token is the denomination in which the deposit is to be made,
   * @param amount is the amount of token to be deposited in the pool by the staker
   * @return shares The amount of new shares in the `pool` created as part of the action.
   * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
   * @dev Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
   *
   * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a pool is not recommended.  This can lead to attack vectors
   *          where the token balance and corresponding pool shares are not in sync upon reentrancy.
   */
  function depositIntoPool(IPool pool, IERC20 token, uint256 amount) external returns (uint256 shares);

  /**
   * @notice Deposits `amount` of `token` into the specified `pool`, with the resultant shares credited to `staker`
   * @param staker Staker address
   * @param pool is the specified pool where deposit is to be made,
   * @param token is the denomination in which the deposit is to be made,
   * @param amount is the amount of token to be deposited in the pool by the staker
   * @return shares The amount of new shares in the `pool` created as part of the action.
   * @dev The `msg.sender` must have previously approved this contract to transfer at least `amount` of `token` on their behalf.
   * @dev Cannot be called by an address that is 'frozen' (this function will revert if the `msg.sender` is frozen).
   *
   * WARNING: Depositing tokens that allow reentrancy (eg. ERC-777) into a pool is not recommended.  This can lead to attack vectors
   *          where the token balance and corresponding pool shares are not in sync upon reentrancy.
   */
  function depositIntoPoolWithStaker(
    address staker,
    IPool pool,
    IERC20 token,
    uint256 amount
  ) external returns (uint256 shares);

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
  ) external returns (uint256 shares);

  /// @notice Used by the DelegationController.sol to remove a Staker's shares from a particular pool when entering the withdrawal queue
  function removeShares(address staker, IPool pool, uint256 shares) external;

  /// @notice Used by the DelegationController.sol to award a Staker some shares that have passed through the withdrawal queue
  function addShares(address staker, IERC20 token, IPool pool, uint256 shares) external;

  /// @notice Used by the DelegationController.sol to convert withdrawn shares to tokens and send them to a recipient
  function withdrawSharesAsTokens(address recipient, IPool pool, uint256 shares, IERC20 token) external;

  /// @notice Returns the current shares of `user` in `pool`
  function stakerPoolShares(address user, IPool pool) external view returns (uint256 shares);

  /**
   * @notice Get all details on the staker's deposits and corresponding shares
   * @return (staker's pools, shares in these pools)
   */
  function getDeposits(address staker) external view returns (IPool[] memory, uint256[] memory);

  /// @notice Simple getter function that returns `stakerPoolList[staker].length`.
  function stakerPoolListLength(address staker) external view returns (uint256);

  /**
   * @notice Owner-only function that adds the provided Pools to the 'whitelist' of pools that stakers can deposit into
   * @param poolsToWhitelist Pools that will be added to the `poolIsWhitelistedForDeposit` mapping (if they aren't in it already)
   * @param thirdPartyTransfersForbiddenValues bool values to set `thirdPartyTransfersForbidden` to for each pool
   */
  function addPoolsToDepositWhitelist(
    IPool[] calldata poolsToWhitelist,
    bool[] calldata thirdPartyTransfersForbiddenValues
  ) external;

  /**
   * @notice Owner-only function that removes the provided Pools from the 'whitelist' of pools that stakers can deposit into
   * @param poolsToRemoveFromWhitelist Pools that will be removed to the `poolIsWhitelistedForDeposit` mapping (if they are in it)
   */
  function removePoolsFromDepositWhitelist(IPool[] calldata poolsToRemoveFromWhitelist) external;

  /// @notice Returns the single, central Delegation contract of BinLayer
  function delegation() external view returns (IDelegationController);

  /// @notice Returns the single, central Slasher contract of BinLayer
  function slasher() external view returns (ISlasher);

  /// @notice Returns the address of the `poolWhitelister`
  function poolWhitelister() external view returns (address);

  /**
   * @notice Returns bool for whether or not `pool` enables credit transfers. i.e enabling
   * depositIntoPoolWithSignature calls or queueing withdrawals to a different address than the staker.
   */
  function thirdPartyTransfersForbidden(IPool pool) external view returns (bool);
}
