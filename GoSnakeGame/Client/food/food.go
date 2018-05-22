package food

import "../point"

type Food struct {
	point.Point
	initialX int
	initialY int
}

func New(x, y int, value rune, writer point.PointWriter) Food {
	return Food{point.New(x, y, value, writer), x, y}
}

func (f *Food) Reset() {
	f.Move(f.initialX, f.initialY)
}
