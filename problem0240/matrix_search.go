package problem0240

/*
https://leetcode.com/problems/search-a-2d-matrix-ii/

Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
This matrix has the following properties:
Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
*/

func searchMatrix(matrix [][]int, target int) bool {
	var r, c int
	for c = len(matrix[0]) - 1; c >= 0 && r < len(matrix); {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			// if target is smaller than current position
			// then the whole collumn is invalid
			c--
		} else {
			// if target is bigger than current position
			// then the whole row is invalid
			r++
		}
	}
	return false
}
