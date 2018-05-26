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
	Close()
}

type game struct {
	frame          ifigure
	food           ifood
	snake          isnake
	screen         dal.IScreen
	timeBuffer     int64
	commandChannel chan Command
}

func NewGame(height int, width int) (IGame, error) {
	dalFactory := dal.CreateDalFactory()
	screen, err := dalFactory.CreateScreen()
	if err != nil {
		return nil, err
	}
	bllFactory := NewBllFactory(dalFactory, screen)
	frame := bllFactory.CrateFrame(height, width, '+')
	food := bllFactory.CreateFood(width/2, height/2, '$', width, height)
	snake := bllFactory.CreateSnake(width/3, height/3, '+')
	return &game{frame, food, snake, screen, 0, keyboardInput(screen)}, nil
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

func (this *game) Close() {
	this.screen.Close()
}
