package bll

import (
	serializer "../../../Shared/serializer"
	"../dal"
)

type HandlerType byte

type IHandler interface {
	Type() HandlerType
	Handle(data []byte, session dal.ISession) ([]byte, bool)
}

type gameStateHandler struct {
}

func (this gameStateHandler) Type() HandlerType {
	return HandlerType(serializer.GameStateType)
}

func Handle(data []byte, session dal.ISession) ([]byte, bool) {
	state := session.GetState()
	return serializer.EncodeGameState(serializer.GameState{state}), true
}
