import React from 'react'
import {Row,  Col} from 'react-bootstrap'
import Card from './Card'
import '../styles/OponentsCards.sass'
class OponentsCards extends React.Component {
    render() {
        return(
            <div className="">
                <Row className="zeroMargin playerCardsRow mb-xs">
                    <div className="playerCardsCol ml-z mr-xs">
                        <Card suit={'S'} rank={'A'} faceUp={false}/>
                    </div>
                    <div className="playerCardsCol ml-z mr-xs">
                        <Card suit={'S'} rank={'A'} faceUp={true}/>
                    </div>
                    <div className="playerCardsCol ml-z mr-z">
                        <Card />
                    </div>
                </Row>
                <Row className="zeroMargin playerCardsRow">
                    <div className="playerCardsCol ml-z mr-xs">
                        <Card suit={'S'} rank={'A'} faceUp={false}/>
                    </div>
                    <div className="playerCardsCol ml-z mr-xs">
                        <Card suit={'S'} rank={'A'} faceUp={true}/>
                    </div>
                    <div className="playerCardsCol ml-z mr-z">
                        <Card />
                    </div>
                </Row>
            </div>
        )
    }
}

export default (OponentsCards)
