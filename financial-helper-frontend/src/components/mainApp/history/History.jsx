import {InfoColumn} from "../infoColumn/InfoColumn.jsx";
import './History.scss'
import {HistoryCard} from "../cards/HistoryCards/HistoryCard.jsx";
import {useEffect, useState} from "react";
import {useModal} from "../../../hooks/useModal.js";
import axios from "axios";
import {AddExpenseModal} from "./ExpenseModal.jsx";
import {useApp} from "../../../hooks/useApp.js";
import {toast} from "react-toastify";
import "../../../utils/modalUtils.js"
import {openConfirm, openModal} from "../../../utils/modalUtils.js";
import {ConfirmItemDeleteModal} from "../../confirmation/Confirm.jsx";

export function History() {
    const {wallets, updateWallets} = useApp()
    const [expenses, setExpenses] = useState([]);
    const {setIsActive, setModal, modal} = useModal();
    const [expenseToDelete, setExpenseToDelete] = useState('');

    const getExpense = expense => {
        // console.log("getExpense: ", expense)
        setExpenses(expenses => [...expenses, {
            id: expense.ID,
            wallet: expense.Wallet,
            wallet_name: expense.WalletName,
            currency: expense.Currency,
            shop_name: expense.Name,
            category: expense.Category,
            price: expense.Value,
            date: new Date(expense.Date).toISOString().split("T")[0],
        }])
    }

    const getExpenses = async () => {
        await axios.get("http://localhost:8080/api/expenses/get_expenses", {withCredentials: true})
            .then(response => {
                setExpenses([])
                // console.log("getExpenses: ", expenses)
                if (response.data.sort(function(a,b) {
                    if (a.Date < b.Date) return 1;
                    if (a.Date > b.Date) return -1;
                    if (a.CreatedAt < b.CreatedAt) return 1;
                    if (a.CreatedAt >= b.CreatedAt) return -1;
                    }))
                {
                    for (let expense of response.data) {
                        getExpense(expense)
                    }
                }
            })
            .catch(error => console.log(error))
    }

    const increaseWalletBalance = async data => {
        await axios.put("http://localhost:8080/api/wallets/increase_balance", data,{withCredentials: true})
            .then(() => {
                updateWallets()
            })
            .catch(error => {
                toast.error("Cannot update wallet balance")
                console.log(error.response)
            })
    }

    async function deleteExpense(expense){
        console.log(expense)
        setExpenses(w => w.filter(item => item.id !== expense.id))
        await axios.delete("http://localhost:8080/api/expenses/delete",{
            data: {
                id: expense.id,
            },
            withCredentials: true},)
            .then(() => {
                increaseWalletBalance({
                    name: expense.wallet_name,
                    value: expense.price.toString(),
                    currency: expense.currency
                })
                toast.success("Successfully delete expense")
            })
            .catch(error => {
                console.log(error.response)
                toast.error("Expense not delete")
            })
    }

    useEffect(() => {
        const timeout = setTimeout(() => {
            // console.log("Expenses useEffect")
            getExpenses().then()
        }, 10)

        return () => clearTimeout(timeout)
    }, [wallets])

    return (
        <InfoColumn>
            <div className={'column__title'}>Совершенные операции</div>
            <div className={'history'}>
                <div className={'history__header'}>
                    <div>Все операции</div>
                    {/*<img src="icons/list.svg"/>*/}
                </div>
                <div className={'history__cards'}>
                    {expenses.map((item, index) => (
                        <HistoryCard expense={item} confirmDelete={openConfirm} setExpenseToDelete={setExpenseToDelete} key={index}/>
                    ))}
                </div>
            </div>
            <button onClick={() => openModal(setIsActive, setModal, 'addExpense')} className={"primary-button"}>Добавить</button>
            <AddExpenseModal open={modal === 'addExpense'} addExpense={getExpense}/>
            <ConfirmItemDeleteModal open={modal === 'confirmDeleteExpense'} item={expenseToDelete} deleteAction={deleteExpense}/>
        </InfoColumn>
    )
}