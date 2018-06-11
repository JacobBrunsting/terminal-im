package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	cursorToStart = "\r"
	cursorUp      = "\033[1A"
	deleteToEnd   = "\033[K"
	prompt        = "> "
)

type OnInput interface {
	OnInput(input string, inputHandler InputHandler)
}

type InputHandler interface {
	StartScanning(callback OnInput)
	Println(message string)
}

type inputHandler struct {
	reader *bufio.Reader
}

func NewInputHandler() InputHandler {
	return &inputHandler{bufio.NewReader(os.Stdin)}
}

func (i *inputHandler) StartScanning(callback OnInput) {
	fmt.Printf(prompt)
	go i.startScanning(callback)
}

func (i *inputHandler) startScanning(callback OnInput) {
	for {
		input, err := i.reader.ReadString('\n')
		if err != nil {
			log.Printf("error: %v\n", err.Error())
			return
		}
		i.Println("")
		callback.OnInput(strings.TrimSuffix(input, "\n"), i)
	}
}

func (i *inputHandler) Println(message string) {
	fmt.Print(cursorUp + cursorToStart + deleteToEnd)
	fmt.Println(message)
	fmt.Print(prompt)
}
