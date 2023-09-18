import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

import AdminConsumer from '../../api/services/admin/AdminConsumer';
import { Consumer } from '../../api/services/admin/models/Consumer';
import { GetAllConsumersResponse } from '../../api/services/admin/responses/GetAllConsumers';
import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import LoadingText from '../../components/utils/LoadingText';
import { delay } from '../../utils/delay';
import Table from './Table';

function Consumers() {
  const [consumers, setConsumers] = useState<GetAllConsumersResponse | null>(null);

  async function loadConsumers() {
    await delay(500);
    const consumers = await AdminConsumer.getAllConsumers();
    setConsumers(consumers.data);
  }

  useEffect(() => {
    loadConsumers();
  }, []);

  return (
    <Container>
      <Header text="gateway consumers" align="left" size="sm" />

      <div className="flex flex-row items-center justify-between mb-3">
        <h3 className="text-xs -mt-2 text-gray-600 ">gateway consumers</h3>
        <Link to="/consumers/register">
          <button className="flex-end text-xs py-1 px-2 bg-blue-500 text-white shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg">
            <span>add consumer</span>
          </button>
        </Link>
      </div>

      {consumers ? (
        <Table consumers={consumers?.consumers as Consumer[]} />
      ) : (
        <LoadingText text="loading gateway services" />
      )}
    </Container>
  );
}

export default Consumers;
