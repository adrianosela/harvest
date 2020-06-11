package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/adrianosela/harvest"
)

// Memory implements the Store
// interface in-memory (mock)
type Memory struct {
	sync.RWMutex

	games map[string]*harvest.Game
}

// NewMemory is the constructor for a Memory type Store
func NewMemory() *Memory {
	return &Memory{
		games: make(map[string]*harvest.Game),
	}
}

// CreateGame writes a new game to the Store
func (m *Memory) CreateGame(game *harvest.Game) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.games[game.ID]; ok {
		return fmt.Errorf("game %s already in store", game.ID)
	}
	m.games[game.ID] = game
	return nil
}

// GetGame reads a game from the Store
func (m *Memory) GetGame(gameID string) (*harvest.Game, error) {
	m.RLock()
	defer m.RUnlock()

	if game, ok := m.games[gameID]; ok {
		return game, nil
	}
	return nil, fmt.Errorf("game %s not in store", gameID)
}

// ListGames gets a list of games from the db
func (m *Memory) ListGames(opts *harvest.ListOpts) ([]*harvest.Game, error) {
	games := []*harvest.Game{}

	for _, game := range m.games {
		if opts != nil {
			if opts.ExcludeStarted && game.Started {
				continue
			}
			if opts.ExcludeEnded && game.Ended {
				continue
			}
			if opts.ExcludeFull && len(game.Players) >= harvest.MaxPlayers {
				continue
			}
		}
		games = append(games, game)
	}

	return games, nil
}

// UpdateGame updates a game in the Store
func (m *Memory) UpdateGame(game *harvest.Game) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.games[game.ID]; !ok {
		return fmt.Errorf("game %s not in store", game.ID)
	}
	m.games[game.ID] = game
	return nil
}

// DeleteGame deletes a game from the Store
func (m *Memory) DeleteGame(gameID string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.games, gameID)
	return nil
}

// WatchGame builds an iterator with a blocking
// function for updates on a game's state
func (m *Memory) WatchGame(ctx context.Context, gameID string) (harvest.UpdateStream, error) {
	return nil, nil // TODO
}
