import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType, EthereumAccount } from "@/utils/types"
import { ethers } from "ethers"
import { WEB3_MNEMONIC, WEB3_NUM_ACCOUNTS, WEB3_PATH } from "@/utils/constants";
import { createEthereumAccount, findAllEthereumAccounts, findEthereumAccount } from "server/mongodb/actions/EthereumAccount";

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async () => {
            const accounts: EthereumAccount[] = await findAllEthereumAccounts(WEB3_MNEMONIC, WEB3_PATH);
            if (accounts.length === WEB3_NUM_ACCOUNTS) {
                return accounts
            }

            for (let i = 0; i < WEB3_NUM_ACCOUNTS; i++) {
                const account = await findEthereumAccount(WEB3_MNEMONIC, WEB3_PATH, i);
                if (!account) {
                    const wallet = ethers.Wallet.fromMnemonic(WEB3_MNEMONIC, WEB3_PATH + `/${i}`);
                    await createEthereumAccount({ privateKey: wallet.privateKey, address: wallet.address, index: i, path: WEB3_PATH, mnemonic: WEB3_MNEMONIC })
                }
            }

            return await findAllEthereumAccounts(WEB3_MNEMONIC, WEB3_PATH);
        },
    },
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;