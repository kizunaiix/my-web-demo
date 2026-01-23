import { createBrowserRouter, Navigate } from "react-router";

import { HomePage } from "@/pages/HomePage";
import { SearchPage } from "@/pages/SearchPage";
import { NotFoundPage } from "@/pages/NotFoundPage";
import { PlaygroundPage } from "@/pages/PlaygroundPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Navigate to="/my-web-demo" />,
  },
  {
    path: "/my-web-demo",
    element: <HomePage />,
  },
  {
    path: "/my-web-demo/search",
    element: <SearchPage />,
  },
  {
    path: "/my-web-demo/playground",
    element: <PlaygroundPage />,
  },
  {
    path: "*",
    element: <NotFoundPage />,
  },
]);

export { router };
