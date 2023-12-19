import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { getNodesByWorkspaceId } from "server/mongodb/actions/ProtocolNode";

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
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;