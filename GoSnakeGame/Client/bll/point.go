package bll

import "../dal"

type point struct {
	x      int
	y      int
	value  rune
	writer dal.IPointWriter
}

type ipoint interface {
	Move(x int, y int)
	Draw()
	Overlaps(other ipoint) bool
	Position() (x int, y int)
}

func newIPoint(x, y int, value rune, writer dal.IPointWriter) ipoint {
	return &point{x, y, value, writer}
}

func (this *point) Move(x int, y int) {
	this.x = x
	this.y = y
}

func (this point) Draw() {
	this.writer.Write(this.x, this.y, this.value)
}

func (this point) Overlaps(other ipoint) bool {
	x, y := other.Position()
	return this.x == x && this.y == y
}
func (this point) Position() (x int, y int) {
	return this.x, this.y
}
