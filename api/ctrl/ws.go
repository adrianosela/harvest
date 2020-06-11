package ctrl

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/adrianosela/harvest"
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

func (c *Controller) watchGameHandler(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaims(r)

	var gameID string
	if gameID = mux.Vars(r)["game_id"]; gameID == "" {
		errStr := "no game id in request URL"
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil) // note: sets HTTP header on fail
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	defer ws.Close()

	ctx := context.Background()
	cs, err := c.games.WatchGame(ctx, gameID)
	if err != nil {
		log.Println(err)
		return
	}
	defer cs.Close(ctx)

	for cs.Next(ctx) {
		var game *harvest.Game
		if err := cs.Decode(&game); err != nil {
			log.Println(err)
			return
		}

		game.PlayerView(claims.Subject)
		message, _ := json.Marshal(&msg.Message{Type: msg.TypeState, Args: game})
		ws.WriteMessage(websocket.TextMessage, message)
	}
}
