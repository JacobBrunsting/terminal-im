package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jbrunsting/terminal-im/models"
	"github.com/jbrunsting/terminal-im/utils"
)

type RoomController interface {
	PostRoom(w http.ResponseWriter, r *http.Request)
	GetRoom(w http.ResponseWriter, r *http.Request)
	DeleteRoom(w http.ResponseWriter, r *http.Request)
}

type roomController struct {
}

func NewRoomController() RoomController {
	return &roomController{}
}

func (c *roomController) PostRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		log.Printf("could not unmarshal PostRoom request body\n%v", err)
		utils.SendError(w, "Could not parse body as JSON", http.StatusBadRequest)
		return
	}

	utils.SendSuccess(w, room, http.StatusOK)
}

func (c *roomController) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomName := mux.Vars(r)["room"]
	if roomName == "" {
		utils.SendError(w, "Room name required", http.StatusBadRequest)
		return
	}

	utils.SendSuccess(w, nil, http.StatusOK)
}

func (c *roomController) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	roomName := mux.Vars(r)["room"]
	if roomName == "" {
		utils.SendError(w, "Room name required", http.StatusBadRequest)
		return
	}

	utils.SendSuccess(w, nil, http.StatusNoContent)
}
