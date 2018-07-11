package tests

import (
	"testing"

	dal "../../src/Server/serverModule/dal"
)

func Test_session_HandleCommand(t *testing.T) {
	got := 0
	want := 1
	session := dal.NewServerDalFactory().CreateSession()
	t.Run("test_session_HandleCommand", func(t *testing.T) {
		session.HandleCommand(1)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func Test_session_GetState(t *testing.T) {

	want := make([][]rune, 1)
	want[0] = make([]rune, 1)
	want[0][0] = 1
	session := dal.NewServerDalFactory().CreateSession()
	t.Run("test_session_HandleCommand", func(t *testing.T) {
		got := session.GetState()
		if !isEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func isEqual(x, y [][]rune) bool {
	if len(x) != len(y) {
		return false
	}
	for i, a := range x {
		if len(a) != len(y[i]) {
			return false
		}
		for j, _ := range a {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}
	return true
}
