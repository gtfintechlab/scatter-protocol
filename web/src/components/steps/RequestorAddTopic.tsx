import { getNodesForWorkspace } from "@/actions/ProtocolNode";
import { ProtocolContext } from "@/contexts/ProtocolContext";
import { DEFAULT_STEP_OPTIONS, STEPS_CONFIG } from "@/utils/constants";
import { ProtocolNode, Step, StepTypes } from "@/utils/types";
import { useContext, useEffect, useState } from "react"
import NodeDropdown from "../NodeDropdown";
import { createStep, getStepsByWorkspace } from "@/actions/Step";

export function RequestorAddTopic({ completionCallback }: { completionCallback: () => void }) {
    const [topicName, setTopicName] = useState<string>("");
    const [trainingJobPath, setTrainingJobPath] = useState<string>("");
    const [reward, setReward] = useState<number | undefined>(undefined);
    const [validationThreshold, setValidationThreshold] = useState<number | undefined>(undefined);
    const [evaluationJobPath, setEvaluationJobPath] = useState<string>("");
    const [evaluationJobDataPath, setEvaluationJobDataPath] = useState<string>("");
    const [selectedNode, setSelectedNode] = useState<string>("");

    const [submitting, setSubmitting] = useState<boolean>(false);
    const [workspaceNodes, setWorkspacesNodes] = useState<ProtocolNode[]>([]);
    const { currentWorkspace } = useContext(ProtocolContext)

    const setDefaultValues = () => {
        setTopicName(DEFAULT_STEP_OPTIONS.topicName)
        setTrainingJobPath(DEFAULT_STEP_OPTIONS.trainingJobPath)
        setReward(DEFAULT_STEP_OPTIONS.reward)
        setValidationThreshold(DEFAULT_STEP_OPTIONS.validationThreshold)
        setEvaluationJobPath(DEFAULT_STEP_OPTIONS.evaluationJobPath)
        setEvaluationJobDataPath(DEFAULT_STEP_OPTIONS.evaluationJobDataPath)
    }

    const idToBlockchainAddress = () => {
        const mapping: { [key: string]: ProtocolNode } = {};
        for (let i = 0; i < workspaceNodes.length; i++) {
            mapping[workspaceNodes[i]._id.toString() as string] = workspaceNodes[i];
        }

        return mapping;
    }
    useEffect(() => {
        const workspaceNodeSetter = async () => {
            if (!currentWorkspace._id?.toString()) {
                return;
            }

            const nodes = await getNodesForWorkspace(currentWorkspace._id?.toString())
            setWorkspacesNodes([...nodes])
        }
        setDefaultValues()
        workspaceNodeSetter().then().catch()
    }, [currentWorkspace])

    const { apiPath, apiMethod } = STEPS_CONFIG[StepTypes.REQUESTOR_ADD_TOPIC]
    const submitStep = async () => {
        setSubmitting(true);
        const steps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string);

        const step: Omit<Step, "_id"> = {
            type: StepTypes.REQUESTOR_ADD_TOPIC,
            apiPath,
            apiMethod,
            body: {
                topicName,
                trainingJobPath,
                reward: reward as number,
                validationThreshold: validationThreshold as number,
                evaluationJobDataPath,
                evaluationJobPath
            },
            workspaceId: currentWorkspace._id?.toString() as string,
            node: selectedNode,
            order: steps.length
        }

        await createStep(step);
        setSubmitting(false);

        completionCallback();
    }

    return (
        <div className="h-full flex flex-col justify-between overflow-y-scroll gap-y-3">
            <div className='flex flex-col gap-x-2 gap-y-3 w-full'>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Choose Node:</label>
                        <NodeDropdown items={idToBlockchainAddress()} selectedCallback={(nodeId: string) => setSelectedNode(nodeId)}></NodeDropdown>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Topic Name:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Topic Name'
                            value={topicName} onChange={(event) => setTopicName(event?.target.value)} disabled={submitting}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Training Job Path:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Training Job Path'
                            value={trainingJobPath} onChange={(event) => setTrainingJobPath(event?.target.value)} disabled={submitting}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Reward:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Reward'
                            value={reward} onChange={(event) => setReward(parseInt(event?.target.value))} type="number" disabled={submitting}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Validation Threshold:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Validation Threshold'
                            value={validationThreshold} onChange={(event) => setValidationThreshold(parseInt(event?.target.value))} type="number" disabled={submitting}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Evaluation Job Path:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Evaluation Job Path'
                            value={evaluationJobPath} onChange={(event) => setEvaluationJobPath(event?.target.value)} disabled={submitting}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Evaluation Job Data Path:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Evaluation Job Data Path'
                            value={evaluationJobDataPath} onChange={(event) => setEvaluationJobDataPath(event?.target.value)} disabled={submitting}></input>
                    </div>
                </div>
            </div>
            <button className={`bg-black self-end w-full p-2 rounded-md text-sm hover:bg-white border-2 border-black hover:text-black ${submitting ? "opacity-50" : ""}`} onClick={async () => {
                await submitStep()
            }} disabled={submitting}>Create Step</button>
        </div>)
}