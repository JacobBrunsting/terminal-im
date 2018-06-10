package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jbrunsting/terminal-im/models"
	"github.com/jbrunsting/terminal-im/utils"
)

func (r *Requester) CreateRoom(name string) (*models.Room, error) {
	room := models.Room{
		Name:        name,
		HistorySize: 100, // TODO: Change to have min/max/default value
	}

	body, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Post(
		r.BaseUrl+"/v1/rooms",
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	respBody, err := utils.GetResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	if respBody.Error != nil {
		if respBody.Error.Code == http.StatusConflict {
			return nil, &models.NameConflict{}
		}

		return nil, fmt.Errorf("%v: %v", respBody.Error.Code, respBody.Error.Message)
	}

	var returnedRoom models.Room
	err = json.Unmarshal(respBody.Data, &returnedRoom)
	if err != nil {
		return nil, fmt.Errorf("got error '%v' when decoding response data %v", err.Error(), string(respBody.Data))
	}

	return &returnedRoom, nil
}

func (r *Requester) RetrieveRoom(name string) (*models.Room, error) {
	resp, err := r.Client.Get(r.BaseUrl + "/v1/rooms/" + name)
	if err != nil {
		return nil, err
	}

	respBody, err := utils.GetResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	if respBody.Error != nil {
		if respBody.Error.Code == http.StatusNotFound {
			return nil, &models.NotFound{name}
		}
		return nil, fmt.Errorf("%v: %v", respBody.Error.Code, respBody.Error.Message)
	}

	var room models.Room
	err = json.Unmarshal(respBody.Data, &room)
	if err != nil {
		return nil, fmt.Errorf("got error '%v' when decoding response data %v", err.Error(), string(respBody.Data))
	}

	return &room, nil
}
