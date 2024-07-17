// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.20;

import '@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol';
import '@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol';
import '@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol';
import '../permissions/Pausable.sol';
import '../libraries/EIP1271SignatureUtils.sol';
import './DelegationControllerStorage.sol';
import '../helpers/Errors.sol';

/**
 * @title DelegationController.sol
 * @notice  this is the delegation contract for BinLayer. The primary functions of this contract are:
 * - allowing anyone to register as an operator in BinLayer
 * - enabling operators to set parameters for stakers who delegate to them
 * - allowing any staker to delegate their stake to the operator of their choice (each staker can only delegate to one operator at a time)
 * - allowing a staker to undelegate their assets from the operator they are delegated to (done as part of the withdrawal process, initiated through the PoolController)
 */
contract DelegationController is Initializable, OwnableUpgradeable, Pausable, DelegationControllerStorage, ReentrancyGuardUpgradeable {
  // @dev Flag index that, when set, pauses new delegations
  uint8 internal constant PAUSED_NEW_DELEGATION = 0;

  // @dev Index for flag that pauses unstake when set.
  uint8 internal constant PAUSED_UNSTAKE = 1;

  // @dev Index for flag that pauses withdraw when set.
  uint8 internal constant PAUSED_WITHDRAW = 2;

  // @dev Chain ID at the time of contract deployment
  uint256 internal immutable ORIGINAL_CHAIN_ID;

  // @dev Maximum Value for `stakerOptOutWindow`. Approximately equivalent to 6 months.
  uint256 public constant MAX_STAKER_OPT_OUT_WINDOW = 180 days;

  // @notice Simple permission for functions that are only callable by the PoolController.sol contract
  modifier onlyPoolController() {
    require(msg.sender == address(poolController), Errors.ONLY_POOL_CONTROLLER);
    _;
  }

  /*******************************************************************************
                            INITIALIZING FUNCTIONS
    *******************************************************************************/

  /**
   * @dev Initializes the immutable addresses of the pool mananger and slasher.
   */
  constructor(IPoolController _poolController, ISlasher _slasher) DelegationControllerStorage(_poolController, _slasher) {
    _disableInitializers();
    ORIGINAL_CHAIN_ID = block.chainid;
  }

  /**
   * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
   * minWithdrawalDelay is set only once here
   */
  function initialize(
    address initialOwner,
    IPauserRegistry _pauserRegistry,
    uint256 initialPausedStatus,
    uint256 _minWithdrawalDelay,
    IPool[] calldata _pools,
    uint256[] calldata _withdrawalDelay
  ) external initializer {
    _initializePauser(_pauserRegistry, initialPausedStatus);
    _DOMAIN_SEPARATOR = _calculateDomainSeparator();
    __ReentrancyGuard_init();
    _transferOwnership(initialOwner);
    _setMinWithdrawalDelay(_minWithdrawalDelay);
    _setPoolWithdrawalDelay(_pools, _withdrawalDelay);
  }

  /*******************************************************************************
                            EXTERNAL FUNCTIONS 
    *******************************************************************************/

  /**
   * @notice Registers the caller as an operator in BinLayer.
   * @param registeringOperatorDetails is the `OperatorDetails` for the operator.
   * @param metadataURI is a URI for the operator's metadata, i.e. a link providing more details on the operator.
   *
   * @dev Once an operator is registered, they cannot 'deregister' as an operator, and they will forever be considered "delegated to themself".
   * @dev This function will revert if the caller is already delegated to an operator.
   * @dev Note that the `metadataURI` is *never stored * and is only emitted in the `OperatorMetadataURIUpdated` event
   */
  function registerAsOperator(OperatorDetails calldata registeringOperatorDetails, string calldata metadataURI) external {
    require(!isDelegated(msg.sender), Errors.OPERATOR_ALREADY_REGISTERED);
    _setOperatorDetails(msg.sender, registeringOperatorDetails);
    SignatureWithExpiry memory emptySignatureAndExpiry;
    // delegate from the operator to themselves
    _delegate(msg.sender, msg.sender, emptySignatureAndExpiry, bytes32(0));
    // emit events
    emit OperatorRegistered(msg.sender, registeringOperatorDetails);
    emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
  }

  /**
   * @notice Updates an operator's stored `OperatorDetails`.
   * @param newOperatorDetails is the updated `OperatorDetails` for the operator, to replace their current OperatorDetails`.
   *
   * @dev The caller must have previously registered as an operator in BinLayer.
   */
  function modifyOperatorDetails(OperatorDetails calldata newOperatorDetails) external {
    require(isOperator(msg.sender), Errors.CALLER_NOT_OPERATOR);
    _setOperatorDetails(msg.sender, newOperatorDetails);
  }

  /**
   * @notice Called by an operator to emit an `OperatorMetadataURIUpdated` event indicating the information has updated.
   * @param metadataURI The URI for metadata associated with an operator
   */
  function updateOperatorMetadataURI(string calldata metadataURI) external {
    require(isOperator(msg.sender), Errors.CALLER_NOT_OPERATOR);
    emit OperatorMetadataURIUpdated(msg.sender, metadataURI);
  }

  /**
   * @notice Caller delegates their stake to an operator.
   * @param operator The account (`msg.sender`) is delegating its assets to for use in serving applications built on BinLayer.
   * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
   * @param approverSalt A unique single use value tied to an individual signature.
   * @dev The approverSignatureAndExpiry is used in the event that:
   *          1) the operator's `delegationApprover` address is set to a non-zero value.
   *                  AND
   *          2) neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator
   *             or their delegationApprover is the `msg.sender`, then approval is assumed.
   * @dev In the event that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
   * in this case to save on complexity + gas costs
   */
  function delegateTo(address operator, SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 approverSalt) external {
    require(!isDelegated(msg.sender), Errors.ALREADY_DELEGATED);
    require(isOperator(operator), Errors.NOT_REGISTERED_IN_BINLAYER);
    // go through the internal delegation flow, checking the `approverSignatureAndExpiry` if applicable
    _delegate(msg.sender, operator, approverSignatureAndExpiry, approverSalt);
  }

  /**
   * @notice Caller delegates a staker's stake to an operator with valid signatures from both parties.
   * @param staker The account delegating stake to an `operator` account
   * @param operator The account (`staker`) is delegating its assets to for use in serving applications built on BinLayer.
   * @param stakerSignatureAndExpiry Signed data from the staker authorizing delegating stake to an operator
   * @param approverSignatureAndExpiry is a parameter that will be used for verifying that the operator approves of this delegation action in the event that:
   * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
   *
   * @dev If `staker` is an EOA, then `stakerSignature` is verified to be a valid ECDSA stakerSignature from `staker`, indicating their intention for this action.
   * @dev If `staker` is a contract, then `stakerSignature` will be checked according to EIP-1271.
   * @dev the operator's `delegationApprover` address is set to a non-zero value.
   * @dev neither the operator nor their `delegationApprover` is the `msg.sender`, since in the event that the operator or their delegationApprover
   * is the `msg.sender`, then approval is assumed.
   * @dev This function will revert if the current `block.timestamp` is equal to or exceeds the expiry
   * @dev In the case that `approverSignatureAndExpiry` is not checked, its content is ignored entirely; it's recommended to use an empty input
   * in this case to save on complexity + gas costs
   */
  function delegateToBySignature(
    address staker,
    address operator,
    SignatureWithExpiry memory stakerSignatureAndExpiry,
    SignatureWithExpiry memory approverSignatureAndExpiry,
    bytes32 approverSalt
  ) external {
    // check the signature expiry
    require(stakerSignatureAndExpiry.expiry >= block.timestamp, Errors.SIGNATURE_EXPIRED);

    require(!isDelegated(staker), Errors.ALREADY_DELEGATED);
    require(isOperator(operator), Errors.NOT_REGISTERED_IN_BINLAYER);

    // calculate the digest hash, then increment `staker`'s nonce
    uint256 currentStakerNonce = stakerNonce[staker];
    bytes32 stakerDigestHash = calculateStakerDelegationDigestHash(staker, currentStakerNonce, operator, stakerSignatureAndExpiry.expiry);
    unchecked {
      stakerNonce[staker] = currentStakerNonce + 1;
    }

    // actually check that the signature is valid
    EIP1271SignatureUtils.checkSignature_EIP1271(staker, stakerDigestHash, stakerSignatureAndExpiry.signature);

    // go through the internal delegation flow, checking the `approverSignatureAndExpiry` if applicable
    _delegate(staker, operator, approverSignatureAndExpiry, approverSalt);
  }

  /**
   * Allows the staker, the staker's operator, or that operator's delegationApprover to undelegate
   * a staker from their operator. Undelegation immediately removes ALL active shares/pools from
   * both the staker and operator, and places the shares and pools in the unstake
   */
  function undelegate(address staker) external onlyWhenNotPaused(PAUSED_UNSTAKE) returns (bytes32[] memory withdrawalRoots) {
    require(isDelegated(staker), Errors.STAKER_MUST_BE_DELEGATED);
    require(!isOperator(staker), Errors.CANNOT_UNDELEGATE_OPERATOR);
    require(staker != address(0), Errors.ZERO_ADDRESS_NOT_VALID);
    address operator = delegatedTo[staker];
    require(
      msg.sender == staker || msg.sender == operator || msg.sender == _operatorDetails[operator].delegationApprover,
      Errors.CALLER_CANNOT_UNDELEGATE
    );

    // Gather pools and shares to remove from staker/operator during undelegation
    // Undelegation removes ALL currently-active pools and shares
    (IPool[] memory pools, uint256[] memory shares) = getDelegatableShares(staker);

    // emit an event if this action was not initiated by the staker themselves
    if (msg.sender != staker) {
      emit StakerForceUndelegated(staker, operator);
    }

    // undelegate the staker
    emit StakerUndelegated(staker, operator);
    delegatedTo[staker] = address(0);

    // if no delegatable shares, return an empty array, and don't unstake
    if (pools.length == 0) {
      withdrawalRoots = new bytes32[](0);
    } else {
      withdrawalRoots = new bytes32[](pools.length);
      for (uint256 i = 0; i < pools.length; i++) {
        IPool[] memory singlePool = new IPool[](1);
        uint256[] memory singleShare = new uint256[](1);
        singlePool[0] = pools[i];
        singleShare[0] = shares[i];

        withdrawalRoots[i] = _removeSharesAndUnstake({
          staker: staker,
          operator: operator,
          withdrawer: staker,
          pools: singlePool,
          shares: singleShare
        });
      }
    }

    return withdrawalRoots;
  }

  /**
   * Enables a staker to withdraw a portion of shares. Withdrawn shares/pools are promptly removed.
   * from the staker. If the staker is delegated, withdrawn shares/pools are also removed from
   * their operator.
   *
   * All withdrawn shares/pools are placed in a queue and can be fully withdrawn after a delay.
   */
  function unstakes(UnstakeParams[] calldata unstakeParams) external onlyWhenNotPaused(PAUSED_UNSTAKE) returns (bytes32[] memory) {
    bytes32[] memory withdrawalRoots = new bytes32[](unstakeParams.length);
    address operator = delegatedTo[msg.sender];

    for (uint256 i = 0; i < unstakeParams.length; i++) {
      require(unstakeParams[i].pools.length == unstakeParams[i].shares.length, Errors.INPUT_LENGTH_MISMATCH);
      require(
        unstakeParams[i].withdrawer == msg.sender || unstakeParams[i].withdrawer == wrappedTokenGateway,
        Errors.WITHDRAWER_NOT_STAKER_OR_GATEWAY
      );

      // Remove shares from staker's pools and place pools/shares in unstake.
      // If the staker is delegated to an operator, the operator's delegated shares are also reduced
      // NOTE: This will fail if the staker doesn't have the shares implied by the input parameters
      withdrawalRoots[i] = _removeSharesAndUnstake({
        staker: msg.sender,
        operator: operator,
        withdrawer: unstakeParams[i].withdrawer,
        pools: unstakeParams[i].pools,
        shares: unstakeParams[i].shares
      });
    }
    return withdrawalRoots;
  }

  /**
   * @notice Used to complete the specified `withdrawal`. The caller must match `withdrawal.withdrawer`
   * @param withdrawal The Withdrawal to complete.
   * @param tokens Array in which the i-th entry specifies the `token` input to the 'withdraw' function of the i-th Pool in the `withdrawal.pools` array.
   * This input can be provided with zero length if `receiveAsTokens` is set to 'false' (since in that case, this input will be unused)
   * @param middlewareTimesIndex is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
   * @param receiveAsTokens If true, the shares specified in the withdrawal will be withdrawn from the specified pools themselves
   * and sent to the caller, through calls to `withdrawal.pools[i].withdraw`. If false, then the shares in the specified pools
   * will simply be transferred to the caller directly.
   * @dev middlewareTimesIndex is unused, but will be used in the Slasher eventually
   * @dev beaconChainETHPool shares are non-transferrable, so if `receiveAsTokens = false` and `withdrawal.withdrawer != withdrawal.staker`, note that
   * any beaconChainETHPool shares in the `withdrawal` will be _returned to the staker_, rather than transferred to the withdrawer, unlike shares in
   * any other pools, which will be transferred to the withdrawer.
   */
  function withdraw(
    Withdrawal calldata withdrawal,
    IERC20[] calldata tokens,
    uint256 middlewareTimesIndex,
    bool receiveAsTokens
  ) external onlyWhenNotPaused(PAUSED_WITHDRAW) nonReentrant {
    _withdraw(withdrawal, tokens, middlewareTimesIndex, receiveAsTokens);
  }

  /**
   * @notice Array-ified version of `withdraw`.
   * Used to complete the specified `withdrawals`. The function caller must match `withdrawals[...].withdrawer`
   * @param withdrawals The Withdrawals to complete.
   * @param tokens Array of tokens for each Withdrawal. See `withdraw` for the usage of a single array.
   * @param middlewareTimesIndexes One index to reference per Withdrawal. See `withdraw` for the usage of a single index.
   * @param receiveAsTokens Whether or not to complete each withdrawal as tokens. See `withdraw` for the usage of a single boolean.
   * @dev See `withdraw` for relevant dev tags
   */
  function withdraws(
    Withdrawal[] calldata withdrawals,
    IERC20[][] calldata tokens,
    uint256[] calldata middlewareTimesIndexes,
    bool[] calldata receiveAsTokens
  ) external onlyWhenNotPaused(PAUSED_WITHDRAW) nonReentrant {
    for (uint256 i = 0; i < withdrawals.length; ++i) {
      _withdraw(withdrawals[i], tokens[i], middlewareTimesIndexes[i], receiveAsTokens[i]);
    }
  }

  /**
   * @notice Increases a staker's delegated share balance in a pool.
   * @param staker The address to increase the delegated shares for their operator.
   * @param pool The pool in which to increase the delegated shares.
   * @param shares The number of shares to increase.
   *
   * @dev *If the staker is actively delegated*, then increases the `staker`'s delegated shares in `pool` by `shares`. Otherwise does nothing.
   * @dev Callable only by the PoolController.sol.
   */
  function increaseDelegatedShares(address staker, IPool pool, uint256 shares) external onlyPoolController {
    // if the staker is delegated to an operator
    if (isDelegated(staker)) {
      address operator = delegatedTo[staker];

      // add pool shares to delegate's shares
      _increaseOperatorShares({operator: operator, staker: staker, pool: pool, shares: shares});
    }
  }

  /**
   * @notice Decreases a staker's delegated share balance in a pool.
   * @param staker The address to increase the delegated shares for their operator.
   * @param pool The pool in which to decrease the delegated shares.
   * @param shares The number of shares to decrease.
   *
   * @dev *If the staker is actively delegated*, then decreases the `staker`'s delegated shares in `pool` by `shares`. Otherwise does nothing.
   * @dev Callable only by the PoolController.sol.
   */
  function decreaseDelegatedShares(address staker, IPool pool, uint256 shares) external onlyPoolController {
    // if the staker is delegated to an operator
    if (isDelegated(staker)) {
      address operator = delegatedTo[staker];

      // subtract pool shares from delegate's shares
      _decreaseOperatorShares({operator: operator, staker: staker, pool: pool, shares: shares});
    }
  }

  /**
   * @notice Owner-only function for modifying the value of the `minWithdrawalDelay` variable.
   * @param newMinWithdrawalDelay new value of `minWithdrawalDelay`.
   */
  function setMinWithdrawalDelay(uint256 newMinWithdrawalDelay) external onlyOwner {
    _setMinWithdrawalDelay(newMinWithdrawalDelay);
  }

  /**
   * @notice Called by owner to set the minimum withdrawal delay for each passed in pool
   * Note that the min cooldown to complete a withdrawal of a pool is
   * MAX(minWithdrawalDelay, poolWithdrawalDelay[pool])
   * @param pools The pools to set the minimum withdrawal delay for
   * @param withdrawalDelay The minimum withdrawal delay to set for each pool
   */
  function setPoolWithdrawalDelay(IPool[] calldata pools, uint256[] calldata withdrawalDelay) external onlyOwner {
    _setPoolWithdrawalDelay(pools, withdrawalDelay);
  }

  /**
   * @notice Called by owner to update the wrapped token gateway
   * @param _newWrappedTokenGateway New wrapped token gateway address
   */
  function updateWrappedTokenGateway(address _newWrappedTokenGateway) external onlyOwner {
    emit UpdateWrappedTokenGateway(wrappedTokenGateway, _newWrappedTokenGateway);
    wrappedTokenGateway = _newWrappedTokenGateway;
  }

  /*******************************************************************************
                            INTERNAL FUNCTIONS
    *******************************************************************************/

  /**
   * @notice Sets operator parameters in the `_operatorDetails` mapping.
   * @param operator The account registered as an operator updating their operatorDetails
   * @param newOperatorDetails The new parameters for the operator
   */
  function _setOperatorDetails(address operator, OperatorDetails calldata newOperatorDetails) internal {
    require(newOperatorDetails.stakerOptOutWindow <= MAX_STAKER_OPT_OUT_WINDOW, Errors.OPT_OUT_WINDOW_EXCEEDS_MAX);
    require(newOperatorDetails.stakerOptOutWindow >= _operatorDetails[operator].stakerOptOutWindow, Errors.DECREASE_OPT_OUT_WINDOW);
    _operatorDetails[operator] = newOperatorDetails;
    emit OperatorDetailsModified(msg.sender, newOperatorDetails);
  }

  /**
   * @notice Delegates *from* a `staker` *to* an `operator`.
   * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
   * @param operator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
   * @param approverSignatureAndExpiry Verifies the operator approves of this delegation
   * @param approverSalt Is a salt used to help guarantee signature uniqueness. Each salt can only be used once by a given approver.
   * @dev Ensures that:
   *          1) the `staker` is not already delegated to an operator
   *          2) the `operator` has indeed registered as an operator in BinLayer
   *          3) if applicable, that the approver signature is valid and non-expired
   */
  function _delegate(
    address staker,
    address operator,
    SignatureWithExpiry memory approverSignatureAndExpiry,
    bytes32 approverSalt
  ) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
    // fetch the operator's `delegationApprover` address and store it in memory in case we need to use it multiple times
    address _delegationApprover = _operatorDetails[operator].delegationApprover;
    /**
     * Check the `_delegationApprover`'s signature, if applicable.
     * If the `_delegationApprover` is the zero address, then the operator allows all stakers to delegate to them and this verification is skipped.
     * If the `_delegationApprover` or the `operator` themselves is the caller, then approval is assumed and signature verification is skipped as well.
     */
    if (_delegationApprover != address(0) && msg.sender != _delegationApprover && msg.sender != operator) {
      // check the signature expiry
      require(approverSignatureAndExpiry.expiry >= block.timestamp, Errors.SIGNATURE_EXPIRED);
      // check that the salt hasn't been used previously, then mark the salt as spent
      require(!delegationApproverSaltIsSpent[_delegationApprover][approverSalt], Errors.SALT_ALREADY_SPENT);
      delegationApproverSaltIsSpent[_delegationApprover][approverSalt] = true;

      // calculate the digest hash
      bytes32 approverDigestHash = calculateDelegationApprovalDigestHash(
        staker,
        operator,
        _delegationApprover,
        approverSalt,
        approverSignatureAndExpiry.expiry
      );

      // actually check that the signature is valid
      EIP1271SignatureUtils.checkSignature_EIP1271(_delegationApprover, approverDigestHash, approverSignatureAndExpiry.signature);
    }

    // record the delegation relation between the staker and operator, and emit an event
    delegatedTo[staker] = operator;
    emit StakerDelegated(staker, operator);

    (IPool[] memory pools, uint256[] memory shares) = getDelegatableShares(staker);

    for (uint256 i = 0; i < pools.length; ) {
      _increaseOperatorShares({operator: operator, staker: staker, pool: pools[i], shares: shares[i]});

      unchecked {
        ++i;
      }
    }
  }

  /**
   * @dev commented-out param (middlewareTimesIndex) is the index in the operator that the staker who triggered the withdrawal was delegated to's middleware times array
   * This param is intended to be passed on to the Slasher contract.
   */
  function _withdraw(
    Withdrawal calldata withdrawal,
    IERC20[] calldata tokens,
    uint256 /*middlewareTimesIndex*/,
    bool receiveAsTokens
  ) internal {
    bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

    require(pendingWithdrawals[withdrawalRoot], Errors.ACTION_NOT_IN_QUEUE);

    require(withdrawal.startTimestamp + minWithdrawalDelay <= block.timestamp, Errors.MIN_WITHDRAWAL_DELAY_NOT_PASSED);

    require(msg.sender == withdrawal.withdrawer, Errors.ONLY_WITHDRAWER_CAN_COMPLETE);

    if (receiveAsTokens) {
      require(tokens.length == withdrawal.pools.length, Errors.INPUT_LENGTH_MISMATCH);
    }

    // Remove `withdrawalRoot` from pending roots
    delete pendingWithdrawals[withdrawalRoot];

    // Finalize action by converting shares to tokens for each pool, or
    // by re-awarding shares in each pool.
    if (receiveAsTokens) {
      for (uint256 i = 0; i < withdrawal.pools.length; ) {
        require(
          withdrawal.startTimestamp + poolWithdrawalDelay[withdrawal.pools[i]] <= block.timestamp,
          Errors.WITHDRAWAL_DELAY_NOT_PASSED
        );

        _withdrawSharesAsTokens({
          staker: withdrawal.staker,
          withdrawer: msg.sender,
          pool: withdrawal.pools[i],
          shares: withdrawal.shares[i],
          token: tokens[i]
        });
        unchecked {
          ++i;
        }
      }
      // Award shares back in PoolController.sol. If withdrawer is delegated, increase the shares delegated to the operator
    } else {
      address currentOperator = delegatedTo[msg.sender];
      for (uint256 i = 0; i < withdrawal.pools.length; ) {
        require(
          withdrawal.startTimestamp + poolWithdrawalDelay[withdrawal.pools[i]] <= block.timestamp,
          Errors.WITHDRAWAL_DELAY_NOT_PASSED
        );

        poolController.addShares(msg.sender, tokens[i], withdrawal.pools[i], withdrawal.shares[i]);
        // Similar to `isDelegated` logic
        if (currentOperator != address(0)) {
          _increaseOperatorShares({
            operator: currentOperator,
            // the 'staker' here is the address receiving new shares
            staker: msg.sender,
            pool: withdrawal.pools[i],
            shares: withdrawal.shares[i]
          });
        }

        unchecked {
          ++i;
        }
      }
    }

    emit WithdrawalCompleted(withdrawalRoot);
  }

  // @notice Increases `operator`s delegated shares in `pool` by `shares` and emits an `OperatorSharesIncreased` event
  function _increaseOperatorShares(address operator, address staker, IPool pool, uint256 shares) internal {
    operatorShares[operator][pool] += shares;
    emit OperatorSharesIncreased(operator, staker, pool, shares);
  }

  // @notice Decreases `operator`s delegated shares in `pool` by `shares` and emits an `OperatorSharesDecreased` event
  function _decreaseOperatorShares(address operator, address staker, IPool pool, uint256 shares) internal {
    // This will revert on underflow, so no check needed
    operatorShares[operator][pool] -= shares;
    emit OperatorSharesDecreased(operator, staker, pool, shares);
  }

  /**
   * @notice Removes `shares` in `pools` from `staker` who is currently delegated to `operator` and unstake to the `withdrawer`.
   * @dev If the `operator` is indeed an operator, then the operator's delegated shares in the `pools` are also decreased appropriately.
   * @dev If `withdrawer` is not the same address as `staker`, then thirdPartyTransfersForbidden[pool] must be set to false in the PoolController.sol.
   */
  function _removeSharesAndUnstake(
    address staker,
    address operator,
    address withdrawer,
    IPool[] memory pools,
    uint256[] memory shares
  ) internal returns (bytes32) {
    require(staker != address(0), Errors.ZERO_ADDRESS_NOT_VALID);
    require(pools.length != 0, Errors.POOLS_CANNOT_BE_EMPTY);

    // Remove shares from staker and operator
    // Each of these operations fail if we attempt to remove more shares than exist
    for (uint256 i = 0; i < pools.length; ) {
      // Similar to `isDelegated` logic
      if (operator != address(0)) {
        _decreaseOperatorShares({operator: operator, staker: staker, pool: pools[i], shares: shares[i]});
      }

      require(staker == withdrawer || !poolController.thirdPartyTransfersForbidden(pools[i]), Errors.THIRD_PARTY_TRANSFERS_FORBIDDEN);
      // this call will revert if `shares[i]` exceeds the Staker's current shares in `pools[i]`
      poolController.removeShares(staker, pools[i], shares[i]);

      unchecked {
        ++i;
      }
    }

    // Create queue entry and increment withdrawal nonce
    uint256 nonce = cumulativeWithdrawalsQueued[staker];
    cumulativeWithdrawalsQueued[staker]++;

    Withdrawal memory withdrawal = Withdrawal({
      staker: staker,
      delegatedTo: operator,
      withdrawer: withdrawer,
      nonce: nonce,
      startTimestamp: uint32(block.timestamp),
      pools: pools,
      shares: shares
    });

    bytes32 withdrawalRoot = calculateWithdrawalRoot(withdrawal);

    // Place withdrawal in queue
    pendingWithdrawals[withdrawalRoot] = true;

    emit WithdrawalQueued(withdrawalRoot, withdrawal);
    return withdrawalRoot;
  }

  /**
   * @notice Withdraws `shares` in `pool` to `withdrawer`. Call is ultimately forwarded to the `pool` with info on the `token`.
   */
  function _withdrawSharesAsTokens(address staker, address withdrawer, IPool pool, uint256 shares, IERC20 token) internal {
    poolController.withdrawSharesAsTokens(withdrawer, pool, shares, token);
  }

  function _setMinWithdrawalDelay(uint256 _minWithdrawalDelay) internal {
    require(_minWithdrawalDelay <= MAX_WITHDRAWAL_DELAY, Errors.MIN_WITHDRAWAL_DELAY_EXCEEDS_MAX);
    emit MinWithdrawalDelaySet(minWithdrawalDelay, _minWithdrawalDelay);
    minWithdrawalDelay = _minWithdrawalDelay;
  }

  /**
   * @notice Sets the withdrawal delay for each pool in `_pools` to `_withdrawalDelay`.
   * gets called when initializing contract or by calling `setPoolWithdrawalDelay`
   */
  function _setPoolWithdrawalDelay(IPool[] calldata _pools, uint256[] calldata _withdrawalDelay) internal {
    require(_pools.length == _withdrawalDelay.length, Errors.INPUT_LENGTH_MISMATCH);
    uint256 numStrats = _pools.length;
    for (uint256 i = 0; i < numStrats; ++i) {
      IPool pool = _pools[i];
      uint256 prevPoolWithdrawalDelay = poolWithdrawalDelay[pool];
      uint256 newPoolWithdrawalDelay = _withdrawalDelay[i];
      require(newPoolWithdrawalDelay <= MAX_WITHDRAWAL_DELAY, Errors.POOL_WITHDRAWAL_DELAY_EXCEEDS_MAX);

      // set the new withdrawal delay
      poolWithdrawalDelay[pool] = newPoolWithdrawalDelay;
      emit PoolWithdrawalDelaySet(pool, prevPoolWithdrawalDelay, newPoolWithdrawalDelay);
    }
  }

  /*******************************************************************************
                            VIEW FUNCTIONS
    *******************************************************************************/

  /**
   * @notice Getter function for the current EIP-712 domain separator for this contract.
   *
   * @dev The domain separator will change in the event of a fork that changes the ChainID.
   * @dev By introducing a domain separator the DApp developers are guaranteed that there can be no signature collision.
   * for more detailed information please read EIP-712.
   */
  function domainSeparator() public view returns (bytes32) {
    if (block.chainid == ORIGINAL_CHAIN_ID) {
      return _DOMAIN_SEPARATOR;
    } else {
      return _calculateDomainSeparator();
    }
  }

  /**
   * @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
   */
  function isDelegated(address staker) public view returns (bool) {
    return (delegatedTo[staker] != address(0));
  }

  /**
   * @notice Returns true is an operator has previously registered for delegation.
   */
  function isOperator(address operator) public view returns (bool) {
    return delegatedTo[operator] == operator;
  }

  /**
   * @notice Returns the OperatorDetails struct associated with an `operator`.
   */
  function operatorDetails(address operator) external view returns (OperatorDetails memory) {
    return _operatorDetails[operator];
  }

  /**
   * @notice Returns the delegationApprover account for an operator
   */
  function delegationApprover(address operator) external view returns (address) {
    return _operatorDetails[operator].delegationApprover;
  }

  /**
   * @notice Returns the stakerOptOutWindow for an operator
   */
  function stakerOptOutWindow(address operator) external view returns (uint256) {
    return _operatorDetails[operator].stakerOptOutWindow;
  }

  /// @notice Given array of pools, returns array of shares for the operator
  function getOperatorShares(address operator, IPool[] memory pools) public view returns (uint256[] memory) {
    uint256[] memory shares = new uint256[](pools.length);
    for (uint256 i = 0; i < pools.length; ++i) {
      shares[i] = operatorShares[operator][pools[i]];
    }
    return shares;
  }

  /**
   * @notice Returns the number of actively-delegatable shares a staker has across all pools.
   * @dev Returns two empty arrays in the case that the Staker has no actively-delegateable shares.
   */
  function getDelegatableShares(address staker) public view returns (IPool[] memory, uint256[] memory) {
    // Get currently active shares and pools for `staker`
    (IPool[] memory poolControllerStrats, uint256[] memory poolControllerShares) = poolController.getDeposits(staker);
    return (poolControllerStrats, poolControllerShares);
  }

  /**
   * @notice Given a list of pools, return the minimum cooldown that must pass to withdraw
   * from all the inputted pools. Return value is >= minWithdrawalDelay as this is the global min withdrawal delay.
   * @param pools The pools to check withdrawal delays for
   */
  function getWithdrawalDelay(IPool[] calldata pools) public view returns (uint256) {
    uint256 withdrawalDelay = minWithdrawalDelay;
    for (uint256 i = 0; i < pools.length; ++i) {
      uint256 currWithdrawalDelay = poolWithdrawalDelay[pools[i]];
      if (currWithdrawalDelay > withdrawalDelay) {
        withdrawalDelay = currWithdrawalDelay;
      }
    }
    return withdrawalDelay;
  }

  /// @notice Returns the keccak256 hash of `withdrawal`.
  function calculateWithdrawalRoot(Withdrawal memory withdrawal) public pure returns (bytes32) {
    return keccak256(abi.encode(withdrawal));
  }

  /**
   * @notice Calculates the digestHash for a `staker` to sign to delegate to an `operator`
   * @param staker The signing staker
   * @param operator The operator who is being delegated to
   * @param expiry The desired expiry time of the staker's signature
   */
  function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) external view returns (bytes32) {
    // fetch the staker's current nonce
    uint256 currentStakerNonce = stakerNonce[staker];
    // calculate the digest hash
    return calculateStakerDelegationDigestHash(staker, currentStakerNonce, operator, expiry);
  }

  /**
   * @notice Calculates the digest hash to be signed and used in the `delegateToBySignature` function
   * @param staker The signing staker
   * @param _stakerNonce The nonce of the staker. In practice we use the staker's current nonce, stored at `stakerNonce[staker]`
   * @param operator The operator who is being delegated to
   * @param expiry The desired expiry time of the staker's signature
   */
  function calculateStakerDelegationDigestHash(
    address staker,
    uint256 _stakerNonce,
    address operator,
    uint256 expiry
  ) public view returns (bytes32) {
    // calculate the struct hash
    bytes32 stakerStructHash = keccak256(abi.encode(STAKER_DELEGATION_TYPEHASH, staker, operator, _stakerNonce, expiry));
    // calculate the digest hash
    bytes32 stakerDigestHash = keccak256(abi.encodePacked('\x19\x01', domainSeparator(), stakerStructHash));
    return stakerDigestHash;
  }

  /**
   * @notice Calculates the digest hash to be signed by the operator's delegationApprove and used in the `delegateTo` and `delegateToBySignature` functions.
   * @param staker The account delegating their stake
   * @param operator The account receiving delegated stake
   * @param _delegationApprover the operator's `delegationApprover` who will be signing the delegationHash (in general)
   * @param approverSalt A unique and single use value associated with the approver signature.
   * @param expiry Time after which the approver's signature becomes invalid
   */
  function calculateDelegationApprovalDigestHash(
    address staker,
    address operator,
    address _delegationApprover,
    bytes32 approverSalt,
    uint256 expiry
  ) public view returns (bytes32) {
    // calculate the struct hash
    bytes32 approverStructHash = keccak256(
      abi.encode(DELEGATION_APPROVAL_TYPEHASH, _delegationApprover, staker, operator, approverSalt, expiry)
    );
    // calculate the digest hash
    bytes32 approverDigestHash = keccak256(abi.encodePacked('\x19\x01', domainSeparator(), approverStructHash));
    return approverDigestHash;
  }

  /**
   * @dev Recalculates the domain separator when the chainid changes due to a fork.
   */
  function _calculateDomainSeparator() internal view returns (bytes32) {
    return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes('BinLayer')), block.chainid, address(this)));
  }
}
