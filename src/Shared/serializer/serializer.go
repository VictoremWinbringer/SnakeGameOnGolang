package serializer

import (
	"bytes"
	"encoding/gob"

	. "../models"
)

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
