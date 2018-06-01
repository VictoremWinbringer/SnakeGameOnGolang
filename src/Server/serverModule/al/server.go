package al

import (
	"fmt"

	udpModule "../../../Shared/udp"
	"../bll"
)

type IServer interface {
	Start() error
}

type server struct {
	listener   udpModule.IUdpListener
	dispatcher bll.IDispatcher
	clients    map[udpModule.Connection]int
}

func NewServer(port int, ip string, dispatcher bll.IDispatcher) (IServer, error) {
	if dispatcher == nil {
		return nil, fmt.Errorf("handler is nil!")
	}
	udpLiscener, err := udpModule.NewUdpServer(port, ip)
	if err != nil {
		return nil, err
	}
	return server{udpLiscener, dispatcher, make(map[udpModule.Connection]int)}, nil
}

func (this server) Start() error {

	go this.listen()
	return nil
}

func (this server) listen() {
	defer func() {
		this.listener.Close()
	}()
	var currentId int = 0
	for {
		data := make([]byte, 4096)
		_, remoteaddr, err := this.listener.Read(data)
		id, ok := this.clients[remoteaddr]
		if !ok {
			id = currentId + 1
			this.clients[remoteaddr] = id
		}
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go this.sendResponse(data, remoteaddr, id)
	}
}

func (this server) sendResponse(data []byte, address udpModule.Connection, clientId int) {
	data, ok = this.dispatcher.Dispatch(data, clientId)
	if !ok {
		return
	}
	_, err := this.listener.Write(data, address)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}
