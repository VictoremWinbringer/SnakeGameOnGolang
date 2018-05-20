package frame

import (
	"../figure"
	"../point"
)

type Frame struct {
	figure.Figure
}

func New(h, w int, value rune, writer point.PointWriter) Frame {

	points := make([]point.Point, 0)
	points = addHorizontal(w, 0, points, value, writer)
	points = addHorizontal(w, h, points, value, writer)
	points = addVertical(h, 0, points, value, writer)
	points = addVertical(h, w, points, value, writer)

	return Frame{figure.New(points)}
}

func addHorizontal(w, y int, points []point.Point, value rune, writer point.PointWriter) []point.Point {
	for i := 0; i <= w; i++ {
		points = append(points, point.New(i, y, value, writer))
	}
	return points
}

func addVertical(h, x int, points []point.Point, value rune, writer point.PointWriter) []point.Point {
	for i := 0; i <= h; i++ {
		points = append(points, point.New(x, i, value, writer))
	}
	return points
}
