import styles from "./HistoryCard.module.scss";
import './HistoryCard.scss'
import {CircleX} from 'lucide-react'
import {useModal} from "../../../../hooks/useModal.js";

export function HistoryCard({expense, confirmDelete, setExpenseToDelete}) {
    const  {setIsActive, setModal} = useModal();

    return (
        <div className={styles.history__card}>
            <div className={"history-card__header"}>
                <div className={"history-card__primary-text history-card__operation-date"}>{expense.date}</div>
                {/*<button onClick={() => closeAction(expense)} className={"wallet__button"}><CircleX/></button>*/}
                <button
                    onClick={() => confirmDelete(setIsActive, setModal, 'confirmDeleteExpense', expense, setExpenseToDelete)}
                    className={"wallet__button wallet__cross"}><CircleX/></button>
            </div>
            <div className={"history-card__container"}>
                <div className={"history-card__info"}>
                    <div className={"history-card__primary-text"}>{expense.category}</div>
                    <div className={"history-card__primary-text history-card__shop-name"}>{expense.shop_name}</div>
                </div>
                <div className={"history-card__price"}>
                    <div className={"history-card__primary-text"}>{expense.price} {expense.currency}</div>
                    <div className={"history-card__primary-text history-card__wallet-name"}>{expense.wallet_name}</div>
                </div>
            </div>
        </div>
    )
}