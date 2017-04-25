package cursor

import (
	"github.com/nsf/termbox-go"
)

type Cursor struct {

	X int
	Y int

	Enabled bool

	FgColor termbox.Attribute
	BgColor termbox.Attribute
}

func New() *Cursor {

	cursor := Cursor{}
	cursor.Enabled = true
	cursor.FgColor = termbox.ColorBlack
	cursor.BgColor = termbox.ColorWhite

	return &cursor
}