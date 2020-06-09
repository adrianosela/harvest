package ctrl

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/adrianosela/harvest/api/auth"
	"github.com/adrianosela/harvest/api/msg"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// FIXME: whitelist r.Header.Get("Origin")
		return true
	},
}

func (c *Controller) wsHandler(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaims(r)

	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		errStr := "no game id in request URL"
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	state, err := c.store.Read(gameID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("game not found"))
		return
	}

	snap := state.Snapshot(claims.Subject)

	ws, err := upgrader.Upgrade(w, r, nil) // note: sets HTTP header on fail
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	defer ws.Close()

	message, _ := json.Marshal(&msg.Message{Type: "STATE", Args: snap})
	ws.WriteMessage(websocket.TextMessage, message)

	message, _ = json.Marshal(&msg.Message{Type: "MOCK",
		Args: map[string]string{
			"arg1": "val1",
			"arg2": "val2",
		}})

	for {
		ws.WriteMessage(websocket.TextMessage, message)
		time.Sleep(time.Second * 1)
	}
}
