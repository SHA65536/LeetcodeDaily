package problem0977

/*
https://leetcode.com/problems/squares-of-a-sorted-array

Given an integer array nums sorted in non-decreasing order,
return an array of the squares of each number sorted in non-decreasing order.
*/

func sortedSquares(nums []int) []int {
	var start, end, cur int
	end = len(nums) - 1
	cur = end
	var res = make([]int, end+1)
	for end >= start {
		if square(nums[start]) > square(nums[end]) {
			res[cur] = square(nums[start])
			start++
		} else {
			res[cur] = square(nums[end])
			end--
		}
		cur--
	}
	return res
}

func square(n int) int { return n * n }
