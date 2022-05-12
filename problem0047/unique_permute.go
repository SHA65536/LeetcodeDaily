package problem0047

/*
https://leetcode.com/problems/permutations-ii

Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
*/

type void interface{}

var member void

func permuteUnique(nums []int) [][]int {
	var length int = len(nums)
	var res = [][]int{}
	var seen = map[int]void{}
	if length == 1 {
		return [][]int{{nums[0]}}
	}
	for i, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = member
		} else {
			continue
		}
		others := append(nums[i+1:], nums[:i]...)
		for _, perm := range permuteUnique(others) {
			cur := append([]int{num}, perm...)
			res = append(res, cur)
		}
	}
	return res
}
