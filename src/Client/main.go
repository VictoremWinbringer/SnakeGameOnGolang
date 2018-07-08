package main

import (
	szr "../Shared/serializer"
	"../Shared/udp"
)

func requesStateFromServer() [][]rune {
	 client,err  := udp.NewUdpClient("127.0.0.1:7788")
	if err  != nil {
		panic(err)
	}
	client.Write(szr.EncodeMessage(szr.Message{
		Id: 1,
		Type: szr.GameStateType,
		Data: make([]byte, 0)}))
	return make([][]rune, 0)
}

func writeStateToBuffer(state [][]rune)  {

}

func readStateFromBuffer() [][]rune  {
	return make([][]rune, 0)
}

func showState(state [][]rune) {

}

func readPressedKey() szr.CommandCode  {

	return szr.ExitGame
}

func writeCommanCodeToBuffer(code szr.CommandCode) {
		
}

func readCommandCodeFromBuffer() szr.CommandCode {
	return szr.ExitGame
}

func creteCommandWithCode(code szr.CommandCode) szr.Command {
return szr.Command{}
}

func sendCommandToServer(command szr.Command) {

}

func main() {
	go func (){
		for {
			writeStateToBuffer(requesStateFromServer())
		}
	}()

	for{
		showState(readStateFromBuffer())
	}

	go func(){
		for {
			writeCommanCodeToBuffer(readPressedKey())
		}
	}()

	go func() {
		for {
			sendCommandToServer(creteCommandWithCode(readCommandCodeFromBuffer()))
		}
	}()
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
