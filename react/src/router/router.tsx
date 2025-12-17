import { createBrowserRouter, Navigate } from "react-router";

import { Home } from "../pages/Home";
import { PageNotFound } from "../pages/PageNotFound";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Navigate to="/my-web-demo" />,
  },
  {
    path: "/my-web-demo",
    element: <Home />,
  },
  {
    path: "*",
    element: <PageNotFound />,
  },
]);
