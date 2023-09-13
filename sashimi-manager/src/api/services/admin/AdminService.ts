import HttpRequest from '../../requests/HttpRequest';
import type { RegisterServiceBody } from './body/RegisterServiceBody';
import type { GetAllServicesResponse } from './responses/GetAllServices';
import type { GetServiceByIdResponse } from './responses/GetServiceById';
import type { RegisterServiceResponse } from './responses/RegisterService';

function getAllServices() {
  return HttpRequest.get<GetAllServicesResponse>(`/service/all`);
}

function getServiceById(id: string) {
  return HttpRequest.get<GetServiceByIdResponse>(`/service/${id}`);
}

function registerService(body: RegisterServiceBody) {
  return HttpRequest.post<RegisterServiceResponse>(`/service`, body);
}

export default {
  getAllServices,
  getServiceById,
  registerService
};
