package main

import (
	"net/http"
	"log"
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
	log.Println("Got request!")
}
