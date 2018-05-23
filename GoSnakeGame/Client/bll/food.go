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

func newIFood(x, y, maxX, maxY int, value rune, writer dal.IPointWriter) ifood {
	return food{newIPoint(x, y, value, writer), maxX, maxY}
}

func (f food) Reset() {
	f.Move(rand.Intn(f.maxX), rand.Intn(f.maxY))
}
