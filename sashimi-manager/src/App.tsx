import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import Footer from './components/layout/Footer';
import Navbar from './components/layout/Navbar';

function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <div>Hello world!</div>
    },
    {
      path: '/about',
      element: <div>About</div>
    }
  ]);

  return (
    <>
      <Navbar />
      <RouterProvider router={router} />
      <Footer />
    </>
  );
}

export default App;
