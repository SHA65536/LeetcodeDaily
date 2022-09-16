package problem0695

/*
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.
*/

func maxAreaOfIsland(grid [][]int) int {
	var maxArea, curArea int
	var visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] {
				curArea = floodArea(grid, visited, i, j)
				if curArea > maxArea {
					maxArea = curArea
				}
			}
		}
	}
	return maxArea
}

func floodArea(grid [][]int, visited [][]bool, curX, curY int) int {
	var subArea int
	if curX < 0 || curY < 0 || curX >= len(grid) || curY >= len(grid[0]) {
		return subArea
	}
	if visited[curX][curY] {
		return subArea
	}
	visited[curX][curY] = true
	if grid[curX][curY] == 1 {
		subArea = 1
		subArea += floodArea(grid, visited, curX+1, curY)
		subArea += floodArea(grid, visited, curX-1, curY)
		subArea += floodArea(grid, visited, curX, curY+1)
		subArea += floodArea(grid, visited, curX, curY-1)
	}
	return subArea
}
