import {Col, Container, Row} from "react-bootstrap";
import {CardSmall} from "../card/cardSmall/CardSmall.jsx";

export function Possibilities() {
    return (
        <Container>
            <div className="block-vertical">
                <div className="title-text">Возможности</div>
                <Row className="card-position gx-4 gy-3">
                    <Col md={6} xs={12}>
                        <CardSmall title="Учет доходов" subtitle="Позволяет вести учет финансов, создавая кошельки с удобными названиями. Доступна возможность использовать кошельки в рублях и долларах."/>
                    </Col>
                    <Col md={6} xs={12}>
                        <CardSmall title='Учет расходов' subtitle="Позволяет вести учет совершенных финансовых операций, с указанием места, в котором была проведена операция, категории расходов и даты совершения операции."/>
                    </Col>
                </Row>
            </div>
        </Container>
    )
}