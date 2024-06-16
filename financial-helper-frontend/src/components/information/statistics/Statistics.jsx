import {Col, Container, Row} from "react-bootstrap";
import {CardLarge} from '../card/cardLarge/CardLarge.jsx'
import './statistics.scss'

export function Statistics() {
    
    return(
        <Container>
            <Row className="statistic">
                <Col lg={6} md={12}>
                    <div className="statistic-title-position title-text">Статистика</div>
                </Col>
                <Col lg={6} md={12}>
                    <CardLarge />
                </Col>
            </Row>
        </Container>
    )
}