package tests

import (
	"../../src/Shared/serializer"
	"testing"
)

func Test_serializer_Encode_Decode_Message(t *testing.T) {

	bytes := make([]byte, 2)
	bytes[0] = 1
	bytes[1] = 2
	want := serializer.Message{
		3,
		serializer.GameStateType,
		bytes}

	t.Run("test_session_HandleCommand", func(t *testing.T) {
		encoded := serializer.EncodeMessage(want)
		got := serializer.DecodeMessage(encoded)
		if !equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		} else {
			t.Logf("got %v, want %v", got, want)
		}
	})
}

func Test_serializer_Encode_Decode_GameState(t *testing.T){
	runes := make([][]rune,1)
	runes[0] = make([]rune,1)
	runes[0][0] = '+'
	want := serializer.GameState{runes}

	t.Run("test_session_HandleCommand", func(t *testing.T) {
		encoded := serializer.EncodeGameState(want)
		got := serializer.DecodeGameState(encoded)
		if got.State[0][0] != want.State[0][0] {
			t.Errorf("got %v, want %v", got, want)
		} else {
			t.Logf("got %v, want %v", got, want)
		}
	})
}

func Test_serializer_Encode_Decode_Command(t *testing.T){

	want := serializer.Command{8,serializer.MoveRight}
	t.Run("test_session_HandleCommand", func(t *testing.T) {
		encoded := serializer.EncodeCommand(want)
		got := serializer.DecodeCommand(encoded)
		if got.Id != want.Id || got.Code != want.Code {
			t.Errorf("got %v, want %v", got, want)
		} else {
			t.Logf("got %v, want %v", got, want)
		}
	})
}

func equal(x, y serializer.Message) bool {
	if len(x.Data) != len(y.Data) {
		return false
	}
	for i, v := range x.Data {
		if v != y.Data[i] {
			return false
		}
	}
	return x.Id == y.Id && x.Type == y.Type
}
