import styles from '../Card.module.scss'
import './CardLarge.scss'

export function CardLarge() {
    return (
        <div className={styles.card}>
            <div className="card-data-lg">
                <div className="card-title-lg">Заголовок про статистику</div>
                <div className="card-subtitle-data">
                    <div className="bullet-data">
                        <div className="bullet"><img src="/icons/bullit.svg" /></div>
                        <div className="card-subtitle-lg">Текст про статистику Текст про статистику Текст про статистику Текст
                            про статистику Текст про статистику
                        </div>
                    </div>
                    <div className="bullet-data">
                        <div className="bullet"><img src="/icons/bullit.svg" /></div>
                        <div className="card-subtitle-lg">Текст про статистику Текст про статистику Текст про статистику</div>
                    </div>
                    <div className="bullet-data">
                        <div className="bullet"><img src="/icons/bullit.svg" /></div>
                        <div className="card-subtitle-lg">Текст про статистику Текст про статистику Текст про статистику Текст
                            про статистику
                        </div>
                    </div>
                </div>
            </div>
        </div>
)
}