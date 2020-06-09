import React from 'react'
import GameBoard from '../../game/GameBoard';
import {openWebSocket, login} from '../../../service/api';

class Main extends React.Component {
    async componentWillMount() {
        await login('MOCK_USERNAME', 'MOCK_PASS')
        // TODO: if login succeeds, we dont need to pass MOCK_TOKEN below
        openWebSocket('MOCK_GAME_ID', "MOCK_TOKEN")
    }

    render() {
        return(
            <div>
                <GameBoard />
            </div>
        );
    }
}

export default (Main)
