import { internalRequest } from "@/utils/requests";
import { HttpMethod, Workspace } from "@/utils/types";
import { urls } from "@/utils/urls";

const workspaceUrl = urls.baseUrl + urls.api.workspaces.general;

export const getWorkspaces = async () => {
    return internalRequest<Workspace[]>({
        url: workspaceUrl,
        method: HttpMethod.GET,
    })
}

export const userCreateWorkspace = async (workspaceName: string) => {
    return internalRequest<Workspace>({
        url: workspaceUrl,
        method: HttpMethod.POST,
        body: {
            workspaceName
        }
    })
}

export const userDeleteWorkspace = async (workspaceId: string) => {
    return internalRequest<Workspace>({
        url: workspaceUrl,
        method: HttpMethod.DELETE,
        body: {
            workspaceId
        }
    })
}