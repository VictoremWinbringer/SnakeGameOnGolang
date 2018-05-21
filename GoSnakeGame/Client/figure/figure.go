package figure

import (
	"../point"
)

type Figure struct {
	Points []point.Point
}

func (this Figure) Draw() {
	for i := range this.Points {
		this.Points[i].Draw()
	}
}

func New(points []point.Point) Figure {
	return Figure{points}
}

func (this Figure) IsHitPoint(point point.Point) bool {
	for _, p := range this.Points {
		if p.Overlaps(point) {
			return true
		}
	}
	return false
}

func (this Figure) IsHitFigure(f Figure) bool {
	for _, p := range this.Points {
		if f.IsHitPoint(p) {
			return true
		}
	}
	return false
}
