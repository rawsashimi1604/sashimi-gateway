import HttpRequest from '../../requests/HttpRequest';
import type { GetAllConsumersResponse } from './responses/GetAllConsumers';

function getAllConsumers() {
  return HttpRequest.get<GetAllConsumersResponse>(`/consumer`);
}

// function registerConsumer(body: RegisterRouteBody) {
//   return HttpRequest.post<RegisterRouteResponse>(`/route`, body);
// }

export default {
  getAllConsumers
  //   registerRoute
};
