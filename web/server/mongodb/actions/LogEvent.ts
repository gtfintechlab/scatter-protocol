import dbConnect from "server/utils/dbConnect"
import { LogEventModel } from "../models/LogEvent"
import { Types } from "mongoose"

export const deleteAllLogEventsByWorkspace = async (workspaceId: string) => {
    await dbConnect()
    const events = await LogEventModel.deleteMany({ workspaceId: workspaceId })
    return events;
}

export const getEventsByWorkspace = async (workspaceId: string | Types.ObjectId) => {
    await dbConnect();
    const events = await LogEventModel.find({ workspaceId: workspaceId })

    return events
}