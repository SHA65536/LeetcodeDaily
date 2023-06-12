package problem0228

import "fmt"

/*
You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).
Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
Each range [a,b] in the list should be output as:
    "a->b" if a != b
    "a" if a == b
*/

func summaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return []string{}
	}
	var start int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if nums[i-1] == start {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
		}
	}
	if nums[len(nums)-1] != start {
		res = append(res, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
	} else {
		res = append(res, fmt.Sprintf("%d", start))
	}

	return res
}
