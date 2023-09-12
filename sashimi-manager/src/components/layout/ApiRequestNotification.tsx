import React from 'react';

import { Request } from '../../api/services/admin/models/Request';
import { ApiMethod } from '../../types/api/ApiMethod.interface';
import { parseDateString } from '../../utils/parseDate';
import ApiMethodTag from '../tags/ApiMethodTag';

interface ApiRequestNotificationProps {
  request: Request;
}

function ApiRequestNotification({ request }: ApiRequestNotificationProps) {
  const bgColorMap = {
    '2xx': {
      bg: 'bg-sashimi-green',
      text: 'text-sashimi-deepgreen'
    },
    '4xx': {
      bg: 'bg-sashimi-pink',
      text: 'text-sashimi-deeppink'
    },
    '5xx': {
      bg: 'bg-sashimi-purple',
      text: 'text-sashimi-deeppurple'
    },
    others: {
      bg: 'bg-sashimi-gray',
      text: 'text-sashimi-deepgray'
    }
  };

  function getColor() {
    if (request.code >= 200 && request.code < 300) {
      return bgColorMap['2xx'];
    } else if (request.code >= 400 && request.code < 500) {
      return bgColorMap['4xx'];
    } else if (request.code >= 500) {
      return bgColorMap['5xx'];
    } else {
      return bgColorMap['others'];
    }
  }

  return (
    <div
      className={`bg-gray-50 animate__animated animate__fadeIn border-sashimi-gray border px-2 py-2.5 rounded-lg text-xs font-sans shadow-md`}
    >
      <div className="flex flex-row items-center justify-between">
        <div className="flex flex-row items-center gap-2">
          <ApiMethodTag method={request.method as ApiMethod} />
          <span className="font-sans tracking-wider">{request.path}</span>
        </div>
        <div className="">
          <span className="tracking-wider">status</span>{' '}
          <span className={`${getColor().text} font-bold`}>{request.code}</span>
        </div>
      </div>
      <div className="text-xs mt-1 text-right font-sans tracking-wider">
        at: <span className="italic text-sashimi-deepgray">{parseDateString(request.time)}</span>
      </div>
    </div>
  );
}

export default ApiRequestNotification;
