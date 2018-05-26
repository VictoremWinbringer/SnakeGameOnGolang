package bll

import (
	"../dal"
	"../model"
)

type IBllFactory interface {
	CrateFrame(h, w int, value rune) IFigure
	CreateFood(x, y int, value rune, maxX, maxY int) IFood
	CreateSnake(x, y int, value rune) ISnake
}

type bllFactory struct {
	dalFactory dal.IDalFactory
	writer     dal.IWriter
}

func NewBllFactory(dalFactory dal.IDalFactory, writer dal.IWriter) IBllFactory {
	return bllFactory{dalFactory, writer}
}

func (this bllFactory) CrateFrame(h, w int, value rune) IFigure {
	list := this.dalFactory.CreatePointRepository(createFrame(h, w, value)...)
	return figure{list, this.writer}
}

func (this bllFactory) CreateSnake(x, y int, value rune) ISnake {
	points := make([]model.Point, 0)
	for i := 0; i < initialLenth; i++ {
		points = append(points, model.Point{x - i, y, value})
	}
	initialPoints := make([]model.Point, initialLenth)
	copy(initialPoints, points)
	return &snake{figure{this.dalFactory.CreatePointRepository(points...), this.writer}, RightDirection, RightDirection, initialPoints}
}

func (this bllFactory) CreateFood(x, y int, value rune, maxX, maxY int) IFood {
	points := make([]model.Point, 0)
	points = append(points, model.Point{x, y, value})
	return food{figure{this.dalFactory.CreatePointRepository(points...), this.writer}, maxX, maxY}
}

const initialLenth = 3

func createFrame(h, w int, value rune) []model.Point {
	points := make([]model.Point, 0)
	points = addHorizontal(w, 0, value, points)
	points = addHorizontal(w, h, value, points)
	points = addVertical(h, 0, value, points)
	points = addVertical(h, w, value, points)
	return points
}

func addHorizontal(w, y int, value rune, points []model.Point) []model.Point {
	for i := 0; i <= w; i++ {
		points = append(points, model.Point{i, y, value})
	}
	return points
}

func addVertical(h, x int, value rune, points []model.Point) []model.Point {
	for i := 0; i <= h; i++ {
		points = append(points, model.Point{x, i, value})
	}
	return points
}
