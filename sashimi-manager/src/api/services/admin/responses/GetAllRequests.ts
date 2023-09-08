import { Request } from '../models/Request';

export type GetAllRequestsResponse = {
  count: number;
  requests: Request[];
};
