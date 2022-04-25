package two_sum

/*
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	var seen = map[int]int{}
	for idx, num := range nums {
		compliment := target - num
		if compIdx, ok := seen[compliment]; ok {
			return []int{compIdx, idx}
		} else {
			seen[num] = idx
		}
	}
	return []int{}
}
