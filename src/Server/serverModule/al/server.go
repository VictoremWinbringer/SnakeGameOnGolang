package al

import (
	"fmt"

	"sync"

	"../bll"
	. "../dal"
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

const COUNT_THREADS_IN_POOL = 1000

type inValue struct {
	count      int
	data       []byte
	remoteaddr Connection
}

type outValue struct {
	data []byte
	remoteaddr Connection
}

func (this server) listen() {
	in := make(chan inValue, 100)
	out := make(chan outValue, 100)
	defer func() {
		this.listener.Close()
		this.dispatcher.Close()
		close(in)
		close(out)
	}()
	for i := 0; i < COUNT_THREADS_IN_POOL; i++ {
		go func() {
			for input := range in {
				bytes, err := this.dispatcher.Dispatch(input.data[:input.count], fmt.Sprintf("%v", input.remoteaddr))
				this.pool.Put(input.data)
				if err != nil {
					fmt.Printf("Error on dispathing %v\n", err)
					continue
				}
				if len(bytes) < 1 {
					continue
				}
				out <- outValue{bytes,input.remoteaddr}
			}
		}()
	}
	go func() {
		for result := range out {
			for i := 0; i < len(result.data); {
				count, err := this.listener.Write(result.data, result.remoteaddr)
				if err != nil {
					fmt.Printf("Couldn't send response - %v \n", err)
					break
				}
				i+=count
			}
		}
	}()

	for {
		data := this.pool.Get().([]byte)
		count, remoteaddr, err := this.listener.Read(data)
		if err != nil {
			fmt.Printf("Error on reading from listener %v\n", err)
			continue
		}
		in <- inValue{count:count,remoteaddr:remoteaddr, data:data}
	}
}
