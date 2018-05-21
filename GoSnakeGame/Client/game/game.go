package game

import (
	_ "../../Shared/serializer"
	foodModule "../food"
	frameModule "../frame"
	pointModule "../point"
	snakeModule "../snake"
	_ "../udpClient"
	tcellModule "github.com/gdamore/tcell"
)

type Game interface {
	Draw()
	Logic(timeDeltaInNanoSeconds int64) bool
}

type game struct {
	frame      frameModule.Frame
	food       foodModule.Food
	snake      snakeModule.Snake
	screen     tcellModule.Screen
	timeBuffer int64
}

func New(height int, width int) (Game, error) {
	screen, e := tcellModule.NewScreen()
	if e != nil {
		return nil, e
	}
	if e := screen.Init(); e != nil {
		return nil, e
	}
	screen.SetStyle(tcellModule.StyleDefault)
	screen.HideCursor()
	writer := pointModule.NewTerminalWriter(screen)
	frame := frameModule.New(height, width, '+', writer)
	food := foodModule.New(10, 10, '$', writer)
	snake := snakeModule.New(8, 8, '+', writer)
	return &game{frame, food, snake, screen, 0}, nil
}

func (game *game) Draw() {

	game.screen.Clear()
	game.frame.Draw()
	game.food.Draw()
	game.snake.Move()
	game.snake.Draw()
	game.screen.Show()
}

func (game *game) Logic(timeDeltaInNanoSeconds int64) bool {
	game.timeBuffer += timeDeltaInNanoSeconds
	event := game.screen.PollEvent()
	switch keyEvent := event.(type) {
	case *tcellModule.EventKey:
		switch keyEvent.Key() {
		case tcellModule.KeyEsc:
			game.screen.Clear()
			game.screen.Show()
			game.screen.ShowCursor(0, 0)
			return false
		case tcellModule.KeyUp:
			game.snake.Go(snakeModule.Up)
		case tcellModule.KeyDown:
			game.snake.Go(snakeModule.Down)
		case tcellModule.KeyLeft:
			game.snake.Go(snakeModule.Left)
		case tcellModule.KeyRight:
			game.snake.Go(snakeModule.Right)
		}
	}
	if game.timeBuffer >= 100000000 {
		game.snake.Move()
		game.timeBuffer = 0
	}
	return true
}
