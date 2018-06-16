package bll

func NewSeverBllFactory() ISeverBllFactory {
	return factory{}
}

type ISeverBllFactory interface {
	CreateGameStateHandler() IHandler
}

type factory struct {
}

func (this factory) CreateGameStateHandler() IHandler {
	return gameStateHandler{}
}
