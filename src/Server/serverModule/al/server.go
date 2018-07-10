package al

import (
	"fmt"

	udpModule "../../../Shared/udp"
	"../bll"
	"../dal"
)

type IServer interface {
	Start() error
}

type server struct {
	listener   udpModule.IUdpListener
	dispatcher bll.IDispatcher
	clients    map[udpModule.Connection]int
	sessions map[int]dal.ISession
}

func NewServer(port int, ip string, dispatcher bll.IDispatcher) (IServer, error) {
	if dispatcher == nil {
		return nil, fmt.Errorf("handler is nil!")
	}
	udpLiscener, err := udpModule.NewUdpServer(port, ip)
	if err != nil {
		return nil, err
	}
	sessions := make(map[int]dal.ISession)
	return server{udpLiscener, dispatcher, make(map[udpModule.Connection]int), sessions}, nil
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
		if err != nil {
			fmt.Printf("Error on reading from listener %v\n",err)
			continue
		}
		fmt.Printf("Remoute address %v\n", remoteaddr)
		id, ok := this.clients[remoteaddr]
		if !ok {
			id = currentId + 1
			this.clients[remoteaddr] = id
		}
		go this.sendResponse(data, remoteaddr, id)
	}
}

func (this server) sendResponse(data []byte, address udpModule.Connection, clientId int) {
	handler, message, err := this.dispatcher.Dispatch(data, clientId)
	if err != nil {
		fmt.Printf("Couldn't create handler %v \n", err)
		return
	}
	fmt.Printf("Message %v\n", message)
   session, ok := this.sessions[clientId]
   if !ok {
	   session = dal.NewServerDalFactory().CreateSession()
	   session.Start()
	   this.sessions[clientId] = session
   }
   result, ok := handler.Handle(message.Data,session)
   if !ok {
   	print("Can not handle")
   	return
   }
fmt.Printf("Result %v\n", result)
	_, err = this.listener.Write(result, address)
	if err != nil {
		fmt.Printf("Couldn't send response %v \n", err)
	}
}
