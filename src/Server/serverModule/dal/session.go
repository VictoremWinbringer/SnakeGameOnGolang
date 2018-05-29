package dal

type ISession interface {
	HandleCommand(command int)
	GetState() [][]rune
}

type session struct {
}

func Mul(x, y int) int {
	return x * y
}
