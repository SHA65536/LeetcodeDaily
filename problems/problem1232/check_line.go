package problem1232

/*
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
Check if these points make a straight line in the XY plane.
*/

func checkStraightLine(dots [][]int) bool {
	// Get initial deltas
	var dx, dy = dots[1][0] - dots[0][0], dots[1][1] - dots[0][1]

	// Loop over all dots
	for i := 2; i < len(dots); i++ {
		// Comparing against initial gradient
		if dx*(dots[i][1]-dots[1][1]) != dy*(dots[i][0]-dots[1][0]) {
			return false
		}
	}

	return true
}
