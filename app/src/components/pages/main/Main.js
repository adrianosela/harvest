import React from 'react'
import GameBoard from '../../game/GameBoard';
import {openWebSocket, login} from '../../../service/api';
import {getToken} from '../../../helpers/storage'

class Main extends React.Component {
    async componentWillMount() {
        await login('MOCK_PLAYER_1', 'MOCK_PASS');
        openWebSocket('da169ebe-7d99-4429-a298-03bb38055a42', getToken());
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
