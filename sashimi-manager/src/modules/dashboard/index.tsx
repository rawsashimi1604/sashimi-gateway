import React, { useEffect, useState } from 'react';

import AdminRequest from '../../api/services/admin/AdminRequest';
import AdminRoute from '../../api/services/admin/AdminRoute';
import AdminService from '../../api/services/admin/AdminService';
import { GetAllRequestsResponse } from '../../api/services/admin/responses/GetAllRequests';
import { GetAllRoutesResponse } from '../../api/services/admin/responses/GetAllRoutes';
import { GetAllServicesResponse } from '../../api/services/admin/responses/GetAllServices';
import { Request } from '../../api/services/admin/responses/Request';
import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import { delay } from '../../utils/delay';
import Card from './Card';
import Information from './Information';
import LoadingCard from './LoadingCard';

type DashboardRequestData = {
  requests: GetAllRequestsResponse;
  services: GetAllServicesResponse;
  routes: GetAllRoutesResponse;
};

function Dashboard() {
  const [data, setData] = useState<DashboardRequestData | null>(null);

  async function loadDashboardRequestData() {
    await delay(2000);
    const services = await AdminService.getAllServices();
    const requests = await AdminRequest.getAllRequests();
    const routes = await AdminRoute.getAllRoutes();

    setData((prev) => {
      return {
        ...prev,
        services: services.data,
        requests: requests.data,
        routes: routes.data
      };
    });
  }

  useEffect(() => {
    loadDashboardRequestData();
  }, []);

  return (
    <Container>
      <Header text="welcome to sashimi gateway" align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">welcome to sashimi gateway</h3>

      {/* Analytics (Requests, Services, Routes, Data transmitted) */}
      <section className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        {data?.requests ? (
          <Card header="total requests" data={data.requests.count.toString()} />
        ) : (
          <LoadingCard header="total requests" />
        )}

        {data?.services ? (
          <Card header="services" data={data.services.count.toString()} />
        ) : (
          <LoadingCard header="services" />
        )}

        {data?.services ? (
          <Card header="routes" data={data.routes.count.toString()} />
        ) : (
          <LoadingCard header="routes" />
        )}

        {data?.services ? <Card header="data transmitted" data="4,096MB" /> : <LoadingCard header="data transmitted" />}
      </section>

      {/* Graph */}
      <Information />
    </Container>
  );
}

export default Dashboard;
