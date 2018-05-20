package food

import "../point"

type Food struct {
	point.Point
}

func New(x, y int, value rune, writer point.PointWriter) Food {
	return Food{point.New(x, y, value, writer)}
}
