import 'animate.css';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Account from './modules/account';
import Consumer from './modules/consumer';
import Dashboard from './modules/dashboard';
import Logs from './modules/gateway-logs';
import Login from './modules/login';
import RegisterConsumer from './modules/register-consumer';
import RegisterRoute from './modules/register-route';
import RegisterService from './modules/register-service';
import Routes from './modules/routes';
import ServiceInformation from './modules/service-details';
import Services from './modules/services';
import Settings from './modules/settings';
import Test from './modules/test';

function App() {
  const router = createBrowserRouter([
    {
      path: '/login',
      element: <Login />
    },
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
      path: '/consumers',
      element: <Consumer />
    },
    {
      path: '/consumers/register',
      element: <RegisterConsumer />
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
    },
    {
      path: '/test',
      element: <Test />
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
