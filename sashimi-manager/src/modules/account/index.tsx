import React from 'react';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';

function Account() {
  return (
    <Container>
      <Header text="account" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">account</h3>
    </Container>
  );
}

export default Account;
