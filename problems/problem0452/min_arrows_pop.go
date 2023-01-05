package problem0452

import "sort"

/*
There are some spherical balloons taped onto a flat wall that represents the XY-plane.
The balloons are represented as a 2D integer array points where
points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between xstart and xend.
You do not know the exact y-coordinates of the balloons.
Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis.
A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend.
There is no limit to the number of arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any balloons in its path.
Given the array points, return the minimum number of arrows that must be shot to burst all balloons.
*/

func findMinArrowShots(points [][]int) int {
	var res, prevEnd int
	// Sorting by start time and then by end time
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	// End of previous balloon
	prevEnd = points[0][1]
	// Looping over balloon
	for i := range points {
		if points[i][0] > prevEnd {
			// If the start of the current balloon is
			// after the end of the last balloon, it means there is no overlap
			res++
			prevEnd = points[i][1]
		} else if points[i][1] < prevEnd {
			// If there is an overlap, pick the leftmost end
			prevEnd = points[i][1]
		}
	}
	return res + 1
}
