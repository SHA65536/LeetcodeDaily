package problem0040

import "sort"

/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.
*/

func combinationSum2(candidates []int, target int) [][]int {
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
			// Skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// Add current candidate to the numbers
			cur = append(cur, candidates[i])
			// Recurse with updated sum
			sum(target-candidates[i], i+1)
			// Remove previous candidate from numbers
			cur = cur[:len(cur)-1]
		}
	}

	sort.Ints(candidates)

	sum(target, 0)

	return res
}
