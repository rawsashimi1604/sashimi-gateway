import HttpRequest from '../../requests/HttpRequest';
import type { RegisterRouteBody } from './body/RegisterRouteBody';
import type { GetAllRoutesResponse } from './responses/GetAllRoutes';
import type { RegisterRouteResponse } from './responses/RegisterRoute';

function getAllRoutes() {
  return HttpRequest.get<GetAllRoutesResponse>(`/route/all`);
}

function registerRoute(body: RegisterRouteBody) {
  return HttpRequest.post<RegisterRouteResponse>(`/route`, body);
}

export default {
  getAllRoutes,
  registerRoute
};
