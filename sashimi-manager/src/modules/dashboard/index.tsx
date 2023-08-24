import React from 'react';

import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import Card from './Card';

function Dashboard() {
  return (
    <Container>
      <Header text="home" align="left" size="sm" />

      {/* Analytics (Requests, Services, Routes, Data transmitted) */}
      <section className="grid grid-cols-2 lg:grid-cols-4 gap-2">
        <Card header="total requests" data="30" />
        <Card header="services" data="4" />
        <Card header="routes" data="32" />
        <Card header="data transmitted" data="4,096MB" />
      </section>
    </Container>
  );
}

export default Dashboard;