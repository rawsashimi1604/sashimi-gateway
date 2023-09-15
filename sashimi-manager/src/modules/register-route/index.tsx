import React from 'react';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import Form from './Form';

function RegisterRoute() {
  return (
    <Container>
      <Header text="register route" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">register route</h3>

      <Form />
    </Container>
  );
}

export default RegisterRoute;
