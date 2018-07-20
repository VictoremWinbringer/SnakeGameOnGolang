package main

import (
	"fmt"
	"log"
	"os"
	"time"

	. "../Shared/commandCodeEnum"
	. "../Shared/messageTypeEnum"
	. "../Shared/models"
	szr "../Shared/serializer"
	"./dal"
)

func sendStateMessage() error {
	err := messagesRepository.Write(
		Message{
			Type: GameStateType,
			Data: make([]byte, 0)})
	if err != nil {
		return err
	}
	return nil
}

func requestStateFromServer() ([][]rune, error) {
	for {
		m, e := messagesRepository.Read()
		if e != nil {
			return make([][]rune, 0), e
		}
		state := szr.DecodeGameState(m.Data)
		return state.State, nil
	}
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

func parseKeyCode(key dal.Key) CommandCode {
	switch key {
	case dal.KeyUp:
		return MoveLeft
	case dal.KeyDown:
		return MoveRight
	case dal.KeyLeft:
		return MoveUp
	case dal.KeyRight:
		return MoveDown
	case dal.KeyEsc:
		return ExitGame
	default:
		return UndefinedCommand
	}
}

func sendCommandToServer(command Command) error {
	err := messagesRepository.Write(
		Message{
			//	Id:   id,
			Type: CommandType,
			Data: szr.EncodeCommand(command)})
	if err != nil {
		return err
	}
	return nil
}

var messagesRepository dal.IMessagesRepository
var screen dal.IScreen
var udpClient dal.IUdpClient
var logger *log.Logger

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	ip := ""
	println("Inter ip:port to liscen")
	_, e := fmt.Scanln(&ip)
	if e != nil {
		println(e.Error())
		fmt.Scanln()
		return
	}

	ipServer := ""
	println("Inter ip:port of server")
	_, e = fmt.Scanln(&ipServer)
	if e != nil {
		println(e.Error())
		fmt.Scanln()
		return
	}

	factory := dal.CreateDalFactory()
	sc, err := factory.CreateScreen()
	if err != nil {
		print(err.Error())
		fmt.Scanln()
		return
	}
	screen = sc
	defer screen.Close()
	client, err := dal.NewUdpClient(ip, ipServer, 5)
	if err != nil {
		print(err.Error())
		fmt.Scanln()
		return
	}
	messagesRepository = factory.CreateMessagesRepository(client)
	defer messagesRepository.Dispose()
	f, err := os.Create("log.txt")
	if err != nil {
		print(err.Error())
		fmt.Scanln()
		return
	}
	defer f.Close()
	l := log.New(f, "main ", log.LstdFlags)
	logger = l

	go func() {
		for {
			s, e := requestStateFromServer()
			if e != nil {
				l.Println(e.Error())
				continue
			}
			showState(s)
		}
	}()

	go func() {
		var id uint64 = 1
		for {
			key := screen.ReadKey()
			code := parseKeyCode(key)
			if code > 0 {
				for i := 0; i < 4; i++ {
					e := sendCommandToServer(Command{id, code})
					id++
					if e != nil {
						logger.Println(e.Error())
					}
				}
			}
		}
	}()

	for {
		time.Sleep(time.Microsecond * 1000000 / 60)
		e := sendStateMessage()
		if e != nil {
			logger.Println(e.Error())
			continue
		}
	}
}
