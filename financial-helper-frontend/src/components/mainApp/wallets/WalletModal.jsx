import {Modal} from '../../modal/Modal.jsx'
import '../../modal/Modal.scss'
import '../../buttons/buttons.scss'
import {useForm} from "react-hook-form";
import {toast} from "react-toastify";
import axios from "axios";
import {useEffect, useState} from "react";
import Select from "react-select";
import {useApp} from "../../../hooks/useApp.js";
import {ErrorMessage} from "@hookform/error-message";
import {isPriceCorrect} from "../../../utils/modalUtils.js";

export function AddWalletModal({open = false, wallets, setWallets}) {

    const [selectedCurrency, setSelectedCurrency] = useState('rub');
    const currency = [{value:'rub', label: 'RUB'}, { value: 'usd', label: 'USD'}];

    const {register, handleSubmit, reset, formState: { errors}} = useForm({
        defaultValues: {
            wallet_name: '',
            balance: '',
            // currency: '',
        }
    })

    function isWalletDataCorrect(data) {
        return !(data.wallet_name === '' || data.balance === '' || data.currency === '');
    }

    const submitHandler = async data => {
        if (!isWalletDataCorrect(data)) {
            toast.error("Заполните все поля")
            return
        }

        if (!isPriceCorrect(data.balance)) {
            toast.error("Введите баланс, содержащую 0,1 или 2 знака после точки")
            return
        }

        data['currency'] = selectedCurrency;
        await axios.post('http://localhost:8080/api/wallets/add', data, {withCredentials: true})
            .then(response => {
                setWallets([...wallets, {
                    id: response.data.id,
                    name: response.data.wallet_name,
                    value: response.data.balance,
                    currency: response.data.currency,
                }])
                toast.success("Кошелек успешно добавлен")
                reset()
            })
            .catch(error => {
                // console.log(error.response)
                if (error.response.status === 403) {
                    toast.error("Кошелек с таким названием уже существует")
                }
                else if (error.response.status === 500) {
                    toast.error("Сеанс закончился")
                }
                else {
                    toast.error("Кошелек не добавлен")
                }
            })
    }

    const getCurrency = () => {
        return selectedCurrency ? currency.find(c => c.value === selectedCurrency) : ''
    }

    const onCurrencyChange = (newValue) => {
        setSelectedCurrency(newValue.value)
    }

    return (
        <Modal open={open}>
            <form className='modal__form' onSubmit={handleSubmit(submitHandler)}>
                <input  {...register('wallet_name')} className='modal__input' placeholder='Название' type='text' maxLength={12}/>
                <input {...register('balance')} className='modal__input' placeholder='Баланс' type='text' maxLength={10}/>
                <Select onChange={onCurrencyChange} value={getCurrency()} options={currency}/>
                <button className='modal-button' type='submit'>Добавить</button>
            </form>
        </Modal>
    )
}

export function ChangeWalletModal({open = false, current}) {
    const {updateWallets} = useApp();

    const {register, handleSubmit, reset, formState: { errors}} = useForm({
        defaultValues: {
            id: current.id,
            name: current.name,
            value: current.value,
            currency: current.currency,
        }
    })

    const submitHandler = async data => {
        if (!isPriceCorrect(data.value)) {
            toast.error("Введите баланс, содержащую 0,1 или 2 знака после точки")
            return
        }
        await axios.put('http://localhost:8080/api/wallets/update_wallet', data, {withCredentials: true})
            .then(() => {
                // console.log(response)
                updateWallets()
                // setWallets([...wallets, {id: response.data.id, name: response.data.wallet_name, value: response.data.balance}])
                toast.success("Баланс успешно изменен")
            })
            .catch(error => {
                console.log(error.response)
                if (error.response.status === 500) {
                    toast.error("Сеанс закончился")
                }
                else {
                    toast.error("Баланс не был изменен")
                }
            })
    }

    useEffect(() => {
        reset(current)
    }, [current, reset])

    return (
        <Modal open={open}>
            <form className='modal__form' onSubmit={handleSubmit(submitHandler)}>
                <input  {...register('name')} className='modal__input'
                        placeholder='Название' type='text' readOnly/>
                <input {...register('value')} className='modal__input' placeholder='Баланс'
                       type='text' maxLength={10}/>
                <input {...register('currency')} type='text' className='modal__input' readOnly/>
                <button className='modal-button' type='submit'>Изменить</button>
                <ErrorMessage errors={errors} name="balance" render={({message}) => <p className='modal__error'>{message}</p>} />
            </form>
        </Modal>
    )
}