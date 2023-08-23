import { RouterProvider, createBrowserRouter } from "react-router-dom";

function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <div>Hello world!</div>,
    },
    {
      path: "/about",
      element: <div>About</div>,
    },
  ]);

  return (
    <RouterProvider router={router} />
  )
}

export default App;
