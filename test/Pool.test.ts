import { config, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import { BigNumber } from 'ethers';
import { waitForTx } from '../helpers/utilities/tx';

describe('Pool', () => {
  const MAX_PER_DEPOSIT = parseUnits('1000', 18);
  const MAX_TOTAL_DEPOSIT = parseUnits('2000', 18);
  async function deployPoolFixture() {
    const [owner, pauser, unpauser, poolController, alice, ...addrs] = await ethers.getSigners();
    const MintableERC20 = await ethers.getContractFactory('MintableERC20');
    const usdt = await MintableERC20.deploy('USDT', 'USDT');
    const PauserRegistry = await ethers.getContractFactory('PauserRegistry');
    const pauserRegistry = await PauserRegistry.deploy([pauser.address], unpauser.address);
    const PoolBaseTVLLimits = await ethers.getContractFactory('PoolBaseTVLLimits');
    const pool = await upgrades.deployProxy(
      PoolBaseTVLLimits,
      [MAX_PER_DEPOSIT, MAX_TOTAL_DEPOSIT, usdt.address, pauserRegistry.address],
      { constructorArgs: [poolController.address] }
    );
    return {
      owner,
      alice,
      pauser,
      unpauser,
      poolController,
      usdt,
      pauserRegistry,
      pool,
    };
  }

  describe('Pool', () => {
    describe('Initialize', async () => {
      it('check initialized roles', async () => {
        const { owner, usdt, pauserRegistry, pool, poolController } = await loadFixture(
          deployPoolFixture
        );
        expect(await pool.poolController()).to.be.equal(poolController.address);
        expect((await pool.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await pool.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await pool.underlyingToken()).to.be.equal(usdt.address);
        expect(await pool.pauserRegistry()).to.be.equal(pauserRegistry.address);
      });
    });
    describe('Deposit', async () => {
      it('deposit will revert when not pool manaager', async () => {
        const { owner, usdt, pauserRegistry, pool, poolController } = await loadFixture(
          deployPoolFixture
        );

        await expect(pool.deposit(usdt.address, parseUnits('1001', 18))).to.revertedWith(
          'PoolBase.onlyPoolController'
        );
      });
      it('deposit will revert when greater than per deposit', async () => {
        const { owner, usdt, pauserRegistry, pool, poolController } = await loadFixture(
          deployPoolFixture
        );

        await expect(
          pool.connect(poolController).deposit(usdt.address, parseUnits('1001', 18))
        ).to.revertedWith('PoolBaseTVLLimits: max per deposit exceeded');
      });
      it('deposit successful', async () => {
        const { owner, usdt, pauserRegistry, pool, poolController } = await loadFixture(
          deployPoolFixture
        );
        // Mock PoolController.sol depositIntoPool
        await usdt.transfer(pool.address, parseUnits('1000'));
        await pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18));
        expect(await pool.totalShares()).to.be.equal(parseUnits('1000', 18));
      });
      it('deposit will revert when greater than max deposit', async () => {
        const { owner, usdt, pauserRegistry, pool, poolController } = await loadFixture(
          deployPoolFixture
        );

        await usdt.transfer(pool.address, parseUnits('1000'));
        await pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18));
        await usdt.transfer(pool.address, parseUnits('1000'));
        await pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18));
        await usdt.transfer(pool.address, parseUnits('1000'));

        await expect(
          pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18))
        ).to.revertedWith('PoolBaseTVLLimits: max deposits exceeded');
      });
    });
    describe('Pause', async () => {
      it('pause will revert when not pauser', async () => {
        const { owner, usdt, pauserRegistry, pool, poolController } = await loadFixture(
          deployPoolFixture
        );
        await expect(pool.pause(1)).to.revertedWith('msg.sender is not permissioned as pauser');
      });
      it('deposit will revert when pause', async () => {
        const { owner, usdt, pauser, pauserRegistry, pool, poolController } =
          await loadFixture(deployPoolFixture);

        await pool.connect(pauser).pause(1);

        await expect(pool.deposit(usdt.address, parseUnits('1000', 18))).to.revertedWith(
          'Pausable: index is paused'
        );
      });
      it('deposit will successful when unpause', async () => {
        const { owner, usdt, pauser, unpauser, pauserRegistry, pool, poolController } =
          await loadFixture(deployPoolFixture);

        await pool.connect(pauser).pause(1);
        await expect(pool.deposit(usdt.address, parseUnits('1000', 18))).to.revertedWith(
          'Pausable: index is paused'
        );
        await pool.connect(unpauser).unpause(0);

        await usdt.transfer(pool.address, parseUnits('1000'));
        await pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18));
        expect(await pool.totalShares()).to.be.equal(parseUnits('1000', 18));
      });
    });
    describe('Withdraw', async () => {
      it('withdraw successful', async () => {
        const { owner, usdt, pool, poolController } = await loadFixture(deployPoolFixture);
        await usdt.transfer(pool.address, parseUnits('1000'));
        await pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18));
        expect(await pool.totalShares()).to.be.equal(parseUnits('1000', 18));
        expect(await usdt.balanceOf(poolController.address)).to.be.equal(0);
        await pool
          .connect(poolController)
          .withdraw(poolController.address, usdt.address, parseUnits('1000', 18));
        expect(await usdt.balanceOf(poolController.address)).to.be.equal(parseUnits('1000', 18));
      });
      it('withdraw will revert when pause', async () => {
        const { owner, usdt, pauser, pool, poolController } = await loadFixture(
          deployPoolFixture
        );
        await usdt.transfer(pool.address, parseUnits('1000'));
        await pool.connect(poolController).deposit(usdt.address, parseUnits('1000', 18));
        expect(await pool.totalShares()).to.be.equal(parseUnits('1000', 18));

        await pool.connect(pauser).pause(2);

        await expect(
          pool
            .connect(poolController)
            .withdraw(poolController.address, usdt.address, parseUnits('1000', 18))
        ).to.revertedWith('Pausable: index is paused');
      });
    });
  });
});
