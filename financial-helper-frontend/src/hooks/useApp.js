import {AppContext} from "../context/AppContext.jsx";
import {useContext} from "react";

export function useApp() {
    return useContext(AppContext);
}