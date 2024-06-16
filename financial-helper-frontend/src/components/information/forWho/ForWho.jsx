import {Col, Container, Row} from "react-bootstrap";
import {CardSmall} from "../card/cardSmall/CardSmall.jsx";
import './blocks.scss'

export function ForWho() {
    return (
        <Container>
            <div className="block-vertical">
                <div className="title-text">Кому подойдет</div>
                <Row className="card-position gx-4 gy-3">
                    <Col md={4} xs={12}>
                        <CardSmall title="Студент" subtitle="Зачастую у студентов имеются проблемы с контролем личных финансов, поскольку их бюджет весьма ограничен. С помощью приложения можно начать отслеживать свои доходы и расходы и научиться их планировать более правильно." />
                    </Col>
                    <Col md={4} xs={12}>
                        <CardSmall title="Семья" subtitle="У семьи имеются общие расходы, которые важно контроллировать всем её членам. Благодаря приложению можно вести учет таких расходов на одном аккаунте, к которому будет доступ у всех членов семьи."/>
                    </Col>
                    <Col md={4} xs={12}>
                        <CardSmall title="Предприниматель" subtitle="Предприниматели имеют различные источники доходов и расходов, за которыми бывает трудно уследить. Благодаря приложению, можно вести их учет в одном месте и всегда иметь доступ к данным."/>
                    </Col>
                </Row>
            </div>
        </Container>
    )
}