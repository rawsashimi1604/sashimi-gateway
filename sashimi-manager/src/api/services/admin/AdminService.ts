import HttpRequest from '../../requests/HttpRequest';
import type { GetAllServicesResponse } from './responses/GetAllServices';
import type { GetServiceByIdResponse } from './responses/GetServiceById';

function getAllServices() {
  return HttpRequest.get<GetAllServicesResponse>(`/service/all`);
}

function getServiceById(id: string) {
  return HttpRequest.get<GetServiceByIdResponse>(`/service/${id}`);
}

export default {
  getAllServices,
  getServiceById
};
