import React from 'react'
import OponentContainer from './OponentContainer'
import '../styles/GameBoard.sass'
import {Row} from 'react-bootstrap';

class GameBoard extends React.Component {
    render() {
        return(
            <div className="gameBoardContainer">
                <Row className="zeroMargin">
                    <OponentContainer />
                    <OponentContainer />
                    <OponentContainer />
                </Row>
            </div>
        )
    }
}

export default (GameBoard);
