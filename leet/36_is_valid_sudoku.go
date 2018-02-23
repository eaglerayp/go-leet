package leet

const nullB = byte('.')

func isValidSudoku(board [][]byte) bool {
	// base solution, check every row, column and block for char once

	// create
	blockSet := [9]map[byte]bool{}
	columnSet := [9]map[byte]bool{}
	rowSet := [9]map[byte]bool{}
	for i := 0; i < 9; i++ {
		blockSet[i] = map[byte]bool{}
		columnSet[i] = map[byte]bool{}
		rowSet[i] = map[byte]bool{}
	}

	// board is 9*9
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			block := (row/3)*3 + column/3
			b := board[row][column]
			if b == nullB {
				continue
			}
			if columnSet[column][b] {
				return false
			}
			columnSet[column][b] = true
			if rowSet[row][b] {
				return false
			}
			rowSet[row][b] = true
			if blockSet[block][b] {
				return false
			}
			blockSet[block][b] = true
		}
	}
	return true
}
