import React from 'react'
import {Row} from 'react-bootstrap';
import Card from './Card'

import '../styles/DeckContainer.sass'

class DeckContainer extends React.Component {
    render() {
        return(
            <div className="mr-auto ml-auto">
                <Row className="zeroMargin mb-xs">
                    <div className="deckCardsCol ml-z mr-xs">
                        <Card suit={'S'} rank={'A'} faceUp={false}/>
                    </div>
                    <div className="deckCardsCol ml-z mr-xs">
                        <Card suit={'S'} rank={'A'} faceUp={true}/>
                    </div>
                </Row>  
            </div>
        )
    }
}

export default (DeckContainer)
