import { AggregatedRequest } from './AggregatedRequest';

export type GetAggregatedRequestResponse = {
  dataPoints: number;
  requests: AggregatedRequest[];
};
