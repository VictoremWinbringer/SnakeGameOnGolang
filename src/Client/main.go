package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	szr "../Shared/serializer"
	"../Shared/udp"
	"./dal"
)

func sendStateMessage() error {
	mtx.Lock()
	defer mtx.Unlock()
	err := messagesRepository.Write(
		szr.Message{
			Id:   messageCurrentId,
			Type: szr.GameStateType,
			Data: make([]byte, 0)})
	messageCurrentId++
	if err != nil {
		return err
	}
	return nil
}

func requestStateFromServer() ([][]rune, error) {
	mtx.Lock()
	defer mtx.Unlock()
	for {
		m, e := messagesRepository.Read()
		if e != nil {
			return make([][]rune, 0), e
		}
		if m.Id <= currentReceivedMessageId {
			continue
		}
		currentReceivedMessageId = m.Id
		state := szr.DecodeGameState(m.Data)
		return state.State, nil
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

func readPressedKey() szr.CommandCode {
	key := screen.ReadKey()
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

func writeCommandCodeToBuffer(code szr.CommandCode) {
	commandBuffer <- code
}

func readCommandCodeFromBuffer() szr.CommandCode {
	return <-commandBuffer
}

func creteCommandWithCode(code szr.CommandCode, id uint64) szr.Command {
	return szr.Command{id, code}
}

func sendCommandToServer(command szr.Command) error {
	mtx.Lock()
	defer mtx.Unlock()
	logger.Println(command.Code)
	logger.Println(command.Id)
	err := messagesRepository.Write(
		szr.Message{
			Id:   messageCurrentId,
			Type: szr.CommandType,
			Data: szr.EncodeCommand(command)})
	messageCurrentId++
	if err != nil {
		return err
	}
	return nil
}

var messagesRepository dal.IMessagesRepository
var messageCurrentId uint64
var currentReceivedMessageId uint64
var stateBuffer chan [][]rune
var screen dal.IScreen
var commandBuffer chan szr.CommandCode
var currentCommandId uint64
var udpClient udp.IUdpClient
var mtx sync.Mutex
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
	currentCommandId = 0
	commandBuffer = make(chan szr.CommandCode)
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
	messageCurrentId = 1
	currentReceivedMessageId = 0
	stateBuffer = make(chan [][]rune, 100)
	f, err := os.Create("log.txt")
	check(err)
	defer f.Close()
	l := log.New(f, "main ", log.LstdFlags)
	logger = l

	go func() {
		for {
			code := readPressedKey()
			if code < 1 {
				continue
			}
			currentCommandId++
			id := currentCommandId
			for i := 0; i < 4; i++ {
				e := sendCommandToServer(creteCommandWithCode(code, id))
				if e != nil {
					l.Println(e.Error())
				}
			}
		}
	}()

	//go func() {
	for {
		e := sendStateMessage()
		if e != nil {
			l.Println(e.Error())
			continue
		}
		s, e := requestStateFromServer()
		if e != nil {
			l.Println(e.Error())
			continue
		}
		showState(s)
	}
	//}()

	//fmt.Scanln()

	// upd, err := u.New("127.0.0.1:8888")
	// if err != nil {
	// 	fmt.Printf("Error on start client %v", err)
	// 	return
	// }
	// defer func() {
	// 	upd.Close()
	// }()
	// _, err = upd.Write(ser.EncodeGameState(ser.GameState{"Hello Sever!"}))
	// if err != nil {
	// 	fmt.Printf("Error on write %v", err)
	// 	return
	// }
	// buffer := make([]byte, 4096)
	// _, err = upd.Read(buffer)
	// if err != nil {
	// 	fmt.Printf("Error on read %v", err)
	// 	return
	// }
	// fmt.Println(ser.DecodeGameState(buffer))
	/*	game, err := al.NewGame(20, 40)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			game.Close()
		}()
		timeCurrent := time.Now()
		c := make(chan int)
		go func() {
			for {
				timeNow := time.Now()
				if !game.Logic(timeNow.UnixNano() - timeCurrent.UnixNano()) {
					c <- 1
					return
				}
				timeCurrent = timeNow
			}
		}()
		for {
			select {
			case <-c:
				return
			default:
				game.Draw()
				time.Sleep(time.Millisecond * 18)
			}
		}
	*/
}
