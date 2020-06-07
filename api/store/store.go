package store

import (
	"github.com/adrianosela/harvest"
)

// Store represents the storage needs
// of the Harvest server
type Store interface {
	Create(*harvest.Game) error
	Read(string) (*harvest.Game, error)
	Update(*harvest.Game) error
	Delete(string) error
}
