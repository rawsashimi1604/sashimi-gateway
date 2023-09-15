import HttpRequest from '../../requests/HttpRequest';
import type { GetAggregatedRequestResponse } from './responses/GetAggregatedRequest';
import type { GetAllRequestsResponse } from './responses/GetAllRequests';

function getAllRequests() {
  return HttpRequest.get<GetAllRequestsResponse>(`/request/all`);
}

function getAggregatedRequest(timespan: number, dataPoints: number) {
  let queryString = '/request/aggregate?';

  if (timespan) {
    queryString += `timespan=${encodeURIComponent(timespan)}&`;
  }

  if (dataPoints) {
    queryString += `dataPoints=${encodeURIComponent(dataPoints)}&`;
  }

  return HttpRequest.get<GetAggregatedRequestResponse>(queryString);
}

export default {
  getAllRequests,
  getAggregatedRequest
};
