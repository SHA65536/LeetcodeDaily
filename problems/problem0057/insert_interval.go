package problem0057

/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval
and intervals is sorted in ascending order by starti.
You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
Insert newInterval into intervals such that intervals is still sorted in ascending order by start i and intervals
still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	var i int

	// All the unrelated intervals at the start of the list
	for i = 0; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	// Merging all overlapping intervals into newInterval
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	// Adding the merged newInterval
	res = append(res, newInterval)

	// Adding all the unrelated intervals at the end of the list
	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
