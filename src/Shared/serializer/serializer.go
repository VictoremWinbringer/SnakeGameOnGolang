package serializer

import (
	"bytes"
	"encoding/gob"
)

type MessageType byte

const UndefinedType MessageType = 0
const GameStateType MessageType = 1
const CommandType MessageType = 2

type Message struct {
	Id   uint64
	Type MessageType
	Data []byte
}

type PartialContent struct {
	PartNubert      int
	TotalPartsCount int
	PartType        byte
	Data            []byte
}

type GameState struct {
	State [][]rune
}

type CommandCode byte

const (
	UndefinedCommand CommandCode = 0
	MoveUp           CommandCode = 1
)

type Command struct {
	Id uint64
	Code CommandCode
}

func EncodeMessage(value Message) []byte {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	e.Encode(value)
	return b.Bytes()
}

func DecodeMessage(data []byte) Message {
	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)
	var input Message
	d.Decode(&input)
	return input
}

func EncodeGameState(value GameState) []byte {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	e.Encode(value)
	return b.Bytes()
}

func DecodeGameState(data []byte) GameState {
	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)
	var input GameState
	d.Decode(&input)
	return input
}

func EncodeCommand(value Command) []byte {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	e.Encode(value)
	return b.Bytes()
}

func DecodeCommand(data []byte) Command {
	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)
	var input Command
	d.Decode(&input)
	return input
}
