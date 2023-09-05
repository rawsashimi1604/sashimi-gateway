import { Service } from './Service';

export type GetAllServicesResponse = {
  count: number;
  services: Service[];
};
