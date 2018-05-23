package bll

import (
	"../dal"
)

type snake struct {
	figure
	direction        Direction
	initialDirection Direction
	initialPoints    []ipoint
}

type isnake interface {
	Draw()
	Go(direction Direction)
	Move()
	TryEat(f ifood)
	IsHit(f iframe) bool
	IsHitTail() bool
	Reset()
}

func newISnake(x, y int, value rune, writer dal.IPointWriter) isnake {

	points := make([]ipoint, 0)
	initialPoints := make([]ipoint, 0)
	points = append(points, newIPoint(x, y, value, writer))
	points = append(points, newIPoint(x-1, y, value, writer))
	points = append(points, newIPoint(x-2, y, value, writer))
	copy(initialPoints, points)
	return &snake{figure{(points)}, RightDirection, RightDirection, initialPoints}
}

type Direction uint8

const RightDirection Direction = 1
const LeftDirection Direction = 2
const UpDirection Direction = 3
const DownDirection Direction = 4

func (s *snake) Go(direction Direction) {
	s.direction = direction
}

func (s snake) Move() {
	var oldX int
	var oldY int
	for i, p := range s.points {
		if i == 0 {
			x, y := p.Position()
			oldX = x
			oldY = y
			switch s.direction {
			case RightDirection:
				x = x + 1
			case LeftDirection:
				x = x - 1
			case UpDirection:
				y = y - 1
			case DownDirection:
				y = y + 1
			}
			p.Move(x, y)
		} else {
			tempX, tempY := p.Position()
			p.Move(oldX, oldY)
			oldX = tempX
			oldY = tempY
		}
		s.points[i] = p
	}
}

func (s *snake) TryEat(f ifood) {
	if s.points[0].Overlaps(f) {
		last := s.points[len(s.points)-1]
		s.points = append(s.points, last)
		f.Reset()
	}
}

func (s *snake) Reset() {
	s.points = make([]ipoint, 0)
	copy(s.points, s.initialPoints)
	s.direction = s.initialDirection
}

func (s *snake) IsHitTail() bool {
	head := s.points[0]
	for i, p := range s.points {
		if i > 0 && head.Overlaps(p) {
			return true
		}
	}
	return false
}

func (s *snake) IsHit(f iframe) bool {
	head := s.points[0]
	return f.isHitPoint(head)
}
