package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/jbrunsting/terminal-im/server/controllers"
	"github.com/jbrunsting/terminal-im/server/router"
)

func main() {
	getConfig()

	port := viper.GetString("server.port")

	r := mux.NewRouter()
	rc := controllers.NewRoomController()
	router.Route(r.PathPrefix("/v1").Subrouter(), rc)

	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}
}
