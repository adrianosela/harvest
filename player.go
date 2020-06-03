package harvest

// Player represents a player in a game of harvest
type Player struct {
	ID   string `json:"player_id"`
	Hand Hand   `json:"hand"`
}
