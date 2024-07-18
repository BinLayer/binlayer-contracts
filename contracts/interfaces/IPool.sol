// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';

/**
 * @title Minimal interface for an `Pool` contract.
 * @notice Custom `Pool` implementations may expand extensively on this interface.
 */
interface IPool {
  /**
   * @notice Used to deposit tokens into this Pool
   * @param token is the ERC20 token being deposited
   * @param amount is the amount of token being deposited
   * @dev This function is only callable by the poolController contract. It is invoked inside of the poolController's
   * `depositIntoPool` function, and individual share balances are recorded in the poolController as well.
   * @return newShares is the number of new shares issued at the current exchange ratio.
   */
  function deposit(IERC20 token, uint256 amount) external returns (uint256);

  /**
   * @notice Used to withdraw tokens from this Pool, to the `recipient`'s address
   * @param recipient is the address to receive the withdrawn funds
   * @param token is the ERC20 token being transferred out
   * @param amountShares is the amount of shares being withdrawn
   * @dev This function is only callable by the poolController contract. It is invoked inside of the poolController's
   * other functions, and individual share balances are recorded in the poolController as well.
   */
  function withdraw(address recipient, IERC20 token, uint256 amountShares) external;

  /**
   * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this pool.
   * @notice In contrast to `sharesToUnderlyingView`, this function **may** make state modifications
   * @param amountShares is the amount of shares to calculate its conversion into the underlying token
   * @return The amount of underlying tokens corresponding to the input `amountShares`
   * @dev Implementation for these functions in particular may vary significantly for different pools
   */
  function sharesToUnderlying(uint256 amountShares) external returns (uint256);

  /**
   * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this pool.
   * @notice In contrast to `underlyingToSharesView`, this function **may** make state modifications
   * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into pool shares
   * @return The amount of underlying tokens corresponding to the input `amountShares`
   * @dev Implementation for these functions in particular may vary significantly for different pools
   */
  function underlyingToShares(uint256 amountUnderlying) external returns (uint256);

  /**
   * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
   * this pool. In contrast to `userUnderlyingView`, this function **may** make state modifications
   */
  function userUnderlying(address user) external returns (uint256);

  /**
   * @notice convenience function for fetching the current total shares of `user` in this pool, by
   * querying the `poolController` contract
   */
  function shares(address user) external view returns (uint256);

  /**
   * @notice Used to convert a number of shares to the equivalent amount of underlying tokens for this pool.
   * @notice In contrast to `sharesToUnderlying`, this function guarantees no state modifications
   * @param amountShares is the amount of shares to calculate its conversion into the underlying token
   * @return The amount of shares corresponding to the input `amountUnderlying`
   * @dev Implementation for these functions in particular may vary significantly for different pools
   */
  function sharesToUnderlyingView(uint256 amountShares) external view returns (uint256);

  /**
   * @notice Used to convert an amount of underlying tokens to the equivalent amount of shares in this pool.
   * @notice In contrast to `underlyingToShares`, this function guarantees no state modifications
   * @param amountUnderlying is the amount of `underlyingToken` to calculate its conversion into pool shares
   * @return The amount of shares corresponding to the input `amountUnderlying`
   * @dev Implementation for these functions in particular may vary significantly for different pools
   */
  function underlyingToSharesView(uint256 amountUnderlying) external view returns (uint256);

  /**
   * @notice convenience function for fetching the current underlying value of all of the `user`'s shares in
   * this pool. In contrast to `userUnderlying`, this function guarantees no state modifications
   */
  function userUnderlyingView(address user) external view returns (uint256);

  /// @notice The underlying token for shares in this Pool
  function underlyingToken() external view returns (IERC20);

  /// @notice The total number of extant shares in this Pool
  function totalShares() external view returns (uint256);

  /// @notice Returns either a brief string explaining the pool's goal & purpose, or a link to metadata that explains in more detail.
  function explanation() external view returns (string memory);
}
