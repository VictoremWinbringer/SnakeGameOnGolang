package dal

import(
	"../../../Shared/udp"
	"io"
)

type IDataReaderWriter interface {
	io.Reader
	io.Writer
	GetConnection() udp.Connection
	SetConnection(connection udp.Connection)
}

func newDataReaderWriter(listener udp.IUdpListener) IDataReaderWriter {
	return &dataRW{nil,listener}
}

type dataRW struct {
	connecton udp.Connection
	listener udp.IUdpListener
}

func (this *dataRW) GetConnection() udp.Connection {
	return this.connecton
}

func (this *dataRW) SetConnection(connection udp.Connection)  {
	this.connecton = connection
}

func (this *dataRW) Read(p []byte) (n int, err error){
	n, this.connecton, err = this.listener.Read(p)
	return  n, err
}

func (this *dataRW) Write(p []byte) (n int, err error) {
	return this.listener.Write(p, this.connecton)
}