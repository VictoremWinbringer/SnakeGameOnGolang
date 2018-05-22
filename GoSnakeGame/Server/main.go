package main

import (
	"fmt"

	ser "../Shared/serializer"
	s "./server"
)

type MyHandler struct {
	handlers map[byte]func([]byte, int) []byte
}

func handleGameData(requestData []byte, clientId int) []byte {
	input := ser.DecodeGameState(requestData)
	fmt.Printf("from client %v \n", input)
	output := ser.GameState{input.State + " From server!"}
	return ser.EncodeGameState(output)
}

func handleCommand(requestData []byte, clientId int) []byte {
	input := ser.DecodeCommand(requestData)
	fmt.Printf("from client %v \n", input)
	input.Code = input.Code + 10
	return ser.EncodeCommand(input)
}

func (h MyHandler) Hanle(requestData []byte, clientId int) []byte {
	return h.handlers[requestData[0]](requestData, clientId)
}

func main() {
	handlers := make(map[byte]func([]byte, int) []byte, 1)
	handlers[ser.GameStateType] = handleGameData
	handlers[ser.CommandType] = handleCommand
	server, err := s.NewServer(8888, "127.0.0.1", MyHandler{handlers})
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
