import HttpRequest from '../../requests/HttpRequest';
import type { GetAllServicesResponse } from './responses/GetAllServices';

function getAllServices() {
  return HttpRequest.get<GetAllServicesResponse>(`/services/all`);
}

export default {
  getAllServices
};
