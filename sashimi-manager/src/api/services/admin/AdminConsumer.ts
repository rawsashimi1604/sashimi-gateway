import HttpRequest from '../../requests/HttpRequest';
import { RegisterConsumerBody } from './body/RegisterConsumerBody';
import type { GetAllConsumersResponse } from './responses/GetAllConsumers';

function getAllConsumers() {
  return HttpRequest.get<GetAllConsumersResponse>(`/consumer`);
}

function registerConsumer(body: RegisterConsumerBody) {
  return HttpRequest.post<RegisterConsumerBody>(`/consumer`, body);
}

export default {
  getAllConsumers,
  registerConsumer
};
