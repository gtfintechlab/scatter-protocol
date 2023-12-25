import { getWorkspaceEvents } from "@/actions/LogEvent";
import { ProtocolContext } from "@/contexts/ProtocolContext";
import { LogEvent, LogTypes, Nodes, PeerType } from "@/utils/types";
import { useContext, useEffect, useState } from "react"
import LineChart from "./LineChart";
import { getNodeByBlockchainAddressAndWorkspace } from "@/actions/ProtocolNode";
import { ROLE_TO_COLOR_MAPPING } from "@/utils/constants";
import { capitalizeWords } from "@/utils/format";
import TokenBalanceGraphs from "./TokenBalanceGraphs";
import LotteryBalanceGraphs from "./LotteryBalanceGraphs";

export default function DataManager({ className, blockchainAddress }: { blockchainAddress: string, className?: string }) {
    const { currentWorkspace } = useContext(ProtocolContext)
    const [eventData, setEventData] = useState<{ [key: string]: LogEvent[][] }>({});
    const [filteredEventData, setFilteredEventData] = useState<{ [key: string]: LogEvent[][] }>({});
    const [initialLoad, setInitialLoad] = useState<boolean>(true);
    const [addressToRole, setAddressToRole] = useState<{ [key: string]: PeerType }>({});
    useEffect(() => {
        const eventSetter = async () => {
            const events = await getWorkspaceEvents(currentWorkspace._id?.toString() as string);
            // Create a new object to store grouped data
            const groupedData: { [key: string]: LogEvent[][] } = {};
            const roleMapping: { [key: string]: PeerType } = {}

            for (let i = 0; i < events.length; i++) {
                const logType = events[i].logType;
                const blockchainAddress = events[i].blockchainAddress;
                groupedData[logType] = groupedData[logType] || [];
                const subArray = groupedData[logType].find((arr) => arr[0].blockchainAddress === blockchainAddress);
                if (subArray) {
                    // If subarray exists, push the current item
                    subArray.push(events[i]);
                } else {
                    // If subarray doesn't exist, create a new one
                    groupedData[logType].push([events[i]]);
                }

                if (roleMapping[blockchainAddress]) {
                    continue;
                }
                const node = await getNodeByBlockchainAddressAndWorkspace(blockchainAddress, currentWorkspace._id?.toString() as string)
                roleMapping[blockchainAddress] = node.peerType;
            }
            setEventData({ ...groupedData })
            setFilteredEventData({ ...groupedData })
            setAddressToRole({ ...roleMapping })
        }

        if (!currentWorkspace._id) {
            setInitialLoad(false);
            return
        }

        eventSetter().then().catch()
        setInitialLoad(false);
    }, [currentWorkspace])

    useEffect(() => {
        if (initialLoad) {
            return
        }

        if (blockchainAddress === Nodes.ALL) {
            setFilteredEventData({ ...eventData })
            return
        }

        // Create a new object for the final result
        const filteredObject: { [key: string]: LogEvent[][] } = {};
        // Iterate through groupedData and extract the specified blockchainAddress
        for (const logType in eventData) {
            const logTypeData = eventData[logType];
            // Find the array of arrays that corresponds to the specified blockchainAddress
            const selectedArray = logTypeData.find((arr) => arr[0].blockchainAddress === blockchainAddress);

            if (selectedArray) {
                // If the array is found, add it to the finalObject
                filteredObject[logType] = [selectedArray];
            }
        }
        setFilteredEventData({ ...filteredObject })

    }, [blockchainAddress])

    return (
        <main className={`h-full bg-white rounded-md p-4 ${className} flex flex-col gap-2 min-w-0 overflow-y-scroll`}>
            {Object.entries(filteredEventData).map(([logType, dataPoints], index: number) => {
                return (
                    <div className="flex flex-col" key={index}>
                        {logType === LogTypes.TOKEN_BALANCE && <TokenBalanceGraphs dataPoints={dataPoints} addressToRole={addressToRole}></TokenBalanceGraphs>}
                        {logType === LogTypes.LOTTERY_BALANCE && <LotteryBalanceGraphs dataPoints={dataPoints} ></LotteryBalanceGraphs>}

                        {/* {dataPoints.map((data: LogEvent[], index: number) => {
                            return (
                                <div key={index}>
                                    <div className="flex flex-row gap-2 items-center">
                                        <h1 className="text-black text-sm font-semibold">{logType}:</h1>
                                        {logType == LogTypes.TOKEN_BALANCE && <h1 className="truncate text-black text-sm">{data.length ? data[0].blockchainAddress : ""}</h1>}
                                        {logType == LogTypes.TOKEN_BALANCE && <h1 className={`p-1 text-sm rounded-md bg-${ROLE_TO_COLOR_MAPPING[data.length ? addressToRole[data[0].blockchainAddress] : PeerType.NO_ROLE]}-500 text-white`}>{capitalizeWords(data.length ? addressToRole[data[0].blockchainAddress] : "")}</h1>}
                                    </div>
                                    <LineChart data={data}></LineChart>
                                </div>)
                        })} */}
                    </div>
                )
            })}
        </main>
    )
}