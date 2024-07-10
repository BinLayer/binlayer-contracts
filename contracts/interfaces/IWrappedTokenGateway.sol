// SPDX-License-Identifier: AGPL-3.0
pragma solidity 0.8.20;

import {IERC20} from '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import {IDelegationController} from '../interfaces/IDelegationController.sol';

interface IWrappedTokenGateway {
  function depositNativeToken(address staker) external payable;

  function withdrawNativeTokens(
    IDelegationController.Withdrawal[] calldata withdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexs,
    bool[] calldata receiveAsTokens
  ) external;
}
