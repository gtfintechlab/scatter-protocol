import { Workspace } from "@/utils/types";
import mongoose, { Schema } from "mongoose";

const WorkspaceSchema = new Schema<Workspace>(
    {
        name: {
            type: String,
            required: true
        }
    },
    {
        timestamps: {
            createdAt: true,
            updatedAt: false,
        },
    }
);

export const WorkspaceModel =
    (mongoose.models.Workspace as mongoose.Model<Workspace>) ??
    mongoose.model<Workspace>('Workspace', WorkspaceSchema);