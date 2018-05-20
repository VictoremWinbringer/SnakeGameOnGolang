package main

import (
	"fmt"

	ser "../Shared/serializer"
	s "./udpServer"
)

type MyHandler struct {
}

func (MyHandler) Hanle(requestData []byte) []byte {

	input := ser.DecodeGameState(requestData)
	fmt.Printf("from client %v \n", input)
	output := ser.GameState{input.State + " From server!"}
	return ser.EncodeGameState(output)
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
	fmt.Println("Started server")
	fmt.Scanln()
}
