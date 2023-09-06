import React from 'react';

import Header from '../../components/typography/Header';
import RequestChart from './RequestChart';

function Information() {
  return (
    <div className="mt-6">
      <Header text="api requests" size="sm" align="left" />
      <RequestChart />
    </div>
  );
}

export default Information;
