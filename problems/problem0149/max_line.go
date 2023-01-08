package problem0149

/*
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane,
return the maximum number of points that lie on the same straight line.
*/

const max_float float64 = 100001

func maxPoints(points [][]int) int {
	var res int
	// For each starting point
	for i := range points {
		var same int = 1
		// sMap[slope] represents how many points are on a line starting at the
		// current point with slope of s
		var sMap = map[float64]int{}
		// For each ending point that we didn't start from
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				// If it's the same point, add to the counter
				same++
			} else if points[i][0] == points[j][0] {
				// If it's same x but different y, the slope is max
				sMap[max_float]++
			} else {
				// Calculate slope and add 1 to the map at the slope
				slope := float64(points[i][1]-points[j][1]) / float64(points[i][0]-points[j][0])
				sMap[slope]++
			}
		}
		var localMax int
		// Find the slope with the most points on it
		for _, v := range sMap {
			localMax = max(localMax, v)
		}
		// Pick the max
		res = max(res, localMax+same)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
