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

	// game operations operations
	r.Methods(http.MethodGet).Path("/games").Handler(c.auth.Wrap(c.listGamesHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/snapshot").Handler(c.auth.Wrap(c.snapshotGameHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/join").Handler(c.auth.Wrap(c.joinGameHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/leave").Handler(c.auth.Wrap(c.leaveGameHandler))
	r.Methods(http.MethodGet).Path("/game/{game_id}/watch").Handler(c.auth.Wrap(c.watchGameHandler))

	return cors(r) // enable cors
}
