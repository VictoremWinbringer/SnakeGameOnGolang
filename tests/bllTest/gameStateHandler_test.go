package tets

import (
	"testing"

	"../../src/Server/serverModule/bll"
	"../../src/Shared/serializer"
)

func Test_gameStateHandler_Type(t *testing.T) {

	want := bll.HandlerType(serializer.GameStateType)
	handler := bll.NewSeverBllFactory().CreateGameStateHandler()
	t.Run("test_session_HandleCommand", func(t *testing.T) {
		got := handler.Type()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func Test_gameStateHandler_Handle(t *testing.T) {
	want := serializer.GameState{makeMatrix()}
	handler := bll.NewSeverBllFactory().CreateGameStateHandler()
	t.Run("test_session_HandleCommand", func(t *testing.T) {
		gotBytes, ok := handler.Handle(make([]byte, 0), moqSession{})
		if !ok {
			t.Error("Not Ok!")
		}
		got := serializer.DecodeGameState(gotBytes)
		if !Equals(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

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
}

func (this moqSession) GetState() [][]rune {
	return makeMatrix()
}

func (this moqSession) HandleCommand(command int) {
	//Do nothing
}

func makeMatrix() [][]rune {
	matrix := make([][]rune, 1)
	matrix[0] = make([]rune, 1)
	matrix[0][0] = 1
	return matrix
}
