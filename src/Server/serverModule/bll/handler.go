package bll

import (
	"../dal"
	. "../../../Shared/serializer"
)

type HandlerType byte

type IHandler interface {
	Type() HandlerType
	Handle(data Message, session dal.ISession) ([]byte, error)
}
