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
	Overlaps(other Point) bool
	Position() (x int, y int)
}

func ipoint newIPoint(x,y int,value rune,writer dal.IPointWriter){
	return point{x,y,value,writer}
}

func (this *point) Move(x int, y int) {
	this.x = x
	this.y = y
}

func (this point) Draw() {
	this.writer.Write(this.x, this.y, this.value)
}

func (this point) Overlaps(other Point) bool {
	return this.x == other.x && this.y == other.y
}
func (this point) Position() (x int, y int) {
	return this.x, this.y
}
