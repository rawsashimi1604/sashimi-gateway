import { Route } from './Route';

export type Service = {
  id: number;
  name: string;
  targetUrl: string;
  path: string;
  description: string;
  routes: Route[];
  createdAt: string;
  updatedAt: string;
  healthCheckEnabled: boolean;
  health: string;
};
