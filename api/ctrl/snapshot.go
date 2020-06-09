package ctrl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adrianosela/harvest/api/auth"
	"github.com/gorilla/mux"
)

// takes a snapshot of a game from the perspective of a given player
func (c *Controller) snapshotHandler(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaims(r)

	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		http.Error(w, "no game id in request URL", http.StatusBadRequest)
		return
	}

	state, err := c.store.Read(gameID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("game not found"))
		return
	}

	// hide private fields
	snap := state.Snapshot(claims.Subject)

	snapBytes, err := json.Marshal(&snap)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize game state"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(snapBytes)
	return
}
