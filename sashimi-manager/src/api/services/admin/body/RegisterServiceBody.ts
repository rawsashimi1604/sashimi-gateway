export type RegisterServiceBody = {
  name: string;
  targetUrl: string;
  path: string;
  description: string;
  healthCheckEnabled: boolean;
};
