package store

import (
	"fmt"
	"sync"

	"github.com/adrianosela/harvest"
)

// Mongo implements the Store
// interface in mongo db
type Mongo struct {
	// TODO
}

// Create writes a new game to the Store
func (m *Mongo) Create(game *harvest.Game) error {
	return nil // TODO
}

// Read reads a game from the Store
func (m *Mongo) Read(gameID string) (*harvest.Game, error) {
	return nil, nil // TODO
}

// Update updates a game in the Store
func (m *Mongo) Update(game *harvest.Game) error {
	return nil // TODO
}

// Delete deletes a game from the Store
func (m *Mongo) Delete(gameID string) error {
	return nil // TODO
}
