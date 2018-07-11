package udp

import (
	"net"
	"time"
)

type IUdpClient interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Close() error
}

type udpClient struct {
	connection net.Conn
	timeOut    uint
}

func NewUdpClient(ip string, timeOut uint) (IUdpClient, error) {
	conn, err := net.Dial("udp", ip)
	if err != nil {
		return nil, err
	}
	return udpClient{conn, timeOut}, nil
}

func (c udpClient) Read(p []byte) (int, error) {
	c.connection.SetReadDeadline(time.Now().Add(time.Second * time.Duration(c.timeOut)))
	return c.connection.Read(p)
}

func (c udpClient) Write(p []byte) (int, error) {
	c.connection.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(c.timeOut)))
	return c.connection.Write(p)
}

func (c udpClient) Close() error {
	return c.connection.Close()
}
