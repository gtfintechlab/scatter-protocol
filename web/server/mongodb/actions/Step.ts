import { Types } from "mongoose";
import dbConnect from "server/utils/dbConnect";
import { StepModel } from "server/mongodb/models/Step";
import { Step } from "@/utils/types";

export const findStepsByWorkspace = async (workspaceId: string | Types.ObjectId) => {
    await dbConnect();
    const steps = await StepModel.find({ workspaceId });

    return steps
}

export const findStepsByNodeId = async (nodeId: string | Types.ObjectId) => {
    await dbConnect();
    const steps = await StepModel.find({ nodeId });

    return steps
}

export const createStep = async (step: Step | Partial<Step>) => {
    await dbConnect();
    const createdStep = await StepModel.create(step);
    return createdStep;
}

export const deletStepById = async (stepId: string | Types.ObjectId) => {
    await dbConnect();
    const createdStep = await StepModel.findByIdAndDelete({ _id: stepId });
    return createdStep;
}

export const updateStep = async (step: Step | Partial<Step>) => {
    await dbConnect();
    const update = await StepModel.findOneAndUpdate({ _id: step._id }, { ...step })
    return update
}