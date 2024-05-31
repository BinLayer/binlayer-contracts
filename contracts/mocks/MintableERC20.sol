// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import {ERC20, IERC20} from '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';

contract MintableERC20 is ERC20, Ownable {
  constructor(string memory name, string memory symbol) ERC20(name, symbol) {
    _mint(msg.sender, 1e6 * 1e18);
  }

  function mint(uint256 value) public onlyOwner {
    _mint(_msgSender(), value);
  }
}
