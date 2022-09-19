package problem0539

import (
	"sort"
	"strconv"
)

/*
Given a list of 24-hour clock time points in "HH:MM" format, return the minimum minutes difference between any two time-points in the list.
*/

func findMinDifference(timePoints []string) int {
	var res int = 1440
	var cur, prev int
	sort.Slice(timePoints, func(i, j int) bool {
		return timePoints[i] < timePoints[j]
	})
	prev = getMinutes(timePoints[len(timePoints)-1]) - 1440
	for i := 0; i < len(timePoints); i++ {
		cur = getMinutes(timePoints[i])
		if diff := abs(cur, prev); diff < res {
			res = diff
		}
		prev = cur
	}
	return res
}

func getMinutes(input string) int {
	hours, _ := strconv.Atoi(input[:2])
	minutes, _ := strconv.Atoi(input[3:])

	return hours*60 + minutes
}

func abs(a, b int) int {
	var res int = a - b
	if res < 0 {
		return res * -1
	}
	return res
}
