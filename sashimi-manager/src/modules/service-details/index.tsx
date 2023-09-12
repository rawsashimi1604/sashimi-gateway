import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import AdminService from '../../api/services/admin/AdminService';
import { GetServiceByIdResponse } from '../../api/services/admin/responses/GetServiceById';
import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import LoadingText from '../../components/utils/LoadingText';
import { delay } from '../../utils/delay';
import ServiceInformation from './ServiceInformation';

function ServiceDetails() {
  const { id } = useParams();
  const [service, setService] = useState<GetServiceByIdResponse | null>(null);

  async function loadServiceData() {
    if (!id) return;
    await delay(500);
    const serviceData = await AdminService.getServiceById(id);
    console.log({ serviceData });
    setService(serviceData.data);
  }

  useEffect(() => {
    loadServiceData();
  }, []);

  return (
    <Container>
      <Header text={service?.service.name || 'loading...'} align="left" size="sm" />
      <h3 className="text-xs -mt-2 text-gray-600 mb-4">{service?.service.description || 'loading...'}</h3>
      {service ? <ServiceInformation data={service} /> : <LoadingText text="loading service details..." />}
    </Container>
  );
}

export default ServiceDetails;
