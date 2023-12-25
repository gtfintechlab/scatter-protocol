import { LogEvent } from "@/utils/types";
import LineChart from "./LineChart";

export default function LotteryBalanceGraphs({ dataPoints }: { dataPoints: LogEvent[][] }) {

    return (
        <div className="flex flex-col">
            {dataPoints.map((data: LogEvent[], index: number) => {
                return (
                    <div key={index}>
                        {index === 0 && <div>
                            <div className="flex flex-row gap-2 items-center">
                                <h1 className="text-black text-sm font-semibold">Lottery Balances:</h1>
                            </div>
                            <LineChart data={data} dataKey="Lottery Balance" xLabel="Time" yLabel="Lottery Balance"></LineChart> </div>
                        }
                    </div>)
            })}

        </div>
    )
} 