package problem0036

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

func isValidSudoku(board [][]byte) bool {
	// Checking rows and cols
	for r := 0; r < 9; r++ {
		row := make([]bool, 9) // row[i] shows wether we have seen i in this row already
		col := make([]bool, 9) // col[i] shows wether we have seen i in this col already
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				if row[board[r][c]-'1'] { // If already seen
					return false // Invalid
				}
				row[board[r][c]-'1'] = true // Set to seen
			}
			if board[c][r] != '.' {
				if col[board[c][r]-'1'] { // If already seen
					return false // Invalid
				}
				col[board[c][r]-'1'] = true // Set to seen
			}
		}
	}
	// Checking squares
	for r := 0; r < 9; r += 3 { // r is the starting row of the inner square
		for c := 0; c < 9; c += 3 { // c is the starting col of the inner square
			row := make([]bool, 9)
			for i := 0; i < 3; i++ { // i is the row of the cell inside the square
				for j := 0; j < 3; j++ { // j is the col of the cell inside the square
					if board[r+i][c+j] == '.' {
						continue
					}
					if row[board[r+i][c+j]-'1'] { // If already seen
						return false // Invalid
					}
					row[board[r+i][c+j]-'1'] = true // Set to seen
				}
			}
		}
	}
	return true
}
