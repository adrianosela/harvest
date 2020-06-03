package manager

// MongoDB implements the rooms Manager
// interface in a mongodb instance
type MongoDB struct {
	/* TODO: complete mongodb implementation */
}

// NewMongoDB returns a mongodb rooms Manager
func NewMongoDB(rooms int) (*MongoDB, error) {
	return &MongoDB{
		/* TODO: complete mongodb implementation */
	}, nil
}

// Join adds a player to a room
func (mgr *MongoDB) Join(player, room string) error {
	/* TODO: complete mongodb implementation */
	return nil
}

// Leave removes a player from a room
func (mgr *MongoDB) Leave(player, room string) error {
	/* TODO: complete mongodb implementation */
	return nil
}

// Status returns the Status object for the rooms manager
func (mgr *MongoDB) Status() (*Status, error) {
	/* TODO: complete mongodb implementation */
	return nil, nil
}
