package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

/*
  Follows the Google JSON style guide
*/

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type httpResponseOutgoing struct {
	NextLink string      `json:"nextLink,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Error    *HttpError  `json:"error,omitempty"`
}

type HttpResponse struct {
	NextLink string
	Data     json.RawMessage
	Error    *HttpError
}

func SendSuccess(w http.ResponseWriter, data interface{}, status int) {
	resp := httpResponseOutgoing{Data: data}

	respBody, err := json.Marshal(resp)
	if err == nil {
		w.WriteHeader(status)
		w.Write(respBody)
	} else {
		log.Printf("error marshalling response body\n %v\n error\n %v", resp, err)
		SendError(w, "Could not convert response to JSON", http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")

	resp := HttpResponse{Error: &HttpError{Code: status, Message: message}}

	respBody, err := json.Marshal(resp)
	if err == nil {
		w.WriteHeader(status)
		w.Write(respBody)
	} else {
		log.Printf("error marshalling response body\n %v\n error\n %v", resp, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetResponse(body io.ReadCloser) (*HttpResponse, error) {
	defer body.Close()

	var resp HttpResponse

	err := json.NewDecoder(body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
