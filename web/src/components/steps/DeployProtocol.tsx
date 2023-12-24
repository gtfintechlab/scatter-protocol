import { useContext, useEffect, useState } from "react"
import NodeDropdown from "../NodeDropdown";
import { getNodesForWorkspace, getSingleNodeById } from "@/actions/ProtocolNode";
import { ProtocolContext, StepContext } from "@/contexts/ProtocolContext";
import { Optional, ProtocolNode, Step, StepTypes } from "@/utils/types";
import { createStep, getStepsByWorkspace, updateStep } from "@/actions/Step";
import { STEPS_CONFIG } from "@/utils/constants";

export default function DeployProtocol({ completionCallback }: { completionCallback: () => void }) {
    const [submitting, setSubmitting] = useState<boolean>(false);
    const [workspaceNodes, setWorkspacesNodes] = useState<ProtocolNode[]>([]);
    const { currentWorkspace } = useContext(ProtocolContext);
    const { setStepUpdateKey, stepUpdateKey, editMode, stepInEdit, setEditMode, setStepInEdit } = useContext(StepContext)
    const [selectedNodeId, setSelectedNodeId] = useState<string>();
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
            setWorkspacesNodes([...nodes.filter((node) => { return (node.isProtocolOwner) })])
        }
        workspaceNodeSetter().then().catch()
    }, [currentWorkspace])

    useEffect(() => {
        if (editMode) {
            setSelectedNodeId(stepInEdit?.nodeId.toString())
        }
    }, [editMode])


    const submitStep = async () => {
        const { apiPath, apiMethod } = STEPS_CONFIG[StepTypes.DEPLOY_PROTOCOL]
        setSubmitting(true);
        const steps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string);
        const node = await getSingleNodeById(selectedNodeId as string);
        const step: Optional<Step, "_id"> = {
            type: StepTypes.DEPLOY_PROTOCOL,
            apiPath,
            apiMethod,
            body: {
                privateKey: node.privateKey,
                blockchainAddress: node.blockchainAddress
            },
            workspaceId: currentWorkspace._id?.toString() as string,
            nodeId: node._id,
            order: editMode ? stepInEdit?.order as number : steps.length
        }

        if (editMode) {
            step._id = (stepInEdit as Step)._id
            await updateStep(step);
            setEditMode(false);
            setStepInEdit(null);
        } else {
            await createStep(step);
        }
        setSubmitting(false);

        completionCallback();
        setStepUpdateKey(stepUpdateKey + 1)

    }
    return (
        <div className="h-full flex flex-col justify-between overflow-y-scroll gap-y-3">
            <div>
                <label className="text-black text-sm font-semibold">Select a Protocol Owner</label>
                <NodeDropdown items={idToBlockchainAddress()} selectedCallback={(nodeId: string) => setSelectedNodeId(nodeId)}
                    initialValue={stepInEdit?.nodeId.toString()}
                ></NodeDropdown>
            </div>
            <button className={`bg-black self-end w-full p-2 rounded-md text-sm hover:bg-white border-2 border-black hover:text-black ${submitting ? "opacity-50" : ""}`} onClick={async () => {
                await submitStep()
            }} disabled={submitting}>{editMode ? "Update Step" : "Create Step"}</button>
        </div>
    )
}