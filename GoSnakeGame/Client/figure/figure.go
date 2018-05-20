package figure

import (
	"../point"
)

type Figure struct {
	points []point.Point
}

func (this Figure) Draw() {
	for i := range this.points {
		this.points[i].Draw()
	}
}

func New(points []point.Point) Figure {
	return Figure{points}
}

func (this Figure) IsHitPoint(point point.Point) bool {
	for _, p := range this.points {
		if p.Overlaps(point) {
			return true
		}
	}
	return false
}

func (this Figure) IsHitFigure(f Figure) bool {
	for _, p := range this.points {
		if f.IsHitPoint(p) {
			return true
		}
	}
	return false
}
