import React from 'react'
import ReactDOM from 'react-dom/client'
import 'bootstrap/dist/css/bootstrap.min.css'
import './index.scss'

import {App} from "./App.jsx";
import {Landing} from "./Landing.jsx";
import {ModalProvider} from "./context/ModalContext.jsx";
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import {ProfileLoader} from "./loaders/ProfileLoader.jsx";
import 'react-toastify/dist/ReactToastify.css';
import {AppProvider} from "./context/AppContext.jsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Landing />
    },
    {
        path: '/app',
        element: <App />,
        loader: ProfileLoader
    }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
      <AppProvider>
          <ModalProvider>
              <App />
          </ModalProvider>
      </AppProvider>
  </React.StrictMode>,
)
