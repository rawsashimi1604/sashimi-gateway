import { Request } from './Request';

export type GetAllRequestsResponse = {
  count: number;
  requests: Request[];
};
