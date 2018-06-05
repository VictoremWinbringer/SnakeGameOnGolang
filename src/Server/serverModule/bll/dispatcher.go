package bll

import (
	"sync"

	serializer "../../../Shared/serializer"
	"../dal"
)

type IDispatcher interface {
	Dispatch(requestData []byte, clientId int) ([]byte, bool)
}

type dispatcher struct {
	lastId   int64
	ids      map[int64]bool
	mxt      sync.Mutex
	sessions map[int]dal.ISession
	handlers map[serializer.MessageType]IHandler
}

func (this *dispatcher) Dispatch(requestData []byte, clientId int) ([]byte, bool) {
	message := serializer.DecodeMessage(requestData)
	if !this.checkAndAddIdTreadSafe(message.Id) {
		return nil, false
	}
	handrler, ok := this.handlers[message.Type]
	if !ok {
		return nil, false
	}
	return handrler.Handle(message.Data, this.sessions[clientId])
}

func (this *dispatcher) checkAndAddIdTreadSafe(id int64) bool {
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
		newHistory := make(map[int64]bool)
		for k, v := range this.ids {
			if this.lastId-k > 100 {
				newHistory[k] = v
			}
		}
		this.ids = newHistory
	}
}

func (this *dispatcher) checkId(id int64) bool {
	if _, ok := this.ids[id]; this.lastId > id || ok {
		return false
	}
	return true
}

func (this *dispatcher) addId(id int64) {
	this.ids[id] = true
}
