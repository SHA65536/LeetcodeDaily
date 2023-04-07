package problem1020

/*
You are given an m x n binary matrix grid, where 0 represents a sea cell and 1 represents a land cell.
A move consists of walking from one land cell to another adjacent (4-directionally) land cell or walking off the boundary of the grid.
Return the number of land cells in grid for which we cannot walk off the boundary of the grid in any number of moves.
*/

const (
	Land  = 1
	Water = 0
)

func numEnclaves(grid [][]int) int {
	var numIslands, cur int
	var found = make([][]bool, len(grid))
	for i := range grid {
		found[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			cur = 0
			if grid[i][j] == Land && !found[i][j] {
				if !discover(grid, found, i, j, &cur) {
					numIslands += cur
				}
			}
		}
	}
	return numIslands
}

func discover(grid [][]int, found [][]bool, x, y int, cur *int) bool {
	var res bool
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
		return false
	}
	if grid[x][y] == Water || found[x][y] {
		return false
	}
	found[x][y] = true
	*cur++
	res = discover(grid, found, x+1, y, cur) || res
	res = discover(grid, found, x, y+1, cur) || res
	res = discover(grid, found, x-1, y, cur) || res
	res = discover(grid, found, x, y-1, cur) || res
	return res || isRim(x, y, grid)
}

func isRim(i, j int, grid [][]int) bool {
	return i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1
}
