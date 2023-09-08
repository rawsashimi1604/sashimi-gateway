import React from 'react';

import { ApiMethod } from '../../types/api/ApiMethod.interface';

interface ApiMethodTagProps {
  method: ApiMethod;
}

function ApiMethodTag({ method }: ApiMethodTagProps) {
  const color: any = {
    GET: 'bg-sashimi-deepgreen',
    POST: 'bg-sashimi-deepyellow'
  };

  const selectedColor = color[method as string];

  return (
    <div className={`inline-block text-white px-1.5 py-1 rounded-lg text-xs shadow-sm ${selectedColor}`}>{method}</div>
  );
}

export default ApiMethodTag;
