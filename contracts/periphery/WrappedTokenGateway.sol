// SPDX-License-Identifier: AGPL-3.0
pragma solidity 0.8.20;

import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';
import {IERC20} from '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import {SafeERC20} from '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import {IWrappedToken} from '../interfaces/IWrappedToken.sol';
import {IPool} from '../interfaces/IPool.sol';
import {IPoolController} from '../interfaces/IPoolController.sol';
import {IDelegationController} from '../interfaces/IDelegationController.sol';
import {IWrappedTokenGateway} from '../interfaces/IWrappedTokenGateway.sol';
import {Errors} from '../helpers/Errors.sol';

contract WrappedTokenGateway is IWrappedTokenGateway, Ownable {
  using SafeERC20 for IERC20;

  IWrappedToken internal immutable wrappedToken;
  IPool internal immutable pool;
  IPoolController internal immutable poolController;
  IDelegationController internal immutable delegationController;

  constructor(
    address _wrappedToken,
    address _owner,
    IPool _pool,
    IPoolController _poolController,
    IDelegationController _delegationController
  ) {
    wrappedToken = IWrappedToken(_wrappedToken);
    pool = _pool;
    poolController = _poolController;
    delegationController = _delegationController;
    _transferOwnership(_owner);
    IWrappedToken(_wrappedToken).approve(address(poolController), type(uint256).max);
  }

  function depositNativeToken(address staker) external payable {
    wrappedToken.deposit{value: msg.value}();
    poolController.depositIntoPoolWithStaker(staker, pool, IERC20(address(wrappedToken)), msg.value);
  }

  function withdrawNativeTokens(
    IDelegationController.Withdrawal[] calldata withdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexs,
    bool[] calldata receiveAsTokens
  ) external {
    for (uint256 i = 0; i < withdrawals.length; i++) {
      require(withdrawals[i].staker == msg.sender, Errors.WITHDRAWER_MUST_BE_STAKER);
      for (uint256 j = 0; j < withdrawals[i].pools.length; j++) {
        require(withdrawals[i].pools[j] == pool, Errors.ONLY_SUPPORT_WRAPPED_TOKEN_POOL);
      }
    }
    uint256 beforeBalance = wrappedToken.balanceOf(address(this));
    delegationController.completeQueuedWithdrawals(withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens);
    uint256 afterBalance = wrappedToken.balanceOf(address(this));
    uint256 amountToWithdraw = afterBalance - beforeBalance;
    wrappedToken.withdraw(amountToWithdraw);
    _safeTransferNative(msg.sender, amountToWithdraw);
  }

  /**
   * @dev transfer Native to an address, revert if it fails.
   * @param to recipient of the transfer
   * @param value the amount to send
   */
  function _safeTransferNative(address to, uint256 value) internal {
    (bool success, ) = to.call{value: value}(new bytes(0));
    require(success, 'NATIVE_TRANSFER_FAILED');
  }

  /**
   * @dev transfer ERC20 from the utility contract, for ERC20 recovery in case of stuck tokens due
   * direct transfers to the contract address.
   * @param token token to transfer
   * @param to recipient of the transfer
   * @param amount amount to send
   */
  function emergencyTokenTransfer(address token, address to, uint256 amount) external onlyOwner {
    IERC20(token).safeTransfer(to, amount);
  }

  /**
   * @dev transfer native token from the utility contract, for native token recovery in case of stuck native
   * due to selfdestructs or Native transfers to the pre-computed contract address before deployment.
   * @param to recipient of the transfer
   * @param amount amount to send
   */
  function emergencyNativeTransfer(address to, uint256 amount) external onlyOwner {
    _safeTransferNative(to, amount);
  }

  /**
   * @dev Get WrappedToken address used by WrappedTokenGateway
   */
  function getWrappedTokenAddress() external view returns (address) {
    return address(wrappedToken);
  }

  /**
   * @dev Get WrappedToken pool address used by WrappedTokenGateway
   */
  function getPoolAddress() external view returns (address) {
    return address(pool);
  }

  /**
   * @dev Only WrappedToken contract is allowed to transfer native token here. Prevent other addresses to send native token to this contract.
   */
  receive() external payable {
    require(msg.sender == address(wrappedToken), 'Receive not allowed');
  }

  /**
   * @dev Revert fallback calls
   */
  fallback() external payable {
    revert('Fallback not allowed');
  }
}
