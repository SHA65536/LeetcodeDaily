package problem2389

import "sort"

/*
You are given an integer array nums of length n, and an integer array queries of length m.
Return an array answer of length m where answer[i] isthe maximum size of a subsequence that
you can take from nums such that the sum of its elements is less than or equal to queries[i].
A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
*/

func answerQueries(nums []int, queries []int) []int {
	var res = make([]int, len(queries))
	// Sorting by value
	sort.Ints(nums)
	// Making prefix sum
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	// Finding the biggest num not over
	for i, num := range queries {
		idx := sort.SearchInts(nums, num)
		if idx < len(nums) && nums[idx] == num {
			// If found exactly
			res[i] = idx + 1
		} else {
			// If not found get one to the left
			res[i] = idx
		}
	}
	return res
}
