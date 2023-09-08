import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

import AdminRoute from '../../api/services/admin/AdminRoute';
import { Route } from '../../api/services/admin/models/Route';
import { GetAllRoutesResponse } from '../../api/services/admin/responses/GetAllRoutes';
import Container from '../../components/layout/Container';
import Header from '../../components/typography/Header';
import LoadingText from '../../components/utils/LoadingText';
import { delay } from '../../utils/delay';
import Table from './Table';

function Routes() {
  const [routes, setRoutes] = useState<GetAllRoutesResponse | null>(null);

  async function loadRoutes() {
    await delay(500);
    const routes = await AdminRoute.getAllRoutes();
    setRoutes(routes.data);
  }

  useEffect(() => {
    loadRoutes();
  }, []);

  return (
    <Container>
      <Header text="gateway routes" align="left" size="sm" />
      <div className="flex flex-row items-center justify-between mb-3">
        <h3 className="text-xs -mt-2 text-gray-600 ">gateway routes</h3>
        <Link to="/routes/register">
          <button className="flex-end text-xs py-1 px-2 bg-blue-500 text-white shadow-md rounded-lg font-sans border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg">
            <span>add route</span>
          </button>
        </Link>
      </div>

      {routes ? <Table routes={routes?.routes as Route[]} /> : <LoadingText text="loading gateway routes" />}
    </Container>
  );
}

export default Routes;
