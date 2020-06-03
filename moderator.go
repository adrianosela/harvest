package harvest

// Moderator represents operations executable
// by the administrator of a game of harvest.
// It essentially manages a set of Game objects.
type Moderator interface {
	Join(playerID, gameID string) error
	Leave(playerID, gameID string) error
	Snapshot(gameID string) (*Game, error)
}
