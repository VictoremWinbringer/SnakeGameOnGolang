package bll

import (
	"math/rand"
)

type food struct {
	figure
	maxX int
	maxY int
}

type ifood interface {
	ifigure
	Reset()
}

func (this food) Reset() {
	x := rand.Intn(this.maxX)
	y := rand.Intn(this.maxY)
	if x < 1 {
		x = 1
	}
	if y < 1 {
		y = 1
	}
	this.move(x, y)
}

func (this food) move(x, y int) {
	head := this.points.First()
	head.X = x
	head.Y = y
}
