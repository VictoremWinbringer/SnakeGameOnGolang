package dal

type ISession interface {
	HandleCommand(command int)
	GetState() [][]rune
}
