import React from 'react';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';

function Settings() {
  return (
    <Container>
      <Header text="settings" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">settings</h3>
    </Container>
  );
}

export default Settings;
