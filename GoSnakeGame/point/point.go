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

type Direction int

const Left Direction = 2
const Right Direction = 1
const Up Direction = 4
const Down Direction = 3

type Point struct {
	x      int
	y      int
	value  rune
	writer PointWriter
}

func New(x int, y int, value rune, writer PointWriter) Point {
	return Point{x, y, value, writer}
}

func (point *Point) Move(direction Direction) {
	switch direction {
	case Right:
		point.x--
	case Left:
		point.x++
	case Down:
		point.y++
	case Up:
		point.y--
	}
}

func (p Point) Draw() {
	p.writer.Write(p.x, p.y, p.value)
}

func (this Point) Overlaps(other Point) bool {
	return this.x == other.x && this.y == other.y
}
