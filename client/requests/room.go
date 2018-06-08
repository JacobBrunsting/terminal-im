package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jbrunsting/terminal-im/models"
)

func (r *Requester) CreateRoom(name string) error {
	room := models.Room{
		Name:        name,
		HistorySize: 100, // TODO: Change to have min/max/default value
	}

	body, err := json.Marshal(room)
	if err != nil {
		return err
	}

	resp, err := r.Client.Post(
		r.BaseUrl+"/v1/rooms",
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusConflict {
		return &models.NameConflict{}
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got bad response %v when posting room %v", resp.StatusCode, room)
	}

	return nil
}

func (r *Requester) RetrieveRoom(name string) (*models.Room, error) {
	resp, err := r.Client.Get(r.BaseUrl + "/v1/rooms/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, &models.NotFound{name}
	}

	var room models.Room
	err = json.NewDecoder(resp.Body).Decode(&room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}
