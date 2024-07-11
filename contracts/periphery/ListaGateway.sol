// SPDX-License-Identifier: AGPL-3.0
pragma solidity 0.8.20;

import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';
import {IERC20} from '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import {SafeERC20} from '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import {IListaStakeManager} from '../interfaces/IListaStakeManager.sol';
import {IPool} from '../interfaces/IPool.sol';
import {IPoolController} from '../interfaces/IPoolController.sol';
import {IDelegationController} from '../interfaces/IDelegationController.sol';
import {IListaGateway} from '../interfaces/IListaGateway.sol';

contract ListaGateway is IListaGateway, Ownable {
  using SafeERC20 for IERC20;

  IERC20 internal immutable slisBNB;
  IListaStakeManager internal immutable listaStakeManager;
  IPool internal immutable strategy;
  IPoolController internal immutable poolController;

  constructor(
    address _slisBNB,
    address _owner,
    IListaStakeManager _listaStakeManager,
    IPool _strategy,
    IPoolController _poolController
  ) {
    slisBNB = IERC20(_slisBNB);
    listaStakeManager = _listaStakeManager;
    strategy = _strategy;
    poolController = _poolController;
    _transferOwnership(_owner);
    IERC20(_slisBNB).approve(address(poolController), type(uint256).max);
  }

  function depositNativeToken() external payable {
    uint256 beforeBalance = slisBNB.balanceOf(address(this));
    listaStakeManager.deposit{value: msg.value}();
    uint256 afterBalance = slisBNB.balanceOf(address(this));
    uint256 amountToStake = afterBalance - beforeBalance;
    poolController.depositIntoPoolWithStaker(msg.sender, strategy, slisBNB, amountToStake);
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
   * @dev Get Lista contracts address used by ListaGateway
   */
  function getListaContracts() external view returns (IListaStakeManager _listaStakeManager, IERC20 _slisBNB) {
    _listaStakeManager = listaStakeManager;
    _slisBNB = slisBNB;
  }

  /**
   * @dev Get slisBNB strategy address used by ListaGateway
   */
  function getPoolAddress() external view returns (address) {
    return address(strategy);
  }

  /**
   * @dev Only ListaStakeManager contract is allowed to transfer native token here. Prevent other addresses to send native token to this contract.
   */
  receive() external payable {
    require(msg.sender == address(listaStakeManager), 'Receive not allowed');
  }

  /**
   * @dev Revert fallback calls
   */
  fallback() external payable {
    revert('Fallback not allowed');
  }
}
