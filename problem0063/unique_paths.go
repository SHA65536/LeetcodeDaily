package problem0063

/*
https://leetcode.com/problems/unique-paths-ii/

You are given an m x n integer array grid.
There is a robot initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m-1][n-1]).
The robot can only move either down or right at any point in time.

An obstacle and space are marked as 1 or 0 respectively in grid.
A path that the robot takes cannot include any square that is an obstacle.

Return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The testcases are generated so that the answer will be less than or equal to 2 * 10^9.
*/

type Point struct {
	X int
	Y int
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var visited = map[Point]int{{len(obstacleGrid) - 1, len(obstacleGrid[0]) - 1}: 1}
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	return search(obstacleGrid, visited, Point{0, 0})
}

func search(grid [][]int, visited map[Point]int, cur Point) int {
	var res int
	if val, ok := visited[cur]; ok {
		return val
	}
	right := Point{cur.X, cur.Y + 1}
	if right.X < len(grid) && right.Y < len(grid[0]) && grid[right.X][right.Y] == 0 {
		res += search(grid, visited, right)
	}
	down := Point{cur.X + 1, cur.Y}
	if down.X < len(grid) && down.Y < len(grid[0]) && grid[down.X][down.Y] == 0 {
		res += search(grid, visited, down)
	}
	visited[cur] = res
	return res
}
