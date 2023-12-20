import { Types } from "mongoose";
import dbConnect from "server/utils/dbConnect";
import { ProtocolNodeModel } from "../models/ProtocolNode";
import { ProtocolNode } from "@/utils/types";

export const getNodesByWorkspaceId = async (workspaceId: string | Types.ObjectId) => {
    await dbConnect();
    const nodes = await ProtocolNodeModel.find({ workspaceId })
    return nodes;
}

export const createProtocolNode = async (nodeInfo: ProtocolNode | Partial<ProtocolNode>) => {
    await dbConnect();
    const node = await ProtocolNodeModel.create(nodeInfo);
    return node;
}