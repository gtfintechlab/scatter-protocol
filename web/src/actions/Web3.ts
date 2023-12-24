import { internalRequest } from "@/utils/requests";
import { EthereumAccount, HttpMethod } from "@/utils/types";
import { urls } from "@/utils/urls";

const web3Url = urls.baseUrl + urls.api.web3.accounts;

export const getAllWeb3Accounts = async () => {
    return internalRequest<EthereumAccount[]>({
        url: web3Url,
        method: HttpMethod.GET,
    })
}