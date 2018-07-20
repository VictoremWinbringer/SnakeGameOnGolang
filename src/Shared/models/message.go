package models

import "../messageTypeEnum"

type Message struct {
	Type messageTypeEnum.Type
	Data []byte
}
