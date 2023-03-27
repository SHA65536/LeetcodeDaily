package problem0064

import "math"

/*
Given a m x n grid filled with non-negative numbers,
find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
Note: You can only move either down or right at any point in time.
*/

func minPathSum(grid [][]int) int {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j == len(grid[0])-1 {
				continue
			}
			var r, d int = math.MaxInt, math.MaxInt
			if i+1 < len(grid) {
				r = grid[i+1][j]
			}
			if j+1 < len(grid[0]) {
				d = grid[i][j+1]
			}
			grid[i][j] += min(r, d)
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
