package bll

import (
	"fmt"
	"sync"

	serializer "../../../Shared/serializer"
)

type IDispatcher interface {
	Dispatch(requestData []byte) (IHandler, serializer.Message, error)
}

type dispatcher struct {
	lastId   uint64
	mxt      sync.Mutex
	handlers map[serializer.MessageType]IHandler
}

func (this *dispatcher) Dispatch(requestData []byte) (IHandler, serializer.Message, error) {
	message := serializer.DecodeMessage(requestData)
	if !this.checkAndAddIdTreadSafe(message.Id) {
		return nil, message, fmt.Errorf("message with id: %v not valid", message.Id)
	}
	handrler, ok := this.handlers[message.Type]
	if !ok {
		return nil, message, fmt.Errorf("handler for type %v not found", message.Type)
	}
	return handrler, message, nil
}

func (this *dispatcher) checkAndAddIdTreadSafe(id uint64) bool {
	this.mxt.Lock()
	defer this.mxt.Unlock()
	if this.lastId < id {
		this.lastId = id
		return true
	}
	return false
}
