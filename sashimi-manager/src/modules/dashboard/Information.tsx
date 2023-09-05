import React from 'react';

import { GetAllRequestsResponse } from '../../api/services/admin/responses/GetAllRequests';
import { Request } from '../../api/services/admin/responses/Request';
import Header from '../../components/typography/Header';
import LoadingText from '../../components/utils/LoadingText';
import RequestChart from './RequestChart';

interface InformationProps {
  requests: Request[];
}

function Information({ requests }: InformationProps) {
  return (
    <div className="mt-6">
      <Header text="api requests" size="sm" align="left" />
      {requests ? (
        <RequestChart requests={requests} />
      ) : (
        <LoadingText text="loading request chart..." />
      )}
    </div>
  );
}

export default Information;
