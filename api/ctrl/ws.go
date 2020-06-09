package ctrl

import (
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
