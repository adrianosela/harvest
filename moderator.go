package harvest

import "time"

const (
	// IntermissionTime is the time before a new turn begins
	IntermissionTime = time.Second * 3

	// MaxTurnTime is the maximum time allowed for a player to make a move
	// before the default move (flipping random board card without picking)
	MaxTurnTime = time.Second * 20
)

// Moderator represents operations executable
// by the administrator of a game of harvest.
// It essentially manages a set of Game objects.
type Moderator interface {
	Join(playerID, gameID string) error
	Leave(playerID, gameID string) error
	Snapshot(gameID string) (*Game, error)
}
