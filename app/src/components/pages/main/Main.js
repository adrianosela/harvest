import React from 'react'
import { w3cwebsocket as ws } from "websocket";
import GameBoard from '../../game/GameBoard';

class Main extends React.Component {
    componentWillMount() {
	    const url = 'ws://localhost:8080'; // FIXME
	    const game_id = 'MOCK_GAME_ID'; // FIXME
	    const token = 'MOCK_TOKEN'; // FIXME: replace with token from HTTP GET /login (basic auth with no password)

	    const wsURL = `${url}/game/${game_id}/watch?token=${token}`;

	    const ws = new WebSocket(wsURL);

	    ws.onopen = () => {
	        console.log('connected');
	    };

	    ws.onmessage = (msg) => {
	        console.log(msg.data); // FIXME: what we sending -- state or moves?
	    };

	    ws.onclose = () => {
            console.log('disconnected');
	    };
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
