package problem0934

/*
You are given an n x n binary matrix grid where 1 represents land and 0 represents water.
An island is a 4-directionally connected group of 1's not connected to any other 1's.
There are exactly two islands in grid.
You may change 0's to 1's to connect the two islands to form one island.
Return the smallest number of 0's you must flip to connect the two islands.
*/

const (
	water = iota
	land
	srcIsland
	dstIsland
)

type Tile struct {
	x, y int
}

func shortestBridge(grid [][]int) int {
	var res int
	var cur, nxt []Tile

	// Finding both island
	cur = findIslands(grid)

	// BFS
	for len(cur) > 0 {
		nxt = nxt[:0]
		for _, c := range cur {
			// Skip out of bounds
			if !isInside(grid, c.x, c.y) {
				continue
			}
			// If we found the target, this must be the shortest
			if grid[c.x][c.y] == dstIsland {
				return res
			}
			// If we didn't visit this,
			if grid[c.x][c.y] == water {
				// Mark as visited
				grid[c.x][c.y] = srcIsland
				// Expands
				nxt = append(nxt,
					Tile{c.x + 1, c.y}, // Down
					Tile{c.x, c.y + 1}, // Right
					Tile{c.x - 1, c.y}, // Up
					Tile{c.x, c.y - 1}) // Left
			}
		}
		nxt, cur = cur, nxt
		res++
	}

	return res
}

// Find Islands marks the two islands with
// srcIsland and dstIsland
func findIslands(grid [][]int) []Tile {
	var res []Tile
	var cur = srcIsland
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if isInside(grid, x, y) && grid[x][y] == land {
			if cur == srcIsland {
				res = append(res, Tile{x + 1, y}, Tile{x, y + 1}, Tile{x - 1, y}, Tile{x, y - 1})
			}
			grid[x][y] = cur
			dfs(x+1, y)
			dfs(x, y+1)
			dfs(x-1, y)
			dfs(x, y-1)
			return true
		}
		return false
	}
	for i := range grid {
		for j := range grid[i] {
			if dfs(i, j) {
				cur = dstIsland
			}
		}
	}
	return res
}

func isInside(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
