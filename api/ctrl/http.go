package ctrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adrianosela/harvest/api/auth"
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

// Handler returns the HTTP handler for the controller
func (c *Controller) Handler() http.Handler {
	r := mux.NewRouter()

	// auth
	r.Methods(http.MethodGet).Path("/login").HandlerFunc(c.loginHandler)

	// game state
	r.Methods(http.MethodGet).Path("/game/{game_id}/snapshot").Handler(c.auth.Wrap(c.snapshotHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/watch").Handler(c.auth.Wrap(c.wsHandler))

	return r
}

func (c *Controller) loginHandler(w http.ResponseWriter, r *http.Request) {
	user, pwd, ok := r.BasicAuth()
	if !ok {
		errStr := "no basic credentials in Authorization header"
		log.Println(errStr)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(errStr))
		return
	}

	if err := c.auth.Basic(user, pwd); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	jwt, err := c.auth.GenerateJWT(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("login failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"token\":\"%s\"}", jwt)))
	return
}

func (c *Controller) wsHandler(w http.ResponseWriter, r *http.Request) {
	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		errStr := "no game id in request URL"
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	// TODO: "watch" game collection on Mongo?

	// note: the upgrade function sets the status header
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error())) // propagate err
		return
	}
	defer ws.Close()

	for {
		ws.WriteMessage(websocket.TextMessage, []byte("{ FIXME: NOT IMPLEMENTED }"))
		time.Sleep(time.Second * 1)
	}
}

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
