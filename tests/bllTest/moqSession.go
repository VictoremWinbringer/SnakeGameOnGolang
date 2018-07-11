package tests

import (
	"../../src/Shared/serializer"
)

func Equals(x, y serializer.GameState) bool {
	if len(x.State) != len(y.State) {
		return false
	}

	for i, s := range x.State {
		for j, r := range s {
			if r != y.State[i][j] {
				return false
			}
		}
	}
	return true
}

type moqSession struct {
	CommandFunc func(int)
}

func (this moqSession) GetState() [][]rune {
	return makeMatrix()
}

func (this moqSession) HandleCommand(command int) {
	this.CommandFunc(command)
}

func (this moqSession) Start()  {

}

func (this moqSession) Stop()  {

}

func makeMatrix() [][]rune {
	matrix := make([][]rune, 1)
	matrix[0] = make([]rune, 1)
	matrix[0][0] = 1
	return matrix
}
