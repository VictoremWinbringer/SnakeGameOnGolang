package dal

type IServerDalFactory interface {
	CreateSession() ISession
}

type serverDalFactory struct {
}
