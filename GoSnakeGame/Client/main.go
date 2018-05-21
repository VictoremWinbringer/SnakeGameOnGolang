package main

import (
	"log"
	"time"

	_ "../Shared/serializer"
	gameModule "./game"
	_ "./udpClient"
)

func main() {
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
	game, err := gameModule.New(50, 50)
	if err != nil {
		log.Fatal(err)
	}
	timeCurrent := time.Now()
	go func() {
		for {
			timeNow := time.Now()
			game.Logic(timeCurrent.UnixNano() - timeNow.UnixNano())
			timeCurrent = timeNow
		}
	}()
	for {
		time.Sleep(time.Millisecond * 100)
		game.Draw()
	}
}
