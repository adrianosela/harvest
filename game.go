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

	// ErrGameStarted occurs when trying to add a player to started game
	ErrGameStarted = errors.New("game started")

	// ErrPlayerInGame occurs when trying to add a player already in a game
	ErrPlayerInGame = errors.New("player in game")

	// ErrPlayerNotInGame occurs when trying to access a player not in a game
	ErrPlayerNotInGame = errors.New("player not in game")
)

// Game contains all state about a game of harvest
type Game struct {
	ID      string    `json:"_id" bson:"_id"`
	Players []*Player `json:"players" bson:"players"`

	Stack   *Deck `json:"stack,omitempty" bson:"stack"`
	Rejects *Deck `json:"rejects" bson:"rejects"`

	Turn  int `json:"turn" bson:"turn"`
	Round int `json:"round" bson:"round"`

	Started   bool      `json:"started" bson:"ongoing"`
	StartedAt time.Time `json:"started_at" bson:"started_at"`

	Ended   bool
	EndedAt time.Time

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewGame returns a new game with no
// players, no rejects, and a new deck
func NewGame() *Game {
	return &Game{
		ID:      uuid.Must(uuid.NewV4()).String(),
		Players: []*Player{},
		Stack:   NewDeckN(4).Shuffle(),
	}
}

// Start kicks off the round and turn timer
func (g *Game) Start() error {
	if g.Started {
		return ErrGameStarted
	}
	if len(g.Players) < MinPlayers {
		return ErrNotEnoughPlayers
	}
	g.deal()
	timestamp := time.Now()
	g.Started, g.StartedAt, g.UpdatedAt = true, timestamp, timestamp
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
	if g.Started {
		return ErrGameStarted
	}
	if len(g.Players) >= MaxPlayers {
		return ErrGameFull
	}
	for _, pl := range g.Players {
		if pl.ID == id {
			return ErrPlayerInGame
		}
	}
	g.Players = append(g.Players, &Player{ID: id})
	g.UpdatedAt = time.Now()
	return nil
}

// RemovePlayer removes a player from a game
func (g *Game) RemovePlayer(id string) error {
	var updated []*Player
	for _, pl := range g.Players {
		if pl.ID == id {
			continue
		}
		updated = append(updated, pl)
	}
	g.Players = updated
	g.UpdatedAt = time.Now()
	return nil
}

// PlayerView takes a copy of the state of the Game
// from the perspective of a given player id
// i.e. hides other players' cards, etc
func (g *Game) PlayerView(requester string) *Game {
	var rejects *Deck

	if g.Rejects != nil && len(g.Rejects.Cards) > 0 {
		rejects = &Deck{Cards: []Card{g.Rejects.Peak()}}
	}

	return &Game{
		ID:        g.ID,
		Players:   g.hideCards(requester),
		Started:   g.Started,
		StartedAt: g.StartedAt,
		Ended:     g.Ended,
		EndedAt:   g.EndedAt,
		UpdatedAt: g.UpdatedAt,
		Turn:      g.Turn,
		Round:     g.Round,
		Stack:     nil,
		Rejects:   rejects,
	}
}

func (g *Game) hideCards(requester string) []*Player {
	var players []*Player
	for _, player := range g.Players {
		pl := &Player{ID: player.ID}

		for i, card := range player.Hand {
			// card is facing up
			if card.FaceUp {
				pl.Hand[i] = card
				continue
			}
			// card belongs to player and is visible
			if card.OwnerVisible && player.ID == requester {
				pl.Hand[i] = card
				continue
			}
			pl.Hand[i] = Card{FaceUp: false}
		}

		players = append(players, pl)
	}

	return players
}
