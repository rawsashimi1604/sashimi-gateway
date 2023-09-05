import { Route } from './Route';

export type Service = {
  id: number;
  name: string;
  targetUrl: string;
  path: string;
  description: string;
  routes: Route[];
};
