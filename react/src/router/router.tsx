import { createBrowserRouter, Navigate } from "react-router";
import { Home } from "@/pages/Home";
import { PageSearch } from "@/pages/Search";
import { PageNotFound } from "@/pages/PageNotFound";
import { Playground } from "@/pages/Playground";

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
    path: "/my-web-demo/search",
    element: <PageSearch />,
  },
  {
    path: "/my-web-demo/playground",
    element: <Playground />,
  },
  {
    path: "*",
    element: <PageNotFound />,
  },
]);
