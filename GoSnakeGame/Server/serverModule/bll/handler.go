package bll

import "../dal"

type HandlerType byte
type IHandler interface {
	Type() HandlerType
	Handle(data []byte, session dal.ISession) []byte
}
