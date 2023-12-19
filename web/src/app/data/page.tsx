"use client"

import NodeManager from "@/components/NodeManager";
import { ScreenContext } from "@/contexts/ScreenContext";
import { ProtocolNode, ScreensURLs } from "@/utils/types";
import { useContext, useEffect } from "react";

export default function DataPage() {
    const { setCurrentScreen, isMobile } = useContext(ScreenContext);

    const selectedNodeCallback = (node: ProtocolNode | null) => {

    }

    useEffect(() => {
        setCurrentScreen(ScreensURLs.DATA)
    }, [])

    return (
        <main className={`flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10 bg-gray-100 h-full ${isMobile ? 'overflow-y-scroll' : ''}`}>
            <div className="grid md:grid-cols-[360px_1fr] gap-5 items-start h-full">
                <NodeManager nodeCallback={selectedNodeCallback}></NodeManager>
            </div>
        </main>
    )
}
