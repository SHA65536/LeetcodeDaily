package problem0867

/*
https://leetcode.com/problems/transpose-matrix/

Given a 2D integer array matrix, return the transpose of matrix.

The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
*/

func transpose(matrix [][]int) [][]int {
	var result = make([][]int, len(matrix[0]))
	for i := range result {
		result[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}
