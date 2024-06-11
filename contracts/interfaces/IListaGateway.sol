// SPDX-License-Identifier: AGPL-3.0
pragma solidity 0.8.20;

import {IERC20} from '@openzeppelin/contracts/token/ERC20/ERC20.sol';

interface IListaGateway {
  function depositNativeToken() external payable;
}
