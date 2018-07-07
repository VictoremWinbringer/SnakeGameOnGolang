package bll

import (
	"fmt"
	"sync"
	serializer "../../../Shared/serializer"
)

type IDispatcher interface {
	Dispatch(requestData []byte, clientId int) (IHandler,serializer.Message, error)
}

type dispatcher struct {
	lastId   uint64
	ids      map[uint64]bool
	mxt      sync.Mutex
	handlers map[serializer.MessageType]IHandler
}

func (this *dispatcher) Dispatch(requestData []byte, clientId int) (IHandler, serializer.Message, error) {
	message := serializer.DecodeMessage(requestData)
	if !this.checkAndAddIdTreadSafe(message.Id) {
		return nil, serializer.Message{}, fmt.Errorf("message with id: %v not valid", message.Id)
	}
	handrler, ok := this.handlers[message.Type]
	if !ok {
		return nil, serializer.Message{}, fmt.Errorf("handler for type %v not found", message.Type)
	}
	return handrler,message, nil
}

func (this *dispatcher) checkAndAddIdTreadSafe(id uint64) bool {
	this.mxt.Lock()
	defer this.mxt.Unlock()
	this.clearHistory()
	if this.checkId(id) {
		this.addId(id)
		return true
	}
	return false
}

func (this *dispatcher) clearHistory() {
	if len(this.ids) > 10000 {
		newHistory := make(map[uint64]bool)
		for k, v := range this.ids {
			if this.lastId-k > 100 {
				newHistory[k] = v
			}
		}
		this.ids = newHistory
	}
}

func (this *dispatcher) checkId(id uint64) bool {
	if _, ok := this.ids[id]; this.lastId > id || ok {
		return false
	}
	return true
}

func (this *dispatcher) addId(id uint64) {
	this.ids[id] = true
}
