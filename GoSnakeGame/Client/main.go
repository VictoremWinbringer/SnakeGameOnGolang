package main

import (
	"fmt"
	"strings"

	// "fmt"
	// "time"

	// fd "./food"
	// f "./frame"
	// p "./point"
	// s "./snake"
	u "./udpClient"
	// tc "github.com/gdamore/tcell"
)

func main() {
	upd, err := u.New("127.0.0.1:8888")
	if err != nil {
		fmt.Printf("Error on start client %v", err)
		return
	}
	defer func() {
		upd.Close()
	}()
	_, err = upd.Write([]byte("Hello Sever!"))
	if err != nil {
		fmt.Printf("Error on write %v", err)
		return
	}
	buffer := make([]byte, 4096)
	_, err = upd.Read(buffer)
	if err != nil {
		fmt.Printf("Error on read %v", err)
		return
	}
	fmt.Println(strings.Trim(string(buffer), " "))
	fmt.Println("Press enter to start game.")
	fmt.Scanln()
	// screen, e := tc.NewScreen()
	// if e != nil {
	// 	fmt.Println(e)
	// 	fmt.Scanln()
	// 	return
	// }
	// if e := screen.Init(); e != nil {
	// 	fmt.Println(e)
	// 	fmt.Scanln()
	// 	return
	// }
	// screen.SetStyle(tc.StyleDefault)
	// screen.HideCursor()
	// h, w := 30, 90
	// writer := p.NewTerminalWriter(screen)
	// frame := f.New(h, w, '+', writer)
	// food := fd.New(10, 10, '$', writer)
	// snake := s.New(8, 8, '+', writer)
	// for {
	// 	screen.Clear()
	// 	frame.Draw()
	// 	food.Draw()
	// 	snake.Draw()
	// 	screen.Show()
	// 	time.Sleep(time.Millisecond * 100)
	// }
}
