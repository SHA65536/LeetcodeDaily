package problem1351

/*
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.
*/

func countNegatives(grid [][]int) int {
	var res, col, row int
	// Start from the bottom left and go up
	for row = len(grid) - 1; row >= 0; row-- {
		// Move right until you find a negative number
		for col < len(grid[0]) && grid[row][col] >= 0 {
			col++
		}
		// All the numbers left to the col are negative
		res += len(grid[0]) - col
	}
	return res
}
