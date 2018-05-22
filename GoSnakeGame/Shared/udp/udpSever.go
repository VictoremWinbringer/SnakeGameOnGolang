package udp

import (
	"fmt"
	"net"
)

type IUpdServer interface {
	Read(buffer []byte) (int, Connection, error)
	Write(buffer []byte, address Connection) (int, error)
	Close() error
}

type server struct {
	connection *net.UDPConn
}

type Connection *net.UDPAddr

func NewUdpServer(port int, ip string) (IUpdServer, error) {
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
	ser, err := net.ListenUDP("udp", &udpAddress)
	if err != nil {
		return nil, err
	}
	return server{ser}, nil
}

func (this server) Read(buffer []byte) (int, Connection, error) {
	return this.connection.ReadFromUDP(buffer)
}
func (this server) Write(buffer []byte, address Connection) (int, error) {
	return this.connection.WriteToUDP(buffer, address)
}

func (this server) Close() error {
	return this.connection.Close()
}
