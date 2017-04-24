package tictactoe

func (game *Game) Loop() {

	game.print()

	for !game.WantsFinish {

		game.handleInput()
		game.update()
		game.print()

		game.postUpdate()
	}
}