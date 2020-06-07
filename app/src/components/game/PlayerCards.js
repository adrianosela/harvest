import React from 'react'
import {Row} from 'react-bootstrap';
import Card from './Card'

import '../styles/PlayerContainer.sass'

class PlayerCards extends React.Component {
    constructor(props) {
        super(props)
        this.renderTopRow = this.renderTopRow.bind(this)
        this.renderBottomRow = this.renderBottomRow.bind(this)
    }

    renderTopRow() {
        const {data} = this.props
        return data.slice(0,3).map((card, index) => {
            return (
                <div className="deckCardsCol ml-z mr-xs">
                    <Card key={index} suit={card.suit} rank={card.rank} faceUp={true}/>
                </div>
            )
        })
    }

    renderBottomRow() {
        const {data} = this.props
        return data.slice(3,6).map((card, index) => {
            return (
                <div className="deckCardsCol ml-z mr-xs">
                    <Card key={index} suit={card.suit} rank={card.rank} faceUp={true}/>
                </div>
            )
        })
    }

    render() {
        return(
            <div className="">
                <Row className="zeroMargin mb-xs">
                    {this.renderTopRow()}
                </Row>  
                <Row className="zeroMargin mb-xs">
                    {this.renderBottomRow()}
                </Row>  
            </div>
        )
    }
}

export default (PlayerCards)
