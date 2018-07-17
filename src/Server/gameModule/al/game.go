package al

import (
	"../bll"
	"../dal"
)

type Command uint8

const Left Command = 1
const Right Command = 2
const Up Command = 3
const Down Command = 4
const Exit Command = 5

type IGame interface {
	Draw() [][]rune
	Logic(timeDeltaInNanoSeconds int64) bool
}

type game struct {
	frame          bll.IFigure
	food           bll.IFood
	snake          bll.ISnake
	screen         dal.IWriter
	timeBuffer     int64
	commandChannel <-chan Command
}

func NewGame(height int, width int, commands <-chan Command) (IGame, error) {
	dalFactory := dal.CreateDalFactory()
	screen := dalFactory.CreateWriter(width+1, height+1)
	bllFactory := bll.NewBllFactory(dalFactory, screen)
	frame := bllFactory.CrateFrame(height, width, '+')
	food := bllFactory.CreateFood(width/2, height/2, '$', width, height)
	snake := bllFactory.CreateSnake(width/3, height/3, '*')
	return &game{frame, food, snake, screen, 0, commands}, nil
}

func (game *game) Draw() [][]rune {
	game.screen.Clear()
	game.frame.Draw()
	game.food.Draw()
	game.snake.Draw()
	return game.screen.Data()
}

const timeDeltaInNanoSecondsAfterThatSnakeMoves int64 = 200000000

func (game *game) Logic(timeDeltaInNanoSeconds int64) bool {
	game.timeBuffer += timeDeltaInNanoSeconds
	select {
	case command := <-game.commandChannel:
		switch command {
		case Up:
			game.snake.Go(bll.UpDirection)
		case Down:
			game.snake.Go(bll.DownDirection)
		case Left:
			game.snake.Go(bll.LeftDirection)
		case Right:
			game.snake.Go(bll.RightDirection)
		case Exit:
			game.screen.Clear()
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
