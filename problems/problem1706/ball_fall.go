package problem1706

/*
You have a 2-D grid of size m x n representing a box, and you have n balls. The box is open on the top and bottom sides.
Each cell in the box has a diagonal board spanning two corners of the cell that can redirect a ball to the right or to the left.
A board that redirects the ball to the right spans the top-left corner to the bottom-right corner and is represented in the grid as 1.
A board that redirects the ball to the left spans the top-right corner to the bottom-left corner and is represented in the grid as -1.
We drop one ball at the top of each column of the box. Each ball can get stuck in the box or fall out of the bottom.
A ball gets stuck if it hits a "V" shaped pattern between two boards or if a board redirects the ball into either wall of the box.
Return an array answer of size n where answer[i] is the column that the ball falls out of at the bottom after dropping the ball from the ith column at the top, or -1 if the ball gets stuck in the box.
*/

func findBall(grid [][]int) []int {
	var res = make([]int, len(grid[0]))
	for i := range res {
		px, py := -1, i
		x, y := 0, i
		for x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			if px < x { // Coming from the top
				px, py = x, y
				if grid[x][y] == 1 {
					y++
				} else {
					y--
				}
			} else if py > y { // Coming from right
				px, py = x, y
				if grid[x][y] == 1 {
					break
				} else {
					x++
				}
			} else { // Coming from left
				px, py = x, y
				if grid[x][y] == 1 {
					x++
				} else {
					break
				}
			}
		}
		if x == len(grid) {
			res[i] = y
		} else {
			res[i] = -1
		}
	}
	return res
}
