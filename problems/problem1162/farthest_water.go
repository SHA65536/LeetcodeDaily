package problem1162

/*
Given an n x n grid containing only values 0 and 1, where 0 represents water and 1 represents land,
find a water cell such that its distance to the nearest land cell is maximized, and return the distance.
If no land or water exists in the grid, return -1.
The distance used in this problem is the Manhattan distance: the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.
*/

type Pos struct {
	X, Y int
}

// Possible directions
var dirs = [4][2]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

// maxDistance uses BFS for find the cell furthest from land
func maxDistance(grid [][]int) int {
	// res is the maximum depth of the search
	var res int
	var cur, nxt []Pos
	// Finding all the land cells
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				cur = append(cur, Pos{i, j})
			}
		}
	}
	// If the grid is all land or all water, return -1
	if len(cur) == 0 || len(cur) == len(grid)*len(grid[0]) {
		return -1
	}
	// While we have 'beach' cells not visited yet
	for len(cur) > 0 {
		res++
		nxt = []Pos{}
		// For each beach cells
		for _, p := range cur {
			// Check each direction
			for _, off := range dirs {
				npx, npy := p.X+off[0], p.Y+off[1]
				// If new cell is water, add to the list
				if npx >= 0 && npy >= 0 && npx < len(grid) && npy < len(grid[0]) && grid[npx][npy] == 0 {
					nxt = append(nxt, Pos{npx, npy})
					// And fill it with land
					grid[npx][npy] = 1
				}
			}
		}
		cur = nxt
	}
	return res - 1
}
