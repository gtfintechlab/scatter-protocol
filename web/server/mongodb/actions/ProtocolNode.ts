import { Types } from "mongoose";
import dbConnect from "server/utils/dbConnect";
import { ProtocolNodeModel } from "../models/ProtocolNode";

export const getNodesByWorkspaceId = async (workspaceId: string | Types.ObjectId) => {
    await dbConnect();
    const nodes = await ProtocolNodeModel.find({ workspaceId })
    return nodes;
}