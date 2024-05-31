import { config, ethers } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import { BigNumber } from 'ethers';
import { waitForTx } from '../helpers/utilities/tx';

describe('PauserRegistry', () => {
  async function deployPauserRegistryFixture() {
    const [owner, pauser, unpauser, alice, ...addrs] = await ethers.getSigners();
    const PauserRegistry = await ethers.getContractFactory('PauserRegistry');
    const pauserRegistry = await PauserRegistry.deploy([pauser.address], unpauser.address);
    return {
      owner,
      alice,
      pauser,
      unpauser,
      pauserRegistry,
    };
  }

  describe('PauserRegistry', () => {
    describe('Initialize', async () => {
      it('check initialized roles', async () => {
        const { owner, pauser, unpauser, alice, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        expect(await pauserRegistry.isPauser(pauser.address)).to.be.equal(true);
        expect(await pauserRegistry.unpauser()).to.be.equal(unpauser.address);
      });
    });
    describe('Pauser', async () => {
      it('set pauser will revert when not unpauser', async () => {
        const { owner, pauser, alice, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        await expect(pauserRegistry.setIsPauser(alice.address, true)).to.revertedWith(
          'msg.sender is not permissioned as unpauser'
        );
      });
      it('set pauser will revert when pauser is zero', async () => {
        const { owner, pauser, unpauser, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        await expect(
          pauserRegistry.connect(unpauser).setIsPauser(ZERO_ADDRESS, true)
        ).to.revertedWith('PauserRegistry._setPauser: zero address input');
      });
      it('set pauser successful', async () => {
        const { owner, unpauser, alice, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        expect(await pauserRegistry.isPauser(alice.address)).to.be.equal(false);
        await pauserRegistry.connect(unpauser).setIsPauser(alice.address, true);
        expect(await pauserRegistry.isPauser(alice.address)).to.be.equal(true);
      });
      it('unset pauser successful', async () => {
        const { owner, pauser, unpauser, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        expect(await pauserRegistry.isPauser(pauser.address)).to.be.equal(true);
        await pauserRegistry.connect(unpauser).setIsPauser(pauser.address, false);
        expect(await pauserRegistry.isPauser(pauser.address)).to.be.equal(false);
      });
    });
    describe('Unpauser', async () => {
      it('set unpauser will revert when not unpauser', async () => {
        const { owner, pauser, alice, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        await expect(pauserRegistry.setUnpauser(alice.address)).to.revertedWith(
          'msg.sender is not permissioned as unpauser'
        );
      });
      it('set unpauser will revert when unpauser is zero', async () => {
        const { owner, pauser, unpauser, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        await expect(pauserRegistry.connect(unpauser).setUnpauser(ZERO_ADDRESS)).to.revertedWith(
          'PauserRegistry._setUnpauser: zero address input'
        );
      });
      it('set unpauser successful', async () => {
        const { owner, unpauser, alice, pauserRegistry } = await loadFixture(
          deployPauserRegistryFixture
        );
        expect(await pauserRegistry.unpauser()).to.be.equal(unpauser.address);
        await pauserRegistry.connect(unpauser).setUnpauser(alice.address);
        expect(await pauserRegistry.unpauser()).to.be.equal(alice.address);
      });
    });
  });
});
