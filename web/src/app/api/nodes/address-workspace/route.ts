import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { getNodeByBlockchainAndWorkspace } from "server/mongodb/actions/ProtocolNode";

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async (req) => {
            const blockchainAddress = req.nextUrl.searchParams.get("blockchainAddress") as string;
            const workspaceId = req.nextUrl.searchParams.get("workspaceId") as string;
            const node = await getNodeByBlockchainAndWorkspace(blockchainAddress, workspaceId);
            return node
        },
    },
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;