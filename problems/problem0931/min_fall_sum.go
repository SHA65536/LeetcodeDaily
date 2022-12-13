package problem0931

/*
Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right.
Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
*/

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}
	// Loop from the bottom up except the last line that stays the same
	for i := len(matrix) - 2; i >= 0; i-- {
		for j := range matrix[i] {
			if j == 0 {
				// Left side can go right and down
				matrix[i][j] += min(matrix[i+1][j], matrix[i+1][j+1])
			} else if j == len(matrix[i])-1 {
				// Right side can go left and down
				matrix[i][j] += min(matrix[i+1][j-1], matrix[i+1][j])
			} else {
				// Everywhere else can go left, right and down
				matrix[i][j] += min(matrix[i+1][j-1], matrix[i+1][j], matrix[i+1][j+1])
			}
		}
	}
	return min(matrix[0]...)
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}
