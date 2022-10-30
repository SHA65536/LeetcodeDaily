package problem1293

/*
You are given an m x n integer matrix grid where each cell is either 0 (empty) or 1 (obstacle).
You can move up, down, left, or right from and to an empty cell in one step.
Return the minimum number of steps to walk from the upper left corner (0, 0) to the lower right corner (m - 1, n - 1)
given that you can eliminate at most k obstacles.
If it is not possible to find such walk return -1.
*/

func shortestPath(grid [][]int, k int) int {
	// Queue containing currently explored paths
	var queue = make(chan [4]int, 8192)
	// Grid containing visited squares to prevent repetition
	var visited = make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}
	// [4]{currentX, currentY, currentLength, removalsLeft}
	queue <- [4]int{0, 0, 0, k}
	for len(queue) > 0 {
		cur := <-queue
		// If current path is outside of bounds, skip
		if cur[0] < 0 || cur[1] < 0 || cur[0] >= len(grid) || cur[1] >= len(grid[0]) {
			continue
		}
		// If arrived at the end, return
		if cur[0] == len(grid)-1 && cur[1] == len(grid[0])-1 {
			return cur[2]
		}
		// If we see an obstacle, try to go through if there are enough
		// removals left
		if grid[cur[0]][cur[1]] == 1 {
			if cur[3] > 0 {
				cur[3]--
			} else {
				continue
			}
		}
		// If already visisted with more removals, skip
		if visited[cur[0]][cur[1]] != -1 && visited[cur[0]][cur[1]] >= cur[3] {
			continue
		}
		visited[cur[0]][cur[1]] = cur[3]
		// Trying all other possible moves
		queue <- [4]int{cur[0] + 1, cur[1], cur[2] + 1, cur[3]}
		queue <- [4]int{cur[0], cur[1] + 1, cur[2] + 1, cur[3]}
		queue <- [4]int{cur[0] - 1, cur[1], cur[2] + 1, cur[3]}
		queue <- [4]int{cur[0], cur[1] - 1, cur[2] + 1, cur[3]}
	}
	return -1
}
