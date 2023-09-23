import { Service } from './Service';

export type Consumer = {
  id: string;
  username: string;
  createdAt: string;
  updatedAt: string;
  services: Service[];
};
