package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// injected at build-time
	version string
)

func router() http.Handler {
	r := mux.NewRouter()

	// spectate a game via websocket
	r.Methods(http.MethodGet).Path("/game/{game_id}").HandlerFunc(spectateHandler)

	return r
}

func main() {
	if err := http.ListenAndServe(":80", router()); err != nil {
		log.Fatal(err)
	}
}
