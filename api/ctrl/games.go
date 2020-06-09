package ctrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adrianosela/harvest"
	"github.com/adrianosela/harvest/api/auth"
	"github.com/gorilla/mux"
)

type listGamesResponse struct {
	Games map[string]seatsAvailable `json:"games"`
}

type seatsAvailable struct {
	SeatsAvailable int `json:"seats_available"`
}

// listGamesHandler lists all not-full, not-started games
func (c *Controller) listGamesHandler(w http.ResponseWriter, r *http.Request) {

	opts := &harvest.ListOpts{
		ExcludeStarted: true,
		ExcludeEnded:   true,
		ExcludeFull:    true,
	}

	list, err := c.games.ListGames(opts)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("game not found"))
		return
	}

	lgr := listGamesResponse{Games: make(map[string]seatsAvailable)}
	for _, game := range list {
		lgr.Games[game.ID] = seatsAvailable{
			SeatsAvailable: harvest.MaxPlayers - len(game.Players),
		}
	}

	respBytes, err := json.Marshal(&lgr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize games list"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

// adds the player to a game
func (c *Controller) joinGameHandler(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaims(r)

	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		http.Error(w, "no game id in request URL", http.StatusBadRequest)
		return
	}

	state, err := c.games.GetGame(gameID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("game not found"))
		return
	}

	if err = state.AddPlayer(claims.Subject); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("could not add player to game: %s", err.Error())))
		return
	}

	if err = c.games.UpdateGame(state); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // should rarely happen
		w.Write([]byte(fmt.Sprintf("could not add update game: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("player %s joined game %s", claims.Subject, gameID)))
	return
}

// takes a snapshot of a game from the perspective of a given player
func (c *Controller) stateHandler(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaims(r)

	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		http.Error(w, "no game id in request URL", http.StatusBadRequest)
		return
	}

	state, err := c.games.GetGame(gameID)
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
