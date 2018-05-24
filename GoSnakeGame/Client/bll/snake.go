package bll

import (
	"../dal"
	"../domainModels"
)

type snake struct {
	figure
	direction        Direction
	initialDirection Direction
	initialPoints    []domainModels.Point
}

type isnake interface {
	ifigure
	Go(direction Direction)
	Move()
	TryEat(f ifood)
	IsHit(f ifigure) bool
	IsHitTail() bool
	Reset()
}

type Direction uint8

const RightDirection Direction = 1
const LeftDirection Direction = 2
const UpDirection Direction = 3
const DownDirection Direction = 4
const speed = 1

func (this *snake) Go(direction Direction) {
	this.direction = direction
}

func (this snake) Move() {
	var oldX int
	var oldY int
	p := this.points.Next()
	x, y := p.X, p.Y
	oldX = x
	oldY = y
	switch this.direction {
	case RightDirection:
		x = x + speed
	case LeftDirection:
		x = x - speed
	case UpDirection:
		y = y - speed
	case DownDirection:
		y = y + speed
	}
	p.X = x
	p.Y = y
	for p = this.points.Next(); p != nil; p = this.points.Next() {
		tempX, tempY := p.X, p.Y
		p.X, p.Y = oldX, oldY
		oldX = tempX
		oldY = tempY
	}
}

func (this *snake) TryEat(f ifood) {
	if f.isHit(*this.points.Head()) {
		this.points.AddToEnd(*this.points.Last())
		f.Reset()
	}
}

func (this *snake) Reset() {
	points := make([]domainModels.Point, len(this.initialPoints))
	copy(points, this.initialPoints)
	this.direction = this.initialDirection
	this.points = dal.NewILinkedListWithData(points)
}

func (this *snake) IsHitTail() bool {
	head := this.points.Next()
	for p := this.points.Next(); p != nil; p = this.points.Next() {
		if p.X == head.X && p.Y == head.Y {
			return true
		}
	}
	return false
}

func (this *snake) IsHit(frame ifigure) bool {
	head := this.points.Head()
	return frame.isHit(*head)
}
