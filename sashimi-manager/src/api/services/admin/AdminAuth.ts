import HttpRequest from '../../requests/HttpRequest';
import type { LoginResponse } from './responses/LoginResponse';

function login(username: string, password: string) {
  return HttpRequest.post<LoginResponse>(`/login`, {
    username,
    password
  });
}

export default {
  login
};
