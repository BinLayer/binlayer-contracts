// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

interface IERC677 {
  function transferAndCall(address to, uint value, bytes calldata data) external returns (bool success);

  event Transfer(address indexed from, address indexed to, uint value, bytes data);
}
