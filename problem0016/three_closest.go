package problem0016

import "sort"

func threeSumClosest(nums []int, target int) int {
	var res int
	sort.Ints(nums)
	res = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-1; i++ {
		// Skipping duplicate starting numbers
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		start, end := i+1, len(nums)-1
		for end > start {
			cur := nums[i] + nums[start] + nums[end]
			res = closer(res, cur, target)
			if cur == target {
				return target
			}
			if cur <= target {
				start++
			} else {
				end--
			}
		}
	}
	return res
}

// Returns which number is closer
func closer(a, b, target int) int {
	ad := target - a
	bd := target - b
	if 0 > ad {
		ad *= -1
	}
	if 0 > bd {
		bd *= -1
	}
	if ad < bd {
		return a
	}
	return b
}

/*
https://leetcode.com/problems/3sum-closest

Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.
*/
