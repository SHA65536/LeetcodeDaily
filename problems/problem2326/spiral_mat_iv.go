package problem2326

import . "leetcodedaily/helpers/listnode"

/*
You are given two integers m and n, which represent the dimensions of a matrix.
You are also given the head of a linked list of integers.
Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise), starting from the top-left of the matrix.
If there are remaining empty spaces, fill them with -1.
Return the generated matrix.
*/

var dirs = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	var res = make([][]int, m)
	var row, col, dir int
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	for i := 0; head != nil && i < m*n; i++ {
		res[row][col] = head.Val
		head = head.Next
		nrow, ncol := row+dirs[dir][0], col+dirs[dir][1]
		if outOfBounds(res, nrow, ncol) || res[nrow][ncol] != -1 {
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
