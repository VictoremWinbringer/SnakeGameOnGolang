package food

import (
	"math/rand"

	"../point"
)

type Food struct {
	point.Point
	maxX int
	maxY int
}

func New(x, y, maxX, maxY int, value rune, writer point.PointWriter) Food {
	return Food{point.New(x, y, value, writer), maxX, maxY}
}

func (f *Food) Reset() {
	f.Move(rand.Intn(f.maxX), rand.Intn(f.maxY))
}
