package bll

import (
	"math/rand"

	"../dal"
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

func NewIFood(x, y, maxX, maxY int, value rune, writer dal.IPointWriter) Food {
	return Food{newIPoint(x, y, value, writer), maxX, maxY}
}

func (f *food) Reset() {
	f.Move(rand.Intn(f.maxX), rand.Intn(f.maxY))
}
