package controller

import (
	"github.com/adrianosela/harvest"
)

// Controller manages a game server's state.
//
// The game service architecture depends largely
// on the implementation of this interface. e.g.:
// - if the implementation is in-memory (mock),
//   players will have to connect to the same server
//   in order to play together. (impl. in memory.go)
// - if the implementation is in a database (e.g. mongoDB),
//   then state can be shared across multiple servers
//   and rooms are server-independent. (impl. in mongo.go)
type Controller interface {
	Join(player, game string) error
	Leave(player, game string) error
	Snapshot(game string) (*harvest.Game, error) // get a snapshot of an ongoing game
}
