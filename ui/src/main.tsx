import React from 'react'
import ReactDOM from 'react-dom/client'
import { ErrorBoundary } from "./pages";
import { ApiProvider } from "./providers";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { routes } from "./routes";

import { BaseStyles, ThemeProvider } from '@primer/react';
import "./components/editor/setup"

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const router = createBrowserRouter(routes, {
  // https://reactrouter.com/en/main/guides/api-development-strategy#current-future-flags
  future: { v7_normalizeFormMethod: true },
})


root.render(
  <React.StrictMode>
    <ThemeProvider>
      <ApiProvider>
        <BaseStyles>
          <ErrorBoundary>
            <RouterProvider router={router} />
          </ErrorBoundary>
        </BaseStyles>
      </ApiProvider>
    </ThemeProvider>
  </React.StrictMode>,
)
