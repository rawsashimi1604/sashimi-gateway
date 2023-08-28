import { ChartConfiguration } from 'c3';
import React, { useState } from 'react';

import SelectInput from '../../components/input/SelectInput';
import Chart from '../../components/utils/Chart';
import { generateTimeData } from '../../utils/utils';

type Timeframe = '1h' | '15m' | '5m' | '1m';

function RequestChart() {
  const [timeframe, setTimeframe] = useState<Timeframe>('15m');

  let chosenTimeSeriesData, format;
  switch (timeframe) {
    case '1h':
      chosenTimeSeriesData = generateTimeData(60, 6);
      format = '%H:%M';
      break;
    case '15m':
      chosenTimeSeriesData = generateTimeData(15, 6);
      format = '%H:%M';
      break;
    case '5m':
      chosenTimeSeriesData = generateTimeData(5, 6);
      format = '%H:%M';
      break;
    case '1m':
      chosenTimeSeriesData = generateTimeData(1, 6);
      format = '%H:%M';
      break;
    default:
      chosenTimeSeriesData = generateTimeData(15, 6);
      format = '%H:%M';
      break;
  }

  const chartConfig: ChartConfiguration = {
    data: {
      x: 'x',
      columns: [
        ['x', ...chosenTimeSeriesData],
        ['2xx', 5, 10, 0, 0, 20, 23],
        ['4xx', 5, 10, 12, 5, 20, 23],
        ['5xx', 5, 10, 0, 20, 20, 23]
      ],
      types: {
        '2xx': 'area-spline',
        '4xx': 'area-spline',
        '5xx': 'area-spline'
      },
      groups: [['2xx', '4xx', '5xx']]
    },
    color: {
      pattern: ['#006400', '#ff7f0e', '#4B0082']
    },
    axis: {
      x: {
        type: 'timeseries',
        tick: {
          format: format,
          culling: false
        }
      }
    }
  };

  return (
    <React.Fragment>
      <div className="text-xs flex items-center justify-end">
        <span className="mr-2 font-md border-r pr-2 border-sashimi-deepgray">
          Error rate (4xx and 5xx Status):{' '}
          <b className="text-sashimi-deeppink font-bold tracking-tighter">
            35.06%
          </b>
        </span>
        <span className="mr-2">select timeframe</span>
        <SelectInput
          options={['1h', '15m', '5m', '1m']}
          onChange={(value) => setTimeframe(value as Timeframe)}
        />
      </div>
      <Chart config={chartConfig} />
    </React.Fragment>
  );
}

export default RequestChart;
