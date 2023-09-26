import { Consumer } from '../models/Consumer';
import { JwtCredential } from '../models/JwtCredential';

export type RegisterConsumerResponse = {
  consumer: Consumer;
  credential: JwtCredential;
};
