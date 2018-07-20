package bll

import (
	. "../../../Shared/models"
	"../dal"
)

type HandlerType byte

type IHandler interface {
	Type() HandlerType
	Handle(data Message, session dal.ISession) ([]byte, error)
}
