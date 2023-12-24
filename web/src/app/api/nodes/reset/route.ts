import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { resetNodeTokenSupply } from "server/mongodb/actions/ProtocolNode";

interface ResetNodeRequest {
    workspaceId: string;
}
const route: APIWrapperType = APIWrapper({
    POST: {
        config: {
        },
        handler: async (req) => {
            const body: ResetNodeRequest = await req.json();
            await resetNodeTokenSupply(body.workspaceId);

        },
    },
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;