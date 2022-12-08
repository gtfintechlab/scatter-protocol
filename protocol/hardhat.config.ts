/** @type import('hardhat/config').HardhatUserConfig */
import * as dotenv from "dotenv";
dotenv.config();
import crypto from "crypto";
import "@nomiclabs/hardhat-ethers";
import "@openzeppelin/hardhat-upgrades";
import "@nomiclabs/hardhat-waffle";

const PRIVATE_KEY = process.env.PRIVATE_KEY
  ? process.env.PRIVATE_KEY
  : crypto.randomBytes(32).toString("hex");

module.exports = {
  solidity: "0.8.17",
  networks: {
    goerli: {
      url: process.env.ETHEREUM_GOERLI_NODE,
      accounts: [`${PRIVATE_KEY}`],
    },
  },
};
