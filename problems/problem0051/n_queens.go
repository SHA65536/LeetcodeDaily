package problem0051

/*
https://leetcode.com/problems/n-queens/

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
*/

func solveNQueens(n int) [][]string {
	var results [][]string
	var board = make([]string, n)
	for idx := range board {
		for i := 0; i < n; i++ {
			board[idx] += "."
		}
	}
	solveQueens(0, n, board, make([]int, n), make([]int, n), make([]int, n), &results)
	return results
}

func solveQueens(cur, max int, board []string, cols, upDiags, downDiags []int, results *[][]string) {
	if cur == max {
		// If we reached the final stage this means we have a valid
		// solution, so we copy the board to the results
		final := make([]string, len(board))
		copy(final, board)
		*results = append(*results, final)
		return
	}
	// Iterating over the possible places in the cur row
	for i := 0; i < max; i++ {
		var valid = true
		// Checking if board[cur][i] is a valid position

		// Checking collumns
		for j := 0; j < cur; j++ {
			if i == cols[j] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		// Checking upward diagonals
		for j := 0; j < cur; j++ {
			if cur-i == upDiags[j] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		// Checking downward diagonals
		for j := 0; j < cur; j++ {
			if cur+i == downDiags[j] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		// If we reached here, the space is valid
		// We place a queen
		board[cur] = board[cur][:i] + "Q" + board[cur][i+1:]
		// We place the threatened cols and diags
		cols[cur] = i
		upDiags[cur] = cur - i
		downDiags[cur] = cur + i
		solveQueens(cur+1, max, board, cols, upDiags, downDiags, results)
		// We remove the queen we place earlier
		board[cur] = board[cur][:i] + "." + board[cur][i+1:]
	}
}
