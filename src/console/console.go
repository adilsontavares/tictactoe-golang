package console

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

func Begin() {
	
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.OutputNormal)
}

func PrintColor(x int, y int, s string, fg termbox.Attribute, bg termbox.Attribute) {
	for i, r := range (s) {
		termbox.SetCell(x + i, y, r, fg, bg)
	}
}

func Printf(x int, y int, format string, a ...interface{}) {

	str := fmt.Sprintf(format, a...)
	Print(x, y, str)
}

func Print(x int, y int, s string) {
	PrintColor(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func PrintCenter(x int, y int, s string) {
	Print(x - len(s) / 2, y, s)
}

func End() {

	termbox.Close()
}