import c3, { ChartConfiguration } from 'c3';
import React, { useEffect, useRef } from 'react';

interface ChartProps {
  config: ChartConfiguration;
}

function Chart({ config }: ChartProps) {
  const chartRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (chartRef && chartRef.current) {
      const chart = c3.generate({
        ...config,
        bindto: chartRef.current
      });
    }
  }, [config]);

  return <div ref={chartRef}></div>;
}

export default Chart;
