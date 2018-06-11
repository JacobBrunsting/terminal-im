package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/jbrunsting/terminal-im/client/cli"
	"github.com/jbrunsting/terminal-im/client/requests"
	"github.com/jbrunsting/terminal-im/models"
)

type onInput struct {
}

func (o *onInput) OnInput(input string, inputHandler cli.InputHandler) {
	inputHandler.Println(fmt.Sprintf("Input is %v", input))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	roomPtr := flag.String("room", "", "room being joined or created")
	hostPtr := flag.String("host", "", "url of terminal-im instance being used")

	flag.Parse()

	if roomPtr == nil || *roomPtr == "" {
		log.Fatalf("Parameter 'room' required")
	}
	room := *roomPtr

	if hostPtr == nil || *hostPtr == "" {
		log.Fatalf("Parameter 'host' required")
	}
	host := *hostPtr
	if last := len(host) - 1; last > 0 && host[last] == '/' {
		host = host[:last]
	}

	requester := requests.Requester{
		BaseUrl: host,
		Client:  &http.Client{},
	}

	isNewRoom := false
	roomObj, err := requester.RetrieveRoom(room)
	if err != nil {
		if !models.IsNotFound(err) {
			log.Fatalf(err.Error())
		}

		roomObj, err = requester.CreateRoom(room)
		isNewRoom = true
		if err != nil {
			if models.IsNameConflict(err) {
				log.Fatalf("Room name taken")
			}
			log.Fatalf(err.Error())
		}
	}

	if isNewRoom {
		log.Printf("Created new room %v\n", roomObj)
	} else {
		log.Printf("Got existing room %v\n", roomObj)
	}

	inputHandler := cli.NewInputHandler()
	inputHandler.StartScanning(&onInput{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
