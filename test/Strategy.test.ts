import { config, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import { BigNumber } from 'ethers';
import { waitForTx } from '../helpers/utilities/tx';

describe('Strategy', () => {
  const MAX_PER_DEPOSIT = parseUnits('1000', 18);
  const MAX_TOTAL_DEPOSIT = parseUnits('2000', 18);
  async function deployStrategyFixture() {
    const [owner, pauser, unpauser, strategyManager, alice, ...addrs] = await ethers.getSigners();
    const MintableERC20 = await ethers.getContractFactory('MintableERC20');
    const usdt = await MintableERC20.deploy('USDT', 'USDT');
    const PauserRegistry = await ethers.getContractFactory('PauserRegistry');
    const pauserRegistry = await PauserRegistry.deploy([pauser.address], unpauser.address);
    const StrategyBaseTVLLimits = await ethers.getContractFactory('StrategyBaseTVLLimits');
    const strategy = await upgrades.deployProxy(
      StrategyBaseTVLLimits,
      [MAX_PER_DEPOSIT, MAX_TOTAL_DEPOSIT, usdt.address, pauserRegistry.address],
      { constructorArgs: [strategyManager.address] }
    );
    return {
      owner,
      alice,
      pauser,
      unpauser,
      strategyManager,
      usdt,
      pauserRegistry,
      strategy,
    };
  }

  describe('Strategy', () => {
    describe('Initialize', async () => {
      it('check initialized roles', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );
        expect(await strategy.strategyManager()).to.be.equal(strategyManager.address);
        expect((await strategy.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await strategy.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await strategy.underlyingToken()).to.be.equal(usdt.address);
        expect(await strategy.pauserRegistry()).to.be.equal(pauserRegistry.address);
      });
    });
    describe('Deposit', async () => {
      it('deposit will revert when not strategy manaager', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );

        await expect(strategy.deposit(usdt.address, parseUnits('1001', 18))).to.revertedWith(
          'StrategyBase.onlyStrategyManager'
        );
      });
      it('deposit will revert when greater than per deposit', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );

        await expect(
          strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1001', 18))
        ).to.revertedWith('StrategyBaseTVLLimits: max per deposit exceeded');
      });
      it('deposit successful', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );
        // Mock StrategyManager depositIntoStrategy
        await usdt.transfer(strategy.address, parseUnits('1000'));
        await strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18));
        expect(await strategy.totalShares()).to.be.equal(parseUnits('1000', 18));
      });
      it('deposit will revert when greater than max deposit', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );

        await usdt.transfer(strategy.address, parseUnits('1000'));
        await strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18));
        await usdt.transfer(strategy.address, parseUnits('1000'));
        await strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18));
        await usdt.transfer(strategy.address, parseUnits('1000'));

        await expect(
          strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18))
        ).to.revertedWith('StrategyBaseTVLLimits: max deposits exceeded');
      });
    });
    describe('Pause', async () => {
      it('pause will revert when not pauser', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );
        await expect(strategy.pause(1)).to.revertedWith('msg.sender is not permissioned as pauser');
      });
      it('deposit will revert when pause', async () => {
        const { owner, usdt, pauser, pauserRegistry, strategy, strategyManager } =
          await loadFixture(deployStrategyFixture);

        await strategy.connect(pauser).pause(1);

        await expect(strategy.deposit(usdt.address, parseUnits('1000', 18))).to.revertedWith(
          'Pausable: index is paused'
        );
      });
      it('deposit will successful when unpause', async () => {
        const { owner, usdt, pauser, unpauser, pauserRegistry, strategy, strategyManager } =
          await loadFixture(deployStrategyFixture);

        await strategy.connect(pauser).pause(1);
        await expect(strategy.deposit(usdt.address, parseUnits('1000', 18))).to.revertedWith(
          'Pausable: index is paused'
        );
        await strategy.connect(unpauser).unpause(0);

        await usdt.transfer(strategy.address, parseUnits('1000'));
        await strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18));
        expect(await strategy.totalShares()).to.be.equal(parseUnits('1000', 18));
      });
    });
    describe('Withdraw', async () => {
      it('withdraw successful', async () => {
        const { owner, usdt, strategy, strategyManager } = await loadFixture(deployStrategyFixture);
        await usdt.transfer(strategy.address, parseUnits('1000'));
        await strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18));
        expect(await strategy.totalShares()).to.be.equal(parseUnits('1000', 18));
        expect(await usdt.balanceOf(strategyManager.address)).to.be.equal(0);
        await strategy
          .connect(strategyManager)
          .withdraw(strategyManager.address, usdt.address, parseUnits('1000', 18));
        expect(await usdt.balanceOf(strategyManager.address)).to.be.equal(parseUnits('1000', 18));
      });
      it('withdraw will revert when pause', async () => {
        const { owner, usdt, pauser, strategy, strategyManager } = await loadFixture(
          deployStrategyFixture
        );
        await usdt.transfer(strategy.address, parseUnits('1000'));
        await strategy.connect(strategyManager).deposit(usdt.address, parseUnits('1000', 18));
        expect(await strategy.totalShares()).to.be.equal(parseUnits('1000', 18));

        await strategy.connect(pauser).pause(2);

        await expect(
          strategy
            .connect(strategyManager)
            .withdraw(strategyManager.address, usdt.address, parseUnits('1000', 18))
        ).to.revertedWith('Pausable: index is paused');
      });
    });
  });
});
