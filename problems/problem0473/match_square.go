package problem0473

import (
	"sort"
)

/*
https://leetcode.com/problems/matchsticks-to-square

You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick.
You want to use all the matchsticks to make one square.
You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.
Return true if you can make this square and false otherwise.
*/

func makesquare(matchsticks []int) bool {
	var length int
	var sides = make([]int, 4)
	for _, val := range matchsticks {
		length += val
	}
	// Square perimeter should divide by 4, and we need more than 4 sticks
	if length%4 != 0 || len(matchsticks) < 4 {
		return false
	}
	// Sorting sticks descending for optimization
	sort.Slice(matchsticks, func(i, j int) bool { return matchsticks[i] > matchsticks[j] })
	return solveDfs(sides, matchsticks, 0, length/4)
}

func solveDfs(sides, m []int, index, target int) bool {
	// If we used all the sticks, check if the sides are equal
	if index == len(m) {
		return sides[0] == sides[1] && sides[1] == sides[2] && sides[2] == sides[3]
	}
	// Attempting adding the matchstick to all sides
Sides:
	for i := range sides {
		// If the current side will be more than the target
		// No need to continue this path
		// This is also why we sorted the sticks descending
		if sides[i]+m[index] > target {
			continue
		}
		// Checking if we've already calculated a side of this length
		for j := i - 1; j >= 0; j-- {
			if sides[i] == sides[j] {
				continue Sides
			}
		}
		sides[i] += m[index]
		if solveDfs(sides, m, index+1, target) {
			return true
		}
		sides[i] -= m[index]
	}
	return false
}
