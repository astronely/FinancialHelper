import 'reactjs-popup/dist/index.css';
import {Modal} from "../modal/Modal.jsx";
import '../modal/Modal.scss'
import {useModal} from "../../hooks/useModal.js";

export function ConfirmItemDeleteModal({open = false, deleteAction, item}) {
    const  {setIsActive, modal} = useModal();
    let text = modal === 'confirmDeleteWallet' ? 'кошелек' : 'операцию'
    return(
        <Modal open={open}>
            <div className='modal__form'>
                <div className='modal__text'>Вы уверены что хотите удалить {text}?</div>
                <div className='modal__buttons'>
                    <button className='modal-button' type='submit' onClick={() => setIsActive(false)}>Отменить</button>
                    <button className='modal-button' type='submit' onClick={() => {
                        deleteAction(item)
                        setIsActive(false)
                    }}>Удалить</button>
                </div>
            </div>
        </Modal>
    )
}