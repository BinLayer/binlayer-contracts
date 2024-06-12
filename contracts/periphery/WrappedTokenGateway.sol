// SPDX-License-Identifier: AGPL-3.0
pragma solidity 0.8.20;

import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';
import {IERC20} from '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import {SafeERC20} from '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import {IWrappedToken} from '../interfaces/IWrappedToken.sol';
import {IStrategy} from '../interfaces/IStrategy.sol';
import {IStrategyManager} from '../interfaces/IStrategyManager.sol';
import {IDelegationManager} from '../interfaces/IDelegationManager.sol';
import {IWrappedTokenGateway} from '../interfaces/IWrappedTokenGateway.sol';

contract WrappedTokenGateway is IWrappedTokenGateway, Ownable {
  using SafeERC20 for IERC20;

  IWrappedToken internal immutable wrappedToken;
  IStrategy internal immutable strategy;
  IStrategyManager internal immutable strategyManager;
  IDelegationManager internal immutable delegationManager;

  constructor(
    address _wrappedToken,
    address _owner,
    IStrategy _strategy,
    IStrategyManager _strategyManager,
    IDelegationManager _delegationManager
  ) {
    wrappedToken = IWrappedToken(_wrappedToken);
    strategy = _strategy;
    strategyManager = _strategyManager;
    delegationManager = _delegationManager;
    _transferOwnership(_owner);
    IWrappedToken(_wrappedToken).approve(address(strategyManager), type(uint256).max);
  }

  function depositNativeToken(address staker) external payable {
    wrappedToken.deposit{value: msg.value}();
    strategyManager.depositIntoStrategyWithStaker(staker, strategy, IERC20(address(wrappedToken)), msg.value);
  }

  function withdrawNativeTokens(
    IDelegationManager.Withdrawal[] calldata withdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexs,
    bool[] calldata receiveAsTokens
  ) external {
    for (uint256 i = 0; i < withdrawals.length; i++) {
      require(withdrawals[i].staker == msg.sender, 'Withdrawer must be staker');
      for (uint256 j = 0; j < withdrawals[i].strategies.length; j++) {
        require(withdrawals[i].strategies[j] == strategy, 'Only support wrapped token strategy');
      }
    }
    uint256 beforeBalance = wrappedToken.balanceOf(address(this));
    delegationManager.completeQueuedWithdrawals(withdrawals, tokens, middlewareTimesIndexs, receiveAsTokens);
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
   * @dev Get WrappedToken strategy address used by WrappedTokenGateway
   */
  function getStrategyAddress() external view returns (address) {
    return address(strategy);
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
