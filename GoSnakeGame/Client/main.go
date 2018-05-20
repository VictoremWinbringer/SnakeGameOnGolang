package main

import (
	"fmt"
	"time"

	_ "../Shared/serializer"
	fd "./food"
	f "./frame"
	p "./point"
	s "./snake"
	_ "./udpClient"
	tc "github.com/gdamore/tcell"
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
	fmt.Println("Press enter to start game.")
	screen, e := tc.NewScreen()
	if e != nil {
		fmt.Println(e)
		fmt.Scanln()
		return
	}
	if e := screen.Init(); e != nil {
		fmt.Println(e)
		fmt.Scanln()
		return
	}
	screen.SetStyle(tc.StyleDefault)
	screen.HideCursor()
	h, w := 30, 90
	writer := p.NewTerminalWriter(screen)
	frame := f.New(h, w, '+', writer)
	food := fd.New(10, 10, '$', writer)
	snake := s.New(8, 8, '+', writer)
	channel := make(chan interface{})
	go func() {
		for {
			event := screen.PollEvent()
			switch keyEvent := event.(type) {
			case *tc.EventKey:
				switch keyEvent.Key() {
				case tc.KeyEsc:
					screen.Clear()
					screen.Show()
					screen.ShowCursor(0, 0)
					channel <- new(interface{})
				case tc.KeyUp:
					snake.Go(s.Up)
				case tc.KeyDown:
					snake.Go(s.Down)
				case tc.KeyLeft:
					snake.Go(s.Left)
				case tc.KeyRight:
					snake.Go(s.Right)
				}
			}
		}
	}()
	for {
		select {
		case <-channel:
			return

		default:
			screen.Clear()
			frame.Draw()
			food.Draw()
			snake.Move()
			snake.Draw()
			screen.Show()
			time.Sleep(time.Millisecond * 100)
		}
	}
}
