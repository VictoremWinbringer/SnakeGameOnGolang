package snake

import (
	"../figure"
	"../point"
)

type Snake struct {
	figure.Figure
}

func New(x, y int, value rune, writer point.PointWriter) Snake {

	points := make([]point.Point, 0)
	points = append(points, point.New(x, y, value, writer))
	points = append(points, point.New(x-1, y, value, writer))
	points = append(points, point.New(x-2, y, value, writer))

	return Snake{figure.New(points)}
}
