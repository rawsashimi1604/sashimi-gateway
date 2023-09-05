// TODO: define type
export type GetGatewayInformationResponse = {
  gateway: GatewayConfig;
};

export type GatewayConfig = {
  gatewayName: string;
  hostName: string;
  dateCreated: string;
  tagLine: string;
  port: string;
};
