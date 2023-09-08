import React, { useEffect, useState } from 'react';
import { AiOutlineLoading3Quarters } from 'react-icons/ai';

import { Request } from '../../api/services/admin/models/Request';
import { socketManager } from '../../api/websockets/socket';

function Notifications() {
  const [isConnected, setIsConnected] = useState(false);
  const [requests, setRequests] = useState<Request[]>([]);

  useEffect(() => {
    function onConnect() {
      console.log('connected via websockets!');
      setIsConnected(true);
    }

    function onDisconnect() {
      console.log('disconnected via websockets!');
      setIsConnected(false);
    }

    function onReceiveNotification(requests: Request[]) {
      console.log(requests);
      setRequests((prev) => [...prev, ...requests]);
    }

    socketManager.on('connect', onConnect);
    socketManager.on('disconnect', onDisconnect);
    socketManager.on('event:apiRequests', onReceiveNotification);

    // Cleanup on unmount
    return () => {
      socketManager.off('connect', onConnect);
      socketManager.off('disconnect', onDisconnect);
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
    </div>
  );
}

export default Notifications;
