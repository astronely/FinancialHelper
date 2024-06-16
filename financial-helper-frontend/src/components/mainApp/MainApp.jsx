import {Container} from "react-bootstrap";
import {Wallets} from "./wallets/Wallets.jsx";
import {History} from "./history/History.jsx";
import './MainApp.scss'

export function MainApp() {

    return (
        <main style={{paddingTop: 40}}>
            <Container className={'main-app'}>
                <Wallets />
                <History />
            </Container>
        </main>

    )
}