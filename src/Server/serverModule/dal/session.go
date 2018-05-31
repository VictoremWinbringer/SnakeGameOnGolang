package dal

type ISession interface {
	HandleCommand(command int)
	GetState() [][]rune
}

type session struct {
	commandFunc func(int)
	stateFunc   func() [][]rune
}

func (this *session) HandleCommand(command int) {
	this.commandFunc(command)
}

func (this session) GetState() [][]rune {
	return this.stateFunc()
}

func Mul(x, y int) int {
	return x * y
}
