package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

func spectateHandler(w http.ResponseWriter, r *http.Request) {
	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		http.Error(w, "no game id in request URL", http.StatusBadRequest)
		return
	}

	// note: the upgrade function sets the status header
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error())) // propagate err
		return
	}

	for {
		// send state periodically
		ws.WriteMessage(websocket.TextMessage,
			[]byte(fmt.Sprintf("{{ state for game %s }}", gameID)))

		time.Sleep(time.Second * 5)
	}
}
