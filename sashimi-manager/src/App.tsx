import 'animate.css';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Account from './modules/account';
import Dashboard from './modules/dashboard';
import Logs from './modules/gateway-logs';
import RegisterRoute from './modules/register-route';
import RegisterService from './modules/register-service';
import Routes from './modules/routes';
import ServiceInformation from './modules/service-details';
import Services from './modules/services';
import Settings from './modules/settings';

function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Dashboard />
    },
    {
      path: '/services',
      element: <Services />
    },
    {
      path: '/services/:id',
      element: <ServiceInformation />
    },
    {
      path: '/services/register',
      element: <RegisterService />
    },
    {
      path: '/routes',
      element: <Routes />
    },
    {
      path: '/routes/register',
      element: <RegisterRoute />
    },
    {
      path: '/logs',
      element: <Logs />
    },
    {
      path: '/account',
      element: <Account />
    },
    {
      path: '/settings',
      element: <Settings />
    }
  ]);

  return (
    <>
      <div className="bg-white font-sans">
        <RouterProvider router={router} />
      </div>
    </>
  );
}

export default App;
