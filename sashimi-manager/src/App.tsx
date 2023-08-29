import 'animate.css';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Account from './modules/account';
import Dashboard from './modules/dashboard';
import Logs from './modules/gateway-logs';
import Routes from './modules/routes';
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
      path: '/routes',
      element: <Routes />
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
