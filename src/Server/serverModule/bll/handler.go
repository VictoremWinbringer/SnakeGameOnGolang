package bll

import "../dal"

type HandlerType byte

type IHandler interface {
	Type() HandlerType
	Handle(requestId int64, data []byte, session dal.ISession) []byte
}
