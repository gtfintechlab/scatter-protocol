// LineChart.tsx

import React from 'react';
import { LogEvent } from '@/utils/types';
import {
    AnimatedAxis, // any of these can be non-animated equivalents
    AnimatedGrid,
    AnimatedLineSeries,
    XYChart,
    Tooltip,
} from '@visx/xychart';

const accessors = {
    xAccessor: (data: LogEvent) => {
        const date = new Date(data.xDataPoint);
        const year = date.getFullYear();
        const month = date.getMonth() + 1; // Month is zero-based, so add 1
        const day = date.getDate();
        const hours = date.getHours();
        const minutes = date.getMinutes();
        const seconds = date.getSeconds();
        return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
    },
    yAccessor: (data: LogEvent) => data.yDataPoint,
};

const LineChart = ({ data, dataKey, xLabel, yLabel }: { data: LogEvent[], dataKey: string, xLabel: string, yLabel: string }) => {

    const getMinY = () => {
        const yPoints = data.map((event: LogEvent) => {
            return event.yDataPoint
        })
        return Math.min(...yPoints) - 100
    }

    const getMaxY = () => {
        const yPoints = data.map((event: LogEvent) => {
            return event.yDataPoint
        })

        return Math.max(...yPoints) + 100
    }
    return (
        <div className='flex flex-row'>
            <h1 className='text-xs text-black self-center rotate-[-90deg] font-semibold text-center'>{yLabel}</h1>
            <div className='w-full'>
                <XYChart height={500} xScale={{ type: 'band' }} yScale={{ type: 'linear', domain: [getMinY(), getMaxY()], zero: false }}>
                    <AnimatedAxis orientation="bottom" numTicks={5} label={xLabel} />
                    <AnimatedAxis orientation="left" numTicks={5} />
                    <AnimatedGrid columns={false} numTicks={0} />
                    <AnimatedLineSeries dataKey={dataKey} data={data} {...accessors} colorAccessor={() => { return 'red' }} />
                    <Tooltip
                        snapTooltipToDatumX
                        snapTooltipToDatumY
                        showVerticalCrosshair
                        showSeriesGlyphs
                        renderTooltip={({ tooltipData, colorScale }: any) => (
                            <div>
                                <div style={{ color: colorScale(tooltipData.nearestDatum.key) }}>
                                    {tooltipData.nearestDatum.key}
                                </div>
                                {accessors.xAccessor(tooltipData.nearestDatum.datum)}
                                {', '}
                                {accessors.yAccessor(tooltipData.nearestDatum.datum)}
                            </div>
                        )}
                    />
                </XYChart>
            </div>
        </div>
    );
};

export default LineChart;
