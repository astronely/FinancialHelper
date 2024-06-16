import React from "react";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {Landing} from "./Landing.jsx";
import {ProfileLoader} from "./loaders/ProfileLoader.jsx";
import {Profile} from "./Profile.jsx";
import {ToastContainer} from "react-toastify";
import 'react-toastify/dist/ReactToastify.min.css'
import {ModalManager} from "./components/modal/ModalManager.jsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Landing />
    },
    {
        path: '/profile',
        element: <Profile />,
        loader: ProfileLoader
    }
]);

export function App() {
    return (
        <>
            <RouterProvider router={router} />
            <ToastContainer position="top-center" autoClose={1000} hideProgressBar={false} newestOnTop={false}
                            closeOnClick
                            pauseOnFocusLoss={false}
                            draggable
                            pauseOnHover={false}
                            theme="colored"
            />
        </>
    )
}