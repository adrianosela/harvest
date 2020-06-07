package harvest

// Player represents a player in a game of harvest.
type Player struct {
	ID   string               `json:"player_id" bson:"player_id"`
	Hand [CardsPerPlayer]Card `json:"hand" bson:"hand"`
}

// ComputeScore computes a player's score
func (pl *Player) ComputeScore() int {
	val := 0
	for _, card := range pl.Hand {
		val += card.Value()
	}
	return val
}
