import React, { useEffect, useState } from 'react';

import AdminGateway from '../../api/services/admin/AdminGateway';
import AdminService from '../../api/services/admin/AdminService';
import { GetAllServicesResponse } from '../../api/services/admin/responses/GetAllServices';
import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import LoadingSpinner from '../../components/utils/LoadingSpinner';
import { delay } from '../../utils/delay';
import Card from './Card';
import Information from './Information';
import LoadingCard from './LoadingCard';

type DashboardRequestData = {
  services: GetAllServicesResponse;
};

function Dashboard() {
  const [data, setData] = useState<DashboardRequestData | null>(null);

  async function loadDashboardRequestData() {
    await delay(5000);
    const services = await AdminService.getAllServices();
    setData((prev) => {
      return { ...prev, services: services.data };
    });
  }

  useEffect(() => {
    loadDashboardRequestData();
  }, []);

  useEffect(() => {
    console.log(data);
  }, [data]);

  return (
    <Container>
      <Header text="welcome to sashimi gateway" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">
        welcome to sashimi gateway
      </h3>

      {/* Analytics (Requests, Services, Routes, Data transmitted) */}
      <section className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        {data?.services ? (
          <Card header="total requests" data="30" />
        ) : (
          <LoadingCard header="total requests" />
        )}

        {data?.services ? (
          <Card header="services" data="4" />
        ) : (
          <LoadingCard header="services" />
        )}

        {data?.services ? (
          <Card header="routes" data="32" />
        ) : (
          <LoadingCard header="routes" />
        )}

        {data?.services ? (
          <Card header="data transmitted" data="4,096MB" />
        ) : (
          <LoadingCard header="data transmitted" />
        )}
      </section>

      {/* Graph */}
      <Information />
    </Container>
  );
}

export default Dashboard;
