import { internalRequest } from "@/utils/requests";
import { HttpMethod, ProtocolNode } from "@/utils/types";
import { urls } from "@/utils/urls";

const nodeUrl = urls.baseUrl + urls.api.nodes.general;
const singleNodeUrl = urls.baseUrl + urls.api.nodes.single;
const resetUrl = urls.baseUrl + urls.api.nodes.reset;

export const getNodesForWorkspace = async (workspaceId: string) => {
    return internalRequest<ProtocolNode[]>({
        url: nodeUrl,
        method: HttpMethod.GET,
        queryParams: {
            workspaceId
        }
    })
}

export const createProtocolNode = async (nodeInfo: ProtocolNode | Partial<ProtocolNode>) => {
    return internalRequest<ProtocolNode>({
        url: nodeUrl,
        method: HttpMethod.POST,
        body: {
            ...nodeInfo
        }
    })
}

export const deleteProtocolNode = async (nodeId: string) => {
    return internalRequest<ProtocolNode>({
        url: nodeUrl,
        method: HttpMethod.DELETE,
        body: {
            nodeId
        }
    })
}

export const updateProtocolNode = async (nodeInfo: ProtocolNode | Partial<ProtocolNode>) => {
    return internalRequest<ProtocolNode>({
        url: nodeUrl,
        method: HttpMethod.PATCH,
        body: {
            ...nodeInfo
        }
    })
}

export const getSingleNodeById = async (nodeId: string) => {
    return internalRequest<ProtocolNode>({
        url: singleNodeUrl,
        method: HttpMethod.GET,
        queryParams: {
            nodeId
        }
    })
}

export const resetProtocolNodesByWorkspace = async (workspaceId: string) => {
    return internalRequest({
        url: resetUrl,
        method: HttpMethod.POST,
        body: {
            workspaceId
        }
    })
}