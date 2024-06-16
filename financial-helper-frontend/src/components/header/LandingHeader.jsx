import './Header.scss'
import '../buttons/buttons.scss'
import {Container, Nav, Navbar} from 'react-bootstrap';
import React from "react";
import {useModal} from "../../hooks/useModal.js";
import {useNavigate} from "react-router-dom";
import axios from "axios";

export function LandingHeader() {
    const  {setIsActive, setModal} = useModal();
    const navigate = useNavigate();

    async function isAuth() {
        let result = false
        await axios.get("http://localhost:8080/user/test", {withCredentials: true})
            .then(response => {
                if (response.status === 200) {
                    navigate("/profile", {replace: true,})
                    result = true
                    localStorage.setItem("username", response.data.username)
                }
            })
            .catch(error => {
                if (error.response.status !== 401) {
                    console.log(error)
                }
            });
        return result
    }

    async function openModal(modalName) {
        if (!await isAuth()) {
            setIsActive(true)
            setModal(modalName)
        }
    }

    return (
        <Navbar expand="lg" className="navbar">
            <Container>
                <img src="/icons/logo.svg" className='main-logo' alt='MainLogo'/>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link className='nav-link nav-text' href="#Home">Главная</Nav.Link>
                        <Nav.Link className='nav-link nav-text' href="#ForWho">Для кого</Nav.Link>
                        <Nav.Link className='nav-link nav-text' href="#Possibilities">Возможности</Nav.Link>
                    </Nav>
                    <div className="header-right">
                        <button className="primary-button btn-navigation nav-text" type="submit"
                                onClick={() => openModal("signIn")}>Войти
                        </button>
                        <button className="primary-button btn-navigation nav-text" type="submit"
                                onClick={() => openModal("signUp")}>Зарегистрироваться
                        </button>
                    </div>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}