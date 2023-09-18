import { Consumer } from '../models/Consumer';

export type GetAllConsumersResponse = {
  count: number;
  consumers: Consumer[];
};
