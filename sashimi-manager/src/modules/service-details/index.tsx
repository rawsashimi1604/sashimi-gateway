import React from 'react';
import { useParams } from 'react-router-dom';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';

function ServiceDetails() {
  const { id } = useParams();

  return (
    <Container>
      <Header text="service details" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">service details for id: {id} </h3>
    </Container>
  );
}

export default ServiceDetails;
