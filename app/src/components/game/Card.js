import React, { Component } from 'react';
import styles from '../styles/Card.css';

class Card extends Component {
  render() {
    const suitMap = {
      'S': 'Spades',
      'H': 'Hearts',
      'C': 'Clubs',
      'D': 'Diamonds'
    };

    const rankMap = {
      'A': 'Ace',
      '2': 'Deuce',
      '3': 'Three',
      '4': 'Four',
      '5': 'Five',
      '6': 'Six',
      '7': 'Seven',
      '8': 'Eight',
      '9': 'Nine',
      'T': 'Ten',
      'J': 'Jack',
      'Q': 'Queen',
      'K': 'King'
    };

    return (
        <div className="p-card">
        {
            this.props.faceUp === true && 
            <>
                <div className={`${this.props.suit} mark dark ml-auto mr-auto`}>{this.props.rank}</div>
            </>
        }
        </div>
    );
  }
}

export default Card;