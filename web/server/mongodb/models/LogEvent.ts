import { LogEvent } from '@/utils/types';
import mongoose, { Schema } from 'mongoose';

const LogEventSchema = new Schema<LogEvent>(
    {
        logType: {
            type: String,
            required: true,
        },
        workspaceId: {
            type: Schema.ObjectId,
            required: true,
        },
        blockchainAddress: {
            type: String,
            required: true,
        },
        xDataPoint: {
            type: Number,
            required: true,
        },
        yDataPoint: {
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

export const LogEventModel =
    (mongoose.models.LogEvent as mongoose.Model<LogEvent>) ??
    mongoose.model<LogEvent>('LogEvent', LogEventSchema);