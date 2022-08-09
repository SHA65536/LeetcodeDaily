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

const One byte = '1'

func isValidSudoku(board [][]byte) bool {
	// Checking rows and cols
	for r := 0; r < 9; r++ {
		row := make([]int, 9)
		col := make([]int, 9)
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				if row[board[r][c]-One] == 1 {
					return false
				}
				row[board[r][c]-One] = 1
			}
			if board[c][r] != '.' {
				if col[board[c][r]-One] == 1 {
					return false
				}
				col[board[c][r]-One] = 1
			}
		}
	}
	// Checking squares
	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			row := make([]int, 9)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[r+i][c+j] == '.' {
						continue
					}
					if row[board[r+i][c+j]-One] == 1 {
						return false
					}
					row[board[r+i][c+j]-One] = 1
				}
			}
		}
	}
	return true
}
