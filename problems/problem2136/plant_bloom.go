package problem2136

import (
	"sort"
)

/*
You have n flower seeds. Every seed must be planted first before it can begin to grow, then bloom.\
Planting a seed takes time and so does the growth of a seed. You are given two 0-indexed integer arrays plantTime and growTime, of length n each:
plantTime[i] is the number of full days it takes you to plant the ith seed.
Every day, you can work on planting exactly one seed.
You do not have to work on planting the same seed on consecutive days,
but the planting of a seed is not complete until you have worked plantTime[i] days on planting it in total.
growTime[i] is the number of full days it takes the ith seed to grow after being completely planted.
After the last day of its growth, the flower blooms and stays bloomed forever.
From the beginning of day 0, you can plant the seeds in any order.
Return the earliest possible day where all seeds are blooming.
*/

func earliestFullBloom(plantTime []int, growTime []int) int {
	var res int

	// plantOrder represents the order in which the plants grow
	var plantOrder = make([]int, len(plantTime))
	for i := range plantOrder {
		plantOrder[i] = i
	}
	// we sort them by grow time
	sort.Slice(plantOrder, func(i, j int) bool { return growTime[plantOrder[i]] < growTime[plantOrder[j]] })
	for _, p := range plantOrder {
		res = max(res, growTime[p]) + plantTime[p]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
