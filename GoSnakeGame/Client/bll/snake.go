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
	this.points.ForEach(func(i int, p *domainModels.Point) error {
		if i == 0 {
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
		} else {
			tempX, tempY := p.X, p.Y
			p.X, p.Y = oldX, oldY
			oldX = tempX
			oldY = tempY
		}
		return nil
	})
}

func (this *snake) TryEat(f ifood) {
	if f.isHit(*this.points.First()) {
		this.points.Add(*this.points.Last())
		f.Reset()
	}
}

func (this *snake) Reset() {
	points := make([]domainModels.Point, len(this.initialPoints))
	copy(points, this.initialPoints)
	this.direction = this.initialDirection
	this.points = dal.NewIPointRepositoryWithData(points)
}

func (this *snake) IsHitTail() bool {
	head := this.points.First()
	isHit := false
	this.points.ForEach(func(i int, p *domainModels.Point) error {
		if i != 0 && p.X == head.X && p.Y == head.Y {
			isHit = true
		}
		return nil
	})
	return isHit
}

func (this *snake) IsHit(frame ifigure) bool {
	head := this.points.First()
	return frame.isHit(*head)
}
