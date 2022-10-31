package problem0766

/*
Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return false.
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same elements.
*/

func isToeplitzMatrix(matrix [][]int) bool {
	for i := len(matrix) - 1; i > 0; i-- {
		var x, y = i, 0
		var prev = matrix[x][y]
		for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
			if matrix[x][y] != prev {
				return false
			}
			prev = matrix[x][y]
			x++
			y++
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		var x, y = 0, i
		var prev = matrix[x][y]
		for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
			if matrix[x][y] != prev {
				return false
			}
			prev = matrix[x][y]
			x++
			y++
		}
	}
	return true
}
