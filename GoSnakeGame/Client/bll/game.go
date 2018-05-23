package bll

import (
	"../dal"
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
	frame          iframe
	food           ifood
	snake          isnake
	screen         dal.IScreen
	timeBuffer     int64
	commandChannel chan Command
}

func NewGame(height int, width int) (IGame, error) {

	writer, err := dal.NewIPointWriter()
	if err != nil {
		return nil, err
	}
	frame := newIFrame(height, width, '+', writer)
	food := newIFood(width/2, height/2, width, height, '$', writer)
	snake := newISnake(width/3, height/3, '+', writer)
	return &game{frame, food, snake, writer, 0, keyboardInput(writer)}, nil
}

func (game *game) Draw() {

	game.screen.Clear()
	game.frame.Draw()
	game.food.Draw()
	game.snake.Draw()
	game.screen.Show()
}

const timeDeltaInNanoSecondsAfterThatSnakeMoves int64 = 200000000

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

func (game *game) Logic(timeDeltaInNanoSeconds int64) bool {
	game.timeBuffer += timeDeltaInNanoSeconds
	select {
	case command := <-game.commandChannel:
		switch command {
		case Up:
			game.snake.Go(UpDirection)
		case Down:
			game.snake.Go(DownDirection)
		case Left:
			game.snake.Go(LeftDirection)
		case Right:
			game.snake.Go(RightDirection)
		case Exit:
			game.screen.Clear()
			game.screen.ShowCursor(0, 0)
			game.screen.Show()
			return false
		}
	default:
	}
	if game.snake.IsHit(game.frame) || game.snake.IsHitTail() {
		game.snake.Reset()
	}
	game.snake.TryEat(game.food)
	if game.timeBuffer >= timeDeltaInNanoSecondsAfterThatSnakeMoves {
		game.snake.Move()
		game.timeBuffer -= timeDeltaInNanoSecondsAfterThatSnakeMoves
	}
	return true
}
