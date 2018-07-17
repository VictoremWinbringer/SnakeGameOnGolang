package bll

import (
	"fmt"

	"../../../Shared/udp"
)

type IDispatcher interface {
	Dispatch(data []byte, connection udp.Connection)
	Close()
}

type dispatcher struct {
	onSuccess func([]byte, udp.Connection)
	onError   func(error)
	clients   map[string]IClient
	factory   ISeverBllFactory
}

func (this *dispatcher) Dispatch(data []byte, connection udp.Connection) {
	this.checkAliveClients()
	ip := fmt.Sprintf("%v", connection)
	c, ok := this.clients[ip]
	if !ok {
		fmt.Printf("Connected new client %s\n", ip)
		this.clients[ip] = this.factory.CreateClient()
		c = this.clients[ip]
	}
	c.UpdateLastActiveTime()
	b, e := c.Accept(data)
	if e != nil {
		this.onError(e)
		return
	}
	this.onSuccess(b, connection)
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
