import {Modal} from '../modal/Modal.jsx'
import '../modal/Modal.scss'
import '../buttons/buttons.scss'
import {useForm} from "react-hook-form";
import axios from "axios";
import {useNavigate} from "react-router-dom";
import {useModal} from "../../hooks/useModal.js";
import {toast} from "react-toastify";

export function SignIn({open = false}) {
    const navigate = useNavigate();
    const  {setIsActive} = useModal();

    const {register, handleSubmit, reset} = useForm({
        defaultValues: {
            email: '',
            password: ''
        }
    })

    const submitHandler = async data => {
        // console.log(data)
        await axios.post('http://localhost:8080/api/signin', data, {withCredentials: true})
            .then(response => {
                // console.log(response)
                if (response.status === 200) {
                    // console.log("Successfully logged in")
                    toast.success("Вы успешно вошли")
                    navigate("/profile", {replace: true})
                    // localStorage.setItem("username", response.data.user.username)
                    setIsActive(false)
                }
            })
            .catch(error => {
                console.log(error.response)
                toast.error("Неправильный логин или пароль")
            })
        reset()
    }

    return (
        <Modal open={open}>
            <form className='modal__form' onSubmit={handleSubmit(submitHandler)}>
                <input  {...register('email')} className='modal__input' placeholder='Email' type='text' maxLength={45} />
                <input {...register('password')} className='modal__input' placeholder='пароль' type='password' />
                <button className='modal-button' type='submit'>Авторизоваться</button>
            </form>
        </Modal>
    )
}