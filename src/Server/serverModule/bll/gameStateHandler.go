package bll

import (
	"../../../Shared/messageTypeEnum"
	. "../../../Shared/models"
	"../../../Shared/serializer"
	"../dal"
)

type gameStateHandler struct {
}

func (this gameStateHandler) Type() HandlerType {
	return HandlerType(messageTypeEnum.GameStateType)
}

func (this gameStateHandler) Handle(data Message, session dal.ISession) ([]byte, error) {
	state := session.GetState()
	messageData := serializer.EncodeGameState(GameState{state})
	return serializer.EncodeMessage(this.createMessage(messageData)), nil
}

func (this gameStateHandler) createMessage(data []byte) Message {
	message := Message{
		messageTypeEnum.GameStateType,
		data}
	return message
}
