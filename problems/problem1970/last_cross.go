package problem1970

/*
There is a 1-based binary matrix where 0 represents land and 1 represents water.
You are given integers row and col representing the number of rows and columns in the matrix, respectively.
Initially on day 0, the entire matrix is land. However, each day a new cell becomes flooded with water.
You are given a 1-based 2D array cells, where cells[i] = [ri, ci] represents that
on the ith day, the cell on the rith row and cith column (1-based coordinates) will be covered with water (i.e., changed to 1).
You want to find the last day that it is possible to walk from the top to the bottom by only walking on land cells.
You can start from any cell in the top row and end at any cell in the bottom row.
You can only travel in the four cardinal directions (left, right, up, and down).
Return the last day where it is possible to walk from the top to the bottom by only walking on land cells.
*/

type Coord struct {
	X, Y int
}

var Start = Coord{-1, -1}
var End = Coord{-2, -2}

var Moves = [4][2]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func latestDayToCross(row, col int, cells [][]int) int {
	var res int
	// swamp[x][y] is 1 if the cells is land, and 0 if water
	var swamp = make([][]int, row)
	// Start with everything as water
	for i := range swamp {
		swamp[i] = make([]int, col)
	}

	// groups[coord] is the parent of coord in the disjoint set
	var groups = map[Coord]Coord{Start: Start, End: End}
	// Add the start and end to the group
	for i := 0; i < col; i++ {
		groups[Coord{0, i}] = Start
		groups[Coord{row - 1, i}] = End
	}

	// Start adding land (instead of removing water)
	for res = len(cells) - 1; res >= 0; res-- {
		cx, cy := cells[res][0]-1, cells[res][1]-1
		// Mark current cells as land
		swamp[cx][cy] = 1
		for _, mov := range Moves {
			// Try to connect all adjacent land tiles
			nx, ny := cx+mov[0], cy+mov[1]
			if isInside(row, col, nx, ny) && swamp[nx][ny] == 1 {
				union(groups, Coord{cx, cy}, Coord{nx, ny})
			}
		}
		// If the parent of start is the parent of end, we are connected
		if find(groups, Start) == find(groups, End) {
			return res
		}
	}
	return res
}

func isInside(row, col, x, y int) bool {
	return x >= 0 && x < row && y >= 0 && y < col
}

func find(uf map[Coord]Coord, x Coord) Coord {
	if uf[x] == x {
		// If x is the parent of itself, it is the root of the group
		return uf[x]
	} else {
		// If x is not the parent of itself, we call this function again
		// to find the real parent, and update the map
		uf[x] = find(uf, uf[x])
		return uf[x]
	}
}

func union(uf map[Coord]Coord, x, y Coord) {
	var rootx, rooty Coord
	if _, found := uf[x]; !found {
		// If this is the first time seeing x, set it as the root of it's group
		uf[x] = x
	}
	if _, found := uf[y]; !found {
		// If this is the first time seeing y, set it as the root of it's group
		uf[y] = y
	}
	// Finding the roots of x and y
	rootx = find(uf, x)
	rooty = find(uf, y)
	// Setting the root of rootx be rooty effectivly merging the groups
	uf[rootx] = rooty
}
