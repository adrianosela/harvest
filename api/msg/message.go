package msg

import "encoding/json"

// Message represents communication over websocket
type Message struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

func unmarshal(data []byte) (*Message, error) {
	var m Message
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
