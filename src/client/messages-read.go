package client

import (
	"../board"
	"log"
)

func (game *Game) interpretMessage(id int, data map[string]interface{}) bool {

	log.Printf("##### Received message with ID %v", id)

	switch id {
	case 2:	return game.gameFinished(data)
	case 3: return game.requestPlay()
	case 4: return game.itemWasPlaced(data)
	case 5: return game.startNewGame(data)
	case 6: return game.invalidPosition(data)
	}

	return false
}

func (game *Game) gameFinished(data map[string]interface{}) bool {

	winner := data["winner"]
	computerScore := data["cscore"]
	playerScore := data["pscore"]

	if winner == nil || computerScore == nil || playerScore == nil {
		return false
	}

	var message string
	win := int(winner.(float64))
	cscore := int(computerScore.(float64))
	pscore := int(playerScore.(float64))

	if win == game.PlayerItem {
		message = "Game over. You won!"
	} else if win == board.OppositeItem(game.PlayerItem) {
		message = "Game over. You lose..."
	} else {
		message = "Game over! That's a tie."
	}

	game.PlayerScore = pscore
	game.ComputerScore = cscore

	game.ShowMessage(message)
	log.Printf("%v\n", message)

	game.state = StateFinished
	game.WantsDisplay()

	return true
}

func (game *Game) requestPlay() bool {

	log.Printf("Player, it is your turn!\n")
	
	game.state = StateWaitingPlay

	// game.ShowMessage("Your turn! Place the %v.", string(board.CharacterForItem(game.PlayerItem)))
	game.ClearMessage()
	game.WantsDisplay()

	return true
}

func (game *Game) itemWasPlaced(data map[string]interface{}) bool {

	player := data["player"]
	line := data["lin"]
	column := data["col"]

	if player == nil || column == nil || line == nil {
		return false
	}

	item := int(player.(float64))
	lin := int(line.(float64))
	col := int(column.(float64))

	log.Printf("Place an %v at (%v, %v).\n", string(board.CharacterForItem(item)), lin, col)

	game.Board.Place(item, lin, col)
	game.WantsDisplay()

	return true
}

func (game *Game) startNewGame(data map[string]interface{}) bool {

	player := data["player"]

	if player == nil {
		return false
	}

	game.state = StateIdle

	game.Reset()
	game.PlayerItem = int(player.(float64))
	game.WantsDisplay()

	return true
}

func (game *Game) invalidPosition(data map[string]interface{}) bool {

	line := data["lin"]
	column := data["col"]

	if line == nil || column == nil {
		return false
	}

	game.ShowMessage("You can not place a %v there.", string(board.CharacterForItem(game.PlayerItem)))
	game.WantsDisplay()

	return true
}