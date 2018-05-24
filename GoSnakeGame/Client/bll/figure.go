package bll

import (
	"../dal"
	"../domainModels"
)

type figure struct {
	points dal.IPointRepository
	writer IWriter
}

type ifigure interface {
	draw()
	isHit(point domainModels.Point) bool
}

func (this figure) draw() {
	this.points.ForEach(func(i int, p *domainModels.Point) error {
		this.writer.Write(p.X, p.Y, p.Symbol)
		return nil
	})
}

func (this figure) isHit(point domainModels.Point) bool {
	isHit := false
	this.points.ForEach(func(i int, p *domainModels.Point) error {
		if p.X == point.X && p.Y == point.Y {
			isHit = true
		}
		return nil
	})
	return isHit
}
