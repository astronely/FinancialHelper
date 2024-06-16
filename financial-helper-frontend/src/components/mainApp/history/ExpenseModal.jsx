import {Modal} from '../../modal/Modal.jsx'
import Select from "react-select"
import '../../modal/Modal.scss'
import '../../buttons/buttons.scss'
import {useForm} from "react-hook-form";
import {toast} from "react-toastify";
import axios from "axios";
import DatePicker from "react-datepicker";
import 'react-datepicker/dist/react-datepicker.min.css'
import {useState} from "react";
import {useApp} from "../../../hooks/useApp.js";
import {isPriceCorrect} from "../../../utils/modalUtils.js";

export function AddExpenseModal({open = false, addExpense}) {
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [selectedCategory, setSelectedCategory] = useState('Супермаркеты');
    const [selectedWallet, setSelectedWallet] = useState({});
    const {wallets, updateWallets} = useApp();

    const categories = [
        {
            value: 'Супермаркеты',
            label: 'Супермаркеты'
        }, {
            value: 'Развлечение',
            label: 'Развлечение'
        }, {
            value: 'Спорт',
            label: 'Спорт'
        }, {
            value: 'Красота',
            label: 'Красота',
        }, {
            value: 'Медицина',
            label: 'Медицина'
        }, {
            value: 'Фастфуд',
            label: 'Фастфуд'
        }, {
            value: 'Рестораны',
            label: 'Рестораны'
        }, {
            value: 'Другое',
            label: 'Другое'
        }]
    let walletsMenu = []

    const updateWalletsMenu = () => {
        walletsMenu = [];
        for (let wallet of wallets) {
            walletsMenu.push({
                label: wallet.name,
                value: wallet.id,
            })
        }
    }

    updateWalletsMenu()

    const {register, handleSubmit, trigger, reset, formState: {errors}} = useForm({
        defaultValues: {
            wallet: '',
            shop_name: '',
            category: '',
            price: '',
            date: '',
        }
    })

    const decreaseWalletBalance = async data => {
        await axios.put('http://localhost:8080/api/wallets/decrease_balance', data, {withCredentials: true})
            .then(() => {
                updateWallets()
            })
            .catch(error => console.log(error.response))
    }

    function isExpenseDataCorrect(data) {
        return !(data.wallet === '' || data.shop_name === '' || data.category === '' || data.price === '' || data.date === '');
    }

    function isDateCorrect() {
        return !!selectedDate
    }

    function isWalletSelected(selectedWallet) {
        return selectedWallet.value !== undefined
    }

    function completeData(data, currentWallet, selectedWallet, selectedCategory, selectedDate) {
        data['wallet'] = currentWallet.id.toString();
        data['wallet_name'] = selectedWallet.label;
        data['category'] = selectedCategory;
        data['wallet_currency'] = currentWallet.currency
        data['date'] = selectedDate.toISOString().split('T')[0];
    }

    const submitHandler = async data => {
        if (!isWalletSelected(selectedWallet)) {
            toast.error("Выберите кошелек")
            return
        }

        if (!isDateCorrect()) {
            toast.error("Выберите дату")
            return
        }
        const currentWallet = wallets.find(w => w.id === selectedWallet.value)
        completeData(data, currentWallet, selectedWallet, selectedCategory, selectedDate)

        if (!isExpenseDataCorrect(data)) {
            toast.error("Заполните все поля")
            return
        }

        if (!isPriceCorrect(data.price)) {
            toast.error("Введите сумму, содержащую 0,1 или 2 знака после точки")
            return
        }

        console.log(currentWallet.value, data.price)
        if (parseFloat(currentWallet.value) - parseFloat(data.price) < 0) {
            toast.error("Недостаточно средств")
            return
        }

        await axios.post('http://localhost:8080/api/expenses/add', data, {withCredentials: true})
            .then(response => {
                addExpense(response.data)
                decreaseWalletBalance({
                    name: data.wallet_name,
                    value: data.price,
                    currency: data.wallet_currency
                })
                toast.success("Операция успешно добавлена")
                reset()
            })
            .catch(error => {
                console.log(error.response)
                if (error.response.status === 500) {
                    toast.error("Сеанс закончился")
                }
                else {
                    toast.error("Операция не добавлена")
                }
            })
    }

    const getCategory = () => {
        return selectedCategory ? categories.find(c => c.value === selectedCategory) : ''
    }

    const onCategoryChange = (newValue) => {
        setSelectedCategory(newValue.value)
    }

    const getWallet = () => {
        // console.log("get wallet")
        return selectedWallet.value ? walletsMenu.find(c => c.value === selectedWallet.value) : ''
    }

    const onWalletChange = (newValue) => {
        // console.log("on change")
        setSelectedWallet(newValue)
    }

    return (
        <Modal open={open}>
            <form className='modal__form' onSubmit={handleSubmit(submitHandler)}>
                <Select onChange={onWalletChange} value={getWallet()} options={walletsMenu}/>
                <input {...register('shop_name')} className='modal__input' placeholder='Магазин' type='text' maxLength={12}/>
                <Select onChange={onCategoryChange} value={getCategory()} options={categories}/>
                <input {...register('price')} className='modal__input' placeholder='Сумма' type='text' maxLength={10}/>
                <DatePicker showIcon selected={selectedDate} onChange={date => setSelectedDate(date)}/>
                <button className='modal-button' type='submit'>Добавить</button>
            </form>
        </Modal>
    )
}