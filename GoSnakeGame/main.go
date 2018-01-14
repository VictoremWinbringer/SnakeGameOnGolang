package main

import (
	"fmt"
	"time"

	"./point"
	tc "github.com/gdamore/tcell"
)

func main() {

	s, e := tc.NewScreen()
	if e != nil {
		fmt.Println(e)
		fmt.Scanln()
		return
	}

	if e := s.Init(); e != nil {
		fmt.Println(e)
		fmt.Scanln()
		return
	}

	s.SetStyle(tc.StyleDefault)

	s.HideCursor()

	w, h := s.Size()

	p := point.New(3, 3, '*', point.NewTerminalWriter(s))

	p.Draw()

	s.Clear()

	st := fmt.Sprint("Current Time:", time.Now().Format(time.RFC1123)+"\n\r", w, h)

	for i, v := range st {
		s.SetContent(i, 2, v, nil, tc.StyleDefault)
	}

	s.Show()

	time.Sleep(time.Second)

	p.Draw()

	s.Show()

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {

		s.Clear()
		if i%2 == 0 {
			p.Move(point.Down)
		} else {
			p.Move(point.Left)
		}

		p.Draw()
		s.Show()
		time.Sleep(time.Second)
	}

	s.ShowCursor(1, 1)
}
