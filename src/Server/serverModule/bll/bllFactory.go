package bll

import (
	"time"

	"../../../Shared/messageTypeEnum"
	"../dal"
)

func NewSeverBllFactory(dalFactory dal.IServerDalFactory) ISeverBllFactory {
	return factory{dalFactory: dalFactory}
}

type ISeverBllFactory interface {
	CreateGameStateHandler() IHandler
	CreateCommandHandler() IHandler
	CreateDispatcher(onError func(error), onSuccess func([]byte, dal.Connection)) IDispatcher
	CreateClient() IClient
}

type factory struct {
	dalFactory dal.IServerDalFactory
}

func (this factory) CreateGameStateHandler() IHandler {
	return &gameStateHandler{}
}

func (this factory) CreateCommandHandler() IHandler {
	return &commandHandler{lastId: 0}
}

func (this factory) CreateDispatcher(onError func(error), onSuccess func([]byte, dal.Connection)) IDispatcher {
	return &dispatcher{onSuccess: onSuccess, onError: onError, clients: make(map[string]IClient), factory: this}
}

func (this factory) CreateClient() IClient {
	session := this.dalFactory.CreateSession()
	session.Start()
	handlers := make(map[messageTypeEnum.Type]IHandler, 0)
	handlers[messageTypeEnum.CommandType] = this.CreateCommandHandler()
	handlers[messageTypeEnum.GameStateType] = this.CreateGameStateHandler()
	return &client{time.Now(), session, handlers}
}
