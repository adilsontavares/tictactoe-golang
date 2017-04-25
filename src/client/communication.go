package client

import (
	"log"
	"../board"
)

func (game *Game) RequestNewGame() {
	game.Socket.SendMessageId(102)
}

func (game *Game) Play() {

	lin := game.Cursor.Y
	col := game.Cursor.X

	log.Printf("Play %v at (%v, %v).\n", string(board.CharacterForItem(game.PlayerItem)), lin, col)

	game.Cursor.Enabled = false

	game.Socket.SendMessage(msgPlay {
		Id: 101,
		Line: lin,
		Column: col,
	})
}