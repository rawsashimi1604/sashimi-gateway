import HttpRequest from '../../requests/HttpRequest';
import { RegisterRouteBody } from './body/RegisterRouteBody';
import type { GetAggregatedRequestResponse } from './responses/GetAggregatedRequest';
import { GetAllRequestsResponse } from './responses/GetAllRequests';
import { RegisterRouteResponse } from './responses/RegisterRoute';

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

function registerRoute(body: RegisterRouteBody) {
  return HttpRequest.post<RegisterRouteResponse>(`/route`, body);
}

export default {
  getAllRequests,
  getAggregatedRequest,
  registerRoute
};
