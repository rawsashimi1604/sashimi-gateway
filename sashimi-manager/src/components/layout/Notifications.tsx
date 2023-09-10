import React, { useEffect, useState } from 'react';
import { AiOutlineLoading3Quarters } from 'react-icons/ai';

import { Request } from '../../api/services/admin/models/Request';
import ApiRequestNotification from './ApiRequestNotification';

function Notifications() {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [requests, setRequests] = useState<Request[]>([]);

  useEffect(() => {
    // Create WebSocket connection.
    const PATH =
      import.meta.env.VITE_BACKEND_DOMAIN +
      import.meta.env.VITE_ADMIN_API_PATH +
      import.meta.env.VITE_WEBSOCKET_API_PATH;
    const websocket = new WebSocket('ws://' + PATH);

    // Connection opened
    websocket.addEventListener('open', (event) => {
      websocket.send(JSON.stringify({ message: 'Hello Server!' }));
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
        <h1 className="font-lora font-light tracking-tight">notifications</h1>
        {ws ? (
          <div className="font-lora flex flex-row border px-2 rounded-lg shadow-sm bg-gray-50 py-0.5 text-sm gap-2 items-center justify-between text-gray-600">
            <span className="font-lora tracking-wider">listening</span>
            <AiOutlineLoading3Quarters className="animate-spin" />
          </div>
        ) : (
          <div className="font-lora flex flex-row border px-2 rounded-lg shadow-sm bg-sashimi-pink py-0.5 text-sm gap-2 items-center justify-between text-sashimi-deeppink">
            <span className="font-lora tracking-wider">not connected</span>
          </div>
        )}
      </div>

      <div className="flex flex-col mt-2 gap-3 overflow-y-scroll grow container">
        {requests?.reverse().map((req) => {
          return <ApiRequestNotification key={req.id} request={req} />;
        })}
      </div>
    </div>
  );
}

export default Notifications;
