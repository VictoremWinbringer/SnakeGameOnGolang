package bll

import (
	serializer "../../../Shared/serializer"
	"../dal"
	"sync"
)

type commandHandler struct {
	lastId   uint64
	ids      map[uint64]bool
	mxt      sync.Mutex
}

func (this commandHandler) Type() HandlerType {
	return HandlerType(serializer.CommandType)
}

func (this commandHandler) Handle(data []byte, session dal.ISession) ([]byte, bool) {
	command := serializer.DecodeCommand(data)

	if ! this.checkAndAddIdTreadSafe(command.Id){
		return make([]byte, 0), false
	}
	session.HandleCommand(int(command.Code))
	return make([]byte, 0), true
}

func (this *commandHandler) checkAndAddIdTreadSafe(id uint64) bool {
	this.mxt.Lock()
	defer this.mxt.Unlock()
	this.clearHistory()
	if this.checkId(id) {
		this.addId(id)
		return true
	}
	return false
}

func (this *commandHandler) clearHistory() {
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

func (this *commandHandler) checkId(id uint64) bool {
	if _, ok := this.ids[id]; this.lastId > id || ok {
		return false
	}
	return true
}

func (this *commandHandler) addId(id uint64) {
	this.ids[id] = true
}
