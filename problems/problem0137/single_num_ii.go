package problem0137

/*
Given an integer array nums where every element appears three times except for one, which appears exactly once.
Find the single element and return it.
You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

func singleNumber(nums []int) int {
	var freq = map[int]int{}
	for i := range nums {
		freq[nums[i]]++
	}
	for k, v := range freq {
		if v == 1 {
			return k
		}
	}
	return -1
}
