import hardhat from "hardhat";

async function deployContract() {
    await hardhat.run('compile');
    const scatterTokenContract = await (hardhat as any).ethers.getContractFactory("ScatterToken");
    const scatterToken = await scatterTokenContract.deploy();
    return {
        "address": scatterToken.address,
        "contract": scatterToken
    }
}

deployContract().then((output) => {
    console.log(output.address);
});
