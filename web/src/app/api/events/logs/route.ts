import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { deleteAllLogEventsByWorkspace } from "server/mongodb/actions/LogEvent";

interface DeleteEventsByWorkspaceIdRequest {
    workspaceId: string
}
const route: APIWrapperType = APIWrapper({
    DELETE: {
        config: {
        },
        handler: async (req) => {
            const body: DeleteEventsByWorkspaceIdRequest = await req.json()
            const events = deleteAllLogEventsByWorkspace(body.workspaceId)
            return events;
        },
    },
});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;