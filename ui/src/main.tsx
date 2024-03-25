import React from 'react'
import ReactDOM from 'react-dom/client'
import { Main } from "./pages";
import { BaseStyles, ThemeProvider } from '@primer/react';
import "./components/editor/setup"
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider>
      <BaseStyles>
        <Main />
      </BaseStyles>
    </ThemeProvider>
  </React.StrictMode>,
)
