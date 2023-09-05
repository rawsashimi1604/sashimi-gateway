import HttpRequest from '../../requests/HttpRequest';
import type { GetAllRoutesResponse } from './responses/GetAllRoutes';

function getAllRoutes() {
  return HttpRequest.get<GetAllRoutesResponse>(`/route/all`);
}

export default {
  getAllRoutes
};
