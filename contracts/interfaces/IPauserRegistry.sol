// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

/**
 * @title Interface for the `PauserRegistry` contract.
 */
interface IPauserRegistry {
  event PauserStatusChanged(address pauser, bool canPause);

  event UnpauserChanged(address previousUnpauser, address newUnpauser);

  /// @notice Mapping of addresses to whether they hold the pauser role.
  function isPauser(address pauser) external view returns (bool);

  /// @notice Unique address that holds the unpauser role. Capable of changing *both* the pauser and unpauser addresses.
  function unpauser() external view returns (address);
}
