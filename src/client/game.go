package client

import (
	"github.com/nsf/termbox-go"
	"../board"
	"../cursor"
	"../sockets"
	"sync"
	"fmt"
)

var (
	mutex sync.Mutex
)

type Game struct {

	Board *board.Board
	Cursor *cursor.Cursor
	Socket *sockets.Socket
	
	ComputerScore int
	PlayerScore int

	message string
	WantsFinish bool

	PlayerItem int
}

func (game *Game) Loop() {

	for !game.WantsFinish {

		// mutex.Lock()
		
		game.update()
		game.print()
		game.postUpdate()
		game.handleInput()

		// mutex.Unlock()
	}
}

func (game *Game) update() {
	
}

func (game *Game) postUpdate() {
	// game.message = ""
}

func (game *Game) Reset() {

	game.Board.Reset()
	game.ShowMessage("New game has started.")
}

func (game *Game) print() {

	defer termbox.Flush()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	w, h := termbox.Size()
	centerX := w / 2
	centerY := h / 2

	game.drawHUD()

	var cursor *cursor.Cursor
	if game.Cursor.Enabled {
		cursor = game.Cursor
	} else {
		cursor = nil
	}

	game.Board.Print(centerX, centerY, cursor)
}

func (game *Game) ShowMessage(format string, a ...interface{}) {

	message := fmt.Sprintf(format, a...)
	game.message = message
}

func (game *Game) WantsDisplay() {

	// IMPLEMENT SOMETHING HEEEEEERE!
}

func NewGame() *Game {
	
	game := Game{}
	game.Board = board.New()
	game.Cursor = cursor.New()

	return &game
}

func (game *Game) handleInput() {

	evt := termbox.PollEvent()

	// mutex.Lock()

	switch evt.Type {
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

		case termbox.KeyCtrlN:
			game.RequestNewGame()

		case termbox.KeyEnter:
			game.Play()
		}
	}

	// mutex.Unlock()
}