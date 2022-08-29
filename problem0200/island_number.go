package problem0200

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/

func numIslands(grid [][]byte) int {
	var numIslands int
	var found = make([][]bool, len(grid))
	for i := range grid {
		found[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !found[i][j] {
				numIslands++
				discover(grid, found, i, j)
			}
		}
	}
	return numIslands
}

func discover(grid [][]byte, found [][]bool, x, y int) {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
		return
	}
	if grid[x][y] == '0' || found[x][y] {
		return
	}
	found[x][y] = true
	discover(grid, found, x+1, y)
	discover(grid, found, x, y+1)
	discover(grid, found, x-1, y)
	discover(grid, found, x, y-1)
}
