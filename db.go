package harvest

import "context"

// DB represents the storage needs
// of a harvest games' manager
type DB interface {
	// save a new game's state
	CreateGame(*Game) error
	// get a snapshot of the game's state
	GetGame(string) (*Game, error)
	// update a game's'state
	UpdateGame(*Game) error
	// delete a game's state
	DeleteGame(string) error
	// subscribe to a game's state updates
	WatchGame(string) (UpdateStream, error)
}

// UpdateStream represents an iterator-like
// object (with a blocking next() function)
// for updates on a game's state
type UpdateStream interface {
	Next(context.Context) bool
}
