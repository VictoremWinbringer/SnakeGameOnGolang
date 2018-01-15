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

func (this Figure) IsHit(p point.Point) bool {
	for _, v := range this.points {
		if p.Overlaps(v) {
			return true
		}
	}
	return false
}
