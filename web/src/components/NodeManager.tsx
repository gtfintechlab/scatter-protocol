"use client"

import { ProtocolContext } from "@/contexts/ProtocolContext";
import { ProtocolNode, ProtocolNodeState, ScreensURLs } from "@/utils/types"
import { getNodesForWorkspace, updateProtocolNode } from "@/actions/ProtocolNode";
import { useContext, useEffect, useState } from "react"
import NodeCard from "./NodeCard";
import { ScreenContext } from "@/contexts/ScreenContext";
import { startNode, stopNode } from "@/actions/SimulationAPI";

export default function NodeManager({ nodeCallback, className }: { nodeCallback: (node: ProtocolNode | null) => void, className?: string }) {
    const { currentWorkspace } = useContext(ProtocolContext);
    const [nodes, setNodes] = useState<ProtocolNode[]>([]);
    const [updateNodes, setUpdateNodes] = useState<number>(0);
    const { currentScreen } = useContext(ScreenContext);
    const [allStarted, setAllStarted] = useState<boolean>(false);
    const [allStopped, setAllStopped] = useState<boolean>(false);
    const [error, setError] = useState<string>("")
    const [disableStart, setDisableStart] = useState<boolean>(false);

    useEffect(() => {
        const workspaceNodesSetter = async () => {
            if (currentWorkspace._id) {
                const workspaceNodes = await getNodesForWorkspace(currentWorkspace._id.toString())
                setNodes([...workspaceNodes])
                const started = workspaceNodes.every((node: ProtocolNode) => { return node.state === ProtocolNodeState.STARTED })
                const stopped = workspaceNodes.every((node: ProtocolNode) => { return node.state === ProtocolNodeState.STOPPED })

                setAllStarted(started)
                setAllStopped(stopped)
            }
        }
        workspaceNodesSetter().then().catch()
    }, [currentWorkspace, updateNodes])

    const runAllNodes = async () => {
        setError("")
        setDisableStart(true);
        for (let i = 0; i < nodes.length; i++) {
            if (nodes[i].state !== ProtocolNodeState.STARTED) {
                try {
                    await startNode(nodes[i]);
                } catch {
                    setError("Failed to start all nodes; Force starting");
                }
                await updateProtocolNode({ ...nodes[i], state: ProtocolNodeState.STARTED })
            }
        }
        setUpdateNodes(updateNodes + 1);
        setDisableStart(false);
    }

    const stopAllNodes = async () => {
        setError("")
        setDisableStart(true);
        for (let i = 0; i < nodes.length; i++) {
            if (nodes[i].state !== ProtocolNodeState.STOPPED) {
                try {
                    await stopNode(nodes[i].blockchainAddress);
                } catch {
                    setError("Failed to stop all nodes; Force Stopping");
                }
                await updateProtocolNode({ ...nodes[i], state: ProtocolNodeState.STOPPED })
            }
        }

        setUpdateNodes(updateNodes + 1);
        setDisableStart(false);
    }
    return (
        <div className={`h-full bg-white w-full rounded-md p-4 ${className} flex flex-col gap-2 h-full overflow-y-scroll`}>
            <div className="flex flex-col gap-y-2">
                {currentScreen === ScreensURLs.NODES && <button
                    className="p-1 text-white w-full border-2 rounded-md border-black bg-black text-center hover:bg-white hover:text-black cursor-pointer"
                    onClick={() => nodeCallback(null)}
                >
                    Add a Node</button>}
                {(currentScreen === ScreensURLs.HOME && !allStarted) && <div className="flex flex-col gap-y-2">
                    <button
                        className={`p-1 text-white w-full border-2 rounded-md border-green-500 bg-green-500 text-center hover:bg-white hover:text-green-500 cursor-pointer ${disableStart ? "opacity-50" : ""}`}
                        onClick={async () => await runAllNodes()}
                        disabled={disableStart}
                    >
                        Start All Nodes</button>
                </div>}
                {(currentScreen === ScreensURLs.HOME && !allStopped) &&
                    <button
                        className={`p-1 text-white w-full border-2 rounded-md border-red-500 bg-red-500 text-center hover:bg-white hover:text-red-500 cursor-pointer ${disableStart ? "opacity-50" : ""}`}
                        onClick={async () => await stopAllNodes()}
                        disabled={disableStart}
                    >
                        Stop All Nodes
                    </button>}
                {error && <p className="text-red-500 text-center text-sm font-semibold">Error: {error}</p>}
            </div>
            <div className="h-full overflow-y-scroll flex flex-col gap-2">
                {nodes.map((node: ProtocolNode, index: number) => {
                    return (
                        <NodeCard node={node} key={index} updateCallback={() => setUpdateNodes(updateNodes + 1)} onEdit={() => nodeCallback(node)}
                            disableStart={disableStart}
                        />)
                })}
            </div>
        </div>

    )
}