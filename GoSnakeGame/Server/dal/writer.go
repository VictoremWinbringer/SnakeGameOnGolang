package dal

import (
	"fmt"
)

type IWriter interface {
	Write(x, y int, value rune)
	String() string
	Clear()
}

type writer struct {
	data   []rune
	height int
	width  int
}

func (this *writer) Write(x, y int, value rune) {
	index := this.width*y + x
	this.data[index] = value
}

func (this *writer) String() string {
	result := ""
	begin := 0
	end := 0
	for i := 1; i < this.height; i++ {
		end = this.width * i
		result += fmt.Sprintln(this.data[begin:end])
		begin = end
	}
	return string(this.data)
}

func (this *writer) Clear() {
	this.data = make([]rune, this.width*this.height)
}
