package client

import (
	"github.com/nsf/termbox-go"
	"strconv"
	"../console"
	"../board"
)

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
	console.Printf(2, h - 2, "[%v] Computer score:", string(board.CharacterForItem(board.OppositeItem(game.PlayerItem))))
	console.Print(w / 2 - 1 - len(computer), h - 2, computer)

	player := strconv.Itoa(game.PlayerScore)
	console.Printf(w / 2 + 2, h - 2, "[%v] Your score:", string(board.CharacterForItem(game.PlayerItem)))
	console.Print(w - 2 - len(player), h - 2, player)
}

func (game *Game) drawHUDInstruction() {

	var instructions string

	if game.state == StateWaitingPlay {
		instructions = "TIP: Press RETURN to select a slot."
	} else if game.state == StateFinished {
		instructions = "TIP: Press RETURN to start a new game."
	} else {
		instructions = "TIP: Hey there! I am a flashing tip :)"
	}

	w, h := termbox.Size()
	console.PrintCenter(w / 2, h - 5, instructions)
}

func (game *Game) drawHUDMessage(message string) {
	
	w, h := termbox.Size()
	console.PrintCenter(w / 2, (h / 2 - 2) / 2, message)
}

func (game *Game) drawHUDAlert(alert string) {
	
	w, h := termbox.Size()
	console.PrintColor(w / 2 - len(alert) / 2, (h / 2 - 2) / 2, alert, termbox.ColorDefault, termbox.AttrBold)
}

func (game *Game) drawHUD() {

	game.drawHUDWindow("TIC TAC TOE")
	game.drawHUDScore()
	game.drawHUDInstruction()

	if game.alert != "" {
		game.drawHUDAlert(game.alert)
	} else if game.message != "" {
		game.drawHUDMessage(game.message)
	}
}