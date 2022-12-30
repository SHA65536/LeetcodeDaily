package problem0054

/*
Given an m x n matrix, return all elements of the matrix in spiral order.
*/

const visited int = -101

var dirs = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	var end = len(matrix)*len(matrix[0]) - 1
	var res = make([]int, end+1)
	var col, row, dir int
	for cur := 0; cur <= end; cur++ {
		res[cur] = matrix[row][col]
		matrix[row][col] = visited
		nrow, ncol := row+dirs[dir][0], col+dirs[dir][1]
		// Switching direction
		if outOfBounds(matrix, nrow, ncol) || matrix[nrow][ncol] == visited {
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
