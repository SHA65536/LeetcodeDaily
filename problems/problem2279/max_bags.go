package problem2279

import "sort"

/*
You have n bags numbered from 0 to n - 1.
You are given two 0-indexed integer arrays capacity and rocks. The ith bag can hold a maximum of capacity[i] rocks and currently contains rocks[i] rocks.
You are also given an integer additionalRocks, the number of additional rocks you can place in any of the bags.
Return the maximum number of bags that could have full capacity after placing the additional rocks in some bags.
*/

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	var res int
	// Number of rocks needed to fill
	var needed []int = make([]int, len(capacity))
	for i := range needed {
		needed[i] = capacity[i] - rocks[i]
	}
	// Sorting from least needed to most
	sort.Ints(needed)
	for _, n := range needed {
		if n == 0 {
			res++
			continue
		}
		if n <= additionalRocks {
			res++
			additionalRocks -= n
		} else {
			break
		}
	}
	return res
}
