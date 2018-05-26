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

// func (this dalFactory) CreateScreen() (IScreen, error) {
// 	s, e := tc.NewScreen()
// 	if e != nil {
// 		return nil, e
// 	}
// 	if e := s.Init(); e != nil {
// 		return nil, e
// 	}
// 	s.SetStyle(tc.StyleDefault)
// 	s.HideCursor()
// 	return screen{s}, nil
// }

func (this dalFactory) CreateWriter(width, height int) IWriter {
	return &writer{make([]rune, width*height), height, width}
}
