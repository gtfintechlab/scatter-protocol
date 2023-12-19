"use client"

import { ProtocolContext } from "@/contexts/ProtocolContext";
import { ProtocolNode } from "@/utils/types"
import { getNodesForWorkspace } from "@/actions/ProtocolNode";
import { useContext, useEffect, useState } from "react"

export default function NodeManager({ nodeCallback, isNodesPage }: { nodeCallback: (node: ProtocolNode | null) => void, isNodesPage?: boolean }) {
    const { currentWorkspace } = useContext(ProtocolContext);
    const [nodes, setNodes] = useState<ProtocolNode[]>([]);
    useEffect(() => {
        const workspaceNodesSetter = async () => {
            if (currentWorkspace._id) {
                const workspaceNodes = await getNodesForWorkspace(currentWorkspace._id.toString())
                setNodes([...workspaceNodes])
                console.log(workspaceNodes)
            }
        }

        workspaceNodesSetter().then().catch()
    }, [])
    return (
        <div className="h-full bg-white w-full rounded-md p-4">
            <div>
                {nodes.map((node: ProtocolNode) => {
                    return (
                        <div></div>
                    )
                })}
            </div>
            {isNodesPage && <div
                className="p-2 text-white w-full border-2 rounded-md border-black bg-black text-center hover:bg-white hover:text-black cursor-pointer"
                onClick={() => nodeCallback(null)}
            >
                Add a Node</div>}
        </div>
    )
}