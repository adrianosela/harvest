package msg

var (
	// TypeState is the message type for state objects
	TypeState = "STATE"
)

// Message represents communication over websocket
type Message struct {
	Type string      `json:"type"`
	Args interface{} `json:"args"`
}
