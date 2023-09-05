import React, { useEffect, useState } from 'react';

import AdminGateway from '../../api/services/admin/AdminGateway';
import AdminService from '../../api/services/admin/AdminService';
import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import Card from './Card';
import Information from './Information';

function Dashboard() {
  async function loadGatewayInformation() {
    const data = await AdminGateway.getGatewayinformation();
    console.log(data);
  }
  async function loadApiRequest() {
    const data = await AdminService.getAllServices();
    console.log(data);
  }

  useEffect(() => {
    loadGatewayInformation();
    loadApiRequest();
  }, []);

  return (
    <Container>
      <Header text="welcome to sashimi gateway" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">
        welcome to sashimi gateway
      </h3>

      {/* Analytics (Requests, Services, Routes, Data transmitted) */}
      <section className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        <Card header="total requests" data="30" />
        <Card header="services" data="4" />
        <Card header="routes" data="32" />
        <Card header="data transmitted" data="4,096MB" />
      </section>

      {/* Graph */}
      <Information />
    </Container>
  );
}

export default Dashboard;
