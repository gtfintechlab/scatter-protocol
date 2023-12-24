import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType, ProtocolNode } from "@/utils/types"
import { createProtocolNode, deleteProtocolNode, getNodesByWorkspaceId, updateProtocolNode } from "server/mongodb/actions/ProtocolNode";

interface DeleteNodeRequest {
    nodeId: string;
}

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async (req) => {
            const workspaceId = req.nextUrl.searchParams.get("workspaceId") as string;
            const nodes = await getNodesByWorkspaceId(workspaceId);
            return nodes
        },
    },
    POST: {
        config: {},
        handler: async (req) => {
            const body: ProtocolNode | Partial<ProtocolNode> = await req.json()
            const node = await createProtocolNode(body);
            return node;
        },
    },
    DELETE: {
        config: {},
        handler: async (req) => {
            const body: DeleteNodeRequest = await req.json()
            const node = await deleteProtocolNode(body.nodeId);
            return node;
        },
    },
    PATCH: {
        config: {},
        handler: async (req) => {
            const body: Partial<ProtocolNode> = await req.json()
            const node = await updateProtocolNode(body);
            return node;
        },
    }

});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;