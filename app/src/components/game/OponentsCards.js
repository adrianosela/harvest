import React from 'react'
import {Row} from 'react-bootstrap'
import Card from './Card'
import '../styles/OponentsCards.sass'
class OponentsCards extends React.Component {
    constructor(props) {
        super(props)
        this.renderTopRow = this.renderTopRow.bind(this)
        this.renderBottomRow = this.renderBottomRow.bind(this)
    }
    renderTopRow() {
        const {data} = this.props
        return data.slice(0,3).map((card, index) => {
            return (
                <div className="playerCardsCol ml-z mr-xs">
                    <Card key={index} suit={card.suit} rank={card.rank} faceUp={true}/>
                </div>
            )
        })
    }

    renderBottomRow() {
        const {data} = this.props
        return data.slice(3,6).map((card, index) => {
            return (
                <div className="playerCardsCol ml-z mr-xs">
                    <Card key={index} suit={card.suit} rank={card.rank} faceUp={true}/>
                </div>
            )
        })
    }

    render() {
        return(
            <div className="">
                <Row className="zeroMargin playerCardsRow mb-xs">
                    {this.renderTopRow()}
                </Row>
                <Row className="zeroMargin playerCardsRow">
                    {this.renderBottomRow()}
                </Row>
            </div>
        )
    }
}

export default (OponentsCards)
