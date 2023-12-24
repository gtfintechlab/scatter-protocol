import { useContext, useEffect, useState } from "react"
import NodeDropdown from "../NodeDropdown";
import { getNodesForWorkspace, getSingleNodeById } from "@/actions/ProtocolNode";
import { ProtocolContext, StepContext } from "@/contexts/ProtocolContext";
import { Optional, ProtocolNode, Step, StepTypes } from "@/utils/types";
import { createStep, getStepsByWorkspace, updateStep } from "@/actions/Step";
import { STEPS_CONFIG } from "@/utils/constants";

export default function ProtocolOwnerTransferInitialSupply({ completionCallback }: { completionCallback: () => void }) {
    const [submitting, setSubmitting] = useState<boolean>(false);
    const [gettingTransfers, setGettingTransfers] = useState<boolean>(false);
    const [workspaceNodes, setWorkspacesNodes] = useState<ProtocolNode[]>([]);
    const { currentWorkspace } = useContext(ProtocolContext);
    const { setStepUpdateKey, stepUpdateKey } = useContext(StepContext)
    const [selectedNodeId, setSelectedNodeId] = useState<string>();
    const { stepInEdit, editMode, setEditMode, setStepInEdit } = useContext(StepContext)
    const [supplyTransfer, setSupplyTransfer] = useState<{ [key: string]: number }>({})

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

        const supplySetter = async () => {
            const transfers = await getNodeToInitialSupply();
            setSupplyTransfer({ ...transfers });
        }
        workspaceNodeSetter().then().catch()
        supplySetter().then().catch()
    }, [currentWorkspace])

    const getNodeToInitialSupply = async () => {
        const allNodes = await getNodesForWorkspace(currentWorkspace._id?.toString() as string)

        const nodeToInitialSupply: { [key: string]: number } = {}
        for (let i = 0; i < allNodes.length; i++) {
            nodeToInitialSupply[allNodes[i].blockchainAddress] = allNodes[i].initialTokenSupply;
        }

        return nodeToInitialSupply;

    }

    const refreshSupplyTransfers = async () => {
        setGettingTransfers(true);
        const transfers = await getNodeToInitialSupply();
        setSupplyTransfer({ ...transfers });
        setGettingTransfers(false);
    }

    const submitStep = async () => {
        const { apiPath, apiMethod } = STEPS_CONFIG[StepTypes.PROTOCOL_OWNER_TRANSFER_TOKEN]

        setSubmitting(true);
        const steps = await getStepsByWorkspace(currentWorkspace._id?.toString() as string);
        const node = await getSingleNodeById(selectedNodeId as string);

        const step: Optional<Step, "_id"> = {
            type: StepTypes.PROTOCOL_OWNER_TRANSFER_TOKEN,
            apiPath,
            apiMethod,
            body: {
                privateKey: node.privateKey,
                blockchainAddress: node.blockchainAddress,
                transferAmounts: supplyTransfer
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

    useEffect(() => {
        if (editMode) {
            setSelectedNodeId(stepInEdit?.nodeId.toString())
        }
    }, [editMode])
    return (
        <div className="h-full flex flex-col justify-between overflow-y-scroll gap-y-3">
            <div className="flex flex-row gap-2 items-center">
                <div>
                    <label className="text-black text-sm font-semibold">Select a Protocol Owner</label>
                    <NodeDropdown items={idToBlockchainAddress()} selectedCallback={(nodeId: string) => setSelectedNodeId(nodeId)} initialValue={stepInEdit?.nodeId.toString()}></NodeDropdown>
                </div>
                <div className="flex flex-col gap-1">
                    <label className="text-black text-sm font-semibold">Refresh</label>
                    <button className={`bg-green-500 text-sm hover:bg-green-600 rounded-md p-2 text-white border-green-500 ${gettingTransfers ? "opacity-50" : ""}`}
                        disabled={gettingTransfers}
                        onClick={async () => { await refreshSupplyTransfers() }}
                    >Refresh Supply Transfers</button>
                </div>
            </div>
            <div className="relative w-full overflow-x-auto h-full overflow-y-scroll pb-12">
                <table className="w-full caption-bottom text-sm border-collapse min-w-full">
                    <thead className="[&amp;_tr]:border-b">
                        <tr className="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted">
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Address</th>
                            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-500">Transfer Amount</th>
                        </tr>
                    </thead>
                    <tbody>
                        {Object.entries(supplyTransfer).map(([address, amount], index: number) => {
                            return (
                                <tr className="border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted" key={index}>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300">{address}</td>
                                    <td className="p-4 align-middle [&amp;:has([role=checkbox])]:pr-0 text-black border-solid border-b-2 border-gray-300 truncate max-w-[30px]">{amount}</td>
                                </tr>
                            )
                        })}
                    </tbody>
                </table>
            </div>

            <button className={`bg-black self-end w-full p-2 rounded-md text-sm hover:bg-white border-2 border-black hover:text-black ${submitting ? "opacity-50" : ""}`} onClick={async () => {
                await submitStep()
            }} disabled={submitting}>{editMode ? "Update Step" : "Create Step"}</button>
        </div>
    )
}