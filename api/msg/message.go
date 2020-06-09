package msg

// Message represents communication over websocket
type Message struct {
	Type string      `json:"type"`
	Args interface{} `json:"args"`
}
