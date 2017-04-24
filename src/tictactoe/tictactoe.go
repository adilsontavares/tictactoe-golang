package tictactoe

import (
	"github.com/nsf/termbox-go"
	"../console"
	"../board"
	"../cursor"
	"strconv"
)

type Game struct {

	Board *board.Board
	Cursor *cursor.Cursor

	ComputerScore int
	PlayerScore int
	
	MouseX int
	MouseY int

	message string
	WantsFinish bool
}

func (game *Game) drawHUDWindow(title string) {

	fg := termbox.ColorDefault
	bg := termbox.ColorDefault

	w, h := termbox.Size()

	// DRAW BORDERS

	for x := 0; x < w; x++ {
		termbox.SetCell(x, 0, '-', fg, bg)
	}

	for x := 0; x < w; x++ {
		termbox.SetCell(x, h - 1, '-', fg, bg)
	}

	for y := 0; y < h; y++ {
		termbox.SetCell(0, y, '|', fg, bg)
	}

	for y := 0; y < h; y++ {
		termbox.SetCell(w - 1, y, '|', fg, bg)
	}

	for x := 0; x < w; x += w - 1 {
		for y := 0; y < h; y += h - 1 {
			termbox.SetCell(x, y, '+', fg, bg)
		}
	}

	// DRAW TITLE

	for x := 1; x < (w - 1); x++ {
		for y := 1; y <= 2; y++ {
			termbox.SetCell(x, y, '-', fg, bg)
		}		
	}

	console.PrintCenter(w / 2, 1, "   " + title + "   ")
}

func (game *Game) drawHUDScore() {

	fg := termbox.ColorDefault
	bg := termbox.ColorDefault

	w, h := termbox.Size()

	for x := 1; x < (w - 1); x++ {
		termbox.SetCell(x, h - 3, '-', fg, bg)
	}

	termbox.SetCell(w / 2, h - 3, '+', fg, bg)
	termbox.SetCell(w / 2, h - 2, '|', fg, bg)

	computer := strconv.Itoa(game.ComputerScore)
	console.Print(2, h - 2, "Computador (X):")
	console.Print(w / 2 - 1 - len(computer), h - 2, computer)

	player := strconv.Itoa(game.PlayerScore)
	console.Print(w / 2 + 2, h - 2, "Jogador (O):")
	console.Print(w - 2 - len(player), h - 2, player)
}

func (game *Game) drawHUDMessage(message string) {
	
	w, h := termbox.Size()
	console.PrintCenter(w / 2, h - 5, message)
}

func (game *Game) drawHUD() {

	game.drawHUDWindow("TIC TAC TOE")
	game.drawHUDScore()
	game.drawHUDMessage(game.message)
}

func (game *Game) update() {
	// UPDATE
}

func (game *Game) postUpdate() {
	game.message = ""
}

func (game *Game) print() {

	defer termbox.Flush()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	w, h := termbox.Size()
	centerX := w / 2
	centerY := h / 2

	game.drawHUD()
	game.Board.Print(centerX, centerY, game.Cursor)
}

func (game *Game) ShowMessage(message string) {
	game.message = message
}

func (game *Game) moveCursor(offx int, offy int) {

	cursor := game.Cursor

	cursor.X += offx
	cursor.Y += offy

	if cursor.X < 0 {
		cursor.X = 0
	} else if cursor.X >= 2 {
		cursor.X = 2
	}

	if cursor.Y < 0 {
		cursor.Y = 0
	} else if cursor.Y >= 2 {
		cursor.Y = 2
	}
}

func (game *Game) moveCursorDown() {
	game.moveCursor(0, -1)
}

func (game *Game) moveCursorUp() {
	game.moveCursor(0, 1)	
}

func (game *Game) moveCursorRight() {
	game.moveCursor(1, 0)
}

func (game *Game) moveCursorLeft() {
	game.moveCursor(-1, 0)
}

func (game *Game) handleInput() {

	switch evt := termbox.PollEvent(); evt.Type {
	case termbox.EventKey:
		switch evt.Key {
		case termbox.KeyEsc:
			game.WantsFinish = true

		case termbox.KeyArrowDown:
			game.moveCursorDown()

		case termbox.KeyArrowUp:
			game.moveCursorUp()

		case termbox.KeyArrowRight:
			game.moveCursorRight()

		case termbox.KeyArrowLeft:
			game.moveCursorLeft()

		case termbox.KeyEnter:
			game.ShowMessage("You've pressed RETURN!")
		}
	}
}

func New() *Game {
	
	game := Game{}
	game.Board = board.New()
	game.Cursor = cursor.New()

	return &game
}