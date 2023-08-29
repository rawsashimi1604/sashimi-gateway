import React from 'react';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import Table from './Table';

function Services() {
  return (
    <Container>
      <Header text="gateway services" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">gateway services</h3>

      <Table />
    </Container>
  );
}

export default Services;
