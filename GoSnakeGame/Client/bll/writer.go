package bll

type IWriter interface {
	Write(x, y int, value rune)
}
