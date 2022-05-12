package problem0047

import "fmt"

/*
https://leetcode.com/problems/permutations-ii

Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
*/

type void interface{}

var member void

func permuteUnique(nums []int) [][]int {
	var res = [][]int{}
	var seen = map[string]void{}
	var sign string
	for _, perm := range permute(nums) {
		sign = fmt.Sprint(perm)
		if _, ok := seen[sign]; !ok {
			seen[sign] = member
			res = append(res, perm)
		}
	}
	return res
}

func permute(nums []int) [][]int {
	var length int = len(nums)
	var res = [][]int{}
	var others []int
	if length == 1 {
		return [][]int{{nums[0]}}
	}
	for i, num := range nums {
		others = make([]int, length)
		copy(others, nums)
		others[i] = others[length-1]
		others = others[:length-1]
		for _, perm := range permute(others) {
			cur := append([]int{num}, perm...)
			res = append(res, cur)
		}
	}
	return res
}
