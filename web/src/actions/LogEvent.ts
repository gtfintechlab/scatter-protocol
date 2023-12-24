import { internalRequest } from "@/utils/requests"
import { HttpMethod, LogEvent } from "@/utils/types"
import { urls } from "@/utils/urls"

const logsUrl = urls.baseUrl + urls.api.events.logs;

export const deleteLogEvents = async (workspaceId: string) => {
    return internalRequest<LogEvent[]>({
        url: logsUrl,
        method: HttpMethod.DELETE,
        body: {
            workspaceId
        }
    })

}

export const getWorkspaceEvents = async (workspaceId: string) => {
    return internalRequest<LogEvent[]>({
        url: logsUrl,
        method: HttpMethod.GET,
        queryParams: {
            workspaceId
        }
    })

}