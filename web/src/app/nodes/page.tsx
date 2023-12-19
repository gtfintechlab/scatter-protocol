"use client"

import NodeDetailViewer from "@/components/NodeDetailViewer";
import NodeManager from "@/components/NodeManager";
import { ScreenContext } from "@/contexts/ScreenContext"
import { ProtocolNode, ScreensURLs } from "@/utils/types";
import { useContext, useEffect, useState } from "react"


export default function NodesPage() {

    const { setCurrentScreen, isMobile } = useContext(ScreenContext);
    const [selectedNode, setSelectedNode] = useState<ProtocolNode | null>(null)
    const selectedNodeCallback = (node: ProtocolNode | null) => {
        setSelectedNode(node)
    }
    useEffect(() => {
        setCurrentScreen(ScreensURLs.NODES)
    }, [])
    return (<main className={`flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10 bg-gray-100 h-full ${isMobile ? 'overflow-y-scroll' : ''}`}>
        <div className="grid md:grid-cols-[360px_1fr] gap-5 items-start h-full">
            <NodeManager nodeCallback={selectedNodeCallback} isNodesPage={true}></NodeManager>
            <NodeDetailViewer node={selectedNode}>

            </NodeDetailViewer>
        </div>
    </main>
    )
}
