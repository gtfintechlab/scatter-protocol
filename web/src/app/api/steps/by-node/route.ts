import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType, Step } from "@/utils/types"
import { findStepsByNodeId } from "server/mongodb/actions/Step";

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async (req) => {
            const nodeId = req.nextUrl.searchParams.get("nodeId")
            const steps = await findStepsByNodeId(nodeId as string)
            return steps
        }
    }
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;