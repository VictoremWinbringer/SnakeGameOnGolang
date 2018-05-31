package dal

type IServerDalFactory interface {
	CreateSession(commandFunc func(int), stateFunc func() [][]rune) ISession
}

type serverDalFactory struct {
}

func (this *serverDalFactory) CreateSession(commandFunc func(int), stateFunc func() [][]rune) ISession {
	return &session{commandFunc, stateFunc}
}

func NewServerDalFactory() IServerDalFactory {
	return &serverDalFactory{}
}
