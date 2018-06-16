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
