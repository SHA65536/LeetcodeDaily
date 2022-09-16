package problem0994

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
*/

func orangesRotting(grid [][]int) int {
	var count int
	var fresh bool
	var cur = [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case 1:
				fresh = true
			case 2:
				cur = append(cur, [2]int{i, j})
			}
		}
	}
	if !fresh {
		return 0
	}
	if fresh && len(cur) == 0 {
		return -1
	}
	for len(cur) > 0 {
		next := [][2]int{}
		count++
		for _, dot := range cur {
			if dot[0]+1 < len(grid) && grid[dot[0]+1][dot[1]] == 1 {
				next = append(next, [2]int{dot[0] + 1, dot[1]})
				grid[dot[0]+1][dot[1]] = 2
			}
			if dot[0]-1 >= 0 && grid[dot[0]-1][dot[1]] == 1 {
				next = append(next, [2]int{dot[0] - 1, dot[1]})
				grid[dot[0]-1][dot[1]] = 2
			}
			if dot[1]+1 < len(grid[0]) && grid[dot[0]][dot[1]+1] == 1 {
				next = append(next, [2]int{dot[0], dot[1] + 1})
				grid[dot[0]][dot[1]+1] = 2
			}
			if dot[1]-1 >= 0 && grid[dot[0]][dot[1]-1] == 1 {
				next = append(next, [2]int{dot[0], dot[1] - 1})
				grid[dot[0]][dot[1]-1] = 2
			}
		}
		cur = next
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count - 1
}
