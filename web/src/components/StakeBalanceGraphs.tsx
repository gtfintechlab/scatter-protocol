import { LogEvent } from "@/utils/types";
import LineChart from "./LineChart";

export default function StakeBalanceGraphs({ dataPoints }: { dataPoints: LogEvent[][] }) {

    return (
        <div className="flex flex-col">
            {dataPoints.map((data: LogEvent[], index: number) => {
                return (
                    <div key={index}>
                        <div>
                            <div className="flex flex-row gap-2 items-center">
                                <h1 className="text-black text-sm font-semibold">Stake Balances:</h1>
                                <h1 className="truncate text-black text-sm">{data.length ? data[0].blockchainAddress : ""}</h1>

                            </div>
                            <LineChart data={data} dataKey="Lottery Balance" xLabel="Time" yLabel="Stake Balance"></LineChart> </div>
                    </div>)
            })}

        </div>
    )
} 