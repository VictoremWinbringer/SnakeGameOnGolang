package bll

import (
	"../dal"
	"../model"
)

type figure struct {
	points dal.IPointRepository
	writer IWriter
}

type IFigure interface {
	Draw()
	isHit(point model.Point) bool
}

func (this figure) Draw() {
	this.points.ForEach(func(i int, p *model.Point) error {
		this.writer.Write(p.X, p.Y, p.Symbol)
		return nil
	})
}

func (this figure) isHit(point model.Point) bool {
	isHit := false
	this.points.ForEach(func(i int, p *model.Point) error {
		if p.X == point.X && p.Y == point.Y {
			isHit = true
		}
		return nil
	})
	return isHit
}
