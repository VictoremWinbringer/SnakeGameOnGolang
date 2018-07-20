package commandCodeEnum

type Type byte

const (
	UndefinedCommand Type = 0
	MoveUp           Type = 1
	MoveDown         Type = 2
	MoveLeft         Type = 3
	MoveRight        Type = 4
	ExitGame         Type = 5
)
