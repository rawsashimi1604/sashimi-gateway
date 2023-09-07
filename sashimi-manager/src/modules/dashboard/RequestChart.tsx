import { ChartConfiguration } from 'c3';
import React, { useEffect, useState } from 'react';

import AdminRequest from '../../api/services/admin/AdminRequest';
import { AggregatedRequest } from '../../api/services/admin/responses/AggregatedRequest';
import { GetAggregatedRequestResponse } from '../../api/services/admin/responses/GetAggregatedRequest';
import SelectInput from '../../components/input/SelectInput';
import Chart from '../../components/utils/Chart';
import LoadingSpinner from '../../components/utils/LoadingSpinner';
import LoadingText from '../../components/utils/LoadingText';
import { delay } from '../../utils/delay';

type Timeframe = '1h' | '15m' | '5m' | '1m';

const timeframeMap = {
  '1h': 60,
  '15m': 15,
  '5m': 5,
  '1m': 1
};

function RequestChart() {
  const NUMBER_OF_DATAPOINTS = 6;
  const [timeframe, setTimeframe] = useState<Timeframe>('1h');
  const [errorRate, setErrorRate] = useState('0');
  const [aggregatedReq, setAggregatedReq] = useState<GetAggregatedRequestResponse | null>(null);
  const [chartConfig, setChartConfig] = useState<ChartConfiguration | null>(null);

  async function loadAggregatedResponses(timespan: number, dataPoints: number) {
    setChartConfig(null);
    await delay(500);
    const requests = await AdminRequest.getAggregatedRequest(timespan, dataPoints);
    setAggregatedReq(requests.data);
  }

  async function calculateAndSetErrorRate(requests: AggregatedRequest[]) {
    let totalErrRequests = 0;
    let totalRequests = 0;
    for (const req of requests) {
      totalRequests += req.count;
      totalErrRequests += req.count_4xx;
      totalErrRequests += req.count_5xx;
    }

    // Ignore divide by 0 error
    if (totalRequests === 0) return;
    const errorRate = (totalErrRequests / totalRequests) * 100;
    setErrorRate(errorRate.toFixed(3));
  }

  useEffect(() => {
    const timeframeNumber = timeframeMap[timeframe];
    loadAggregatedResponses(timeframeNumber, NUMBER_OF_DATAPOINTS);
  }, [timeframe]);

  useEffect(() => {
    const requests = aggregatedReq?.requests;
    if (!requests) return;
    calculateAndSetErrorRate(requests);
    setChartConfig({
      data: {
        x: 'x',
        columns: [
          [
            'x',
            ...requests.map((request) => {
              const date = new Date(request.timeBucket);
              return date;
            })
          ],
          ['2xx', ...requests.map((request) => request.count_2xx)],
          ['4xx', ...requests.map((request) => request.count_4xx)],
          ['5xx', ...requests.map((request) => request.count_5xx)]
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
            format: '%H:%M',
            culling: false
          }
        }
      }
    });
  }, [aggregatedReq]);

  return (
    <React.Fragment>
      <div className="text-xs flex items-center justify-end">
        <span className="mr-2 font-md border-r pr-2 border-sashimi-deepgray flex items-center gap-2">
          Error rate (4xx and 5xx Status):{' '}
          {chartConfig ? (
            <b className="text-sashimi-deeppink font-bold tracking-tighter">{errorRate}%</b>
          ) : (
            <span>
              <LoadingSpinner size={12} />
            </span>
          )}
        </span>
        <span className="mr-2">select timeframe</span>
        <SelectInput options={['1h', '15m', '5m', '1m']} onChange={(value) => setTimeframe(value as Timeframe)} />
      </div>
      {chartConfig ? <Chart config={chartConfig} /> : <LoadingText text="loading request chart..." />}
    </React.Fragment>
  );
}

export default RequestChart;
