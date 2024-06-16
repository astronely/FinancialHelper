import {LandingHeader} from "./components/header/LandingHeader.jsx";
import {Information} from "./components/information/Information.jsx";
import {ModalManager} from "./components/modal/ModalManager.jsx";
import {ToastContainer} from "react-toastify";
import React from "react";

export function Landing() {
    return (
        <>
            <LandingHeader />
            <main>
                <Information />
            </main>
            <ModalManager />
        </>
    )
}