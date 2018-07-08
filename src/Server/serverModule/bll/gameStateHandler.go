package bll

import (
	"../../../Shared/serializer"
	"../dal"
)

var currentMessageId uint64 = 1

type gameStateHandler struct {
}

func (this gameStateHandler) Type() HandlerType {
	return HandlerType(serializer.GameStateType)
}

func (this gameStateHandler) Handle(data []byte, session dal.ISession) ([]byte, bool) {
	state := session.GetState()
	messageData := serializer.EncodeGameState(serializer.GameState{state})
	message := serializer.Message{
		currentMessageId,
		serializer.GameStateType,
		messageData}
	currentMessageId++
	return serializer.EncodeMessage(message), true
}
