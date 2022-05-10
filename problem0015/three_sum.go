package problem0015

import (
	"sort"
)

/*
https://leetcode.com/problems/3sum/

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

type void struct{}

var member void

func threeSum(nums []int) [][]int {
	var res = [][]int{}
	var seen = map[[3]int]void{}
	var cur [3]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// Skipping duplicate starting numbers
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		start := i + 1
		end := len(nums) - 1
		comp := nums[i] * -1
		for end > start {
			if nums[start]+nums[end] == comp {
				cur = [3]int{nums[i], nums[start], nums[end]}
				if _, ok := seen[cur]; !ok {
					seen[cur] = member
					res = append(res, []int{nums[i], nums[start], nums[end]})
					end = len(nums) - 1
				}
				start++
			} else if nums[start]+nums[end] >= comp {
				end--
			} else {
				start++
			}
		}
	}
	return res
}
