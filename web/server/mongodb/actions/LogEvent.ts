import dbConnect from "server/utils/dbConnect"
import { LogEventModel } from "../models/LogEvent"

export const deleteAllLogEventsByWorkspace = async (workspaceId: string) => {
    await dbConnect()
    const events = await LogEventModel.deleteMany({ workspaceId: workspaceId })
    return events;
}