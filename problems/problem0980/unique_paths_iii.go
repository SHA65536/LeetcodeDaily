package problem0980

/*
You are given an m x n integer array grid where grid[i][j] could be:
    1 representing the starting square. There is exactly one starting square.
    2 representing the ending square. There is exactly one ending square.
    0 representing empty squares we can walk over.
    -1 representing obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square,
that walk over every non-obstacle square exactly once.
*/

func uniquePathsIII(grid [][]int) int {
	var total, res int
	var row, col int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				total++
			}
			if grid[i][j] == 1 {
				row, col = i, j
			}
		}
	}
	pathFinder(grid, row, col, 0, total+1, &res)
	return res
}

func pathFinder(grid [][]int, row, col, curLen, maxLen int, res *int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return
	}
	if grid[row][col] == -1 {
		return
	}
	if grid[row][col] == 2 {
		if curLen == maxLen {
			*res++
		}
		return
	}
	curLen++
	grid[row][col] = -1
	pathFinder(grid, row+1, col, curLen, maxLen, res)
	pathFinder(grid, row, col+1, curLen, maxLen, res)
	pathFinder(grid, row-1, col, curLen, maxLen, res)
	pathFinder(grid, row, col-1, curLen, maxLen, res)
	grid[row][col] = 0
}
