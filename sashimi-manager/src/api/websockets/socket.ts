import socket from 'socket.io-client';

const HOST = import.meta.env.VITE_BACKEND_URL;
const PATH = import.meta.env.VITE_ADMIN_API_PATH + import.meta.env.VITE_WEBSOCKET_API_PATH;

export const socketManager = new socket.Manager(HOST, {
  path: PATH,
  transports: ['websocket']
});
