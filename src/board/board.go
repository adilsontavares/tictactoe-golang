package board

import (
	"github.com/nsf/termbox-go"
	"../console"
	"../cursor"
)

var (
	ItemNone int = 0
	ItemX int = 1
	ItemO int = 2
)

type Board struct {

	items [3][3]int

	Cursor *cursor.Cursor
	
	DivisionsFgColor termbox.Attribute
	PlayersFgColor termbox.Attribute
	BorderFgColor termbox.Attribute

	DivisionsBgColor termbox.Attribute
	PlayersBgColor termbox.Attribute
	BorderBgColor termbox.Attribute
}

func New() *Board {
	
	board := Board{}
	board.items[0][0] = ItemX
	board.items[1][1] = ItemX
	board.items[2][2] = ItemX
	board.items[0][1] = ItemO
	board.items[0][2] = ItemO
	board.items[2][0] = ItemX
	board.items[2][1] = ItemO
	board.items[1][2] = ItemX

	board.Cursor = cursor.New()

	board.DivisionsFgColor = termbox.ColorDefault
	board.PlayersFgColor = termbox.ColorDefault
	board.BorderFgColor = termbox.ColorDefault

	board.DivisionsBgColor = termbox.ColorDefault
	board.PlayersBgColor = termbox.ColorDefault
	board.BorderBgColor = termbox.ColorDefault

	return &board
}

func characterForItem(item int) rune {

	return 'X'

	switch item {
	case ItemNone:
		return ' '

	case ItemX:
		return 'X'

	case ItemO:
		return 'O'
	}

	return '?'
}

func (board *Board) drawBorder(x int, y int) {

	fg := board.BorderFgColor
	bg := board.BorderBgColor

	w := 20
	h := 10

	w_2 := w / 2
	h_2 := h / 2

	for dx := -w_2; dx <= w_2; dx++ {
		termbox.SetCell(x + dx, y + h_2, '-', fg, bg)
	}

	for dx := -w_2; dx <= w_2; dx++ {
		termbox.SetCell(x + dx, y - h_2, '-', fg, bg)
	}

	for dy := -h_2; dy <= h_2; dy++ {
		termbox.SetCell(x - w_2, y + dy, '|', fg, bg)
	}

	for dy := -h_2; dy <= h_2; dy++ {
		termbox.SetCell(x + w_2, y + dy, '|', fg, bg)
	}

	for dx := -w_2; dx <= w_2; dx += w {
		for dy := -h_2; dy <= h_2; dy += h {
			termbox.SetCell(x + dx, y + dy, '+', fg, bg)
		}
	}
} 

func (board *Board) drawDivisions(x int, y int) {

	fg := board.DivisionsFgColor
	bg := board.DivisionsBgColor

	for i := 0; i < 2; i++ {

		for j := 0; j < 3; j++ {

			dx := -2 + i * 4
			dy := -2 + j * 2

			termbox.SetCell(x + dx, y + dy, '|', fg, bg)
		}
	}

	for i := 0; i < 2; i++ {

		for j := 0; j < 3; j++ {

			dx := -5 + j * 4
			dy := -1 + i * 2

			console.Print(x + dx, y + dy, "---")
		}
	}

	for i := 0; i < 2; i++ {

		for j := 0; j < 2; j++ {

			dx := -2 + j * 4
			dy := -1 + i * 2

			termbox.SetCell(x + dx, y + dy, '+', fg, bg)
		}
	}
}

func (board *Board) drawPlayers(x int, y int) {

	fg := board.PlayersFgColor
	bg := board.PlayersBgColor

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			dx := -4 + j * 4
			dy := 2 + i * -2

			termbox.SetCell(x + dx, y + dy, characterForItem(board.items[i][j]), fg, bg)
		}
	}
}

func (board *Board) drawCursor(x int, y int, cursor *cursor.Cursor) {

	fg := cursor.FgColor
	bg := cursor.BgColor

	sw, _ := termbox.Size()

	bx := x - 4 + cursor.X * 4
	by := y + 2 + cursor.Y * -2

	for j := -1; j <= 1; j++ {
		for i := -2; i <= 2; i++ {

			fx := bx + i
			fy := by + j

			cell := termbox.CellBuffer()[fy * sw + fx]
			console.PrintColor(fx, fy, string(cell.Ch), fg, bg)
		}
	}
}

func (board *Board) Print(x int, y int, cursor *cursor.Cursor) {

	board.drawBorder(x, y)
	board.drawDivisions(x, y)
	board.drawPlayers(x, y)

	if cursor != nil {
		board.drawCursor(x, y, cursor)
	}
}