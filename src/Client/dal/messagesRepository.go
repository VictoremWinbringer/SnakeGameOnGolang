package dal

import (
	"../../Shared/serializer"
	"../../Shared/udp"
)

type IMessagesRepository interface {
	Read() (serializer.Message, error)
	Write(message serializer.Message) error
	Dispose()
}

func newIMessageRepository(client udp.IUdpClient) IMessagesRepository {
	return messagesRepository{
        udpClient: client}
}

type messagesRepository struct {
udpClient udp.IUdpClient
}

func (this messagesRepository) Read() (serializer.Message, error) {
	p := make([]byte, 4096)
 _ , e :=	this.udpClient.Read(p)
 m := serializer.DecodeMessage(p)
 return m, e
	//message := serializer.Message{}
	//err := this.decoder.Decode(&message)
	//return message, err
}

func (this messagesRepository) Write(message serializer.Message) error {
	p := serializer.EncodeMessage(message)
	_, e := this.udpClient.Write(p)
	return e
//return this.encoder.Encode(message)
}

func (this messagesRepository) Dispose()  {
this.udpClient.Close()
}
