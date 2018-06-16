package bll

import (
	serializer "../../../Shared/serializer"
	"../dal"
)

type commandHandler struct {
}

func (this commandHandler) Type() HandlerType {
	return HandlerType(serializer.CommandType)
}

func (this commandHandler) Handle(data []byte, session dal.ISession) ([]byte, bool) {
	command := serializer.DecodeCommand(data)
	session.HandleCommand(int(command.Code))
	return make([]byte, 0), true
}
