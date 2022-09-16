package problem0018

import "sort"

/*
 https://leetcode.com/problems/4sum/

Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.
*/

type void struct{}

var member void

func fourSum(nums []int, target int) [][]int {
	var res = [][]int{}
	var seen = map[[4]int]void{}
	var cur [4]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			start := j + 1
			end := len(nums) - 1
			comp := target - (nums[i] + nums[j])
			for end > start {
				if nums[start]+nums[end] == comp {
					cur = [4]int{nums[i], nums[j], nums[start], nums[end]}
					if _, ok := seen[cur]; !ok {
						seen[cur] = member
						res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
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
	}
	return res
}
