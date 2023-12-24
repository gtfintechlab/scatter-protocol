import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType } from "@/utils/types"
import { deleteAllLogEventsByWorkspace, getEventsByWorkspace } from "server/mongodb/actions/LogEvent";

interface DeleteEventsByWorkspaceIdRequest {
    workspaceId: string
}
const route: APIWrapperType = APIWrapper({
    GET: {
        config: {},
        handler: async (req) => {
            const workspaceId = req.nextUrl.searchParams.get("workspaceId") as string;
            const events = await getEventsByWorkspace(workspaceId);
            return events
        }
    },
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