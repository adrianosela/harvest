package harvest

import (
	"errors"
	"time"

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

	// IntermissionTime is the time before a new turn begins
	IntermissionTime = time.Second * 3

	// MaxTurnTime is the maximum time allowed for a player to make a move
	// before the default move (flipping random board card without picking)
	MaxTurnTime = time.Second * 20
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
	ID string `json:"_id" bson:"_id,omitempty"`

	Players []*Player `json:"players" bson:"players"`

	Stack   *Deck `json:"stack,omitempty" bson:"stack"`
	Rejects *Deck `json:"rejects" bson:"rejects"`

	Ongoing bool `json:"ongoing" bson:"ongoing"`
	Turn    int  `json:"turn" bson:"turn"`
	Round   int  `json:"round" bson:"round"`
}

// NewGame returns a new game
func NewGame() *Game {
	return &Game{
		ID: uuid.Must(uuid.NewV4()).String(),

		Players: []*Player{},

		Ongoing: false,
		Turn:    0,
		Round:   0,

		Stack: NewDeckN(4).Shuffle(),

		// Rejects get set by deal()
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

	// TODO

	return nil
}

func (g *Game) deal() {
	for i := 0; i < CardsPerPlayer; i++ {
		for _, player := range g.Players {
			card := g.Stack.Pop()
			card.OwnerVisible = (i >= CardsPerPlayer/2)
			player.Hand[i] = card
		}
	}
	g.Rejects = &Deck{Cards: []Card{g.Stack.Pop()}}
}

// AddPlayer adds a player to a game
func (g *Game) AddPlayer(id string) error {
	if g.Ongoing {
		return ErrGameOngoing
	}
	if len(g.Players) >= MaxPlayers {
		return ErrGameFull
	}
	g.Players = append(g.Players, &Player{ID: id})
	return nil
}

// Obfuscate takes a snapshot of a game from the perspective
// of a given player (i.e. hides other players' cards)
func (g *Game) Obfuscate(playerID string) *Game {
	var copied Game

	copied.ID = g.ID
	copied.Ongoing = g.Ongoing

	copied.Round = g.Round
	copied.Turn = g.Turn

	copied.Stack = nil                                      // hide stack
	copied.Rejects = &Deck{Cards: []Card{g.Rejects.Peak()}} // show top reject

	for _, player := range g.Players {
		copiedPl := &Player{ID: player.ID}

		for i, card := range player.Hand {
			// card is facing up
			if card.FaceUp {
				copiedPl.Hand[i] = card
				continue
			}
			// card belongs to player and is visible
			if card.OwnerVisible && player.ID == playerID {
				copiedPl.Hand[i] = card
				continue
			}
			copiedPl.Hand[i] = Card{FaceUp: false}
		}

		copied.Players = append(copied.Players, copiedPl)
	}

	return &copied
}
