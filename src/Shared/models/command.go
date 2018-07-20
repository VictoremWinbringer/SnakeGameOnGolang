package models

import "../commandCodeEnum"

type Command struct {
	Id   uint64
	Code commandCodeEnum.Type
}
