import { config, deployments, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import TransparentUpgradeableProxy from './build/TransparentUpgradeableProxy.json';
import ProxyAdminArtifact from './build/ProxyAdmin.json';
import DelegationControllerABI from '../abis/DelegationController.json';
import PoolControllerABI from '../abis/PoolController.json';
import { BigNumberish } from 'ethers';
import { increaseTime } from '../helpers';

describe('DelegationController', () => {
  const MAX_PER_DEPOSIT = parseUnits('1000', 18);
  const MAX_TOTAL_DEPOSIT = parseUnits('2000', 18);
  async function deployPoolControllerFixture() {
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
    const PoolControllerProxy = await ethers.getContractFactory(
      TransparentUpgradeableProxy.abi,
      TransparentUpgradeableProxy.bytecode
    );
    let poolControllerProxy = await PoolControllerProxy.deploy(
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
      poolControllerProxy.address,
      slasherProxy.address
    );
    const PoolController = await ethers.getContractFactory('PoolController');
    const poolController = await PoolController.deploy(
      delegationControllerProxy.address,
      slasherProxy.address
    );

    const PoolBaseTVLLimits = await ethers.getContractFactory('PoolBaseTVLLimits');
    const pool = await upgrades.deployProxy(
      PoolBaseTVLLimits,
      [MAX_PER_DEPOSIT, MAX_TOTAL_DEPOSIT, usdt.address, pauserRegistry.address],
      { constructorArgs: [poolControllerProxy.address] }
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
        [pool.address],
        [360],
      ])
    );
    const ifacePoolController = new ethers.utils.Interface(PoolControllerABI);
    await proxyAdmin.upgradeAndCall(
      poolControllerProxy.address,
      poolController.address,
      ifacePoolController.encodeFunctionData('initialize', [
        owner.address,
        owner.address,
        pauserRegistry.address,
        0,
      ])
    );
    poolControllerProxy = await ethers.getContractAt(
      'PoolController',
      poolControllerProxy.address
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
      pool,
      poolControllerProxy,
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
          pool,
          poolControllerProxy,
          delegationControllerProxy,
        } = await loadFixture(deployPoolControllerFixture);
        expect(await pool.poolController()).to.be.equal(poolControllerProxy.address);
        expect((await pool.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await pool.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await pool.underlyingToken()).to.be.equal(usdt.address);
        expect(await pool.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await poolControllerProxy.owner()).to.be.equal(owner.address);
        expect(await poolControllerProxy.poolWhitelister()).to.be.equal(owner.address);
        expect(await poolControllerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await poolControllerProxy.delegation()).to.be.equal(delegationControllerProxy.address);
        expect(await delegationControllerProxy.owner()).to.be.equal(owner.address);
        expect(await delegationControllerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await delegationControllerProxy.poolController()).to.be.equal(
          poolControllerProxy.address
        );
        expect(await delegationControllerProxy.minWithdrawalDelay()).to.be.equal(180);
      });
    });
    describe('Delegation', async () => {
      it('add operator will revert when pause', async () => {
        const { owner, delegationControllerProxy } = await loadFixture(deployPoolControllerFixture);
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
          deployPoolControllerFixture
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
        const { owner, usdt, pauser, pool, poolControllerProxy, delegationControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        // Pause
        await delegationControllerProxy.connect(pauser).pause(3);

        expect(await usdt.balanceOf(pool.address)).to.be.equal(0);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await poolControllerProxy.depositIntoPool(
          pool.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          pools: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          pools: [pool.address],
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
          pool,
          poolControllerProxy,
          delegationControllerProxy,
        } = await loadFixture(deployPoolControllerFixture);
        expect(await usdt.balanceOf(pool.address)).to.be.equal(0);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await poolControllerProxy.depositIntoPool(
          pool.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          pools: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          pools: [pool.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await pool.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
        expect(await pool.shares(owner.address)).to.be.equal(0);
      });
    });
    describe('CompleteWithdraw', async () => {
      it('complete withdraw revert when not reach withdraw delay', async () => {
        const { owner, usdt, pauser, pool, poolControllerProxy, delegationControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        expect(await usdt.balanceOf(pool.address)).to.be.equal(0);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await poolControllerProxy.depositIntoPool(
          pool.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          pools: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          pools: [pool.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await pool.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        const tx = await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
        const receipt = await tx.wait();
        const eventResp = receipt.events[0].args.withdrawal;
        expect(await pool.shares(owner.address)).to.be.equal(0);

        const completeQueuedWithdrawalParams: {
          staker: string;
          delegatedTo: string;
          withdrawer: string;
          nonce: string;
          startTimestamp: string;
          pools: string[];
          shares: BigNumberish[];
        } = {
          staker: eventResp.staker,
          delegatedTo: eventResp.delegatedTo,
          withdrawer: eventResp.withdrawer,
          nonce: eventResp.nonce,
          startTimestamp: eventResp.startTimestamp,
          pools: eventResp.pools,
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
        const { owner, usdt, pauser, pool, poolControllerProxy, delegationControllerProxy } =
          await loadFixture(deployPoolControllerFixture);

        // Pause
        await delegationControllerProxy.connect(pauser).pause(5);

        expect(await usdt.balanceOf(pool.address)).to.be.equal(0);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await poolControllerProxy.depositIntoPool(
          pool.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          pools: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          pools: [pool.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await pool.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        const tx = await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
        const receipt = await tx.wait();
        const eventResp = receipt.events[0].args.withdrawal;
        expect(await pool.shares(owner.address)).to.be.equal(0);

        const completeQueuedWithdrawalParams: {
          staker: string;
          delegatedTo: string;
          withdrawer: string;
          nonce: string;
          startTimestamp: string;
          pools: string[];
          shares: BigNumberish[];
        } = {
          staker: eventResp.staker,
          delegatedTo: eventResp.delegatedTo,
          withdrawer: eventResp.withdrawer,
          nonce: eventResp.nonce,
          startTimestamp: eventResp.startTimestamp,
          pools: eventResp.pools,
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
          pool,
          poolControllerProxy,
          delegationControllerProxy,
        } = await loadFixture(deployPoolControllerFixture);
        expect(await usdt.balanceOf(pool.address)).to.be.equal(0);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await poolControllerProxy.depositIntoPool(
          pool.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(parseUnits('1000', 18));
        const queueWithdrawParams: {
          pools: string[];
          shares: BigNumberish[];
          withdrawer: string;
        } = {
          pools: [pool.address],
          shares: [parseUnits('1000', 18)],
          withdrawer: owner.address,
        };
        expect(await pool.shares(owner.address)).to.be.equal(parseUnits('1000', 18));
        const tx = await delegationControllerProxy.queueWithdrawals([queueWithdrawParams]);
        const receipt = await tx.wait();
        const eventResp = receipt.events[0].args.withdrawal;
        expect(await pool.shares(owner.address)).to.be.equal(0);

        const completeQueuedWithdrawalParams: {
          staker: string;
          delegatedTo: string;
          withdrawer: string;
          nonce: string;
          startTimestamp: string;
          pools: string[];
          shares: BigNumberish[];
        } = {
          staker: eventResp.staker,
          delegatedTo: eventResp.delegatedTo,
          withdrawer: eventResp.withdrawer,
          nonce: eventResp.nonce,
          startTimestamp: eventResp.startTimestamp,
          pools: eventResp.pools,
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
