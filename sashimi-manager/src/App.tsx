import 'animate.css';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Dashboard from './modules/dashboard';
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
