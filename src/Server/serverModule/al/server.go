package al

import (
	"fmt"

	"../bll"
	. "../dal"
)

type IServer interface {
	Start() error
}

type server struct {
	listener   IUdpListener
	dispatcher bll.IDispatcher
}

func NewServer(port int, ip string, factory bll.ISeverBllFactory) (IServer, error) {
	if factory == nil {
		return nil, fmt.Errorf("factory is nil!")
	}
	udpLiscener, err := NewUdpServer(port, ip)
	if err != nil {
		return nil, err
	}
	dispatcher := factory.CreateDispatcher(func(e error) {
		fmt.Printf("Error on dispathing %v\n", err)
	}, func(bytes []byte, connection Connection) {
		if len(bytes) < 1 {
			return
		}
		_, e := udpLiscener.Write(bytes, connection)
		if e != nil {
			fmt.Printf("Couldn't send response - %v \n", e)
		}
	})
	return server{udpLiscener, dispatcher}, nil
}

func (this server) Start() error {
	go this.listen()
	return nil
}

func (this server) listen() {
	defer func() {
		this.listener.Close()
		this.dispatcher.Close()
	}()
	for {
		data := make([]byte, 4096)
		_, remoteaddr, err := this.listener.Read(data)
		if err != nil {
			fmt.Printf("Error on reading from listener %v\n", err)
			continue
		}
		go this.dispatcher.Dispatch(data, remoteaddr)
	}
}
