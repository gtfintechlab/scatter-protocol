import { internalRequest } from "@/utils/requests";
import { HttpMethod, ProtocolNode } from "@/utils/types";
import { urls } from "@/utils/urls";

const nodeUrl = urls.baseUrl + urls.api.nodes.general;

export const getNodesForWorkspace = async (workspaceId: string) => {
    return internalRequest<ProtocolNode[]>({
        url: nodeUrl,
        method: HttpMethod.GET,
        queryParams: {
            workspaceId
        }
    })
}