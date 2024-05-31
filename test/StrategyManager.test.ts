import { config, deployments, ethers, upgrades } from 'hardhat';
import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { MAX_UINT_AMOUNT, ONE_ADDRESS, ONE_HOUR, ZERO_ADDRESS } from '../helpers/constants';
import { parseUnits } from 'ethers/lib/utils';
import TransparentUpgradeableProxy from './build/TransparentUpgradeableProxy.json';
import ProxyAdminArtifact from './build/ProxyAdmin.json';
import DelegationManagerABI from '../abis/DelegationManager.json';
import StrategyManagerABI from '../abis/StrategyManager.json';

describe('StrategyManager', () => {
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

  describe('StrategyManager', () => {
    describe('Initialize', async () => {
      it('check initialized', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManagerProxy } = await loadFixture(
          deployStrategyManagerFixture
        );
        expect(await strategy.strategyManager()).to.be.equal(strategyManagerProxy.address);
        expect((await strategy.getTVLLimits())[0]).to.be.equal(MAX_PER_DEPOSIT);
        expect((await strategy.getTVLLimits())[1]).to.be.equal(MAX_TOTAL_DEPOSIT);
        expect(await strategy.underlyingToken()).to.be.equal(usdt.address);
        expect(await strategy.pauserRegistry()).to.be.equal(pauserRegistry.address);
        expect(await strategyManagerProxy.owner()).to.be.equal(owner.address);
        expect(await strategyManagerProxy.strategyWhitelister()).to.be.equal(owner.address);
        expect(await strategyManagerProxy.pauserRegistry()).to.be.equal(pauserRegistry.address);
      });
    });
    describe('Whitelister', async () => {
      it('config third party forbidden will revert when not whitelister', async () => {
        const { owner, alice, usdt, pauserRegistry, strategy, strategyManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        await expect(
          strategyManagerProxy
            .connect(alice)
            .setThirdPartyTransfersForbidden(strategy.address, false)
        ).to.revertedWith('StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister');
      });
      it('add strategy whitelist will revert when not whitelister', async () => {
        const { owner, alice, usdt, pauserRegistry, strategy, strategyManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        await expect(
          strategyManagerProxy
            .connect(alice)
            .addStrategiesToDepositWhitelist([strategy.address], [false])
        ).to.revertedWith('StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister');
      });
      it('remove strategy whitelist will revert when not whitelister', async () => {
        const { owner, alice, usdt, pauserRegistry, strategy, strategyManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await expect(
          strategyManagerProxy
            .connect(alice)
            .removeStrategiesFromDepositWhitelist([strategy.address])
        ).to.revertedWith('StrategyManager.onlyStrategyWhitelister: not the strategyWhitelister');
      });
      it('config third party forbidden successful', async () => {
        const { owner, alice, usdt, pauserRegistry, strategy, strategyManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        expect(
          await strategyManagerProxy.thirdPartyTransfersForbidden(strategy.address)
        ).to.be.equal(false);
        await strategyManagerProxy.setThirdPartyTransfersForbidden(strategy.address, true);
        expect(
          await strategyManagerProxy.thirdPartyTransfersForbidden(strategy.address)
        ).to.be.equal(true);
      });
      it('add strategy whitelist successful', async () => {
        const { owner, alice, usdt, pauserRegistry, strategy, strategyManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        expect(
          await strategyManagerProxy.strategyIsWhitelistedForDeposit(strategy.address)
        ).to.be.equal(false);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        expect(
          await strategyManagerProxy.strategyIsWhitelistedForDeposit(strategy.address)
        ).to.be.equal(true);
      });
      it('remove strategy whitelist successful', async () => {
        const { owner, alice, usdt, pauserRegistry, strategy, strategyManagerProxy } =
          await loadFixture(deployStrategyManagerFixture);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        expect(
          await strategyManagerProxy.strategyIsWhitelistedForDeposit(strategy.address)
        ).to.be.equal(true);
        await strategyManagerProxy.removeStrategiesFromDepositWhitelist([strategy.address]);
        expect(
          await strategyManagerProxy.strategyIsWhitelistedForDeposit(strategy.address)
        ).to.be.equal(false);
      });
    });
    describe('Deposit', async () => {
      it('deposit will revert when not config whitelist', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManagerProxy } = await loadFixture(
          deployStrategyManagerFixture
        );
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await expect(
          strategyManagerProxy.depositIntoStrategy(
            strategy.address,
            usdt.address,
            parseUnits('1000', 18)
          )
        ).to.revertedWith(
          'StrategyManager.onlyStrategiesWhitelistedForDeposit: strategy not whitelisted'
        );
      });
      it('deposit will revert when pause', async () => {
        const { owner, usdt, pauser, strategy, strategyManagerProxy } = await loadFixture(
          deployStrategyManagerFixture
        );
        await strategyManagerProxy.connect(pauser).pause(1);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await expect(
          strategyManagerProxy.depositIntoStrategy(
            strategy.address,
            usdt.address,
            parseUnits('1000', 18)
          )
        ).to.revertedWith('Pausable: index is paused');
      });
      it('deposit successful', async () => {
        const { owner, usdt, pauserRegistry, strategy, strategyManagerProxy } = await loadFixture(
          deployStrategyManagerFixture
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(0);
        await strategyManagerProxy.addStrategiesToDepositWhitelist([strategy.address], [false]);
        await usdt.approve(strategyManagerProxy.address, MAX_UINT_AMOUNT);
        await strategyManagerProxy.depositIntoStrategy(
          strategy.address,
          usdt.address,
          parseUnits('1000', 18)
        );
        expect(await usdt.balanceOf(strategy.address)).to.be.equal(parseUnits('1000', 18));
      });
    });
  });
});
