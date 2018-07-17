package bll

import (
	"sync"

	serializer "../../../Shared/serializer"
)

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
	return commandHandler{lastId: 0, mxt: &sync.Mutex{}}
}

func (this factory) CreateDispatcher() IDispatcher {
	handlers := make(map[serializer.MessageType]IHandler, 0)
	handlers[serializer.CommandType] = this.CreateCommandHandler()
	handlers[serializer.GameStateType] = this.CreateGameStateHandler()
	return &dispatcher{lastId: 0, mxt: &sync.Mutex{}, handlers: handlers}
}
