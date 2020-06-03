package table

import (
	"errors"
	"github.com/adrianosela/harvest/game/player"
	"sync"
)

const (
	// MaxPlayers represents the maximum number of players per table
	MaxPlayers = 4

	// MinPlayers represents the minimum number of player per table
	MinPlayers = 2

	// MaxSpectators represents the maximum spectators per table
	MaxSpectators = 100
)

var (
	// ErrTableFull occurs when trying to add a player to a full table
	ErrTableFull = errors.New("table is full")
)

// Table represents data about the table where the game is held
type Table struct {
	sync.RWMutex

	ID         string                    `json:"table_id"`
	Players    map[string]*player.Player `json:"players"`
	Spectators int                       `json:"spectators"`
}
