import 'animate.css';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Dashboard from './modules/dashboard';
import Logs from './modules/logs';
import Routes from './modules/routes';
import Services from './modules/services';

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
