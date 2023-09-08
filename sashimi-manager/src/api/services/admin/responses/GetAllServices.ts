import { Service } from '../models/Service';

export type GetAllServicesResponse = {
  count: number;
  services: Service[];
};
