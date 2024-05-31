import { config, deployments, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import TransparentUpgradeableProxy from './build/TransparentUpgradeableProxy.json';
import ProxyAdminArtifact from './build/ProxyAdmin.json';
import DelegationManagerABI from '../abis/DelegationManager.json';
import StrategyManagerABI from '../abis/StrategyManager.json';
import { BigNumberish } from 'ethers';
import { increaseTime } from '../helpers';

describe('DelegationManager', () => {
  const MAX_PER_DEPOSIT = parseUnits('1000', 18);
  const MAX_TOTAL_DEPOSIT = parseUnits('2000', 18);
  async function deployStrategyManagerFixture() {
    const [owner, pauser, unpauser, slasherProxy, alice, ...addrs] = await ethers.getSigners();
    const MintableERC20 = await ethers.getContractFactory('MintableERC20');
    const usdt = await MintableERC20.deploy('USDT', 'USDT');
    const PauserRegistry = await ethers.getContractFactory('PauserRegistry');
    const pauserRegistry = await PauserRegistry.deploy([pauser.address], unpauser.address);

    const EmptyContract = await ethers.getContractFactory('EmptyContract');
    const emptyContract = await EmptyContract.deploy();
    const ProxyAdmin = await ethers.getContractFactory(
      ProxyAdminArtifact.abi,
      ProxyAdminArtifact.bytecode
    );
    const proxyAdmin = await ProxyAdmin.deploy();
    const StrategyManagerProxy = await ethers.getContractFactory(
      TransparentUpgradeableProxy.abi,
      TransparentUpgradeableProxy.bytecode
    );
    let strategyManagerProxy = await StrategyManagerProxy.deploy(
      emptyContract.address,
      proxyAdmin.address,
      '0x'
    );
    const DelegationManagerProxy = await ethers.getContractFactory(
      TransparentUpgradeableProxy.abi,
      TransparentUpgradeableProxy.bytecode
    );
    let delegationManagerProxy = await DelegationManagerProxy.deploy(
      emptyContract.address,
      proxyAdmin.address,
      '0x'
    );
    const DelegationManager = await ethers.getContractFactory('DelegationManager');
    const delegationManager = await DelegationManager.deploy(
      strategyManagerProxy.address,
      slasherProxy.address
    );
    const StrategyManager = await ethers.getContractFactory('StrategyManager');
    const strategyManager = await StrategyManager.deploy(
      delegationManagerProxy.address,
      slasherProxy.address
    );

    const StrategyBaseTVLLimits = await ethers.getContractFactory('StrategyBaseTVLLimits');
    const strategy = await upgrades.deployProxy(
      StrategyBaseTVLLimits,
      [MAX_PER_DEPOSIT, MAX_TOTAL_DEPOSIT, usdt.address, pauserRegistry.address],
      { constructorArgs: [strategyManagerProxy.address] }
    );

    const ifaceDelegation = new ethers.utils.Interface(DelegationManagerABI);
    await proxyAdmin.upgradeAndCall(
      delegationManagerProxy.address,
      delegationManager.address,
      ifaceDelegation.encodeFunctionData('initialize', [
        owner.address,
        pauserRegistry.address,
        1,
        180,
        [strategy.address],
        [360],
      ])
    );
    const ifaceStrategyManager = new ethers.utils.Interface(StrategyManagerABI);
    await proxyAdmin.upgradeAndCall(
      strategyManagerProxy.address,
      strategyManager.address,
      ifaceStrategyManager.encodeFunctionData('initialize', [
        owner.address,
        owner.address,
        pauserRegistry.address,
        0,
      ])
    );
    strategyManagerProxy = await ethers.getContractAt(
      'StrategyManager',
      strategyManagerProxy.address
    );
    delegationManagerProxy = await ethers.getContractAt(
      'DelegationManager',
      delegationManagerProxy.address
    );
    return {
      owner,
      alice,
      pauser,
      unpauser,
      usdt,
      pauserRegistry,
      strategy,
      strategyManagerProxy,
      delegationManagerProxy,
    };
  }

  describe('DelegationManager', () => {
    describe('Initialize', async () => {
      it('check initialized', async () => {
        const {
          owner,
          usdt,
          pauserRegistry,
          strategy,
          strategyManagerProxy,
          delegationManagerProxy,
        } = await loadFixture(deployStrategyManagerFixture);
        expect(await strategy.strategyManager()).to.be.equal(strategyManagerProxy.address);
        expect((await strategy.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await strategy.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await strategy.underlyingToken()).to.be.equal(usdt.address);
        expect(await strategy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await strategyManagerProxy.owner()).to.be.equal(owner.address);
        expect(await strategyManagerProxy.strategyWhitelister()).to.be.equal(owner.address);
        expect(await strategyManagerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await strategyManagerProxy.delegation()).to.be.equal(delegationManagerProxy.address);
        expect(await delegationManagerProxy.owner()).to.be.equal(owner.address);
        expect(await delegationManagerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await delegationManagerProxy.strategyManager()).to.be.equal(
          strategyManagerProxy.address
        );
        expect(await delegationManagerProxy.minWithdrawalDelay()).to.be.equal(180);
      });
    });
    describe('Delegation', async () => {
      it('add operator will revert when pause', async () => {
        const { owner, delegationManagerProxy } = await loadFixture(deployStrategyManagerFixture);
        const operatorDetailsParams: {
          earningsReceiver: string;
          delegationApprover: string;
          stakerOptOutWindow: number;
        } = {
          earningsReceiver: owner.address,
          delegationApprover: owner.address,
          stakerOptOutWindow: 10,
        };
        await expect(
          delegationManagerProxy.registerAsOperator(
            operatorDetailsParams,
            'https://raw.githubusercontent.com/.../main/metadata.json'
          )
        ).to.revertedWith('Pausable: index is paused');
      });
      it('add operator successful', async () => {
        const { owner, unpauser, delegationManagerProxy } = await loadFixture(
          deployStrategyManagerFixture
        );
        await delegationManagerProxy.connect(unpauser).unpause(0);
        const operatorDetailsParams: {
          earningsReceiver: string;
          delegationApprover: string;
          stakerOptOutWindow: number;
        } = {
          earningsReceiver: owner.address,
          delegationApprover: owner.address,
          stakerOptOutWindow: 10,
        };
        await delegationManagerProxy.registerAsOperator(
          operatorDetailsParams,
          'https://raw.githubusercontent.com/.../main/metadata.json'
        );
        expect(await delegationManagerProxy.isOperator(owner.address)).to.be.equal(true);
        expect(await delegationManagerProxy.earningsReceiver(owner.address)).to.be.equal(
          owner.address
        );
        expect(await delegationManagerProxy.delegationApprover(owner.address)).to.be.equal(
          owner.address
        );
        expect(await delegationManagerProxy.stakerOptOutWindow(owner.address)).to.be.equal(10);
      });
    });
    describe('QueueWithdraw', async () => {
      it('queue withdraw will revert when pause', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy, delegationManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        // Pause
        await delegationManagerProxy.connect(pauser).pause(3);

        expect(await usdt.balanceOf(strategy.address)).to.be.equal(0);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await strategyManagerProxy.depositIntoStrategy(
          strategy.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          strategies: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          strategies: [strategy.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };

        await expect(
          delegationManagerProxy.queueWithdrawals([queueWithdrawParams])
        ).to.revertedWith('Pausable: index is paused');
      });
      it('queue withdraw successful', async () => {
        const {
          owner,
          usdt,
          pauserRegistry,
          strategy,
          strategyManagerProxy,
          delegationManagerProxy,
        } = await loadFixture(deployStrategyManagerFixture);
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(0);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await strategyManagerProxy.depositIntoStrategy(
          strategy.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          strategies: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          strategies: [strategy.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await strategy.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        await delegationManagerProxy.queueWithdrawals([queueWithdrawParams]);
        expect(await strategy.shares(owner.address)).to.be.equal(0);
      });
    });
    describe('CompleteWithdraw', async () => {
      it('complete withdraw revert when not reach withdraw delay', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy, delegationManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(0);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await strategyManagerProxy.depositIntoStrategy(
          strategy.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          strategies: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          strategies: [strategy.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await strategy.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        const tx = await delegationManagerProxy.queueWithdrawals([queueWithdrawParams]);
        const receipt = await tx.wait();
        const eventResp = receipt.events[0].args.withdrawal;
        expect(await strategy.shares(owner.address)).to.be.equal(0);

        const completeQueuedWithdrawalParams: {
          staker: string;
          delegatedTo: string;
          withdrawer: string;
          nonce: string;
          startTimestamp: string;
          strategies: string[];
          shares: BigNumberish[];
        } = {
          staker: eventResp.staker,
          delegatedTo: eventResp.delegatedTo,
          withdrawer: eventResp.withdrawer,
          nonce: eventResp.nonce,
          startTimestamp: eventResp.startTimestamp,
          strategies: eventResp.strategies,
          shares: eventResp.shares,
        };
        await expect(
          delegationManagerProxy.completeQueuedWithdrawals(
            [completeQueuedWithdrawalParams],
            [[usdt.address]],
            [0],
            [true]
          )
        ).to.revertedWith(
          'DelegationManager._completeQueuedWithdrawal: minWithdrawalDelay period has not yet passed'
        );
      });
      it('complete withdraw revert when pause', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy, delegationManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);

        // Pause
        await delegationManagerProxy.connect(pauser).pause(5);

        expect(await usdt.balanceOf(strategy.address)).to.be.equal(0);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await strategyManagerProxy.depositIntoStrategy(
          strategy.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          strategies: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          strategies: [strategy.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await strategy.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        const tx = await delegationManagerProxy.queueWithdrawals([queueWithdrawParams]);
        const receipt = await tx.wait();
        const eventResp = receipt.events[0].args.withdrawal;
        expect(await strategy.shares(owner.address)).to.be.equal(0);

        const completeQueuedWithdrawalParams: {
          staker: string;
          delegatedTo: string;
          withdrawer: string;
          nonce: string;
          startTimestamp: string;
          strategies: string[];
          shares: BigNumberish[];
        } = {
          staker: eventResp.staker,
          delegatedTo: eventResp.delegatedTo,
          withdrawer: eventResp.withdrawer,
          nonce: eventResp.nonce,
          startTimestamp: eventResp.startTimestamp,
          strategies: eventResp.strategies,
          shares: eventResp.shares,
        };
        await increaseTime(360);
        await expect(
          delegationManagerProxy.completeQueuedWithdrawals(
            [completeQueuedWithdrawalParams],
            [[usdt.address]],
            [0],
            [true]
          )
        ).to.revertedWith('Pausable: index is paused');
      });
      it('complete withdraw successful', async () => {
        const {
          owner,
          usdt,
          pauserRegistry,
          strategy,
          strategyManagerProxy,
          delegationManagerProxy,
        } = await loadFixture(deployStrategyManagerFixture);
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(0);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await strategyManagerProxy.depositIntoStrategy(
          strategy.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          strategies: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          strategies: [strategy.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await strategy.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        const tx = await delegationManagerProxy.queueWithdrawals([queueWithdrawParams]);
        const receipt = await tx.wait();
        const eventResp = receipt.events[0].args.withdrawal;
        expect(await strategy.shares(owner.address)).to.be.equal(0);

        const completeQueuedWithdrawalParams: {
          staker: string;
          delegatedTo: string;
          withdrawer: string;
          nonce: string;
          startTimestamp: string;
          strategies: string[];
          shares: BigNumberish[];
        } = {
          staker: eventResp.staker,
          delegatedTo: eventResp.delegatedTo,
          withdrawer: eventResp.withdrawer,
          nonce: eventResp.nonce,
          startTimestamp: eventResp.startTimestamp,
          strategies: eventResp.strategies,
          shares: eventResp.shares,
        };
        await increaseTime(360);
        expect(await usdt.balanceOf(owner.address)).to.be.equal(parseUnits('999000', 18));
        await delegationManagerProxy.completeQueuedWithdrawals(
          [completeQueuedWithdrawalParams],
          [[usdt.address]],
          [0],
          [true]
        );
        expect(await usdt.balanceOf(owner.address)).to.be.equal(parseUnits('1000000', 18));
      });
    });
  });
});
