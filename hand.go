package harvest

// Hand represents a player's hand
type Hand [6]Card

// ComputeScore computes the score of a hand
func (h *Hand) ComputeScore() int {
	val := 0
	for _, card := range h {
		val += card.Value()
	}
	return val
}
