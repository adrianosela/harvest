package harvest

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

const (
	// MaxPlayers represents the maximum number of players per game
	MaxPlayers = 4

	// MinPlayers represents the minimum number of player per game
	MinPlayers = 2

	// RoundsPerGame represents the number of rounds in a game
	RoundsPerGame = 6

	// CardsPerPlayer represents the number of cards dealt out to each player
	CardsPerPlayer = 6
)

var (
	// ErrGameFull occurs when trying to add a player to a full game
	ErrGameFull = errors.New("game is full")

	// ErrNotEnoughPlayers occurs when there aren't enough players for a game
	ErrNotEnoughPlayers = errors.New("not enough players")

	// ErrGameOngoing occurs when trying to add a player to an ongoing game
	ErrGameOngoing = errors.New("game is ongoing")

	// ErrNotInGame occurs when trying to deactivate a player not in the game
	ErrNotInGame = errors.New("player not in game")
)

// Game contains all state about a game of harvest
type Game struct {
	ID string `json:"game_id"`

	Players map[string]Player `json:"players,omitempty"`

	Deck *Deck `json:"deck"`

	Ongoing bool `json:"ongoing"`
	Turn    int  `json:"turn"`
	Round   int  `json:"round"`
}

// NewGame returns a new game
func NewGame() *Game {
	return &Game{
		ID:      uuid.Must(uuid.NewV4()).String(),
		Players: make(map[string]Player),
		Deck:    NewDeckN(4),
		Ongoing: false,
		Turn:    0,
		Round:   0,
	}
}

// Start kicks off the round and turn timer
func (g *Game) Start() error {
	if g.Ongoing {
		return ErrGameOngoing
	}
	if len(g.Players) < MinPlayers {
		return ErrNotEnoughPlayers
	}
	g.deal()
	g.Ongoing = true
	return nil
}

func (g *Game) deal() {
	for i := 0; i < CardsPerPlayer; i++ {
		for _, player := range g.Players {
			card := g.Deck.Pick()
			// first half are dealt
			// visible to the owner
			if i < CardsPerPlayer/2 {
				card.VisibleToOwner = true
			}
			// deal card
			player.Hand[i] = card
		}
	}
}

// AddPlayer adds a player to a game
func (g *Game) AddPlayer(id string) error {
	if g.Ongoing {
		return ErrGameOngoing
	}
	if len(g.Players) >= MaxPlayers {
		return ErrGameFull
	}
	g.Players[id] = Player{}
	return nil
}
