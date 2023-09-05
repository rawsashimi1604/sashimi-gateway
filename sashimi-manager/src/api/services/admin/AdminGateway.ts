import HttpRequest from '../../requests/HttpRequest';
import type { GetGatewayInformationResponse } from './responses/GetGatewayInformation';

function getGatewayinformation() {
  return HttpRequest.get<GetGatewayInformationResponse>(`/general`);
}

export default {
  getGatewayinformation
};
