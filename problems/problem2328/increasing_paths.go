package problem2328

import "math"

/*
You are given an m x n integer matrix grid, where you can move from a cell to any adjacent cell in all 4 directions.
Return the number of strictly increasing paths in the grid such that you can start from any cell and end at any cell.
Since the answer may be very large, return it modulo 109 + 7.
Two paths are considered different if they do not have exactly the same sequence of visited cells.
*/
const mod int = 1e9 + 7

func countPaths(grid [][]int) int {
	var res int

	// Cache[x][y] is the number of strictly increasing paths
	// that start at grid[x][y]
	var cache = make([][]int, len(grid))
	for i := range cache {
		cache[i] = make([]int, len(grid[0]))
	}

	// Calculate number of paths starting at each point
	for i := range grid {
		for j := range grid[i] {
			res = (res + getPathsFrom(i, j, math.MinInt, grid, cache)) % mod
		}
	}
	return res
}

// getPathsFrom returns the number of strictly increasing paths starting at grid[x][y]
// p is the value of the previous tile, it's used to make sure we're strictly increasing
func getPathsFrom(x, y, p int, grid, cache [][]int) int {
	var res int = 1
	// Check out of bounds and strictly increasing
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 || grid[x][y] <= p {
		return 0
	}
	// If already calculated, return from cache
	if cache[x][y] != 0 {
		return cache[x][y]
	}

	// Add possible paths for each side
	res = (res + getPathsFrom(x+1, y, grid[x][y], grid, cache)) % mod
	res = (res + getPathsFrom(x, y+1, grid[x][y], grid, cache)) % mod
	res = (res + getPathsFrom(x-1, y, grid[x][y], grid, cache)) % mod
	res = (res + getPathsFrom(x, y-1, grid[x][y], grid, cache)) % mod

	// Save in cache
	cache[x][y] = res
	return res
}
