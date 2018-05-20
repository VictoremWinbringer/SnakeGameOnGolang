package serializer

import (
	"bytes"
	"encoding/gob"
)

const GameStateType byte = 0
const CommandType byte = 1

type GameState struct {
	State string
}

type Command struct {
	Code uint8
}

func EncodeGameState(value GameState) []byte {
	b := new(bytes.Buffer)
	b.WriteByte(GameStateType)
	e := gob.NewEncoder(b)
	e.Encode(value)
	return b.Bytes()
}

func DecodeGameState(data []byte) GameState {
	b := bytes.NewBuffer(data[1:])
	d := gob.NewDecoder(b)
	var input GameState
	d.Decode(&input)
	return input
}

func EncodeCommand(value Command) []byte {
	b := new(bytes.Buffer)
	b.WriteByte(CommandType)
	e := gob.NewEncoder(b)
	e.Encode(value)
	return b.Bytes()
}

func DecodeCommand(data []byte) Command {
	b := bytes.NewBuffer(data[1:])
	d := gob.NewDecoder(b)
	var input Command
	d.Decode(&input)
	return input
}
