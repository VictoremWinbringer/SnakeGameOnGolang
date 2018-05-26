package dal

type IWriter interface {
	Write(x, y int, value rune)
	Data() [][]rune
	Clear()
}

type writer struct {
	data   [][]rune
	height int
	width  int
}

func (this *writer) Write(x, y int, value rune) {
	this.data[x][y] = value
}

func (this *writer) Data() [][]rune {
	matrix := this.createMatrix()
	for i, a := range this.data {
		for j, r := range a {
			matrix[i][j] = r
		}
	}
	return matrix
}

func (this *writer) Clear() {
	this.data = this.createMatrix()
}

func (this *writer) createMatrix() [][]rune {
	matrix := make([][]rune, this.width)
	for i := 0; i < this.width; i++ {
		matrix[i] = make([]rune, this.height)
	}
	return matrix
}
