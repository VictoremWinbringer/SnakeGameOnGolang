package main

import (
	"log"
	"time"

	"./bll"
	"./dal"
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
	dalFactory := dal.CreateDalFactory()
	screen, err := dalFactory.CreateScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		screen.Close()
	}()
	bllFactory := bll.NewBllFactory(dalFactory, screen)
	game := bll.NewGame(20, 40, bllFactory, screen)
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
}
