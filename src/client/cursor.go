package client

func (game *Game) moveCursor(offx int, offy int) {

	if game.state != StateWaitingPlay {
		return
	}

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
	game.moveCursor(0, 1)
}

func (game *Game) moveCursorUp() {
	game.moveCursor(0, -1)	
}

func (game *Game) moveCursorRight() {
	game.moveCursor(1, 0)
}

func (game *Game) moveCursorLeft() {
	game.moveCursor(-1, 0)
}