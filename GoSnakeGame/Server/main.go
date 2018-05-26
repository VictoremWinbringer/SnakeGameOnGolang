package main

import (
	"fmt"
	"time"

	ser "../Shared/serializer"
	"./al"
	"./dal"
	tc "github.com/gdamore/tcell"
)

func main() {
	// screen, _ := CreateScreen()
	// screen.HideCursor()
	// defer func() {
	// 	screen.Fini()
	// }()
	game, _ := al.NewGame(20, 40, make(<-chan al.Command))
	old := time.Now().UnixNano()
	for {
		new := time.Now().UnixNano()
		if !game.Logic(new - old) {
			//	screen.ShowCursor(0, 0)
			return
		}

		result := game.Draw()
		println(result)
		return
		// screen.Show()
		// screen.Clear()
	}
	// handlers := make(map[byte]func([]byte, int) []byte, 1)
	// handlers[ser.GameStateType] = handleGameData
	// handlers[ser.CommandType] = handleCommand
	// server, err := s.NewServer(8888, "127.0.0.1", MyHandler{handlers})
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	// serverStartError := server.Start()
	// if serverStartError != nil {
	// 	fmt.Printf("%v", serverStartError)
	// 	return
	// }
	// fmt.Println("Started server")
	// fmt.Scanln()
}

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

func keyboardInput(screen tc.Screen) chan al.Command {
	commandChannel := make(chan al.Command)
	go func() {
		for {
			key := ReadKey(screen)
			switch key {
			case dal.KeyUp:
				commandChannel <- al.Up
			case dal.KeyDown:
				commandChannel <- al.Down
			case dal.KeyLeft:
				commandChannel <- al.Left
			case dal.KeyRight:
				commandChannel <- al.Right
			case dal.KeyEsc:
				commandChannel <- al.Exit
			}
		}
	}()
	return commandChannel
}

func ReadKey(screen tc.Screen) dal.Key {
	event := screen.PollEvent()
	keyEvent, ok := event.(*tc.EventKey)
	if ok {
		switch keyEvent.Key() {
		case tc.KeyUp:
			return dal.KeyUp
		case tc.KeyDown:
			return dal.KeyDown
		case tc.KeyLeft:
			return dal.KeyLeft
		case tc.KeyRight:
			return dal.KeyRight
		case tc.KeyEsc:
			return dal.KeyEsc
		}
	}
	return dal.KeyUndefined
}

func CreateScreen() (tc.Screen, error) {
	s, e := tc.NewScreen()
	if e != nil {
		return nil, e
	}
	if e := s.Init(); e != nil {
		return nil, e
	}
	s.SetStyle(tc.StyleDefault)
	return s, nil
}
