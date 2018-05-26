package al

import (
	"fmt"

	udpModule "../../Shared/udp"
)

type IServer interface {
	Start() error
}

type IHandler interface {
	Hanle(requestData []byte, clientId int) []byte
}
type server struct {
	udpServer udpModule.IUdpListener
	handler   IHandler
	clients   map[udpModule.Connection]int
}

func NewServer(port int, ip string, handler IHandler) (IServer, error) {
	if handler == nil {
		return nil, fmt.Errorf("handler is nil!")
	}
	udpLiscener, err := udpModule.NewUdpServer(port, ip)
	if err != nil {
		return nil, err
	}
	return server{udpLiscener, handler, make(map[udpModule.Connection]int)}, nil
}

func (s server) Start() error {

	go liscen(s.udpServer, s.handler, s.clients)
	return nil
}

func liscen(ser udpModule.IUdpListener, h IHandler, clients map[udpModule.Connection]int) {
	defer func() {
		ser.Close()
	}()
	var currentId int = 0
	for {
		p := make([]byte, 4096)
		_, remoteaddr, err := ser.Read(p)
		id, ok := clients[remoteaddr]
		if !ok {
			id = currentId + 1
			clients[remoteaddr] = id
		}
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, p, h, remoteaddr, id)
	}
}

func sendResponse(ser udpModule.IUdpListener, p []byte, h IHandler, addr udpModule.Connection, id int) {
	_, err := ser.Write(h.Hanle(p, id), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}
