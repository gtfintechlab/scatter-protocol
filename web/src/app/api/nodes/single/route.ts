import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { getNodeById } from "server/mongodb/actions/ProtocolNode";

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async (req) => {
            const nodeId = req.nextUrl.searchParams.get("nodeId") as string;
            const node = await getNodeById(nodeId);
            return node
        },
    },
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;