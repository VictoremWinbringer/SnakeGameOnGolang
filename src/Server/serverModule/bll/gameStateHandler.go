package bll

import (
	serializer "../../../Shared/serializer"
	"../dal"
)

type gameStateHandler struct {
}

func (this gameStateHandler) Type() HandlerType {
	return HandlerType(serializer.GameStateType)
}

func (this gameStateHandler) Handle(data []byte, session dal.ISession) ([]byte, bool) {
	state := session.GetState()
	return serializer.EncodeGameState(serializer.GameState{state}), true
}
