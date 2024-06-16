import {SignIn} from "../auth/SignIn.jsx";
import {SignUp} from "../auth/SignUp.jsx";
import {useModal} from "../../hooks/useModal.js";
import {AddWalletModal} from "../mainApp/wallets/WalletModal.jsx";

export function ModalManager() {
    const {modal} = useModal();
    return (
        <>
            <SignIn open={modal === 'signIn'} />
            <SignUp open={modal === 'signUp'} />
        </>
    )
}