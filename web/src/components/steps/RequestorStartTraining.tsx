import { useContext, useEffect, useState } from "react"
import NodeDropdown from "../NodeDropdown";
import { getNodesForWorkspace, getSingleNodeById } from "@/actions/ProtocolNode";
import { ProtocolContext, StepContext } from "@/contexts/ProtocolContext";
import { PeerType, ProtocolNode, Step, StepTypes } from "@/utils/types";
import { createStep, getStepsByNode, getStepsByWorkspace } from "@/actions/Step";
import { STEPS_CONFIG } from "@/utils/constants";
import GenericDropdownMenu from "../GenericDropdownMenu";

export default function RequestorStartTraining({ completionCallback }: { completionCallback: () => void }) {
    const [submitting, setSubmitting] = useState<boolean>(false);
    const [workspaceNodes, setWorkspacesNodes] = useState<ProtocolNode[]>([]);
    const { currentWorkspace } = useContext(ProtocolContext);
    const { setStepUpdateKey, stepUpdateKey } = useContext(StepContext)
    const [selectedNodeId, setSelectedNodeId] = useState<string>();
    const [requestorTopics, setRequestorTopics] = useState<{ [key: string]: string }>({});
    const [topicName, setTopicName] = useState<string>("");

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
            setWorkspacesNodes([...nodes.filter((node) => { return (node.peerType === PeerType.REQUESTOR) })])
        }
        workspaceNodeSetter().then().catch()
    }, [currentWorkspace])

    const { apiPath, apiMethod } = STEPS_CONFIG[StepTypes.REQUESTOR_START_TRAINING]

    const submitStep = async () => {
        setSubmitting(true);
        const steps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string);
        const node = await getSingleNodeById(selectedNodeId as string);
        const step: Omit<Step, "_id"> = {
            type: StepTypes.REQUESTOR_START_TRAINING,
            apiPath,
            apiMethod,
            body: {
                blockchainAddress: node.blockchainAddress,
                topicName: topicName
            },
            workspaceId: currentWorkspace._id?.toString() as string,
            nodeId: node._id,
            order: steps.length
        }

        await createStep(step);
        setSubmitting(false);

        completionCallback();
        setStepUpdateKey(stepUpdateKey + 1)

    }

    useEffect(() => {
        if (!selectedNodeId) {
            return;
        }

        const topicSetter = async () => {
            const steps = await getStepsByNode(selectedNodeId);
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

    }, [selectedNodeId])


    return (
        <div className="h-full flex flex-col justify-between overflow-y-scroll gap-y-3">
            <div className="flex flex-row items-center gap-2">
                <div className='flex flex-col gap-2'>
                    <label className="text-black text-sm font-semibold">Select a Protocol Owner</label>
                    <NodeDropdown className='w-fit' items={idToBlockchainAddress()} selectedCallback={(nodeId: string) => setSelectedNodeId(nodeId)}></NodeDropdown>
                </div>
                <div className='flex flex-col gap-2 shrink grow'>
                    <label className="text-black font-semibold text-sm">Topic Name:</label>
                    <GenericDropdownMenu items={requestorTopics} selectedCallback={(item: string) => { setTopicName(item) }}></GenericDropdownMenu>
                </div>
            </div>
            <button className={`bg-black self-end w-full p-2 rounded-md text-sm hover:bg-white border-2 border-black hover:text-black ${submitting ? "opacity-50" : ""}`} onClick={async () => {
                await submitStep()
            }} disabled={submitting}>Create Step</button>
        </div>
    )
}