package point

import "../matrix"

type Direction int

var Left Direction = 1
var Right Direction = 2
var Up Direction = 3
var Down Direction = 4

type Point struct {
	x      int
	y      int
	value  rune
	matrix *matrix.Matrix
}

func New(x int, y int, value rune, matrix *matrix.Matrix) Point {
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

func (p Point) Draw() error {
	return p.matrix.Set(p.y, p.x, p.value)
}
