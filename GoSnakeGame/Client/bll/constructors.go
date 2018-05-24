package bll

import (
	"../dal"
	"../domainModels"
)

func NewGame(height int, width int) (IGame, error) {

	writer, err := dal.NewIScreen()
	if err != nil {
		return nil, err
	}
	frame := newIFigure(createFigure(height, width, '+'), writer)
	food := newIFood(createFood(width/2, height/2, '$'), writer, width, height)
	snake := newISnake(createSnake(width/3, height/3, '+'), writer)
	return &game{frame, food, snake, writer, 0, keyboardInput(writer)}, nil
}

func keyboardInput(screen dal.IScreen) chan Command {
	commandChannel := make(chan Command)
	go func() {
		for {
			key := screen.ReadKey()
			switch key {
			case dal.KeyUp:
				commandChannel <- Up
			case dal.KeyDown:
				commandChannel <- Down
			case dal.KeyLeft:
				commandChannel <- Left
			case dal.KeyRight:
				commandChannel <- Right
			case dal.KeyEsc:
				commandChannel <- Exit
			}
		}
	}()
	return commandChannel
}

func newISnake(list dal.ILinkedList, writer IWriter) isnake {
	initialPoints := make([]domainModels.Point, 0)
	for i := list.Next(); i != nil; i = list.Next() {
		initialPoints = append(initialPoints, domainModels.Point{i.X, i.Y, i.Symbol})
	}
	return &snake{figure{list, writer}, RightDirection, RightDirection, initialPoints}
}

func newIFigure(list dal.ILinkedList, writer IWriter) ifigure {
	return figure{list, writer}
}

func newIFood(list dal.ILinkedList, writer IWriter, maxX, maxY int) ifood {
	return food{figure{list, writer}, maxX, maxY}
}

const initialLenth = 3

func createSnake(x, y int, value rune) dal.ILinkedList {
	points := make([]domainModels.Point, 0)
	for i := 0; i < initialLenth; i++ {
		points = append(points, domainModels.Point{x - i, y, value})
	}
	return dal.NewILinkedListWithData(points)
}

func createFood(x, y int, value rune) dal.ILinkedList {

	points := make([]domainModels.Point, 0)
	points = append(points, domainModels.Point{x, y, value})
	return dal.NewILinkedListWithData(points)
}

func createFigure(h, w int, value rune) dal.ILinkedList {
	points := make([]domainModels.Point, 0)
	points = addHorizontal(w, 0, value, points)
	points = addHorizontal(w, h, value, points)
	points = addVertical(h, 0, value, points)
	points = addVertical(h, w, value, points)
	return dal.NewILinkedListWithData(points)
}

func addHorizontal(w, y int, value rune, points []domainModels.Point) []domainModels.Point {
	for i := 0; i <= w; i++ {
		points = append(points, domainModels.Point{i, y, value})
	}
	return points
}

func addVertical(h, x int, value rune, points []domainModels.Point) []domainModels.Point {
	for i := 0; i <= h; i++ {
		points = append(points, domainModels.Point{x, i, value})
	}
	return points
}
