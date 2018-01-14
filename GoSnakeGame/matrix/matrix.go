package matrix

import (
	"fmt"
	"sync"

	tc "github.com/gdamore/tcell"
)

type Matrix struct {
	value     [][]rune
	mutex     *sync.Mutex
	maxHeight int
	maxWidth  int
}

func New(height, width int) Matrix {
	values := make([][]rune, height)

	for i := range values {
		values[i] = make([]rune, width)
	}
	return Matrix{value: values, mutex: &sync.Mutex{}, maxHeight: height - 1, maxWidth: width - 1}
}

func (m *Matrix) Set(height, width int, value rune) error {

	if height > m.maxHeight {
		return fmt.Errorf("Height is greater than %v", m.maxHeight)
	}

	if width > m.maxWidth {
		return fmt.Errorf("Height is greater than %v", m.maxWidth)
	}

	defer m.mutex.Unlock()

	m.mutex.Lock()

	m.value[height][width] = value

	return nil
}

func (m Matrix) String() string {

	defer m.mutex.Unlock()

	m.mutex.Lock()

	result := ""

	for i := range m.value {
		result += string(m.value[i]) + "\n\r"
	}

	return result
}

func (m *Matrix) Clear() {
	for i, v := range m.value {
		for j := range v {
			m.value[i][j] = ' '
		}
	}
}

func (m *Matrix) Draw(s tc.Screen) {
	for i, v := range m.value {
		for j, r := range v {
			s.SetContent(j, i, r, nil, tc.StyleDefault)
		}
	}
}
