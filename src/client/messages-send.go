package client

import (
	"log"
	"../board"
)

func (game *Game) requestNewGame() {
	game.Socket.SendMessageId(102)
}

func (game *Game) sendPlay() {

	lin := game.Cursor.Y
	col := game.Cursor.X

	log.Printf("Play %v at (%v, %v).\n", string(board.CharacterForItem(game.PlayerItem)), lin, col)

	game.ShowMessage("Waiting computer to play...")
	game.state = StateIdle

	game.Socket.SendMessage(msgPlay {
		Id: 101,
		Line: lin,
		Column: col,
	})
}