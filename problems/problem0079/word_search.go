package problem0079

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.
The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.
*/

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			// Checking if the word starts at each square
			if findAt(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// findAt returns true if the word can be found at i,j starting with the l-th letter
func findAt(board [][]byte, i, j int, word string, l int) bool {
	var result bool
	if len(word) == l {
		// If we reached the end of the word
		return true
	}
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[l] {
		// If coords are outside the board or if current letter isn't in the word
		return false
	}
	// Marking the current cell as visited
	board[i][j] = 0
	// Check adjacent squares
	result = result || findAt(board, i+1, j, word, l+1)
	result = result || findAt(board, i, j+1, word, l+1)
	result = result || findAt(board, i-1, j, word, l+1)
	result = result || findAt(board, i, j-1, word, l+1)
	// Unmark the cell
	board[i][j] = word[l]
	return result
}
