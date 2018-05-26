package dal

import "../model"

type pointRepository struct {
	points []model.Point
	count  int
}

type IPointRepository interface {
	Add(model.Point)
	First() *model.Point
	Last() *model.Point
	Count() int
	ForEach(func(i int, point *model.Point) error) error
	Clear()
}

func (this *pointRepository) Clear() {
	this.points = make([]model.Point, 0)
	this.count = 0
}

func (this *pointRepository) Count() int {
	return this.count
}
func (this *pointRepository) Add(point model.Point) {
	this.points = append(this.points, point)
	this.count += 1
}

func (this *pointRepository) First() *model.Point {
	if this.count < 1 {
		return nil
	}
	return &this.points[0]
}

func (this *pointRepository) Last() *model.Point {
	if this.count < 1 {
		return nil
	}
	return &this.points[this.count-1]
}

func (this *pointRepository) ForEach(action func(i int, point *model.Point) error) error {
	for i, p := range this.points {
		err := action(i, &p)
		if err != nil {
			return err
		}
		this.points[i] = p
	}
	return nil
}
