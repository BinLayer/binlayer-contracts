import { config, deployments, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import TransparentUpgradeableProxy from './build/TransparentUpgradeableProxy.json';
import ProxyAdminArtifact from './build/ProxyAdmin.json';
import DelegationControllerABI from '../abis/DelegationController.json';
import PoolControllerABI from '../abis/PoolController.json';

describe('PoolController', () => {
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

  describe('PoolController', () => {
    describe('Initialize', async () => {
      it('check initialized', async () => {
        const { owner, usdt, pauserRegistry, pool, poolControllerProxy } = await loadFixture(
          deployPoolControllerFixture
        );
        expect(await pool.poolController()).to.be.equal(poolControllerProxy.address);
        expect((await pool.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await pool.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await pool.underlyingToken()).to.be.equal(usdt.address);
        expect(await pool.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await poolControllerProxy.owner()).to.be.equal(owner.address);
        expect(await poolControllerProxy.poolWhitelister()).to.be.equal(owner.address);
        expect(await poolControllerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
      });
    });
    describe('Whitelister', async () => {
      it('config third party forbidden will revert when not whitelister', async () => {
        const { owner, alice, usdt, pauserRegistry, pool, poolControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        await expect(
          poolControllerProxy
            .connect(alice)
            .setThirdPartyTransfersForbidden(pool.address, false)
        ).to.revertedWith('PoolController.onlyPoolWhitelister: not the poolWhitelister');
      });
      it('add pool whitelist will revert when not whitelister', async () => {
        const { owner, alice, usdt, pauserRegistry, pool, poolControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        await expect(
          poolControllerProxy
            .connect(alice)
            .addPoolsToDepositWhitelist([pool.address], [false])
        ).to.revertedWith('PoolController.onlyPoolWhitelister: not the poolWhitelister');
      });
      it('remove pool whitelist will revert when not whitelister', async () => {
        const { owner, alice, usdt, pauserRegistry, pool, poolControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await expect(
          poolControllerProxy
            .connect(alice)
            .removePoolsFromDepositWhitelist([pool.address])
        ).to.revertedWith('PoolController.onlyPoolWhitelister: not the poolWhitelister');
      });
      it('config third party forbidden successful', async () => {
        const { owner, alice, usdt, pauserRegistry, pool, poolControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        expect(
          await poolControllerProxy.thirdPartyTransfersForbidden(pool.address)
        ).to.be.equal(false);
        await poolControllerProxy.setThirdPartyTransfersForbidden(pool.address, true);
        expect(
          await poolControllerProxy.thirdPartyTransfersForbidden(pool.address)
        ).to.be.equal(true);
      });
      it('add pool whitelist successful', async () => {
        const { owner, alice, usdt, pauserRegistry, pool, poolControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        expect(
          await poolControllerProxy.poolIsWhitelistedForDeposit(pool.address)
        ).to.be.equal(false);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        expect(
          await poolControllerProxy.poolIsWhitelistedForDeposit(pool.address)
        ).to.be.equal(true);
      });
      it('remove pool whitelist successful', async () => {
        const { owner, alice, usdt, pauserRegistry, pool, poolControllerProxy } =
          await loadFixture(deployPoolControllerFixture);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        expect(
          await poolControllerProxy.poolIsWhitelistedForDeposit(pool.address)
        ).to.be.equal(true);
        await poolControllerProxy.removePoolsFromDepositWhitelist([pool.address]);
        expect(
          await poolControllerProxy.poolIsWhitelistedForDeposit(pool.address)
        ).to.be.equal(false);
      });
    });
    describe('Deposit', async () => {
      it('deposit will revert when not config whitelist', async () => {
        const { owner, usdt, pauserRegistry, pool, poolControllerProxy } = await loadFixture(
          deployPoolControllerFixture
        );
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await expect(
          poolControllerProxy.depositIntoPool(
            pool.address,
            usdt.address,
            parseUnits('1000', 18)
          )
        ).to.revertedWith(
          'PoolController.onlyPoolsWhitelistedForDeposit: pool not whitelisted'
        );
      });
      it('deposit will revert when pause', async () => {
        const { owner, usdt, pauser, pool, poolControllerProxy } = await loadFixture(
          deployPoolControllerFixture
        );
        await poolControllerProxy.connect(pauser).pause(1);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await expect(
          poolControllerProxy.depositIntoPool(
            pool.address,
            usdt.address,
            parseUnits('1000', 18)
          )
        ).to.revertedWith('Pausable: index is paused');
      });
      it('deposit successful', async () => {
        const { owner, usdt, pauserRegistry, pool, poolControllerProxy } = await loadFixture(
          deployPoolControllerFixture
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(0);
        await poolControllerProxy.addPoolsToDepositWhitelist([pool.address], [false]);
        await usdt.approve(poolControllerProxy.address, MAX_UINT_AMOUNT);
        await poolControllerProxy.depositIntoPool(
          pool.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(pool.address)).to.be.equal(parseUnits('1000', 18));
      });
    });
  });
});
