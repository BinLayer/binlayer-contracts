import { task } from 'hardhat/config';
import { eNetwork, FORK } from '../../helpers';
import {
  POOL_PROXY_ID,
  LISTA_GATEWAY_ID,
  POOL_CONTROLLER_PROXY_ID,
  DELEGATION_CONTROLLER_PROXY_ID,
  PROXY_ADMIN_ID,
  AVS_DIRECTORY_IMPL_ID,
  AVS_DIRECTORY_PROXY_ID,
  PAUSER_REGISTRY_ID,
} from '../../helpers/deploy-ids';
import { Configs } from '../../helpers/config';
import { getParamPerNetwork } from '../../helpers/config-helpers';
import { ethers } from 'ethers';

task(`deploy-avs-directory`, `Deploys the AVSDirectory contract`).setAction(async (_, hre) => {
  if (!hre.network.config.chainId) {
    throw new Error('INVALID_CHAIN_ID');
  }

  await hre.run('compile');

  console.log(`\n- AVSDirectory deployment`);

  const network = (FORK ? FORK : hre.network.name) as eNetwork;
  const owner = getParamPerNetwork(Configs.Owner, network);

  const { deployer } = await hre.getNamedAccounts();
  const { address: pauserRegistryAddress } = await hre.deployments.get(PAUSER_REGISTRY_ID);
  const { address: delegationControllerProxyAddress } = await hre.deployments.get(
    DELEGATION_CONTROLLER_PROXY_ID
  );
  const { address: proxyAdminAddress } = await hre.deployments.get(PROXY_ADMIN_ID);

  const AVSDirectoryImplArtifact = await hre.deployments.deploy(AVS_DIRECTORY_IMPL_ID, {
    from: deployer,
    contract: 'AVSDirectory',
    args: [delegationControllerProxyAddress],
  });

  const ifaceEjectionManager = new ethers.utils.Interface(AVSDirectoryImplArtifact.abi);

  const AVSDirectoryProxyArtifact = await hre.deployments.deploy(AVS_DIRECTORY_PROXY_ID, {
    from: deployer,
    contract:
      '@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy',
    args: [
      AVSDirectoryImplArtifact.address,
      proxyAdminAddress,
      ifaceEjectionManager.encodeFunctionData('initialize', [owner, pauserRegistryAddress, 0]),
    ],
  });
  console.log(`[Deployment][INFO] AVSDirectory deployed ${AVSDirectoryProxyArtifact.address}`);
});
