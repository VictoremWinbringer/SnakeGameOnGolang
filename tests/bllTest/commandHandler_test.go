package tests

import (
	"testing"

	"../../src/Server/serverModule/bll"
	"../../src/Shared/serializer"
)

//TODO: need negative tests

func Test_commandHandler_Type(t *testing.T) {

	want := bll.HandlerType(serializer.CommandType)
	handler := bll.NewSeverBllFactory().CreateCommandHandler()
	t.Run("test_session_HandleCommand", func(t *testing.T) {
		got := handler.Type()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

// func Test_commandHandler_Handle(t *testing.T) {
// 	want := serializer.GameState{makeMatrix()}
// 	handler := bll.NewSeverBllFactory().CreateGameStateHandler()
// 	t.Run("test_session_HandleCommand", func(t *testing.T) {
// 		gotBytes, ok := handler.Handle(make([]byte, 0), moqSession{})
// 		if !ok {
// 			t.Error("Not Ok!")
// 		}
// 		got := serializer.DecodeGameState(gotBytes)
// 		if !Equals(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
