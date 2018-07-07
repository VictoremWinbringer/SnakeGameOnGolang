package dal

type IServerDalFactory interface {
	CreateSession() ISession
}

type serverDalFactory struct {
}

func (this *serverDalFactory) CreateSession() ISession {



	return &session{}
}

func NewServerDalFactory() IServerDalFactory {
	return &serverDalFactory{}
}
