package problem0576

/*
https://leetcode.com/problems/out-of-boundary-paths/

There is an m x n grid with a ball.
The ball is initially at the position [startRow, startColumn].
You are allowed to move the ball to one of the four adjacent cells in the grid (possibly out of the grid crossing the grid boundary).
You can apply at most maxMove moves to the ball.
Given the five integers m, n, maxMove, startRow, startColumn, return the number of paths to move the ball out of the grid boundary.
Since the answer can be very large, return it modulo 10^9 + 7.
*/

var paths = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

const maxres = 1000000007

func findPathsNaive(m int, n int, maxMove int, startRow int, startColumn int) int {
	var res int
	if maxMove <= 0 {
		return 0
	}
	for _, path := range paths {
		var newR, newC = startRow + path[0], startColumn + path[1]
		if newR > m-1 || newR < 0 || newC > n-1 || newC < 0 {
			res += 1
		} else {
			res += findPathsNaive(m, n, maxMove-1, newR, newC)
		}
	}
	return res % maxres
}

func findPathsCache(m int, n int, maxMove int, startRow int, startColumn int) int {
	var cache = map[[3]int]int{}
	return findPathsCacheSolve(cache, m, n, maxMove, startRow, startColumn)
}

func findPathsCacheSolve(cache map[[3]int]int, m, n, move, row, col int) int {
	var res int
	if move <= 0 {
		return 0
	}
	if val, ok := cache[[3]int{row, col, move}]; ok {
		return val
	}
	for _, path := range paths {
		var newR, newC = row + path[0], col + path[1]
		if newR > m-1 || newR < 0 || newC > n-1 || newC < 0 {
			res += 1
		} else {
			res += findPathsCacheSolve(cache, m, n, move-1, newR, newC)
		}
	}
	cache[[3]int{row, col, move}] = res % maxres
	return res % maxres
}
