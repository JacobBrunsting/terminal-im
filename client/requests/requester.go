package requests

import (
	"net/http"
)

type Requester struct {
	BaseUrl string
	Client  *http.Client
}
