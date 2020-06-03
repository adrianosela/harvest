package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	// injected at build-time
	version string
)

func roomHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: add check origin to upgrader
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	// get room reference
	// subscribe websocket conn to room state
	wsConn.WriteMessage(websocket.TextMessage, []byte("connected"))
	for {
	}
}

func main() {
	rtr := mux.NewRouter()
	rtr.Methods(http.MethodGet).Path("/room/{room_id}").HandlerFunc(roomHandler)

	if err := http.ListenAndServe(":80", rtr); err != nil {
		log.Fatal(err)
	}
}
