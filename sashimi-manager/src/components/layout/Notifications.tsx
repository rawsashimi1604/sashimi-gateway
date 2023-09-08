import React, { useEffect, useState } from 'react';
import { AiOutlineLoading3Quarters } from 'react-icons/ai';

import { Request } from '../../api/services/admin/models/Request';
import ApiRequestNotification from './ApiRequestNotification';

function Notifications() {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);
  const [requests, setRequests] = useState<Request[]>([]);

  useEffect(() => {
    // Create WebSocket connection.
    const websocket = new WebSocket('ws://localhost:8080/api/admin/ws');

    // Connection opened
    websocket.addEventListener('open', (event) => {
      websocket.send(JSON.stringify({ message: 'Hello Server!' }));
      setIsConnected(true);
    });

    // Listen for messages
    websocket.addEventListener('message', (event) => {
      const parsed = JSON.parse(event.data);

      if (parsed.requests) {
        const requests: Request[] = [];
        for (const req of parsed.requests) {
          const request: Request = {
            id: req.id,
            serviceId: req.serviceId,
            routeId: req.routeId,
            path: req.path,
            method: req.method,
            time: req.time,
            code: req.code
          };
          requests.push(request);
        }
        setRequests((prev) => [...prev, ...requests]);
      }
    });

    setWs(websocket);

    return () => {
      websocket.close();
    };
  }, []);

  return (
    <div className="p-4 pt-5 flex flex-col relative z-0 w-full h-full">
      {/* border */}
      <div className="absolute left-0 border-l border-gray-200 h-full w-full -z-10"></div>
      <div className="flex flex-row items-center justify-between">
        <h1 className="font-cabin font-light tracking-tight">notifications</h1>
        {isConnected ? (
          <div className="font-cabin flex flex-row border px-2 rounded-lg shadow-sm bg-gray-50 py-0.5 text-sm gap-2 items-center justify-between text-gray-600">
            <span className="font-cabin tracking-wider">listening</span>
            <AiOutlineLoading3Quarters className="animate-spin" />
          </div>
        ) : (
          <div className="font-cabin flex flex-row border px-2 rounded-lg shadow-sm bg-sashimi-pink py-0.5 text-sm gap-2 items-center justify-between text-sashimi-deeppink">
            <span className="font-cabin tracking-wider">not connected</span>
          </div>
        )}
      </div>

      <div className="flex flex-col mt-2 gap-3 overflow-y-scroll grow container">
        {requests?.map((req) => {
          return <ApiRequestNotification key={req.id} request={req} />;
        })}
      </div>
    </div>
  );
}

export default Notifications;
