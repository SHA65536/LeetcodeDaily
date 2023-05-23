package problem1391

/*
You are given an m x n grid. Each cell of grid represents a street. The street of grid[i][j] can be:
    1 which means a street connecting the left cell and the right cell.
    2 which means a street connecting the upper cell and the lower cell.
    3 which means a street connecting the left cell and the lower cell.
    4 which means a street connecting the right cell and the lower cell.
    5 which means a street connecting the left cell and the upper cell.
    6 which means a street connecting the right cell and the upper cell.
You will initially start at the street of the upper-left cell (0, 0).
A valid path in the grid is a path that starts from the upper left cell (0, 0) and
ends at the bottom-right cell (m - 1, n - 1). The path should only follow the streets.
Notice that you are not allowed to change any street.
Return true if there is a valid path in the grid or false otherwise.
*/

const (
	Up = iota
	Down
	Right
	Left
)

var roadDirs = [6][4]int{
	{-1, -1, Right, Left},
	{Up, Down, -1, -1},
	{Left, -1, Down, -1},
	{Right, -1, -1, Down},
	{-1, Left, Up, -1},
	{-1, Right, -1, Up},
}

var dirDelta = [4][2]int{
	{-1, 0}, {1, 0}, {0, 1}, {0, -1},
}

func hasValidPath(grid [][]int) bool {
	// Destination
	var dstX, dstY = len(grid) - 1, len(grid[0]) - 1

	// Edge case
	if dstX == 0 && dstY == 0 {
		return true
	}

	// Already visited tiles
	var seen = make([][]bool, len(grid))

	// Inside check
	var isInside = func(x, y int) bool {
		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
	}

	// Different starting directions
	for i := Up; i <= Left; i++ {
		// Starting position
		var curX, curY int
		var curDir = i

		// Restting seen
		for i := range seen {
			seen[i] = make([]bool, len(grid[i]))
		}

		// Loop until found
		for isInside(curX, curY) && !seen[curX][curY] {
			// Marking seen
			seen[curX][curY] = true

			// Getting next direction
			curDir = roadDirs[grid[curX][curY]-1][curDir]

			// Invalid direction
			if curDir == -1 {
				break
			}

			// Found target
			if curX == dstX && curY == dstY {
				return true
			}

			// Moving
			curX += dirDelta[curDir][0]
			curY += dirDelta[curDir][1]
		}

	}

	return false
}
