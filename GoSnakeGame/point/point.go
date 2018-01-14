package point

import tc "github.com/gdamore/tcell"

type Direction int

var Left Direction = 1
var Right Direction = 2
var Up Direction = 3
var Down Direction = 4

type Point struct {
	x      int
	y      int
	value  rune
	matrix tc.Screen
}

func New(x int, y int, value rune, matrix tc.Screen) Point {
	return Point{x, y, value, matrix}
}

func (point *Point) Move(direction Direction) {
	switch direction {
	case 1:
		point.x--
	case 2:
		point.x++
	case 3:
		point.y++
	case 4:
		point.y--
	}
}

func (p Point) Draw() {
	p.matrix.SetContent(p.x, p.y, p.value, nil, tc.StyleDefault)
}

func (this Point) Overlaps(other Point) bool {
	return this.x == other.x && this.y == other.y
}
