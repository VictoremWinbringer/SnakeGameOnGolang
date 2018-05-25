package dal

import tc "github.com/gdamore/tcell"

type IScreen interface {
	Write(x, y int, value rune)
	Show()
	Clear()
	ShowCursor(x, y int)
	ReadKey() Key
}

type screen struct {
	screen tc.Screen
}

type Key uint8

const KeyUndefined Key = 0
const KeyEsc Key = 1
const KeyUp Key = 2
const KeyDown Key = 3
const KeyLeft Key = 4
const KeyRight Key = 5

func (this screen) Write(x, y int, value rune) {
	this.screen.SetContent(x, y, value, nil, tc.StyleDefault)
}

func (this screen) Clear() {
	this.screen.Clear()
}

func (this screen) Show() {
	this.screen.Show()
}

func (this screen) ReadKey() Key {
	event := this.screen.PollEvent()
	keyEvent, ok := event.(*tc.EventKey)
	if ok {
		switch keyEvent.Key() {
		case tc.KeyUp:
			return KeyUp
		case tc.KeyDown:
			return KeyDown
		case tc.KeyLeft:
			return KeyLeft
		case tc.KeyRight:
			return KeyRight
		case tc.KeyEsc:
			return KeyEsc
		}
	}
	return KeyUndefined
}

func (this screen) ShowCursor(x, y int) {
	this.screen.ShowCursor(x, y)
}
