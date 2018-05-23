package game

import (
	_ "../../Shared/serializer"
	foodModule "../food"
	frameModule "../frame"
	pointModule "../point"
	snakeModule "../snake"
	tcellModule "github.com/gdamore/tcell"
)

type Command uint8

const Left Command = 1
const Right Command = 2
const Up Command = 3
const Down Command = 4
const Exit Command = 5

type IGame interface {
	Draw()
	Logic(timeDeltaInNanoSeconds int64) bool
}

type game struct {
	frame          *frameModule.Frame
	food           *foodModule.Food
	snake          *snakeModule.Snake
	screen         tcellModule.Screen
	timeBuffer     int64
	commandChannel chan Command
}

func New(height int, width int) (IGame, error) {
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
	food := foodModule.New(10, 10, width, height, '$', writer)
	snake := snakeModule.New(8, 8, '+', writer)
	return &game{&frame, &food, &snake, screen, 0, keyboardInput(screen)}, nil
}

func (game *game) Draw() {

	game.screen.Clear()
	game.frame.Draw()
	game.food.Draw()
	game.snake.Draw()
	game.screen.Show()
}

const timeDeltaInNanoSecondsAfterThatSnakeMoves int64 = 100000000

func keyboardInput(screen tcellModule.Screen) chan Command {
	commandChannel := make(chan Command)
	go func() {
		for {
			event := screen.PollEvent()
			keyEvent, ok := event.(*tcellModule.EventKey)
			if ok {
				switch keyEvent.Key() {
				case tcellModule.KeyUp:
					commandChannel <- Up
				case tcellModule.KeyDown:
					commandChannel <- Down
				case tcellModule.KeyLeft:
					commandChannel <- Left
				case tcellModule.KeyRight:
					commandChannel <- Right
				case tcellModule.KeyEsc:
					commandChannel <- Exit
				}
			}
		}
	}()
	return commandChannel
}

func (game *game) Logic(timeDeltaInNanoSeconds int64) bool {
	game.timeBuffer += timeDeltaInNanoSeconds
	select {
	case command := <-game.commandChannel:
		switch command {
		case Up:
			game.snake.Go(snakeModule.Up)
		case Down:
			game.snake.Go(snakeModule.Down)
		case Left:
			game.snake.Go(snakeModule.Left)
		case Right:
			game.snake.Go(snakeModule.Right)
		case Exit:
			game.screen.Clear()
			game.screen.ShowCursor(0, 0)
			game.screen.Show()
			return false
		}
	default:
	}
	game.snake.TryEat(game.food)
	if game.snake.IsHit(game.frame.Figure) || game.snake.IsHitTail() {
		game.snake.Reset()
	}
	if game.timeBuffer >= timeDeltaInNanoSecondsAfterThatSnakeMoves {
		game.snake.Move()
		game.timeBuffer -= timeDeltaInNanoSecondsAfterThatSnakeMoves
	}
	return true
}
