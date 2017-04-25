package server

import (
	"../board"
)

func (client *Client) interpretMessage(id int, data map[string]interface{}) bool {

	switch id {
	case 101:	return client.wantsToPlay(data)
	case 102:	return client.startNewGame()
	}

	return false
}

func (client *Client) wantsToPlay(data map[string]interface{}) bool {
	
	line := data["lin"]
	column := data["col"]

	if line == nil || column == nil {
		return false
	}

	lin := int(line.(float64))
	col := int(column.(float64))

	if client.Board.Place(client.Item, lin, col) {

		client.log("Player placed %v at (%v, %v).", string(board.CharacterForItem(client.Item)), lin, col)

		success := client.sendMessage(msgPlacePlayer {
			Id: 4,
			Player: client.Item,
			Line: lin,
			Column: col,
		}) 

		if !client.Board.IsGameOver() {
			success = success && client.computerPlays()
		}

		success = success && client.validatesGameOver()

		if !client.Board.IsGameOver() {
			return success && client.requestPlay()
		} else {
			return success
		}
	}

	return client.invalidPosition(lin, col)
}

func (client *Client) computerPlays() bool {

	lin, col := client.Board.GetFreePos()
	
	if lin == -1 || col == -1 {
		return true
	}

	player := board.OppositeItem(client.Item)
	client.Board.Place(player, lin, col)

	client.log("Computer placed %v at (%v, %v).", string(board.CharacterForItem(player)), lin, col)

	return client.sendMessage(msgPlacePlayer {
		Id: 4,
		Player: player,
		Line: lin,
		Column: col,
	})
}

func (client *Client) invalidPosition(line int, column int) bool {

	client.log("Invalid position for %v at (%v, %v).", string(board.CharacterForItem(client.Item)), line, column)

	return client.sendMessage(msgInvalidPosition {
		Id: 6,
		Line: line,
		Column: column,
	})
}

func (client *Client) requestPlay() bool {

	client.log("Requesting play.")

	return client.sendMessage(msgRequestPlay {
		Id: 3,
	})
}

func (client *Client) validatesGameOver() bool {

	if !client.Board.IsGameOver() {
		return true
	}

	winner := client.Board.FindWinner()

	if winner == client.Item {
		client.log("Game Over! Player won.")
	} else if winner == board.OppositeItem(client.Item) {
		client.log("Game Over! Player lose.")
	} else {
		client.log("Game Over! It's a tie.")
	}

	return client.sendMessage(msgGameOver {
		Id: 2,
		Winner: winner,
	})
}

func (client *Client) startNewGame() bool {

	client.log("Starting a new game.")
	client.Reset()

	return client.sendMessage(msgNewGame {
		Id: 5,
		Player: board.ItemX,
	}) && client.requestPlay()
}