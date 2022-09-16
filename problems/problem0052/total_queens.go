package problem0052

/*
https://leetcode.com/problems/n-queens-ii/

The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.
*/

func totalNQueens(n int) int {
	var results int
	solveQueens(0, n, make([]int, n), make([]int, n), make([]int, n), &results)
	return results
}

func solveQueens(cur, max int, cols, upDiags, downDiags []int, results *int) {
	if cur == max {
		// If we reached the final stage this means we have a valid
		// solution, so we copy the board to the results
		*results++
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
		// We place the threatened cols and diags
		cols[cur] = i
		upDiags[cur] = cur - i
		downDiags[cur] = cur + i
		solveQueens(cur+1, max, cols, upDiags, downDiags, results)
	}
}
