package point

import tc "github.com/gdamore/tcell"

type PointWriter interface {
	Write(x, y int, value rune)
}
type terminalWriter struct {
	screen tc.Screen
}

func NewTerminalWriter(s tc.Screen) PointWriter {
	return terminalWriter{s}
}

func (t terminalWriter) Write(x, y int, value rune) {
	t.screen.SetContent(x, y, value, nil, tc.StyleDefault)
}

type Point struct {
	x      int
	y      int
	value  rune
	writer PointWriter
}

func New(x int, y int, value rune, writer PointWriter) Point {
	return Point{x, y, value, writer}
}

func (point *Point) Move(x int, y int) {
	point.x = x
	point.y = y
}

func (this Point) Draw() {
	this.writer.Write(this.x, this.y, this.value)
}

func (this Point) Overlaps(other Point) bool {
	return this.x == other.x && this.y == other.y
}
func (this Point) Position() (x int, y int) {
	return this.x, this.y
}
