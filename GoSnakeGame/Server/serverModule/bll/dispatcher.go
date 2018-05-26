package bll

import "../dal"

type IDispatcher interface {
	Dispatch(requestData []byte, clientId int) []byte
}

type dispatcher struct {
	sessions map[int]dal.ISession
	handlers map[HandlerType]IHandler
}
