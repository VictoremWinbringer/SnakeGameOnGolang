package messageTypeEnum

type Type byte

const (
	UndefinedType Type = 0
	GameStateType Type = 1
	CommandType   Type = 2
)
