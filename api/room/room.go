package room

import (
	"errors"
	"github.com/adrianosela/harvest/api/player"
	"sync"
)

const (
	// MaxPlayers represents the maximum number of players per game
	MaxPlayers = 4
	// MinPlayers represents the minimum number of player per game
	MinPlayers = 2
)

var (
	// ErrRoomFull occurs when trying to add a player to a full room
	ErrRoomFull = errors.New("Room is full")
)

// Room represents a room holding state
// to be streamed to each of its spectators
type Room struct {
	sync.RWMutex

	ID         string                    `json:"room_id"`
	Players    map[string]*player.Player `json:"players"`
	Spectating int                       `json:"spectating"`
}

// TODO: complete room
