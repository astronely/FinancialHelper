import styles from "./WalletCard.module.scss";
import './WalletCard.scss'
import {CircleX, Pen} from 'lucide-react'
import {useModal} from "../../../../hooks/useModal.js";

export function WalletCard({wallet, confirmDelete, setWalletToDelete, changeAction}) {
    const  {setIsActive, setModal} = useModal();
    // console.log(wallet)
    return (
        <div className={styles.wallet__card}>
            <div className={"wallet__info"}>
                <div className={'wallet__text'}>
                    {wallet.name}
                    <button onClick={() => changeAction(wallet)} className={"wallet__button wallet__pen"}><Pen/></button>
                </div>
                <div className={'wallet__text'}>{wallet.value} {wallet.currency}</div>
            </div>
            <button onClick={() => confirmDelete(setIsActive, setModal, 'confirmDeleteWallet', wallet.name, setWalletToDelete)} className={"wallet__button wallet__cross"}><CircleX/></button>
        </div>
    )
}