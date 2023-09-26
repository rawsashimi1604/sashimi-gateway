import { Consumer } from './Consumer';

export type JwtCredential = {
  id: string;
  key: string;
  secret: string;
  name: string;
  consumer: Consumer;
};
