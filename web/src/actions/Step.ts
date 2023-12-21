import { internalRequest } from "@/utils/requests";
import { HttpMethod, Step } from "@/utils/types";
import { urls } from "@/utils/urls";

const stepUrl = urls.baseUrl + urls.api.steps.general;

export const getStepsByWorkspace = async (workspaceId: string) => {
    return internalRequest<Step[]>({
        url: stepUrl,
        method: HttpMethod.GET,
        queryParams: {
            workspaceId
        }
    })
}

export const createStep = async (step: Step | Partial<Step>) => {
    return internalRequest<Step[]>({
        url: stepUrl,
        method: HttpMethod.POST,
        body: {
            ...step
        }
    })
}

export const deleteStep = async (stepId: string) => {
    return internalRequest<Step[]>({
        url: stepUrl,
        method: HttpMethod.DELETE,
        body: {
            stepId
        }
    })
}