import { useContext, useEffect, useState } from "react"
import { StepAdder } from "./StepAdder";
import { deleteStep, getStepsByWorkspace } from "@/actions/Step";
import { ProtocolContext, StepContext } from "@/contexts/ProtocolContext";
import { ProtocolNode, Step } from "@/utils/types";
import { getNodesForWorkspace } from "@/actions/ProtocolNode";
import { ROLE_TO_COLOR_MAPPING } from "@/utils/constants";
import { capitalizeWords } from "@/utils/format";
import { externalRequest } from "@/utils/requests";
import { urls } from "@/utils/urls";

export default function StepManager({ className }: { className?: string }) {
    const [stepAddMode, setStepAddMode] = useState<boolean>(false);
    const [steps, setSteps] = useState<Step[]>([]);
    const [workspaceNodes, setWorkspaceNodes] = useState<{ [key: string]: ProtocolNode }>({});
    const { currentWorkspace } = useContext(ProtocolContext);
    const { stepUpdateKey, setStepUpdateKey } = useContext(StepContext);
    const [completionStatus, setCompletionStatus] = useState<{ [key: string]: boolean }>({})
    const [isExecuting, setIsExecuting] = useState<boolean>(false)

    useEffect((
    ) => {
        const stepSetter = async () => {
            const workspaceSteps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string)
            setSteps([...workspaceSteps])
        }

        const nodeSetter = async () => {
            const nodes = await getNodesForWorkspace(currentWorkspace._id?.toString() as string);
            const mapping: { [key: string]: ProtocolNode } = {}
            for (let i = 0; i < nodes.length; i++) {
                mapping[nodes[i]._id.toString() as string] = nodes[i];
            }
            setWorkspaceNodes({ ...mapping })
        }

        if (!currentWorkspace._id) {
            return;
        }

        stepSetter().then().catch();
        nodeSetter().then().catch();
    }, [currentWorkspace, stepUpdateKey])

    const deleteWorkspaceStep = async (stepId: string) => {
        await deleteStep(stepId)
        setStepUpdateKey(stepUpdateKey + 1)
    }

    const executeSteps = async () => {
        setIsExecuting(true);
        for (let step of steps) {
            await externalRequest({
                url: urls.externalUrl + step.apiPath,
                method: step.apiMethod,
                body: step.body
            })
            completionStatus[step._id.toString()] = true;
            setCompletionStatus({ ...completionStatus })
        }
        setCompletionStatus({});
        setIsExecuting(false);
    }


    return (
        <main className={`h-full bg-white rounded-md p-4 ${className} flex flex-col gap-2 min-w-0 overflow-y-scroll`}>
            <div className="flex items-center justify-between">
                <button className="justify-center p-1 bg-black text-sm rounded-md hover:bg-white hover:text-black border-2 border-black"
                    type="button" onClick={() => setStepAddMode(!stepAddMode)}>
                    <span className="text-xs">{stepAddMode ? "Back" : "Add New Step"}</span>
                </button>
                {!stepAddMode && <button className={`${isExecuting ? "opacity-50" : ""} justify-center p-1 bg-green-500 text-sm rounded-md hover:bg-white hover:text-green-500 border-2 border-green-500`} type="button"
                    onClick={async () => await executeSteps()}
                    disabled={isExecuting}>
                    <span className="text-xs">Start Execution</span>
                </button>}
            </div>
            {!stepAddMode &&
                <div className="relative w-full overflow-x-auto h-full overflow-y-scroll pb-12">
                    <table className="w-full caption-bottom text-sm border-collapse min-w-full">                    <thead className="[&amp;_tr]:border-b">
                        <tr className="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Step Number</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Address</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Node Role</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Step Name</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">View Step Details</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Delete Step</th>
                            {isExecuting && <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Status</th>}
                        </tr>
                    </thead>
                        <tbody>
                            {steps.map((step: Step, index: number) => {
                                return (
                                    <tr className="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted" key={index}>
                                        <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{step.order}</td>
                                        <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300 truncate max-w-[30px]">{workspaceNodes[step.nodeId.toString() as string]?.blockchainAddress}</td>
                                        <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">
                                            <div className={`p-1 bg-${ROLE_TO_COLOR_MAPPING[workspaceNodes[step.nodeId.toString() as string]?.peerType]}-500 rounded-md text-white text-center`}>
                                                {capitalizeWords(workspaceNodes[step.nodeId.toString() as string]?.peerType)}
                                            </div>
                                        </td>
                                        <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{
                                            step.type
                                        }</td>
                                        <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{
                                            <button className={`p-1 bg-blue-500 rounded-md text-white text-center w-fit border-2 border-blue-500 hover:bg-white hover:text-blue-500`}>
                                                View Details
                                            </button>

                                        }</td>
                                        <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">
                                            <button className={`p-1 bg-red-500 rounded-md text-white text-center w-fit border-2 border-red-500 hover:bg-white hover:text-red-500`}
                                                onClick={async () => await deleteWorkspaceStep(step._id.toString() as string)}>
                                                Delete
                                            </button>
                                        </td>
                                        {isExecuting && <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300 truncate max-w-[30px]">{completionStatus[step._id.toString() as string] === true ? "Complete" : ""}</td>}
                                    </tr>
                                )
                            })}
                        </tbody>
                    </table>
                </div>}
            {stepAddMode && <StepAdder completeAdditionCallback={() => setStepAddMode(false)}></StepAdder>}
        </main>
    )
}