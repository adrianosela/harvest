package ctrl

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router returns the HTTP routes handler
func (c *Controller) Router() http.Handler {
	r := mux.NewRouter()

	// auth
	r.Methods(http.MethodGet).Path("/login").HandlerFunc(c.loginHandler)

	// player operations
	r.Methods(http.MethodGet).Path("/games").Handler(c.auth.Wrap(c.listGamesHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/state").Handler(c.auth.Wrap(c.stateHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/join").Handler(c.auth.Wrap(c.joinGameHandler))

	// spectator operations
	r.Methods(http.MethodGet).Path("/game/{game_id}/watch").Handler(c.auth.Wrap(c.wsHandler))

	return cors(r) // enable cors
}
