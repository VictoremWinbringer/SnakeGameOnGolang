package bll

func NewSeverBllFactory() ISeverBllFactory {
	return factory{}
}

type ISeverBllFactory interface {
	CreateGameStateHandler() IHandler
	CreateCommandHandler() IHandler
	CreateDispatcher() IDispatcher
}

type factory struct {
}

func (this factory) CreateGameStateHandler() IHandler {
	return gameStateHandler{}
}

func (this factory) CreateCommandHandler() IHandler {
	return commandHandler{}
}

func (this factory) CreateDispatcher() IDispatcher {
	return &dispatcher{}
}
