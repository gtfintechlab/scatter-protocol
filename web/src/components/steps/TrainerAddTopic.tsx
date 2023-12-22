import { getNodesForWorkspace, getSingleNodeById } from "@/actions/ProtocolNode";
import { ProtocolContext, StepContext } from "@/contexts/ProtocolContext";
import { DEFAULT_STEP_OPTIONS, STEPS_CONFIG } from "@/utils/constants";
import { PeerType, ProtocolNode, Step, StepTypes } from "@/utils/types";
import { useContext, useEffect, useState } from "react"
import NodeDropdown from "../NodeDropdown";
import { createStep, getStepsByNode, getStepsByWorkspace } from "@/actions/Step";
import GenericDropdownMenu from "../GenericDropdownMenu";

export function TrainerAddTopic({ completionCallback }: { completionCallback: () => void }) {
    const [selectedRequestor, setSelectedRequestor] = useState<string>("");
    const [selectedTrainer, setSelectedTrainer] = useState<string>("");
    const [topicName, setTopicName] = useState<string>("");
    const [trainingDataPath, setTrainingDataPath] = useState<string>("");
    const [stake, setStake] = useState<number | undefined>(undefined);
    const [requestorTopics, setRequestorTopics] = useState<{ [key: string]: string }>({});

    const [submitting, setSubmitting] = useState<boolean>(false);
    const [trainerNodes, setTrainerNodes] = useState<ProtocolNode[]>([]);
    const [requestorNodes, setRequestorNodes] = useState<ProtocolNode[]>([]);
    const { currentWorkspace } = useContext(ProtocolContext)
    const { setStepUpdateKey, stepUpdateKey } = useContext(StepContext);

    const setDefaultValues = () => {
        setTopicName(DEFAULT_STEP_OPTIONS.topicName)
        setTrainingDataPath(DEFAULT_STEP_OPTIONS.trainingDataPath)
        setStake(DEFAULT_STEP_OPTIONS.stake)
    }

    const idToBlockchainAddress = (nodeList: ProtocolNode[]) => {
        const mapping: { [key: string]: ProtocolNode } = {};
        for (let i = 0; i < nodeList.length; i++) {
            mapping[nodeList[i]._id.toString() as string] = nodeList[i];
        }

        return mapping;
    }
    useEffect(() => {
        const workspaceNodeSetter = async () => {
            if (!currentWorkspace._id?.toString()) {
                return;
            }

            const nodes = await getNodesForWorkspace(currentWorkspace._id?.toString())
            setRequestorNodes([...nodes.filter((node) => { return (node.peerType === PeerType.REQUESTOR) })])
            setTrainerNodes([...nodes.filter((node) => { return (node.peerType === PeerType.TRAINER) })])
        }
        setDefaultValues()
        workspaceNodeSetter().then().catch()
    }, [currentWorkspace])

    const { apiPath, apiMethod } = STEPS_CONFIG[StepTypes.TRAINER_ADD_TOPIC]
    const submitStep = async () => {
        setSubmitting(true);
        const steps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string);
        const requestorNode = await getSingleNodeById(selectedRequestor);
        const trainerNode = await getSingleNodeById(selectedTrainer);

        const step: Omit<Step, "_id"> = {
            type: StepTypes.TRAINER_ADD_TOPIC,
            apiPath,
            apiMethod,
            body: {
                topicName,
                requestorAddress: requestorNode.blockchainAddress,
                stake: stake as number,
                trainingDataPath: trainingDataPath,
                blockchainAddress: trainerNode.blockchainAddress
            },
            workspaceId: currentWorkspace._id?.toString() as string,
            nodeId: selectedTrainer,
            order: steps.length
        }

        await createStep(step);
        setSubmitting(false);

        completionCallback();
        setStepUpdateKey(stepUpdateKey + 1)
    }

    useEffect(() => {
        if (!selectedRequestor) {
            return;
        }

        const topicSetter = async () => {
            const steps = await getStepsByNode(selectedRequestor);
            const mapping: { [key: string]: string } = {};
            for (let i = 0; i < steps.length; i++) {
                if (steps[i].type === StepTypes.REQUESTOR_ADD_TOPIC) {
                    const body = steps[i].body;
                    const topicName = body.topicName as string;
                    mapping[topicName] = topicName;
                }
            }
            setRequestorTopics({ ...mapping })
        }

        topicSetter().then().catch()

    }, [selectedRequestor])

    return (
        <div className="h-full flex flex-col justify-between overflow-y-scroll gap-y-3">
            <div className='flex flex-col gap-x-2 gap-y-3 w-full'>
                <div className='flex flex-row gap-4 w-full items-center flex-1 flex-wrap'>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Choose a Node:</label>
                        <NodeDropdown items={idToBlockchainAddress(trainerNodes)} selectedCallback={(nodeId: string) => setSelectedTrainer(nodeId)}></NodeDropdown>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Select Rquestor Node:</label>
                        <NodeDropdown items={idToBlockchainAddress(requestorNodes)} selectedCallback={(nodeId: string) => setSelectedRequestor(nodeId)}></NodeDropdown>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Topic Name:</label>
                        <GenericDropdownMenu items={requestorTopics} selectedCallback={(item: string) => { setTopicName(item) }}></GenericDropdownMenu>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Training Job Path:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Training Job Path'
                            value={trainingDataPath} onChange={(event) => setTrainingDataPath(event?.target.value)} disabled={submitting}></input>
                    </div>
                    <div className='flex flex-col gap-2 shrink grow'>
                        <label className="text-black font-semibold text-sm">Stake:</label>
                        <input className='text-black p-2 border-none outline-none bg-gray-100 rounded-md' placeholder='Reward'
                            value={stake} onChange={(event) => setStake(parseInt(event?.target.value))} type="number" disabled={submitting}></input>
                    </div>
                </div>
            </div>
            <button className={`bg-black self-end w-full p-2 rounded-md text-sm hover:bg-white border-2 border-black hover:text-black ${submitting ? "opacity-50" : ""}`} onClick={async () => {
                await submitStep()
            }} disabled={submitting}>Create Step</button>
        </div>)
}