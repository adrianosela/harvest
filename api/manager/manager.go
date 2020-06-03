package manager

import (
	"github.com/adrianosela/harvest/game/table"
)

// Manager manages a game server's state.
//
// The game service architecture depends largely
// on the implementation of this interface. e.g.:
// - if the implementation is in-memory (mock),
//   players will have to connect to the same server
//   in order to play together. (impl. in memory.go)
// - if the implementation is in a database (e.g. mongoDB),
//   then state can be shared across multiple servers
//   and rooms are server-independent. (impl. in mongo.go)
type Manager interface {
	Join(player, table string) error
	Leave(player, table string) error
	Status() (*Status, error)
}

// Status is the state of rooms Manager
type Status struct {
	Tables map[string]*table.Table `json:"tables"`
}