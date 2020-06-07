import React from 'react'
import { w3cwebsocket as ws } from "websocket";
import GameBoard from '../../game/GameBoard';

const client = new ws('ws://localhost:8080/game/MOCK_ID');

class Main extends React.Component {
    componentWillMount() {
	client.onopen = () => {
	    console.log('websocket client connected');
	};

	// for now simply log the message
	// to test websocket server conn.
	client.onmessage = (message) => {
	    console.log(message.data);
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
