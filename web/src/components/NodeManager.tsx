"use client"

import { ProtocolContext } from "@/contexts/ProtocolContext";
import { ProtocolNode } from "@/utils/types"
import { getNodesForWorkspace } from "@/actions/ProtocolNode";
import { useContext, useEffect, useState } from "react"
import NodeCard from "./NodeCard";
import { ScreenContext } from "@/contexts/ScreenContext";

export default function NodeManager({ nodeCallback, isNodesPage, className }: { nodeCallback: (node: ProtocolNode | null) => void, isNodesPage?: boolean, className?: string }) {
    const { currentWorkspace } = useContext(ProtocolContext);
    const [nodes, setNodes] = useState<ProtocolNode[]>([]);
    const [updateNodes, setUpdateNodes] = useState<number>(0);
    const { isMobile } = useContext(ScreenContext);

    useEffect(() => {
        const workspaceNodesSetter = async () => {
            if (currentWorkspace._id) {
                const workspaceNodes = await getNodesForWorkspace(currentWorkspace._id.toString())
                setNodes([...workspaceNodes])

            }
        }

        workspaceNodesSetter().then().catch()
    }, [currentWorkspace, updateNodes])
    return (
        <div className={`h-full bg-white w-full rounded-md p-4 ${className} flex flex-col gap-2 h-full ${!isMobile ? "overflow-y-scroll" : ""}`}>
            {isNodesPage && <button
                className="p-2 text-white w-full border-2 rounded-md border-black bg-black text-center hover:bg-white hover:text-black cursor-pointer"
                onClick={() => nodeCallback(null)}
            >
                Add a Node</button>}
            <div className="h-full overflow-y-scroll flex flex-col gap-2">
                {nodes.map((node: ProtocolNode, index: number) => {
                    return (<NodeCard node={node} key={index} updateCallback={() => setUpdateNodes(updateNodes + 1)} onEdit={() => nodeCallback(node)} />)
                })}
            </div>
        </div>

    )
}