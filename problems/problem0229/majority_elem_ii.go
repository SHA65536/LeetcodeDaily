package problem0229

/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
*/

func majorityElement(nums []int) []int {
	var freq = map[int]int{}
	var res []int
	var maj int = len(nums) / 3
	for i := range nums {
		freq[nums[i]]++
	}
	for k, v := range freq {
		if v > maj {
			res = append(res, k)
		}
	}
	return res
}
