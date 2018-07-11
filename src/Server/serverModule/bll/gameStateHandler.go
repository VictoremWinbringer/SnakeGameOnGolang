package bll

import (
	"sync"

	"../../../Shared/serializer"
	"../dal"
)

var currentMessageId uint64 = 1

type gameStateHandler struct {
	mxt sync.Mutex
}

func (this gameStateHandler) Type() HandlerType {
	return HandlerType(serializer.GameStateType)
}

func (this gameStateHandler) Handle(data []byte, session dal.ISession) ([]byte, bool) {
	state := session.GetState()
	messageData := serializer.EncodeGameState(serializer.GameState{state})
	return serializer.EncodeMessage(this.createMessage(messageData)), true
}

func (this gameStateHandler) createMessage(data []byte) serializer.Message {
	this.mxt.Lock()
	defer this.mxt.Unlock()
	message := serializer.Message{
		currentMessageId,
		serializer.GameStateType,
		data}
	currentMessageId++
	return message
}
