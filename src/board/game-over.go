package board

func (board *Board) findWinnerOnHorizontal(line int) (winner int, indexes []int) {

	for i := 1; i < 3; i++ {
		if board.items[line][i] != board.items[line][i - 1] {
			return -1, nil
		}
	}

	if winner := board.items[line][0]; winner != ItemNone {
		return winner, []int{ line * 3, line * 3 + 1, line * 3 + 2 }
	}

	return -1, nil
}

func (board *Board) findWinnerOnVertical(column int) (winner int, indexes []int) {

	for i := 1; i < 3; i++ {
		if board.items[i][column] != board.items[i - 1][column] {
			return -1, nil
		}
	}

	if winner := board.items[0][column]; winner != ItemNone {
		return winner, []int{ column, column + 3, column + 6 }
	}

	return -1, nil
}

func (board *Board) findWinnerOnDiagonal() (winner int, indexes []int) {

	for i := 1; i < 3; i++ {
		if board.items[i][i] != board.items[i - 1][i - 1] {
			return -1, nil
		}
	}

	if winner := board.items[0][0]; winner != ItemNone {
		return winner, []int{ 0, 4, 8 } 
	}

	return -1, nil
}

func (board *Board) findWinnerOnAntiDiagonal() (winner int, indexes []int) {

	for i := 1; i < 3; i++ {
		if board.items[i][2 - i] != board.items[i - 1][2 - (i - 1)] {
			return -1, nil
		}
	}

	if winner := board.items[2][0]; winner != ItemNone {
		return winner, []int{ 2, 4, 6 } 
	}

	return -1, nil
}

func (board *Board) FindWinnerIndexes() []int {

	if board.FindWinner() == -1 {
		return nil
	}

	var winner int
	var indexes []int

	for i := 0; i < 3; i++ {
		if winner, indexes = board.findWinnerOnHorizontal(i); winner != -1 { 
			return indexes
		}
	}

	for i := 0; i < 3; i++ {
		if winner, indexes = board.findWinnerOnVertical(i); winner != -1 { 
			return indexes
		}
	}

	if winner, indexes = board.findWinnerOnDiagonal(); winner != -1 { 
		return indexes
	}

	if winner, indexes = board.findWinnerOnAntiDiagonal(); winner != -1 { 
		return indexes
	}

	return nil
}

func (board *Board) FindWinner() int {

	var winner int

	for i := 0; i < 3; i++ {
		winner, _ = board.findWinnerOnHorizontal(i)
		if winner != -1 { return winner }
	}

	for i := 0; i < 3; i++ {
		winner, _ = board.findWinnerOnVertical(i)
		if winner != -1 { return winner }
	}

	winner, _ = board.findWinnerOnDiagonal()
	if winner != -1 { return winner }

	winner, _ = board.findWinnerOnAntiDiagonal()
	if winner != -1 { return winner }

	return -1
}

func (board *Board) IsFull() bool {
	
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.items[i][j] == ItemNone {
				return false
			}
		}
	}

	return true
}

func (board *Board) IsGameOver() bool {
	return board.FindWinner() != -1 || board.IsFull()
}