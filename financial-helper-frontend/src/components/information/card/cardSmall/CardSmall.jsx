import styles from '../Card.module.scss'
import './CardSmall.scss'


export function CardSmall({title, subtitle}) {
    return (
        <div className={styles.card}>
            <div className='card-small-data'>
                <div className='card-small-title'>{title}</div>
                <div className='card-small-subtitle'>{subtitle}</div>
            </div>
        </div>
    )
}