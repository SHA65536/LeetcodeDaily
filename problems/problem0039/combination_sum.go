package problem0039

import "sort"

/*
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.
The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.
The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var sum func(target, iter int)

	sum = func(target, start int) {
		// If we missed the target
		if target < 0 {
			return
		}
		// If we get to the target
		if target == 0 {
			// Copy the current numbers to the result
			res = append(res, make([]int, len(cur)))
			copy(res[len(res)-1], cur)
			return
		}

		// Loop over candidates
		for i := start; i < len(candidates); i++ {
			// Add current candidate to the numbers
			cur = append(cur, candidates[i])
			// Recurse with updated sum
			sum(target-candidates[i], i)
			// Remove previous candidate from numbers
			cur = cur[:len(cur)-1]
		}
	}

	sort.Ints(candidates)

	sum(target, 0)

	return res
}
