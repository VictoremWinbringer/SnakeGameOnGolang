package dal

import (
	"../../Shared/udp"

	tc "github.com/gdamore/tcell"
)

type IDalFactory interface {
	CreateMessagesRepository(client udp.IUdpClient) IMessagesRepository
	CreateScreen() (IScreen, error)
}

type dalFactory struct {
}

func CreateDalFactory() IDalFactory {
	return dalFactory{}
}
func (this dalFactory) CreateScreen() (IScreen, error) {
	s, e := tc.NewScreen()
	if e != nil {
		return nil, e
	}
	if e := s.Init(); e != nil {
		return nil, e
	}
	s.SetStyle(tc.StyleDefault)
	s.HideCursor()
	return screen{s}, nil
}

func (this dalFactory) CreateMessagesRepository(client udp.IUdpClient) IMessagesRepository {
	return newIMessageRepository(client)
}
