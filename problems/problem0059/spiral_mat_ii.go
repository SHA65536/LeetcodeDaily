package problem0059

/*
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
*/

var dirs = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func generateMatrix(n int) [][]int {
	var res = make([][]int, n)
	var row, col, dir int
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 1; i <= n*n; i++ {
		res[row][col] = i
		nrow, ncol := row+dirs[dir][0], col+dirs[dir][1]
		if outOfBounds(res, nrow, ncol) || res[nrow][ncol] != 0 {
			dir = (dir + 1) % 4
			nrow, ncol = row+dirs[dir][0], col+dirs[dir][1]
		}
		row, col = nrow, ncol
	}
	return res
}

func outOfBounds(mat [][]int, row, col int) bool {
	return row < 0 || col < 0 || row >= len(mat) || col >= len(mat[0])
}
