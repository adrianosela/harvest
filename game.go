package harvest

import (
	"errors"
)

const (
	// MaxPlayers represents the maximum number of players per game
	MaxPlayers = 4

	// MinPlayers represents the minimum number of player per game
	MinPlayers = 2

	// CardsPerPlayer represents the number of cards dealt out to each player
	CardsPerPlayer = 6
)

var (
	// ErrGameFull occurs when trying to add a player to a full game
	ErrGameFull = errors.New("game is full")

	// ErrGameOngoing occurs when trying to add a player to an ongoing game
	ErrGameOngoing = errors.New("game is ongoing")
)

// Game contains all state about a game of harvest
type Game struct {
	ID      string            `json:"game_id"`
	Players map[string]Player `json:"players"`
	Ongoing bool              `json:"ongoing"`
}
