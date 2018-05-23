package bll

import (
	"math/rand"
)

type food struct {
	ipoint
	maxX int
	maxY int
}

type ifood interface {
	ipoint
	Reset()
}

func newIFood(x, y, maxX, maxY int, value rune, writer IWriter) ifood {
	return food{newIPoint(x, y, value, writer), maxX, maxY}
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
	this.Move(x, y)
}
