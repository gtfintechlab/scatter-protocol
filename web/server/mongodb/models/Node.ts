import { PeerType, ProtocolNode } from "@/utils/types";
import mongoose, { Schema } from "mongoose";

const ProtocolNodeSchema = new Schema<ProtocolNode>(
    {
        peerType: {
            type: String,
            required: true,
            enum: Object.values(PeerType),
        },
        apiPort: {
            type: String,
            required: true,
        },
        postgresUsername: {
            type: String,
            required: true,
        },
        postgresPassword: {
            type: String,
            required: true,
        },
        databasePort: {
            type: String,
            required: true,
        },
        blockchainAddress: {
            type: String,
            required: true,
        },
        privateKey: {
            type: String,
            required: true,
        },
        dummyLoad: {
            type: Boolean,
            required: true,
        },
        useMdns: {
            type: Boolean,
            required: true,
        },
        tokenSupply: {
            type: Number,
            required: true,
        },
        isProtocolOwner: {
            type: Boolean,
            required: true,
        },
    },
    {
        timestamps: {
            createdAt: true,
            updatedAt: false,
        },
    }
);

export const ProtocolNodeModel =
    (mongoose.models.ProtocolNode as mongoose.Model<ProtocolNode>) ??
    mongoose.model<ProtocolNode>('ProtocolNode', ProtocolNodeSchema);