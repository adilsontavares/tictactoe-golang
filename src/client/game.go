package client

import (
	"github.com/nsf/termbox-go"
	"../board"
	"../cursor"
	"../sockets"
	"time"
	"fmt"
)

const (
	StateIdle 			= 0
	StateWaitingPlay 	= 1
	StateFinished	 	= 2
)

type Game struct {

	Board *board.Board
	Cursor *cursor.Cursor
	Socket *sockets.Socket
	
	ComputerScore int
	PlayerScore int

	alert string
	message string
	state int

	WantsFinish bool

	PlayerItem int
}

func (game *Game) Loop() {

	eventQueue := make(chan termbox.Event)
	loopTick := time.NewTicker(40 * time.Millisecond)

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

mainLoop:
	for {

		select {
		case ev := <- eventQueue:
			game.handleEvent(ev)

		case <- loopTick.C:
			game.update()
			game.print()
			game.postUpdate()

			if game.WantsFinish {
				break mainLoop
			}	
		}
	}
}

func (game *Game) update() {
	
}

func (game *Game) postUpdate() {
	

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
	if game.state == StateWaitingPlay {
		cursor = game.Cursor
	} else {
		cursor = nil
	}

	game.Board.Print(centerX, centerY, cursor)
}

func (game *Game) ClearMessage() {
	game.ShowMessage("")
}

func (game *Game) ShowMessage(format string, a ...interface{}) {

	message := fmt.Sprintf(format, a...)
	game.message = message

	game.alert = ""
}

func (game *Game) ShowAlert(format string, a ...interface{}) {

	alert := fmt.Sprintf(format, a...)
	game.alert = alert

	game.message = ""
}

func (game *Game) WantsDisplay() {
	termbox.Interrupt()
}

func (game *Game) Quit() {
	game.WantsFinish = true
}

func (game *Game) StartNew() {
	game.requestNewGame()
}

func (game *Game) Play() {
	game.sendPlay()
}

func NewGame() *Game {
	
	game := Game{}
	game.state = StateIdle
	game.Board = board.New()
	game.Cursor = cursor.New()

	return &game
}

func (game *Game) handleEvent(evt termbox.Event) {

	switch evt.Type {
	case termbox.EventKey:

		switch evt.Ch {
		case 'q', 'Q':
			game.Quit()

		case 'n', 'N':
			game.StartNew()
		}

		switch evt.Key {
		case termbox.KeyArrowDown:
			game.moveCursorDown()

		case termbox.KeyArrowUp:
			game.moveCursorUp()

		case termbox.KeyArrowRight:
			game.moveCursorRight()

		case termbox.KeyArrowLeft:
			game.moveCursorLeft()

		case termbox.KeyEnter:
			if game.state == StateWaitingPlay {
				game.Play()
			} else if game.state == StateFinished {
				game.StartNew()
			}
		}
	}
}