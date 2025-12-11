import { createBrowserRouter } from "react-router";

import Home from "../pages/home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <div>Hello World</div>,
  },
  {
    path: "/my-web-demo/static/vite",
    element: <Home />,
  },
  {
    path: "*",
    element: <div>404</div>,
  },

]);

export default router;

