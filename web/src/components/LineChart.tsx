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
    xAccessor: (data: LogEvent) => data.xDataPoint,
    yAccessor: (data: LogEvent) => data.yDataPoint,
};

const LineChart = ({ data }: { data: LogEvent[] }) => {

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
        <XYChart height={175} xScale={{ type: 'band' }} yScale={{ type: 'linear', domain: [getMinY(), getMaxY()], zero: false }}>
            <AnimatedAxis orientation="bottom" />
            <AnimatedAxis orientation="left" numTicks={5} />
            <AnimatedGrid columns={false} numTicks={0} />
            <AnimatedLineSeries dataKey="Line 1" data={data} {...accessors} />
            <AnimatedLineSeries dataKey="Line 2" data={data} {...accessors} />
        </XYChart>
    );
};

export default LineChart;
