package bll

import (
	"sync"

	serializer "../../../Shared/serializer"
	"../dal"
)

type commandHandler struct {
	lastId uint64
	mxt    sync.Mutex
}

func (this commandHandler) Type() HandlerType {
	return HandlerType(serializer.CommandType)
}

func (this commandHandler) Handle(data []byte, session dal.ISession) ([]byte, bool) {
	println("CommandHandler")
	command := serializer.DecodeCommand(data)
	println("Command code")
	println(command.Code)
	if !this.checkAndAddIdTreadSafe(command.Id) {
		return make([]byte, 0), false
	}
	session.HandleCommand(int(command.Code))
	return make([]byte, 0), true
}

func (this *commandHandler) checkAndAddIdTreadSafe(id uint64) bool {
	this.mxt.Lock()
	defer this.mxt.Unlock()
	if this.lastId < id {
		this.lastId = id
		return true
	}
	return false
}
