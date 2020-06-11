package harvest

import "context"

// DB represents the storage needs
// of a harvest games' manager
type DB interface {
	// save a new game's state
	CreateGame(*Game) error
	// get a snapshot of the game's state
	GetGame(string) (*Game, error)
	// list all games and apply certain filters
	ListGames(*ListOpts) ([]*Game, error)
	// update a game's'state
	UpdateGame(*Game) error
	// delete a game's state
	DeleteGame(string) error
	// subscribe to a game's state updates
	WatchGame(context.Context, string) (UpdateStream, error)
}

// ListOpts represents applicable filters
// when querying the db for a list of games
type ListOpts struct {
	ExcludeStarted bool
	ExcludeEnded   bool
	ExcludeFull    bool
}

// UpdateStream represents an iterator-like
// object (with a blocking next() function)
// for updates on a game's state
type UpdateStream interface {
	Next(context.Context) bool
	Close(context.Context) error
	Decode(interface{}) error
}
