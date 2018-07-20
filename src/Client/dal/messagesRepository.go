package dal

import (
	. "../../Shared/models"
	"../../Shared/serializer"
)

type IMessagesRepository interface {
	Read() (Message, error)
	Write(message Message) error
	Dispose()
}

func newIMessageRepository(client IUdpClient) IMessagesRepository {
	return messagesRepository{
		udpClient: client}
}

type messagesRepository struct {
	udpClient IUdpClient
}

func (this messagesRepository) Read() (Message, error) {
	p := make([]byte, 4096)
	_, e := this.udpClient.Read(p)
	m := serializer.DecodeMessage(p)
	return m, e
}

func (this messagesRepository) Write(message Message) error {
	p := serializer.EncodeMessage(message)
	_, e := this.udpClient.Write(p)
	return e
}

func (this messagesRepository) Dispose() {
	this.udpClient.Close()
}
