import { EthereumAccount } from "@/utils/types"
import { EthereumAccountModel } from "../models/EthereumAccount"
import dbConnect from "server/utils/dbConnect"

export const createEthereumAccount = async (ethereumAccount: EthereumAccount | Partial<EthereumAccount>) => {
    await dbConnect();
    const account = await EthereumAccountModel.create(ethereumAccount);
    return account;
}

export const findEthereumAccount = async (mnemonic: string, path: string, index: number) => {
    await dbConnect();
    const account = await EthereumAccountModel.findOne({ mnemonic, path, index });
    return account;
}

export const findAllEthereumAccounts = async (mnemonic: string, path: string) => {
    await dbConnect();
    const accounts = await EthereumAccountModel.find({ mnemonic, path });
    return accounts;
}