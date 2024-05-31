import { task } from 'hardhat/config';
import { eNetwork } from '../../helpers';

import { spawn } from 'child_process';

async function runCommand(command: string) {
  return new Promise<void>((resolve, reject) => {
    const process = spawn(command, [], { shell: true });

    process.stdout.on('data', (data: any) => {
      console.log(data.toString());
    });

    process.stderr.on('data', (data: any) => {
      console.error(data.toString());
    });

    process.on('close', (code) => {
      if (code === 0) {
        resolve();
      } else {
        reject(`Command failed with exit code ${code}`);
      }
    });
  });
}

task(`verify-contracts`)
  .addOptionalParam('contractName', 'The contract name')
  .setAction(async ({ contractName }, { deployments, getNamedAccounts, ...hre }) => {
    const network = hre.network.name as eNetwork;

    const contracts = await deployments.all();

    for (let [name, contract] of Object.entries(contracts)) {
      if (contractName && contractName != name) {
        continue;
      }
      const argStr = contract.args?.map((i) => `"${i}"`).join(' ');
      const cmdStr = `npx hardhat verify --network ${network} ${contract.address} ${argStr}`;
      console.log(name, cmdStr);
      try {
        await runCommand(cmdStr);
        console.log('Command ran successfully!');
      } catch (error) {
        console.error(error);
      }
    }
  });
