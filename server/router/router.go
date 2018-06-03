package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jbrunsting/terminal-im/server/controllers"
)

func Route(
	r *mux.Router,
	room controllers.RoomController) {

	r.HandleFunc("/health",
		GetHealth,
	).Methods(http.MethodGet)

	r.HandleFunc("/rooms", room.PostRoom).Methods(http.MethodPost)
	r.HandleFunc("/rooms/{room}", room.GetRoom).Methods(http.MethodGet)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
