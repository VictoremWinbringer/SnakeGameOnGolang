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
	frame          ifigure
	food           ifood
	snake          isnake
	screen         dal.IScreen
	timeBuffer     int64
	commandChannel chan Command
}

func (game *game) Draw() {

	game.screen.Clear()
	game.frame.draw()
	game.food.draw()
	game.snake.draw()
	game.screen.Show()
}

const timeDeltaInNanoSecondsAfterThatSnakeMoves int64 = 200000000

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
