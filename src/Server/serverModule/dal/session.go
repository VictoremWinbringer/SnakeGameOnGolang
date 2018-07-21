package dal

import (
	"time"

	"../../gameModule/al"
)

type ISession interface {
	HandleCommand(command int)
	GetState() [][]rune
	Start()
	Stop()
}

type session struct {
	game           al.IGame
	commandChannel chan al.Command
	done           chan bool
}

func (this *session) HandleCommand(command int) {
	this.commandChannel <- al.Command(command)
}

func (this session) GetState() [][]rune {
	return this.game.Draw()
}

func (this *session) Start() {
	commandChannel := make(chan al.Command)
	game, _ := al.NewGame(20, 40, commandChannel)
	this.game = game
	this.commandChannel = commandChannel
	this.done = make(chan bool,10)
	go func() {
		old := time.Now().UnixNano()
		for {
			new := time.Now().UnixNano()
			select {
			case <-this.done:
				return
			default:
				if !game.Logic(new - old) {
					return
				}
				old = new
			}
		}
	}()
}

func (this *session) Stop() {
	this.done <- true
}
