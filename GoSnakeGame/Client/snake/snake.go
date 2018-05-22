package snake

import (
	"../figure"
	"../food"
	"../point"
)

type Snake struct {
	figure.Figure
	direction        Direction
	initialDirection Direction
	initialPoints    []point.Point
}

func New(x, y int, value rune, writer point.PointWriter) Snake {

	points := make([]point.Point, 0)
	initialPoints := make([]point.Point, 0)
	points = append(points, point.New(x, y, value, writer))
	points = append(points, point.New(x-1, y, value, writer))
	points = append(points, point.New(x-2, y, value, writer))
	copy(initialPoints, points)
	return Snake{figure.New(points), Right, Right, initialPoints}
}

type Direction uint8

const Right Direction = 1
const Left Direction = 2
const Up Direction = 3
const Down Direction = 4

func (s *Snake) Go(direction Direction) {
	s.direction = direction
}

func (s *Snake) Move() {
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

func (s *Snake) TryEat(f *food.Food) {
	if s.Figure.Points[0].Overlaps(f.Point) {
		last := s.Points[len(s.Points)-1]
		s.Points = append(s.Points, last)
		f.Reset()
	}
}

func (s *Snake) Reset() {
	s.Points = make([]point.Point, 0)
	copy(s.Points, s.initialPoints)
	s.direction = s.initialDirection
}

func (s *Snake) IsHitTail() bool {
	head := s.Points[0]
	for i, p := range s.Points {
		if i > 0 && head.Overlaps(p) {
			return true
		}
	}
	return false
}

func (s *Snake) IsHit(f figure.Figure) bool {
	head := s.Points[0]
	return f.IsHitPoint(head)
}
