import { Route } from '../models/Route';

export type GetAllRoutesResponse = {
  count: number;
  routes: Route[];
};
