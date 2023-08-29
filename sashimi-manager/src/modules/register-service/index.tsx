import React from 'react';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import Form from './Form';

function RegisterService() {
  return (
    <Container>
      <Header text="register service" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">register service</h3>

      <Form />
    </Container>
  );
}

export default RegisterService;
