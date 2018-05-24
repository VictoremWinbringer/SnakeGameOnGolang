package dal

import "../domainModels"

type pointRepository struct {
	points []domainModels.Point
	count  int
}

type IPointRepository interface {
	Add(domainModels.Point)
	First() *domainModels.Point
	Last() *domainModels.Point
	Count() int
	ForEach(func(i int, point *domainModels.Point) error) error
}

func NewIPointRepositoryWithData(points []domainModels.Point) IPointRepository {
	repository := NewIPointRepository()
	for _, p := range points {
		repository.Add(p)
	}
	return repository
}

func NewIPointRepository() IPointRepository {
	return &pointRepository{make([]domainModels.Point, 0), 0}
}

func (this *pointRepository) Count() int {
	return this.count
}
func (this *pointRepository) Add(point domainModels.Point) {
	this.points = append(this.points, point)
	this.count += 1
}

func (this *pointRepository) First() *domainModels.Point {
	if this.count < 1 {
		return nil
	}
	return &this.points[0]
}

func (this *pointRepository) Last() *domainModels.Point {
	if this.count < 1 {
		return nil
	}
	return &this.points[this.count-1]
}

func (this *pointRepository) ForEach(action func(i int, point *domainModels.Point) error) error {
	for i, p := range this.points {
		err := action(i, &p)
		if err != nil {
			return err
		}
		this.points[i] = p
	}
	return nil
}
