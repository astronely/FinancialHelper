import './InformationText.scss'
import './InformationBlock.scss'
import {Home} from './home/Home'
import {ForWho} from './forWho/ForWho.jsx'
import {Possibilities} from "./possibilities/Possibilities.jsx";
import {Statistics} from "./statistics/Statistics.jsx";

export function Information() {
    return (
        <div className='wrap' style={{paddingTop: 50}}>
            <div id="Home" className="bg-odd">
                <Home />
            </div>
            <div id="ForWho" className="bg-even">
                <ForWho />
            </div>
            <div id="Possibilities" className="bg-odd">
                <Possibilities />
            </div>
            {/*<div id="Statistics" className="bg-even">*/}
            {/*    <Statistics />*/}
            {/*</div>*/}
        </div>
    )
}