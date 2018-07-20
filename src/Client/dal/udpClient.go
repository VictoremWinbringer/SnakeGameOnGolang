package dal

import (
	"fmt"
	"net"
)

type IUdpClient interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Close() error
}

type udpClient struct {
	connection     *net.UDPConn
	localAddress   *net.UDPAddr
	remouteAddress *net.UDPAddr
	timeOut        uint
}

func NewUdpClient(clientIp, serverIp string, timeOut uint) (IUdpClient, error) {
	client, e := net.ResolveUDPAddr("udp4", clientIp)
	if e != nil {
		return nil, e
	}
	server, e := net.ResolveUDPAddr("udp4", serverIp)
	if e != nil {
		return nil, e
	}
	conn, err := net.ListenUDP("udp", client)
	if err != nil {
		return nil, err
	}
	return udpClient{conn, client, server, timeOut}, nil
}

func (c udpClient) Read(p []byte) (int, error) {
	//c.connection.SetReadDeadline(time.Now().Add(time.Second * time.Duration(c.timeOut)))
	i, a, e := c.connection.ReadFromUDP(p)
	if e != nil {
		return 0, e
	}
	if a.String() != c.remouteAddress.String() {
		return 0, fmt.Errorf("Unnown addres %v", a)
	}
	return i, nil
}

func (c udpClient) Write(p []byte) (int, error) {
	//	c.connection.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(c.timeOut)))
	return c.connection.WriteToUDP(p, c.remouteAddress)
}

func (c udpClient) Close() error {
	return c.connection.Close()
}
