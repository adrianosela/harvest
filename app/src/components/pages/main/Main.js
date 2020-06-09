import React from 'react'
import GameBoard from '../../game/GameBoard';
import {openWebSocket, login} from '../../../service/api';
import {getToken} from '../../../helpers/storage'

class Main extends React.Component {
    async componentWillMount() {
        await login('MOCK_USERNAME', 'MOCK_PASS');
        openWebSocket('MOCK_GAME_ID', getToken());
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
