import { externalRequest } from "@/utils/requests";
import { HttpMethod, ProtocolNode } from "@/utils/types";
import { urls } from "@/utils/urls";

const startNodeUrl = urls.externalUrl + urls.external.startNode;
const stopNodeUrl = urls.externalUrl + urls.external.stopNode;
export const startNode = async (node: ProtocolNode) => {
    return externalRequest({
        url: startNodeUrl,
        method: HttpMethod.POST,
        body: {
            ...node
        }
    })
}

export const stopNode = async (blockchainAddress: string) => {
    return externalRequest({
        url: stopNodeUrl,
        method: HttpMethod.POST,
        body: {
            blockchainAddress
        }
    })
}