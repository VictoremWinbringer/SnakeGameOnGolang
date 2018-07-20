package al

import (
	"fmt"

	"../bll"
	. "../dal"
	"sync"
)

type IServer interface {
	Start() error
}

type server struct {
	listener   IUdpListener
	dispatcher bll.IDispatcher
	pool       *sync.Pool
}

func NewServer(port int, ip string, factory bll.ISeverBllFactory) (IServer, error) {
	if factory == nil {
		return nil, fmt.Errorf("factory is nil!")
	}
	pool := &sync.Pool{New: func() interface{} {
		return make([]byte, 4096)
	}}
	udpLiscener, err := NewUdpServer(port, ip)
	if err != nil {
		return nil, err
	}
	dispatcher := factory.CreateDispatcher()
	return server{udpLiscener, dispatcher, pool}, nil
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
		data := this.pool.New().([]byte)
		count, remoteaddr, err := this.listener.Read(data)
		if err != nil {
			fmt.Printf("Error on reading from listener %v\n", err)
			continue
		}
		go this.dispatcher.Dispatch(data[:count],fmt.Sprintf("$v",remoteaddr), func(bytes []byte, e error) {
			if bytes != nil{
				this.pool.Put(bytes)
			}
			if e != nil {
				fmt.Printf("Error on dispathing %v\n", err)
				return
			}
			if len(bytes) < 1 {
				return
			}
			_, err := this.listener.Write(bytes, remoteaddr)
			if err != nil {
				fmt.Printf("Couldn't send response - %v \n", e)
			}
		})
	}
}
