import { StepTypes, Step, HttpMethod } from "@/utils/types";
import mongoose, { Schema } from "mongoose";

const StepSchema = new Schema<Step>(
    {
        type: {
            type: String,
            required: true,
            enum: Object.values(StepTypes),
        },
        apiPath: {
            type: String,
            required: true,
        },
        apiMethod: {
            type: String,
            required: true,
            enum: Object.values(HttpMethod),
        },
        body: {
            type: Schema.Types.Mixed,
            required: true,
        },
        workspaceId: {
            type: Schema.Types.ObjectId,
            required: true,
        },
        node: {
            type: Schema.Types.ObjectId,
            required: true,
        },
        order: {
            type: Number,
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

export const StepModel =
    (mongoose.models.Step as mongoose.Model<Step>) ??
    mongoose.model<Step>('Step', StepSchema);