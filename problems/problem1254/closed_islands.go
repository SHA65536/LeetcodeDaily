package problem1254

/*
Given a 2D grid consists of 0s (land) and 1s (water).
An island is a maximal 4-directionally connected group of 0s and a closed island is an island totally (all left, top, right, bottom) surrounded by 1s.
Return the number of closed islands.
*/

const (
	Land  = 0
	Water = 1
)

func closedIsland(grid [][]int) int {
	var numIslands int
	var found = make([][]bool, len(grid))
	for i := range grid {
		found[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == Land && !found[i][j] {
				if !discover(grid, found, i, j) {
					numIslands++
				}
			}
		}
	}
	return numIslands
}

func discover(grid [][]int, found [][]bool, x, y int) bool {
	var res bool
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
		return false
	}
	if grid[x][y] == Water || found[x][y] {
		return false
	}
	found[x][y] = true
	res = discover(grid, found, x+1, y) || res
	res = discover(grid, found, x, y+1) || res
	res = discover(grid, found, x-1, y) || res
	res = discover(grid, found, x, y-1) || res
	return res || isRim(x, y, grid)
}

func isRim(i, j int, grid [][]int) bool {
	return i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1
}
