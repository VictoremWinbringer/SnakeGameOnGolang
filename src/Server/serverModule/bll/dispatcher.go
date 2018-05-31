package bll

import (
	serializer "../../../Shared/serializer"
	"../dal"
)

type IDispatcher interface {
	Dispatch(requestData []byte, clientId int) []byte
}

type dispatcher struct {
	sessions map[int]dal.ISession
	handlers map[serializer.MessageType]IHandler
}

func (this *dispatcher) Dispatch(requestData []byte, clientId int) []byte {
	message := serializer.DecodeMessage(requestData)
	return this.handlers[message.Type].Handle(message.Id, message.Data, this.sessions[clientId])
}
