package main

import (
	"fmt"
	"log"
	"os"

	szr "../Shared/serializer"
	"../Shared/udp"
	"./dal"
)

func sendStateMessage(id uint64) error {
	err := messagesRepository.Write(
		szr.Message{
			Id:   id,
			Type: szr.GameStateType,
			Data: make([]byte, 0)})
	if err != nil {
		return err
	}
	return nil
}

func requestStateFromServer(currentReceivedMessageId uint64) ([][]rune, error, uint64) {
	for {
		m, e := messagesRepository.Read()
		if e != nil {
			return make([][]rune, 0), e, 0
		}
		if m.Id <= currentReceivedMessageId {
			continue
		}
		currentReceivedMessageId = m.Id
		state := szr.DecodeGameState(m.Data)
		return state.State, nil, currentReceivedMessageId
	}
}

func createTestState() [][]rune {
	state := make([][]rune, 0)
	for i := 0; i < 10; i++ {
		state = append(state, make([]rune, 10))
		for j := 0; j < 10; j++ {
			state[i][j] = '+'
		}
	}
	return state
}

func writeStateToBuffer(state [][]rune) {
	stateBuffer <- state
}

func readStateFromBuffer() [][]rune {
	return <-stateBuffer
}

func showState(state [][]rune) {
	screen.Clear()
	for i, a := range state {
		for j, r := range a {
			screen.Write(i, j, r)
		}
	}
	screen.Show()
}

func readPressedKey() chan szr.CommandCode {
	out := make(chan szr.CommandCode)
	go func() {
		key := screen.ReadKey()
		command := parseKeyCode(key)
		out <- command
	}()
	return out
}

func parseKeyCode(key dal.Key) szr.CommandCode {
	switch key {
	case dal.KeyUp:
		return szr.MoveUp
	case dal.KeyDown:
		return szr.MoveDown
	case dal.KeyLeft:
		return szr.MoveLeft
	case dal.KeyRight:
		return szr.MoveRight
	case dal.KeyEsc:
		return szr.ExitGame
	default:
		return szr.UndefinedCommand
	}
}

func sendCommandToServer(command szr.Command, id uint64) error {
	err := messagesRepository.Write(
		szr.Message{
			Id:   id,
			Type: szr.CommandType,
			Data: szr.EncodeCommand(command)})
	if err != nil {
		return err
	}
	return nil
}

var messagesRepository dal.IMessagesRepository
var stateBuffer chan [][]rune
var screen dal.IScreen
var udpClient udp.IUdpClient
var logger *log.Logger

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	factory := dal.CreateDalFactory()
	sc, err := factory.CreateScreen()
	if err != nil {
		print(err.Error())
		return
	}
	screen = sc
	defer screen.Close()
	client, err := udp.NewUdpClient("127.0.0.1:7788", 5)
	if err != nil {
		print(err.Error())
		return
	}
	messagesRepository = factory.CreateMessagesRepository(client)
	defer messagesRepository.Dispose()
	stateBuffer = make(chan [][]rune, 100)
	f, err := os.Create("log.txt")
	check(err)
	defer f.Close()
	l := log.New(f, "main ", log.LstdFlags)
	logger = l
	var currentCommandId uint64 = 10
	var currentReceivedMessageId uint64 = 10
	var currentMessageId uint64 = 10
	codeChan := readPressedKey()
	for {
		select {
		case code := <-codeChan:
			if code > 0 {
				currentCommandId++
				for i := 0; i < 4; i++ {
					currentMessageId++
					e := sendCommandToServer(szr.Command{currentCommandId, code}, currentMessageId)
					if e != nil {
						logger.Println(e.Error())
					}
				}
			}
		default:
			//Do nothing
		}
		e := sendStateMessage(currentMessageId)
		if e != nil {
			logger.Println(e.Error())
		}
		s, e, i := requestStateFromServer(currentReceivedMessageId)
		if e != nil {
			l.Println(e.Error())
			continue
		}
		currentReceivedMessageId = i
		showState(s)
	}
}
