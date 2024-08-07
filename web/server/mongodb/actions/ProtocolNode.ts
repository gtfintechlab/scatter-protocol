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

export const deleteProtocolNode = async (nodeId: string | Types.ObjectId) => {
    await dbConnect();
    const node = await ProtocolNodeModel.findOneAndDelete({ _id: nodeId });
    return node;
}


export const updateProtocolNode = async (nodeInfo: ProtocolNode | Partial<ProtocolNode>) => {
    await dbConnect();
    const node = await ProtocolNodeModel.findOneAndUpdate({ _id: nodeInfo._id }, nodeInfo);
    return node;
}

export const getNodeById = async (nodeId: string | Types.ObjectId) => {
    await dbConnect();
    const node = await ProtocolNodeModel.findOne({ _id: nodeId })
    return node;
}


export const getNodeByBlockchainAndWorkspace = async (blockchainAddress: string, workspaceId: string | Types.ObjectId) => {
    await dbConnect();
    const node = await ProtocolNodeModel.findOne({ blockchainAddress, workspaceId })
    return node;
}
export const resetNodeTokenSupply = async (workspaceId: string) => {
    await dbConnect();
    const nodes = await ProtocolNodeModel.find({ workspaceId: workspaceId });

    for (let i = 0; i < nodes.length; i++) {
        await ProtocolNodeModel.findOneAndUpdate({ _id: nodes[i]._id }, { tokenSupply: 0 })
    }
}