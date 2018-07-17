package bll

import (
	"../../../Shared/serializer"
	"../dal"
)

type gameStateHandler struct {
}

func (this gameStateHandler) Type() HandlerType {
	return HandlerType(serializer.GameStateType)
}

func (this gameStateHandler) Handle(data serializer.Message, session dal.ISession) ([]byte, error) {
	state := session.GetState()
	messageData := serializer.EncodeGameState(serializer.GameState{state})
	return serializer.EncodeMessage(this.createMessage(messageData)), nil
}

func (this gameStateHandler) createMessage(data []byte) serializer.Message {
	message := serializer.Message{
		serializer.GameStateType,
		data}
	return message
}
