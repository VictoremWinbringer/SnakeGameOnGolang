package bll

type figure struct {
	points []Point
}

func (this figure) Draw() {
	for i := range this.Points {
		this.Points[i].Draw()
	}
}

// func (this figure) isHitPoint(point Point) bool {
// 	for _, p := range this.Points {
// 		if p.Overlaps(point) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (this figure) isHitFigure(f Figure) bool {
// 	for _, p := range this.Points {
// 		if f.IsHitPoint(p) {
// 			return true
// 		}
// 	}
// 	return false
// }
