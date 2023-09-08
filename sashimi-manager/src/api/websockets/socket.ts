import { io } from 'socket.io-client';

const URL = import.meta.env.VITE_BACKEND_URL + import.meta.env.VITE_WEBSOCKET_API_PATH;

export const socket = io(URL, {
  autoConnect: false,
  withCredentials: true
});
