package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jbrunsting/terminal-im/models"
	"github.com/jbrunsting/terminal-im/utils"
	"github.com/jbrunsting/terminal-im/server/memstore"
)

type RoomController interface {
	PostRoom(w http.ResponseWriter, r *http.Request)
	GetRoom(w http.ResponseWriter, r *http.Request)
}

type roomController struct {
	rooms memstore.RoomStore
}

func NewRoomController(rooms memstore.RoomStore) RoomController {
	return &roomController{rooms}
}

func (c *roomController) PostRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		log.Printf("could not unmarshal PostRoom request body\n%v", err)
		utils.SendError(w, "Could not parse body as JSON", http.StatusBadRequest)
		return
	}

	err = c.rooms.StoreRoom(&room)
	if err != nil {
		if models.IsNameConflict(err) {
			utils.SendError(w, "Room name taken", http.StatusConflict)
			return
		} else {
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	utils.SendSuccess(w, nil, http.StatusOK)
}

func (c *roomController) GetRoom(w http.ResponseWriter, r *http.Request) {
	roomName := mux.Vars(r)["room"]
	if roomName == "" {
		utils.SendError(w, "Room name required", http.StatusBadRequest)
		return
	}

	room, err := c.rooms.RetrieveRoom(roomName)
	if err != nil {
		if models.IsNotFound(err) {
			utils.SendError(w, "Room not found", http.StatusNotFound)
			return
		} else {
			utils.SendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	utils.SendSuccess(w, room, http.StatusOK)
}
