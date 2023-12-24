"use client"

import { getWorkspaceEvents } from "@/actions/LogEvent";
import NodeManager from "@/components/NodeManager";
import { ProtocolContext } from "@/contexts/ProtocolContext";
import { ScreenContext } from "@/contexts/ScreenContext";
import { LogEvent, ProtocolNode, ScreensURLs } from "@/utils/types";
import { useContext, useEffect, useState } from "react";

export default function DataPage() {
    const { setCurrentScreen, isMobile } = useContext(ScreenContext);
    const { currentWorkspace } = useContext(ProtocolContext)
    const [eventsForWorkspace, setEventsForWorkspace] = useState<LogEvent[]>()

    const selectedNodeCallback = (node: ProtocolNode | null) => {

    }

    useEffect(() => {
        setCurrentScreen(ScreensURLs.DATA)

        const eventSetter = async () => {
            const events = await getWorkspaceEvents(currentWorkspace._id?.toString() as string);
            setEventsForWorkspace([...events]);
            const groupedData = events.reduce((acc: { [key: string]: number[] }, obj: LogEvent) => {
                const blockchainAddress = obj.blockchainAddress;
                if (!acc[blockchainAddress]) {
                    acc[blockchainAddress] = [];
                }
                acc[blockchainAddress].push(obj.yDataPoint);
                return acc;
            }, {});
            console.log(groupedData)

        }

        if (!currentWorkspace._id) {
            return
        }

        eventSetter().then().catch()
    }, [currentWorkspace])

    return (
        <main className={`flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10 bg-gray-100 h-full ${isMobile ? 'overflow-y-scroll' : ''}`}>
            <div className="grid md:grid-cols-[360px_1fr] gap-5 items-start h-full">
                <NodeManager nodeCallback={selectedNodeCallback}></NodeManager>
            </div>
        </main>
    )
}
