import { config, deployments, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import TransparentUpgradeableProxy from './build/TransparentUpgradeableProxy.json';
import ProxyAdminArtifact from './build/ProxyAdmin.json';
import DelegationControllerABI from '../abis/DelegationController.json';
import StrategyManagerABI from '../abis/StrategyManager.json';
import { BigNumberish } from 'ethers';
import { increaseTime } from '../helpers';

describe('DelegationController', () => {
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
    const DelegationControllerProxy = await ethers.getContractFactory(
      TransparentUpgradeableProxy.abi,
      TransparentUpgradeableProxy.bytecode
    );
    let delegationControllerProxy = await DelegationControllerProxy.deploy(
      emptyContract.address,
      proxyAdmin.address,
      '0x'
    );
    const DelegationController = await ethers.getContractFactory('DelegationController');
    const delegationController = await DelegationController.deploy(
      strategyManagerProxy.address,
      slasherProxy.address
    );
    const StrategyManager = await ethers.getContractFactory('StrategyManager');
    const strategyManager = await StrategyManager.deploy(
      delegationControllerProxy.address,
      slasherProxy.address
    );

    const StrategyBaseTVLLimits = await ethers.getContractFactory('StrategyBaseTVLLimits');
    const strategy = await upgrades.deployProxy(
      StrategyBaseTVLLimits,
      [MAX_PER_DEPOSIT, MAX_TOTAL_DEPOSIT, usdt.address, pauserRegistry.address],
      { constructorArgs: [strategyManagerProxy.address] }
    );

    const ifaceDelegation = new ethers.utils.Interface(DelegationControllerABI);
    await proxyAdmin.upgradeAndCall(
      delegationControllerProxy.address,
      delegationController.address,
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
    delegationControllerProxy = await ethers.getContractAt(
      'DelegationController',
      delegationControllerProxy.address
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
      delegationControllerProxy,
    };
  }

  describe('DelegationController', () => {
    describe('Initialize', async () => {
      it('check initialized', async () => {
        const {
          owner,
          usdt,
          pauserRegistry,
          strategy,
          strategyManagerProxy,
          delegationControllerProxy,
        } = await loadFixture(deployStrategyManagerFixture);
        expect(await strategy.strategyManager()).to.be.equal(strategyManagerProxy.address);
        expect((await strategy.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await strategy.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await strategy.underlyingToken()).to.be.equal(usdt.address);
        expect(await strategy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await strategyManagerProxy.owner()).to.be.equal(owner.address);
        expect(await strategyManagerProxy.strategyWhitelister()).to.be.equal(owner.address);
        expect(await strategyManagerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await strategyManagerProxy.delegation()).to.be.equal(delegationControllerProxy.address);
        expect(await delegationControllerProxy.owner()).to.be.equal(owner.address);
        expect(await delegationControllerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await delegationControllerProxy.strategyManager()).to.be.equal(
          strategyManagerProxy.address
        );
        expect(await delegationControllerProxy.minWithdrawalDelay()).to.be.equal(180);
      });
    });
    describe('Delegation', async () => {
      it('add operator will revert when pause', async () => {
        const { owner, delegationControllerProxy } = await loadFixture(deployStrategyManagerFixture);
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
          delegationControllerProxy.registerAsOperator(
            operatorDetailsParams,
            'https://raw.githubusercontent.com/.../main/metadata.json'
          )
        ).to.revertedWith('Pausable: index is paused');
      });
      it('add operator successful', async () => {
        const { owner, unpauser, delegationControllerProxy } = await loadFixture(
          deployStrategyManagerFixture
        );
        await delegationControllerProxy.connect(unpauser).unpause(0);
        const operatorDetailsParams: {
          earningsReceiver: string;
          delegationApprover: string;
          stakerOptOutWindow: number;
        } = {
          earningsReceiver: owner.address,
          delegationApprover: owner.address,
          stakerOptOutWindow: 10,
        };
        await delegationControllerProxy.registerAsOperator(
          operatorDetailsParams,
          'https://raw.githubusercontent.com/.../main/metadata.json'
        );
        expect(await delegationControllerProxy.isOperator(owner.address)).to.be.equal(true);
        expect(await delegationControllerProxy.earningsReceiver(owner.address)).to.be.equal(
          owner.address
        );
        expect(await delegationControllerProxy.delegationApprover(owner.address)).to.be.equal(
          owner.address
        );
        expect(await delegationControllerProxy.stakerOptOutWindow(owner.address)).to.be.equal(10);
      });
    });
    describe('QueueWithdraw', async () => {
      it('queue withdraw will revert when pause', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy, delegationControllerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        // Pause
        await delegationControllerProxy.connect(pauser).pause(3);

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
          delegationControllerProxy.queueWithdrawals([queueWithdrawParams])
        ).to.revertedWith('Pausable: index is paused');
      });
      it('queue withdraw successful', async () => {
        const {
          owner,
          usdt,
          pauserRegistry,
          strategy,
          strategyManagerProxy,
          delegationControllerProxy,
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
        await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
        expect(await strategy.shares(owner.address)).to.be.equal(0);
      });
    });
    describe('CompleteWithdraw', async () => {
      it('complete withdraw revert when not reach withdraw delay', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy, delegationControllerProxy } =
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
        const tx = await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
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
          delegationControllerProxy.completeQueuedWithdrawals(
            [completeQueuedWithdrawalParams],
            [[usdt.address]],
            [0],
            [true]
          )
        ).to.revertedWith(
          'DelegationController._completeQueuedWithdrawal: minWithdrawalDelay period has not yet passed'
        );
      });
      it('complete withdraw revert when pause', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy, delegationControllerProxy } =
          await loadFixture(deployStrategyManagerFixture);

        // Pause
        await delegationControllerProxy.connect(pauser).pause(5);

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
        const tx = await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
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
          delegationControllerProxy.completeQueuedWithdrawals(
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
          delegationControllerProxy,
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
        const tx = await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
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
        await delegationControllerProxy.completeQueuedWithdrawals(
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
