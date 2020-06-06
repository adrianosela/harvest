package harvest

import (
	"math/rand"
	"time"
)

// Deck holds the cards in the deck
type Deck struct {
	Cards []Card `json:"cards"`
}

// NewDeck creates a deck of cards to be used
func NewDeck() *Deck {
	deck := &Deck{}
	deck.fill()
	return deck
}

// NewDeckN creates a deck composed
// of n regular 52 card decks
func NewDeckN(n int) *Deck {
	deck := &Deck{}
	for i := 0; i < n; i++ {
		deck.fill()
	}
	return deck
}

// fill populates a Deck (52 cards)
func (d *Deck) fill() {
	for _, val := range Types {
		for _, suit := range Suits {
			d.Cards = append(d.Cards, Card{
				Rank:   val,
				Suit:   suit,
				FaceUp: false,
			})
		}
	}
}

// Pick removes and returns the card at the top of the deck
func (d *Deck) Pick() Card {
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}

// Shuffle scrambles the cards in a deck
func (d *Deck) Shuffle() {
	//seed our random functions with CPUs time
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d.Cards); i++ {
		// random int up to the number of cards
		r := rand.Intn(i + 1)
		// If card doesn't match the random int, switch with card at random int
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}
}
