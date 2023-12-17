import { createRoot } from 'react-dom/client';
import { Provider } from "react-redux";
import React from 'react';
import ApiList from './ApiList';
import { store } from './store';

const root = createRoot(document.getElementById('root'));

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <ApiList />
    </Provider>
  </React.StrictMode>
);