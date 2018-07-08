package dal

import (
	"../../Shared/serializer"
	"encoding/gob"
	"../../Shared/udp"
)

type IMessagesRepository interface {
	Read() (serializer.Message, error)
	Write(message serializer.Message) error
	Dispose()
}

func newIMessageRepository(client udp.IUdpClient) IMessagesRepository {
	return messagesRepository{
		encoder:   gob.NewEncoder(client),
		decoder:   gob.NewDecoder(client),
        udpClient: client}
}

type messagesRepository struct {
encoder *gob.Encoder
decoder *gob.Decoder
udpClient udp.IUdpClient
}

func (this messagesRepository) Read() (serializer.Message, error) {
	message := serializer.Message{}
	err := this.decoder.Decode(&message)
	return message, err
}

func (this messagesRepository) Write(message serializer.Message) error {
return this.encoder.Encode(message)
}

func (this messagesRepository) Dispose()  {
this.udpClient.Close()
}
