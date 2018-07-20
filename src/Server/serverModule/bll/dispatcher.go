package bll

import (
	"fmt"
)

type IDispatcher interface {
	Dispatch(data []byte, clientId string, callback func([]byte, error))
	Close()
}

type dispatcher struct {
	clients map[string]IClient
	factory ISeverBllFactory
}

func (this *dispatcher) Dispatch(data []byte, clientId string, callback func([]byte, error)) {
	this.checkAliveClients()
	c, ok := this.clients[clientId]
	if !ok {
		fmt.Printf("Connected new client %s\n", clientId)
		this.clients[clientId] = this.factory.CreateClient()
		c = this.clients[clientId]
	}
	c.UpdateLastActiveTime()
	b, e := c.Accept(data)
	callback(b, e)
}

func (this *dispatcher) Close() {
	for _, v := range this.clients {
		v.Close()
	}
}

func (this *dispatcher) checkAliveClients() {
	for k, v := range this.clients {
		if !v.IsAlive() {
			this.clients[k].Close()
			delete(this.clients, k)
			fmt.Printf("Client with ip : %v disconnected!\n", k)
		}
	}
}
