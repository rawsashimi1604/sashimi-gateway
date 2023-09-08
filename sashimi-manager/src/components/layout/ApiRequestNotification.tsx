import React from 'react';

import { Request } from '../../api/services/admin/models/Request';
import { ApiMethod } from '../../types/api/ApiMethod.interface';
import { parseDateString } from '../../utils/parseDate';
import ApiMethodTag from '../tags/ApiMethodTag';

interface ApiRequestNotificationProps {
  request: Request;
}

function ApiRequestNotification({ request }: ApiRequestNotificationProps) {
  const colorMap = {
    '2xx': 'bg-sashimi-green',
    '4xx': 'bg-sashimi-pink',
    '5xx': 'bg-sashimi-purple',
    others: 'bg-sashimi-gray'
  };

  function getColor() {
    if (request.code >= 200 && request.code < 300) {
      return colorMap['2xx'];
    } else if (request.code >= 400 && request.code < 500) {
      return colorMap['4xx'];
    } else if (request.code >= 500) {
      return colorMap['5xx'];
    } else {
      return colorMap['others'];
    }
  }

  return (
    <div
      className={`${getColor()} animate__animated animate__fadeIn border-sashimi-gray border px-2 py-2.5 rounded-lg text-xs font-sans shadow-md`}
    >
      <div className="flex flex-row items-center justify-between">
        <div className="flex flex-row items-center gap-2">
          <ApiMethodTag method={request.method as ApiMethod} />
          <span className="font-sans tracking-wider">{request.path}</span>
        </div>
        <div className="font-bold">CODE {request.code}</div>
      </div>
      <div className="text-xs mt-1 text-right font-sans tracking-wider">
        at: <span className="italic text-sashimi-deepgray">{parseDateString(request.time)}</span>
      </div>
    </div>
  );
}

export default ApiRequestNotification;
