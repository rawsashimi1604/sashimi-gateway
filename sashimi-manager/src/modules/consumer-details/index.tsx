import React from 'react';
import { useParams } from 'react-router-dom';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';

function ConsumerDetails() {
  const { id } = useParams();

  return (
    <Container>
      <Header text="consumer details" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">consumer details</h3>
      <h2>Id : {id}</h2>
    </Container>
  );
}

export default ConsumerDetails;
