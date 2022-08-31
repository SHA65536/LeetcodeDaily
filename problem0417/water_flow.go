package problem0417

/*
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean.
The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.
The island is partitioned into a grid of square cells.
You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west
if the neighboring cell's height is less than or equal to the current cell's height.
Water can flow from any cell adjacent to an ocean into the ocean.
Return a 2D list of grid coordinates result
where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
*/

const (
	empty = iota
	pacific
	atlantic
	both
)

func pacificAtlantic(heights [][]int) [][]int {
	var flows = make([][]int, len(heights))
	var results = make([][]int, 0)
	for i := range heights {
		flows[i] = make([]int, len(heights[0]))
	}
	for i := range heights {
		for j := range heights[i] {
			if i == 0 || j == 0 {
				// Propogate "reverse flow" from cells touching pacific
				checkFlow(flows, heights, i, j, pacific, 0)
			}
			if i+1 == len(heights) || j+1 == len(heights[0]) {
				// Propogate "reverse flow" from cells touching atlantic
				checkFlow(flows, heights, i, j, atlantic, 0)
			}
		}
	}
	for i := range heights {
		for j := range heights[i] {
			if flows[i][j] == both {
				results = append(results, []int{i, j})
			}
		}
	}
	return results
}

func checkFlow(flows, heights [][]int, x, y int, side, prev int) {
	// If we're out of bounds, exit
	if x < 0 || x >= len(heights) || y < 0 || y >= len(heights[0]) {
		return
	}
	// If we're going downhill, exit
	if heights[x][y] < prev {
		return
	}
	// If we already visited, exit
	if flows[x][y]&side == side {
		return
	}
	flows[x][y] |= side
	checkFlow(flows, heights, x+1, y, side, heights[x][y])
	checkFlow(flows, heights, x, y+1, side, heights[x][y])
	checkFlow(flows, heights, x-1, y, side, heights[x][y])
	checkFlow(flows, heights, x, y-1, side, heights[x][y])
}
