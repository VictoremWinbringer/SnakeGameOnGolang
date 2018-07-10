package model

import (
	"../../../Shared/serializer"
	"../../../Shared/udp"
)

type Data struct {
	Message    serializer.Message
	Connection udp.Connection
}
