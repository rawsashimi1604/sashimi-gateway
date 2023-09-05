import HttpRequest from '../../requests/HttpRequest';
import type { GetAllRequestsResponse } from './responses/GetAllRequests';

function getAllRequests() {
  return HttpRequest.get<GetAllRequestsResponse>(`/request/all`);
}

export default {
  getAllRequests
};
