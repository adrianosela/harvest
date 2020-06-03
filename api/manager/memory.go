package manager

// Memory implements the rooms Manager
// interface in-memory, such that all state
// pertaining to its rooms can only be accessed
// on the server hosting the rooms
type Memory struct {
	/* TODO: complete memory implementation */
}

// NewMemory returns an in-memory rooms Manager
func NewMemory(rooms int) *Memory {
	return &Memory{
		/* TODO: complete memory implementation */
	}
}

// Join adds a player to a room
func (mgr *Memory) Join(player, room string) error {
	/* TODO: complete memory implementation */
	return nil
}

// Leave removes a player from a room
func (mgr *Memory) Leave(player, room string) error {
	/* TODO: complete memory implementation */
	return nil
}

// Status returns the Status object for the rooms manager
func (mgr *Memory) Status() (*Status, error) {
	/* TODO: complete memory implementation */
	return nil, nil
}
