import {Col, Container, Row} from "react-bootstrap";
import './blocks.scss'
import './text.scss'
import './buttons.scss'
import {useModal} from "../../../hooks/useModal.js";
import {openModal} from "../../../utils/modalUtils.js";

export function Home() {
    const  {setIsActive, setModal} = useModal();
    return (
        <Container>
            <div className="base-info">
                <Row className="base-position gx-4 gy-3">
                    <Col md="6" sm="12" className="base-text-block">
                        <div className="main-title-block">
                            <span className="header-text">Лучшая инвестиция – это инвестиция в себя через накопления</span>
                            <span className="subtitle-text">Учет доходов и финансов, помощь в накоплении средств и достижении финансовых целей</span>
                        </div>
                        <button className="btn-action" onClick={() => openModal(setIsActive, setModal, "signUp")}>Начать пользоваться</button>
                    </Col>
                    <Col md="6" sm="12" className="picture-block">
                        <img className="col-12 picture" src="/pictures/title-picture.png"/>
                    </Col>
                </Row>
            </div>
        </Container>
    )
}