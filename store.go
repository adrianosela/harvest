package harvest

// Store represents a database
// for storage of games' state
type Store interface {
	Create(*Game) error
	Read(string) (*Game, error)
	Update(*Game) error
	Delete(string) error
}
