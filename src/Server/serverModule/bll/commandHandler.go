package bll

import (
	"fmt"

	serializer "../../../Shared/serializer"
	"../dal"
)

type commandHandler struct {
	lastId     uint64
	inCount    float32
	validCount float32
}

func (this *commandHandler) Type() HandlerType {
	return HandlerType(serializer.CommandType)
}

func (this *commandHandler) Handle(data serializer.Message, session dal.ISession) ([]byte, error) {
	command := serializer.DecodeCommand(data.Data)
	if !this.checkAndChangeId(command.Id) {
		return make([]byte, 0), nil
	}
	session.HandleCommand(int(command.Code))
	return make([]byte, 0), nil
}

func (this *commandHandler) checkAndChangeId(id uint64) bool {
	this.inCount++
	if this.inCount > 100 {
		fmt.Printf("Packet loss = %v%%\n", ((this.inCount/this.validCount)/4)*100)
		this.inCount -= 100
	}
	if this.lastId < id {
		this.lastId = id
		this.validCount++
		return true
	}
	return false
}
