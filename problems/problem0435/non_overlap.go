package problem0435

import "sort"

/*
Given an array of intervals intervals where intervals[i] = [starti, endi],
return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
*/

func eraseOverlapIntervals(inter [][]int) int {
	sort.Slice(inter, func(i, j int) bool {
		return inter[i][1] < inter[j][1]
	})

	var curEnd = inter[0][1]
	var res int = 1

	for _, cur := range inter {
		if cur[0] >= curEnd {
			curEnd = cur[1]
			res++
		}
	}
	return len(inter) - res
}
