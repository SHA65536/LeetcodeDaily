package problem1658

/*
https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

You are given an integer array nums and an integer x.
In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x.
Note that this modifies the array for future operations.

Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.
*/

func minOperations(nums []int, x int) int {
	var sum, cur, start, end int
	var maxLen = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for ; end < len(nums); end++ {
		cur += nums[end]
		for end >= start && cur > sum-x {
			cur -= nums[start]
			start++
		}
		if cur == sum-x {
			if maxLen < end-start+1 {
				maxLen = end - start + 1
			}
		}
	}
	if maxLen == -1 {
		return -1
	} else {
		return len(nums) - maxLen
	}
}
