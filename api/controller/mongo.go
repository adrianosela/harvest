package controller

// MongoDB implements the games Controller
// interface in a mongodb instance
type MongoDB struct {
	/* TODO: complete mongodb implementation */
}

// NewMongoDB returns a mongodb games Controller
func NewMongoDB(games int) (*MongoDB, error) {
	return &MongoDB{
		/* TODO: complete mongodb implementation */
	}, nil
}

// Join adds a player to a game
func (ctrl *MongoDB) Join(player, game string) error {
	/* TODO: complete mongodb implementation */
	return nil
}

// Leave removes a player from a game
func (ctrl *MongoDB) Leave(player, game string) error {
	/* TODO: complete mongodb implementation */
	return nil
}
