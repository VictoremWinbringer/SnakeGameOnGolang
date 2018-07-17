package bll

import (
	"time"
	"../dal"
	"../../../Shared/serializer"
	"fmt"
)

type IClient interface {
	IsAlive() bool
	UpdateLastActiveTime()
	Accept(data []byte) ([]byte, error)
	Close()
}

const MAX_SECONDS_TO_UNACTIV = 5

type client struct {
	lastActive time.Time
	session  dal.ISession
	handlers map[serializer.MessageType]IHandler
}

func (this *client) IsAlive() bool {
	oldTime := this.lastActive
	nowTime := time.Now()
	timeDelta := nowTime.Unix() - oldTime.Unix()
	return timeDelta < MAX_SECONDS_TO_UNACTIV;
}

func (this *client) UpdateLastActiveTime() {
	this.lastActive = time.Now()
}

func (this *client) Accept(data []byte) ([]byte, error) {
	message := serializer.DecodeMessage(data)
	handler, ok := this.handlers[message.Type]
	if !ok {
		return nil, fmt.Errorf("handler for type %v not found", message.Type)
	}
	return handler.Handle(message, this.session)
}

func (this *client) Close() {
	this.session.Stop()
}
