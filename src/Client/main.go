package main

import (
	"fmt"

	szr "../Shared/serializer"
	"../Shared/udp"
	"./dal"
)

func sendStateMessage() error {
	err := messagesRepository.Write(
		szr.Message{
			Id:   messageCurrentId,
			Type: szr.GameStateType,
			Data: make([]byte, 0)})
	if err != nil {
		return err
	}
	messageCurrentId++
	return nil
}

func requestStateFromServer() ([][]rune, error) {
	//return createTestState(), nil
	for {
		m, e := messagesRepository.Read()
		if e != nil {
			return make([][]rune, 0), e
		}
		fmt.Printf("Recived message %v\n", m)
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
	for i := 0; i < 30; i++ {
		state = append(state, make([]rune, 30))
		for j := 0; j < 30; j++ {
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

	return szr.ExitGame
}

func writeCommandCodeToBuffer(code szr.CommandCode) {

}

func readCommandCodeFromBuffer() szr.CommandCode {
	return szr.ExitGame
}

func creteCommandWithCode(code szr.CommandCode) szr.Command {
	return szr.Command{}
}

func sendCommandToServer(command szr.Command) {

}

var messagesRepository dal.IMessagesRepository
var messageCurrentId uint64
var currentReceivedMessageId uint64
var stateBuffer chan [][]rune
var screen dal.IScreen

func main() {
	factory := dal.CreateDalFactory()
	client, err := udp.NewUdpClient("127.0.0.1:7788")
	if err != nil {
		print(err)
		return
	}
	screen, err = factory.CreateScreen()
	if err != nil {
		print(err)
		return
	}
	defer screen.Close()
	messagesRepository = factory.CreateMessagesRepository(client)
	messageCurrentId = 1
	currentReceivedMessageId = 0
	stateBuffer = make(chan [][]rune, 100)

	go func() {
		for {
			e := sendStateMessage()
			if e != nil {
				println(e.Error())
			}
		}
	}()
	go func() {
		for {
			s, e := requestStateFromServer()
			if e != nil {
				print(e.Error())
				continue
			}
			writeStateToBuffer(s)
		}
	}()

	go func() {
		for {
			writeCommandCodeToBuffer(readPressedKey())
		}
	}()

	go func() {
		for {
			sendCommandToServer(creteCommandWithCode(readCommandCodeFromBuffer()))
		}
	}()

	for {
		showState(readStateFromBuffer())
	}

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
