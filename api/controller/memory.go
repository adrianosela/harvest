package controller

// Memory implements the games Controller
// interface in-memory, such that all state
// pertaining to its games can only be accessed
// on the server hosting the games
type Memory struct {
	/* TODO: complete memory implementation */
}

// NewMemory returns an in-memory games Controller
func NewMemory(games int) *Memory {
	return &Memory{
		/* TODO: complete memory implementation */
	}
}

// Join adds a player to a game
func (ctrl *Memory) Join(player, game string) error {
	/* TODO: complete memory implementation */
	return nil
}

// Leave removes a player from a game
func (ctrl *Memory) Leave(player, game string) error {
	/* TODO: complete memory implementation */
	return nil
}
