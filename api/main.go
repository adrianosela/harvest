package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
		http.Error(w, "no game id in request URL", http.StatusBadRequest)
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// FIXME: whitelist r.Header.Get("Origin")
			return true
		},
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
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{ state for game %s }", gameID)))
		time.Sleep(time.Second * 5)
	}
}

func main() {
	rtr := mux.NewRouter()
	rtr.Methods(http.MethodGet).Path("/game/{game_id}").HandlerFunc(spectateGameHandler)

	if err := http.ListenAndServe(":80", rtr); err != nil {
		log.Fatal(err)
	}
}
