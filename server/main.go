package main

import (
	"github.com/jbrunsting/terminal-im/models"
	"github.com/jbrunsting/terminal-im/utils"

	"net/http"
	"log"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", testHandler)

	log.Printf("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	var room models.Room

	d := json.NewDecoder(r.Body)
	if err := d.Decode(&room); err != nil {
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

    utils.SendSuccess(w, room, http.StatusOK)
}
