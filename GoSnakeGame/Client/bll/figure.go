package bll

type figure struct {
	points []ipoint
}

func (this figure) Draw() {
	for i := range this.points {
		this.points[i].Draw()
	}
}
