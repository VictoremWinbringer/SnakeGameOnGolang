package udpClient

import (
	"net"
)

type UdpClient interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Close() error
}

type udpClient struct {
	connection net.Conn
}

func New(ip string) (UdpClient, error) {
	conn, err := net.Dial("udp", ip)
	if err != nil {
		return nil, err
	}
	return udpClient{conn}, nil
}

func (c udpClient) Read(p []byte) (int, error) {
	return c.connection.Read(p)
}

func (c udpClient) Write(p []byte) (int, error) {
	return c.connection.Write(p)
}

func (c udpClient) Close() error {
	return c.connection.Close()
}
