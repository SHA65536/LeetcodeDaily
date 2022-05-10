package problem0015

import (
	"fmt"
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
	var dupes = map[string]void{}
	var seenOne = map[int]void{}
	var res = [][]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := seenOne[nums[i]]; !ok {
			seenOne[nums[i]] = member
		} else {
			continue
		}
		seenTwo := map[int]void{}
		target := -1 * nums[i]
		for _, num := range nums[i+1:] {
			compliment := target - num
			if _, ok := seenTwo[compliment]; ok {
				trio := []int{nums[i], num, compliment}
				hash := hashNums(trio)
				if _, ok = dupes[hash]; !ok {
					dupes[hash] = member
					res = append(res, trio)
				}
			} else {
				seenTwo[num] = member
			}
		}
	}
	return res
}

func hashNums(nums []int) string {
	sort.Ints(nums)
	return fmt.Sprint(nums)
}
