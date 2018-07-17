package dal

import (
	"../model"
)

type IDalFactory interface {
	CreatePointRepository(points ...model.Point) IPointRepository
	//CreateScreen() (IScreen, error)
	CreateWriter(width, height int) IWriter
}

type dalFactory struct {
}

func CreateDalFactory() IDalFactory {
	return dalFactory{}
}

func (this dalFactory) CreatePointRepository(points ...model.Point) IPointRepository {
	repository := pointRepository{make([]model.Point, 0), 0}
	for _, p := range points {
		repository.Add(p)
	}
	return &repository
}

func (this dalFactory) CreateWriter(width, height int) IWriter {
	matrix := make([][]rune, width)
	for i := 0; i < width; i++ {
		matrix[i] = make([]rune, height)
	}
	return &writer{matrix, height, width}
}
