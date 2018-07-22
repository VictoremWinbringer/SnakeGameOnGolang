package bll

import (
	"fmt"
	"sync"

	"math"

	"../../../Shared/messageTypeEnum"
	. "../../../Shared/models"
	"../../../Shared/serializer"
	"../dal"
)

type commandHandler struct {
	lastId     uint64
	startId    uint64
	validCount uint64
	mtx        *sync.Mutex
}

func (this *commandHandler) Type() HandlerType {
	return HandlerType(messageTypeEnum.CommandType)
}

func (this *commandHandler) Handle(data Message, session dal.ISession) ([]byte, error) {
	command := serializer.DecodeCommand(data.Data)
	if !this.checkAndChangeId(command.Id) {
		return make([]byte, 0), nil
	}
	session.HandleCommand(int(command.Code))
	return make([]byte, 0), nil
}

func (this *commandHandler) checkAndChangeId(id uint64) bool {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	if this.validCount > 1000 {
		fmt.Printf("Packet loss = %v%%\n", math.Max(0, 100-(float64(this.validCount)/float64(this.lastId-this.startId))*100))
		this.startId = this.lastId
		this.validCount = 0
	}
	if this.lastId < id {
		this.lastId = id
		this.validCount++
		return true
	}
	return false
}
