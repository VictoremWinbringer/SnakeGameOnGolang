package main

import (
	"fmt"

	fd "./food"
	f "./frame"
	p "./point"
	s "./snake"
	tc "github.com/gdamore/tcell"
)

func main() {
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
	for {
		screen.Clear()
		h, w := 30, 90
		writer := p.NewTerminalWriter(screen)
		frame := f.New(h, w, '+', writer)
		food := fd.New(10, 10, '$', writer)
		snake := s.New(8, 8, '+', writer)
		frame.Draw()
		food.Draw()
		snake.Draw()
		screen.Show()
	}
}
