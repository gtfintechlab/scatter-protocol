import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { createWorkspace, deleteWorkspace, getAllWorkspaces } from "server/mongodb/actions/Workspace";

interface CreateWorkspaceRequest {
    workspaceName: string;
}


interface DeleteWorkspaceRequest {
    workspaceId: string;
}
const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async (req) => {
            const workspaces = await getAllWorkspaces();
            return workspaces
        },
    },
    POST: {
        config: {},
        handler: async (req) => {
            const body: CreateWorkspaceRequest = await req.json();
            const workspace = await createWorkspace(body.workspaceName)
            return workspace
        },
    },
    DELETE: {
        config: {},
        handler: async (req) => {
            const body: DeleteWorkspaceRequest = await req.json();
            const workspace = await deleteWorkspace(body.workspaceId)
            return workspace
        },
    }

});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;