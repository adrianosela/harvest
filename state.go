package state

// State represents the data that will be
// constantly streamed to spectating clients
type State interface {
	Serialize() ([]byte, error)
}
