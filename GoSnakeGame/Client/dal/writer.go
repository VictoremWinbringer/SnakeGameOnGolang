package dal

import tc "github.com/gdamore/tcell"

type IPointWriter interface {
	Write(x, y int, value rune)
}

type terminalWriter struct {
	screen tc.Screen
}

func NewIPointWriter(s tc.Screen) IPointWriter {
	return terminalWriter{s}
}

func (this terminalWriter) Write(x, y int, value rune) {
	this.screen.SetContent(x, y, value, nil, tc.StyleDefault)
}
