package board

// import (

// )

func (board *Board) findWinnerOnHorizontal(line int) int {

	for i := 1; i < 3; i++ {
		if board.items[line][i] != board.items[line][i - 1] {
			return -1
		}
	}

	if winner := board.items[line][0]; winner != ItemNone {
		return winner
	}

	return -1
}

func (board *Board) findWinnerOnVertical(column int) int {

	for i := 1; i < 3; i++ {
		if board.items[i][column] != board.items[i - 1][column] {
			return -1
		}
	}

	if winner := board.items[0][column]; winner != ItemNone {
		return winner
	}

	return -1
}

func (board *Board) findWinnerOnDiagonal() int {

	for i := 1; i < 3; i++ {
		if board.items[i][i] != board.items[i - 1][i - 1] {
			return -1
		}
	}

	if winner := board.items[0][0]; winner != ItemNone {
		return winner
	}

	return -1
}

func (board *Board) findWinnerOnAntiDiagonal() int {

	for i := 1; i < 3; i++ {
		if board.items[i][2 - i] != board.items[i - 1][2 - (i - 1)] {
			return -1
		}
	}

	if winner := board.items[2][0]; winner != ItemNone {
		return winner
	}

	return -1
}

func (board *Board) FindWinner() int {

	var winner int

	for i := 0; i < 3; i++ {
		winner = board.findWinnerOnHorizontal(i)
		if winner != -1 { return winner }
	}

	for i := 0; i < 3; i++ {
		winner = board.findWinnerOnVertical(i)
		if winner != -1 { return winner }
	}

	winner = board.findWinnerOnDiagonal()
	if winner != -1 { return winner }

	winner = board.findWinnerOnAntiDiagonal()
	if winner != -1 { return winner }

	return -1
}

func (board *Board) IsGameOver() bool {
	return board.FindWinner() != -1
}