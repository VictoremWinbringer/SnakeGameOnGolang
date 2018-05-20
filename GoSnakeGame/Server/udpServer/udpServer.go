package udpServer

import (
	"fmt"
	"net"
)

type Server interface {
	Start() error
}

type Handler interface {
	Hanle(requestData []byte) []byte
}
type server struct {
	address net.UDPAddr
	handler Handler
}

func New(port int, ip string, handler Handler) (Server, error) {
	if handler == nil {
		return nil, fmt.Errorf("handler is nil!")
	}
	if port <= 0 {
		return nil, fmt.Errorf("port <= 0 and is is %v", port)
	}
	typedIP := net.ParseIP(ip)
	if typedIP == nil {
		return nil, fmt.Errorf("wrong ip address %v", ip)
	}
	udpAddress := net.UDPAddr{
		Port: port,
		IP:   typedIP,
		Zone: "",
	}
	return server{address: udpAddress, handler: handler}, nil
}

func (s server) Start() error {
	ser, err := net.ListenUDP("udp", &s.address)
	if err != nil {
		return err
	}
	go liscen(ser, s.handler)
	return nil
}

func liscen(ser *net.UDPConn, h Handler) {
	defer func() {
		ser.Close()
	}()
	p := make([]byte, 4096)
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, p, h, remoteaddr)
	}
}

func sendResponse(ser *net.UDPConn, p []byte, h Handler, addr *net.UDPAddr) {
	_, err := ser.WriteToUDP(h.Hanle(p), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}
