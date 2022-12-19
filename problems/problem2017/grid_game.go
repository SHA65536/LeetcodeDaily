package problem2017

/*
You are given a 0-indexed 2D array grid of size 2 x n, where grid[r][c] represents the number of points at position (r, c) on the matrix.
Two robots are playing a game on this matrix. Both robots initially start at (0, 0) and want to reach (1, n-1).
Each robot may only move to the right ((r, c) to (r, c + 1)) or down ((r, c) to (r + 1, c)).
At the start of the game, the first robot moves from (0, 0) to (1, n-1), collecting all the points from the cells on its path.
For all cells (r, c) traversed on the path, grid[r][c] is set to 0.
Then, the second robot moves from (0, 0) to (1, n-1), collecting the points on its path. Note that their paths may intersect with one another.
The first robot wants to minimize the number of points collected by the second robot.
In contrast, the second robot wants to maximize the number of points it collects.
If both robots play optimally, return the number of points collected by the second robot.
*/

func gridGame(grid [][]int) int64 {
	var top, bot, res int64
	res = 9223372036854775807
	for i := range grid[0] {
		top += int64(grid[0][i])
	}
	for i := range grid[0] {
		top -= int64(grid[0][i])
		res = min(res, max(top, bot))
		bot += int64(grid[1][i])
	}
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
