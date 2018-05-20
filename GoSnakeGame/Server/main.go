package main

import (
	"fmt"
	"strings"

	s "./udpServer"
)

type MyHandler struct {
}

func (MyHandler) Hanle(requestData []byte) []byte {
	fmt.Println(strings.Trim(string(requestData), " "))
	return requestData
}

func main() {
	server, err := s.New(8888, "127.0.0.1", MyHandler{})

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	serverStartError := server.Start()

	if serverStartError != nil {
		fmt.Printf("%v", serverStartError)
		return
	}

	fmt.Printf("Started server")
	fmt.Scanln()
}
