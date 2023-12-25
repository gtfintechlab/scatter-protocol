import { LogEvent, PeerType } from "@/utils/types";
import LineChart from "./LineChart";
import { ROLE_TO_COLOR_MAPPING } from "@/utils/constants";
import { capitalizeWords } from "@/utils/format";

export default function TokenBalanceGraphs({ dataPoints, addressToRole }: { dataPoints: LogEvent[][], addressToRole: { [key: string]: PeerType } }) {

    return (
        <div className="flex flex-col">
            {dataPoints.map((data: LogEvent[], index: number) => {
                return (
                    <div key={index}>
                        <div className="flex flex-row gap-2 items-center">
                            <h1 className="text-black text-sm font-semibold">Token Balances:</h1>
                            <h1 className="truncate text-black text-sm">{data.length ? data[0].blockchainAddress : ""}</h1>
                            <h1 className={`p-1 text-sm rounded-md bg-${ROLE_TO_COLOR_MAPPING[(data.length ? addressToRole[data[0].blockchainAddress] : PeerType.NO_ROLE) as PeerType]}-500 text-white`}>{capitalizeWords(data.length ? addressToRole[data[0].blockchainAddress] : "")}</h1>
                        </div>
                        <LineChart data={data} dataKey="Token Balance" xLabel="Time" yLabel="Scatter Token Balance"></LineChart>
                    </div>)
            })}

        </div>
    )
} 