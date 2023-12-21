import { deleteProtocolNode, updateProtocolNode } from "@/actions/ProtocolNode";
import { startNode, stopNode } from "@/actions/SimulationAPI";
import { ScreenContext } from "@/contexts/ScreenContext";
import { ROLE_TO_COLOR_MAPPING } from "@/utils/constants";
import { capitalizeWords } from "@/utils/format";
import { ProtocolNode, ProtocolNodeState, ScreensURLs } from "@/utils/types";
import { useContext, useState } from "react";


export default function NodeCard({ node, onEdit, updateCallback }: { node: ProtocolNode, updateCallback: () => void, onEdit?: () => void }) {
    const { currentScreen } = useContext(ScreenContext);
    const [inUpdateState, setInUpdateState] = useState<boolean>(false);

    const deleteNode = async (nodeId: string) => {
        setInUpdateState(true);
        await deleteProtocolNode(nodeId);
        updateCallback()
        setInUpdateState(false);
    }

    const startProtocolNode = async () => {
        await startNode(node)
        await updateProtocolNode({ ...node, state: ProtocolNodeState.STARTED })
        updateCallback()
    }

    const stopProtocolNode = async () => {
        await stopNode(node.blockchainAddress)
        await updateProtocolNode({ ...node, state: ProtocolNodeState.STOPPED })
        updateCallback()
    }

    return (
        <div className="border text-card-foreground bg-gray-100 rounded-md grid grid-cols-2 gap-x-2 gap-y-4 p-2 text-xs text-black">
            <div className="flex flex-row gap-2 items-center">
                <label className="font-semibold">Node Type:</label>
                <h1 className={`p-1 rounded-md bg-${ROLE_TO_COLOR_MAPPING[node.peerType]}-500 text-white`}>{capitalizeWords(node.peerType)}</h1>
            </div>
            <div className="flex flex-row gap-2 items-center">
                <label className="font-semibold">API Port:</label>
                <h1 className={`p-1 rounded-md text-black`}>{node.apiPort}</h1>
            </div>
            <div className="flex flex-row gap-2 items-center">
                <label className="font-semibold">Blockchain Address:</label>
                <h1 className={`p-1 rounded-md truncate`}>{node.blockchainAddress}</h1>
            </div>
            <div className="flex flex-row gap-2 items-center">
                <label className="font-semibold">Private Key:</label>
                <h1 className={`p-1 rounded-md text-black truncate`}>{node.privateKey}</h1>
            </div>
            <div className="flex flex-row gap-2 items-center">
                <label className="font-semibold">Initial Token Supply:</label>
                <h1 className={`p-1 rounded-md text-black truncate`}>{node.initialTokenSupply}</h1>
            </div>
            <div className="flex flex-row gap-2 items-center">
                <label className="font-semibold">Token Supply:</label>
                <h1 className={`p-1 rounded-md text-black truncate`}>{node.tokenSupply}</h1>
            </div>
            {currentScreen === ScreensURLs.NODES && <div className={`${inUpdateState ? "opacity-50" : ""} flex flex-row gap-2 items-center bg-red-500 text-white rounded-md p-1 border-2 border-red-500 hover:text-red-500 hover:bg-white`}>
                <button className="w-full text-center" onClick={() => deleteNode(node._id.toString())} disabled={inUpdateState}>Delete Node</button>
            </div>}
            {currentScreen === ScreensURLs.NODES && <div className={`${inUpdateState ? "opacity-50" : ""} flex flex-row gap-2 items-center bg-blue-500 text-white rounded-md p-1 border-2 border-blue-500 hover:text-blue-500 hover:bg-white`}>
                <button className="w-full text-center" onClick={() => {
                    setInUpdateState(true);
                    if (onEdit) {
                        onEdit()
                    }
                    setInUpdateState(false);
                }} disabled={inUpdateState}>Edit Node</button>
            </div>}
            {(currentScreen === ScreensURLs.HOME && node.state === ProtocolNodeState.STOPPED) && <div className={`bg-green-500 text-white rounded-md p-1 border-2 border-green-500 hover:text-green-500 hover:bg-white`}>
                <button className="w-full text-center" onClick={async () => {
                    await startProtocolNode()
                }} disabled={inUpdateState}>Start Node</button>
            </div>}
            {(currentScreen === ScreensURLs.HOME && node.state === ProtocolNodeState.STARTED) && <div className={`bg-red-500 text-white rounded-md p-1 border-2 border-red-500 hover:text-red-500 hover:bg-white`}>
                <button className="w-full text-center" onClick={async () => {
                    await stopProtocolNode()
                }} disabled={inUpdateState}>Stop Node</button>
            </div>}
            {(currentScreen === ScreensURLs.HOME) && <div className={`bg-blue-500 text-white rounded-md p-1 border-2 border-blue-500 hover:text-blue-500 hover:bg-white`}>
                <button className="w-full text-center" onClick={async () => {
                }} disabled={inUpdateState}>View Steps</button>
            </div>}

        </div>
    )
}