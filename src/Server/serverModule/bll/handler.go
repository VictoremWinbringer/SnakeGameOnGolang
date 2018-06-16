package bll

import (
	"../dal"
)

type HandlerType byte

type IHandler interface {
	Type() HandlerType
	Handle(data interface{}, session dal.ISession) ([]byte, bool)
}
