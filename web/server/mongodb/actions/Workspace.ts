import dbConnect from "server/utils/dbConnect"
import { WorkspaceModel } from "server/mongodb/models/Workspace"

export const getAllWorkspaces = async () => {
    await dbConnect()
    return await WorkspaceModel.find({})
}

export const createWorkspace = async (workspaceName: string) => {
    await dbConnect();
    return await WorkspaceModel.create({ name: workspaceName });
}

export const deleteWorkspace = async (workspaceId: string) => {
    await dbConnect();
    return await WorkspaceModel.findOneAndDelete({ _id: workspaceId });
} 