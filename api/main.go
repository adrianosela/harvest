package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	// injected at build-time
	version string
)

func spectateGameHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: session details from session token if any

	// get game_id from request URL
	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no game id in request URL"))
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // FIXME
		},
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err))) // FIXME
		return
	}

	// subscribe websocket conn to table state
	wsConn.WriteMessage(websocket.TextMessage,
		[]byte(fmt.Sprintf("connected to game %s", gameID)))

	for {
		// FIXME
	}
}

func main() {
	rtr := mux.NewRouter()
	rtr.Methods(http.MethodGet).Path("/game/{game_id}").HandlerFunc(spectateGameHandler)

	if err := http.ListenAndServe(":80", rtr); err != nil {
		log.Fatal(err)
	}
}
