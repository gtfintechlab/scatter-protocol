import APIWrapper from "server/utils/APIWrapper";
import { APIWrapperType, Step } from "@/utils/types"
import { createStep, deletStepById, findStepsByWorkspace } from "server/mongodb/actions/Step";
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
            const workspace = await createStep(body);
            return workspace
        },
    },
    DELETE: {
        config: {},
        handler: async (req) => {
            const body: DeleteStepRequest = await req.json();
            const workspace = await deletStepById(body.stepId)
            return workspace
        },
    }

});

export let GET: APIWrapperType, POST: APIWrapperType, PATCH: APIWrapperType, DELETE: APIWrapperType, PUT: APIWrapperType;
GET = POST = PATCH = POST = DELETE = PUT = route;