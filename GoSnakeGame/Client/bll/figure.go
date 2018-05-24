package bll

import (
	"../dal"
	"../domainModels"
)

type figure struct {
	points dal.ILinkedList
	writer IWriter
}

type ifigure interface {
	draw()
	isHit(point domainModels.Point) bool
}

func (this figure) draw() {
	for p := this.points.Next(); p != nil; p = this.points.Next() {
		this.writer.Write(p.X, p.Y, p.Symbol)
	}
}

func (this figure) isHit(point domainModels.Point) bool {
	for p := this.points.Next(); p != nil; p = this.points.Next() {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}
