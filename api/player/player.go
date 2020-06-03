package player

import (
	"sync"

	"github.com/adrianosela/harvest/api/deck"
)

// Player represents a player in a game of harvest
type Player struct {
	sync.RWMutex

	PlayerID string       `json:"player_id"`
	Hand     [6]deck.Card `json:"hand"`
}

func (p *Player) ComputeScore() int {
	val := 0
	for _, card := range p.Hand {
		val += card.Value()
	}
	return val
}
