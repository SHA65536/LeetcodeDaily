package problem0885

/*
You start at the cell (rStart, cStart) of an rows x cols grid facing east.
The northwest corner is at the first row and column in the grid, and the southeast corner is at the last row and column.
You will walk in a clockwise spiral shape to visit every position in this grid.
Whenever you move outside the grid's boundary, we continue our walk outside the grid (but may return to the grid boundary later).
Eventually, we reach all rows * cols spaces of the grid.
Return an array of coordinates representing the positions of the grid in the order you visited them.
*/

var dirs = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func spiralMatrixIII(rows int, cols int, row int, col int) [][]int {
	var res = make([][]int, rows*cols)
	var cnt, dir, cur int
	cur = 1
	for cnt < rows*cols {
		// Move one
		for i := 0; i < cur; i++ {
			if !outOfBounds(rows, cols, row, col) {
				res[cnt] = []int{row, col}
				cnt++
			}
			row, col = row+dirs[dir][0], col+dirs[dir][1]
			if cnt >= rows*cols {
				return res
			}
		}
		// Switch dir
		dir = (dir + 1) % 4
		// Move Two
		for i := 0; i < cur; i++ {
			if !outOfBounds(rows, cols, row, col) {
				res[cnt] = []int{row, col}
				cnt++
			}
			row, col = row+dirs[dir][0], col+dirs[dir][1]
			if cnt >= rows*cols {
				return res
			}
		}
		// Switch dir
		dir = (dir + 1) % 4
		cur++
	}
	return res
}

func outOfBounds(rm, cm, row, col int) bool {
	return row < 0 || col < 0 || row >= rm || col >= cm
}
