import { Contract, ContractFactory } from "ethers";
import hardhat from "hardhat";
import "@nomiclabs/hardhat-ethers";

async function deployContract() {
  await hardhat.run("compile");
  const scatterTokenContract: ContractFactory =
    (await hardhat.ethers.getContractFactory(
      "ScatterToken"
    )) as ContractFactory;
  const scatterToken: Contract = (await scatterTokenContract.deploy(
    BigInt(100000000000 * Math.pow(10, 18)),
    100
  )) as Contract;
  return {
    address: scatterToken.address,
    contract: scatterToken,
  };
}

deployContract().then((output) => {
  console.log(output.address);
});
