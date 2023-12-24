import { EthereumAccount } from '@/utils/types';
import mongoose, { Schema } from 'mongoose';

const EthereumAccountSchema = new Schema<EthereumAccount>(
    {
        privateKey: {
            type: String,
            required: true,
        },
        address: {
            type: String,
            required: true,
        },
        mnemonic: {
            type: String,
            required: true,
        },
        path: {
            type: String,
            required: true,
        },
        index: {
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

export const EthereumAccountModel =
    (mongoose.models.EthereumAccount as mongoose.Model<EthereumAccount>) ??
    mongoose.model<EthereumAccount>('EthereumAccount', EthereumAccountSchema);