import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType, Step } from "@/utils/types"
import { createStep, deletStepById, findStepsByWorkspace, updateStep } from "server/mongodb/actions/Step";
import { Types } from "mongoose";

interface DeleteStepRequest {
    stepId: string | Types.ObjectId;
}

const route: APIWrapperType = APIWrapper({
    GET: {
        config: {
        },
        handler: async (req) => {
            const workspaceId = req.nextUrl.searchParams.get("workspaceId")
            const steps = await findStepsByWorkspace(workspaceId as string)
            return steps
        },
    },
    POST: {
        config: {},
        handler: async (req) => {
            const body: Step | Partial<Step> = await req.json();
            const step = await createStep(body);
            return step
        },
    },
    DELETE: {
        config: {},
        handler: async (req) => {
            const body: DeleteStepRequest = await req.json();
            const step = await deletStepById(body.stepId)
            return step
        },
    },
    PATCH: {
        config: {},
        handler: async (req) => {
            const body: Step | Partial<Step> = await req.json();
            const step = await updateStep(body);
            return step;
        },
    }

});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;