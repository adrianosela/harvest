import React from 'react';
import OponentContainer from './OponentContainer';
import DeckContainer from './DeckContainer';
import PlayerContainer from './PlayerContainer';
import '../styles/GameBoard.sass'
import {Row} from 'react-bootstrap';

class GameBoard extends React.Component {

    constructor(props) {
        super(props)
        this.renderOponents = this.renderOponents.bind(this)
    }

    game = {
        "game_id": "6e8a2924-1d42-4926-89f3-954fedc29578",
        "players": [
          {
            "player_id": "aaaaa@gmail.com",
            "hand": [
              {
                "rank": "3",
                "suit": "H"
              },
              {
                "rank": "J",
                "suit": "S"
              },
              {
                "rank": "5",
                "suit": "D"
              },
              {
                "rank": "7",
                "suit": "C"
              },
              {
                "rank": "5",
                "suit": "H"
              },
              {
                "rank": "K",
                "suit": "D"
              }
            ]
          },
          {
            "player_id": "bbbbb@gmail.com",
            "hand": [
              {
                "rank": "7",
                "suit": "S"
              },
              {
                "rank": "4",
                "suit": "S"
              },
              {
                "rank": "K",
                "suit": "D"
              },
              {
                "rank": "3",
                "suit": "H"
              },
              {
                "rank": "9",
                "suit": "H"
              },
              {
                "rank": "9",
                "suit": "D"
              }
            ]
          },
          {
            "player_id": "ccccc@gmail.com",
            "hand": [
              {
                "rank": "K",
                "suit": "C"
              },
              {
                "rank": "6",
                "suit": "S"
              },
              {
                "rank": "4",
                "suit": "C"
              },
              {
                "rank": "2",
                "suit": "C"
              },
              {
                "rank": "5",
                "suit": "S"
              },
              {
                "rank": "Q",
                "suit": "H"
              }
            ]
          },
          {
            "player_id": "ddddd@gmail.com",
            "hand": [
              {
                "rank": "10",
                "suit": "C"
              },
              {
                "rank": "3",
                "suit": "H"
              },
              {
                "rank": "7",
                "suit": "D"
              },
              {
                "rank": "5",
                "suit": "D"
              },
              {
                "rank": "3",
                "suit": "D"
              },
              {
                "rank": "7",
                "suit": "H"
              }
            ]
          }
        ],
        "stack": {
          "cards": null
        },
        "rejects": {
          "cards": null
        },
        "ongoing": true,
        "turn": 0,
        "round": 0
      }

    renderOponents(players) {
        const me = {
            id: "aaaaa@gmail.com"
        } 
        return players.filter(player => player.player_id !== me.id).map((player, index) => {
            return(
                <OponentContainer data={player} key={index} />
            )
        })
    }

    render() {
        const me = {
            id: "aaaaa@gmail.com"
        } 
        const my_game = this.game.players.find(player => player.player_id === me.id)
        return(
            <div className="gameBoardContainer">
                <Row className="zeroMargin mt-xl">
                    {this.renderOponents(this.game.players)}
                </Row>
                <Row className="zeroMargin gameBoardDeckRow">
                    <DeckContainer />
                </Row>
                <Row className="zeroMargin gameBoardPlayerRow">
                    <PlayerContainer data={my_game}/>
                </Row>
            </div>
        )
    }
}

export default (GameBoard);
