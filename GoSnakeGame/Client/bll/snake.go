package bll

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

func newISnake(x, y int, value rune, writer IWriter) isnake {

	points := make([]ipoint, 0)
	initialPoints := make([]ipoint, 0)
	for i := 0; i < initialLenth; i++ {
		points = append(points, newIPoint(x-i, y, value, writer))
		initialPoints = append(initialPoints, newIPoint(x-i, y, value, writer))
	}
	return &snake{figure{(points)}, RightDirection, RightDirection, initialPoints}
}

type Direction uint8

const RightDirection Direction = 1
const LeftDirection Direction = 2
const UpDirection Direction = 3
const DownDirection Direction = 4

const initialLenth = 3
const speed = 1

func (this *snake) Go(direction Direction) {
	this.direction = direction
}

func (this snake) Move() {
	var oldX int
	var oldY int
	for i, p := range this.points {
		if i == 0 {
			x, y := p.Position()
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
			p.Move(x, y)
		} else {
			tempX, tempY := p.Position()
			p.Move(oldX, oldY)
			oldX = tempX
			oldY = tempY
		}
		this.points[i] = p
	}
}

func (this *snake) TryEat(f ifood) {
	if this.points[0].Overlaps(f) {
		last := this.points[len(this.points)-1]
		this.points = append(this.points, last.Copy())
		f.Reset()
	}
}

func (this *snake) Reset() {
	this.points = make([]ipoint, 0)
	for _, point := range this.initialPoints {
		this.points = append(this.points, point.Copy())
	}
	this.direction = this.initialDirection
}

func (this *snake) IsHitTail() bool {
	head := this.points[0]
	for i, p := range this.points {
		if i > 0 && head.Overlaps(p) {
			return true
		}
	}
	return false
}

func (this *snake) IsHit(frame iframe) bool {
	head := this.points[0]
	return frame.isHitPoint(head)
}
