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
}

func newISnake(x, y int, value rune, writer dal.IPointWriter) isnake {

	points := make([]ipoint, 0)
	initialPoints := make([]ipoint, 0)
	points = append(points, newIPoint(x, y, value, writer))
	points = append(points, newIPoint(x-1, y, value, writer))
	points = append(points, newIPoint(x-2, y, value, writer))
	copy(initialPoints, points)
	return snake{figure{(points)}, Right, Right, initialPoints}
}

type Direction uint8

const Right Direction = 1
const Left Direction = 2
const Up Direction = 3
const Down Direction = 4

func (s *isnake) Go(direction Direction) {
	s.direction = direction
}

func (s *isnake) Move() {
	var oldX int
	var oldY int
	for i, p := range s.Points {
		if i == 0 {
			x, y := p.Position()
			oldX = x
			oldY = y
			switch s.direction {
			case Right:
				x = x + 1
			case Left:
				x = x - 1
			case Up:
				y = y - 1
			case Down:
				y = y + 1
			}
			p.Move(x, y)
		} else {
			tempX, tempY := p.Position()
			p.Move(oldX, oldY)
			oldX = tempX
			oldY = tempY
		}
		s.Points[i] = p
	}
}

func (s *snake) TryEat(f ifood) {
	if s.Points[0].Overlaps(f.Point) {
		last := s.Points[len(s.Points)-1]
		s.Points = append(s.Points, last)
		f.Reset()
	}
}

func (s *snake) Reset() {
	s.Points = make([]ipoint, 0)
	copy(s.Points, s.initialPoints)
	s.direction = s.initialDirection
}

func (s *snake) IsHitTail() bool {
	head := s.Points[0]
	for i, p := range s.Points {
		if i > 0 && head.Overlaps(p) {
			return true
		}
	}
	return false
}

func (s *snake) IsHit(f figure) bool {
	head := s.Points[0]
	return f.IsHitPoint(head)
}
