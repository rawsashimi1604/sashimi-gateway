import 'animate.css';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Dashboard from './modules/dashboard';

function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Dashboard />
    },
    {
      path: '/about',
      element: <div>About</div>
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
