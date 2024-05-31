// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

interface IERC677Receiver {
  function onTokenTransfer(address sender, uint256 amount, bytes calldata data) external;
}
