package store

import (
	"fmt"
	"sync"

	"github.com/adrianosela/harvest"
)

// Memory implements the Store
// interface in-memory (mock)
type Memory struct {
	sync.RWMutex

	Games map[string]*harvest.Game
}

// Create writes a new game to the Store
func (m *Memory) Create(game *harvest.Game) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Games[game.ID]; ok {
		return fmt.Errorf("game %s already in store", game.ID)
	}
	m.Games[game.ID] = game
	return nil
}

// Read reads a game from the Store
func (m *Memory) Read(gameID string) (*harvest.Game, error) {
	m.RLock()
	defer m.RUnlock()

	if game, ok := m.Games[gameID]; ok {
		return game, nil
	}
	return nil, fmt.Errorf("game %s not in store", gameID)
}

// Update updates a game in the Store
func (m *Memory) Update(game *harvest.Game) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Games[game.ID]; !ok {
		return fmt.Errorf("game %s not in store", game.ID)
	}
	m.Games[game.ID] = game
	return nil
}

// Delete deletes a game from the Store
func (m *Memory) Delete(gameID string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.Games, gameID)
	return nil
}
