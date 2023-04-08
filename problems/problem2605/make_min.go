package problem2605

/*
Given two arrays of unique digits nums1 and nums2, return the smallest number that contains at least one digit from each array.
*/

func minNumber(nums1 []int, nums2 []int) int {
	var seen = map[int]bool{}
	var res = 10
	for _, v := range nums1 {
		seen[v] = true
	}
	for _, v := range nums2 {
		if seen[v] && v < res {
			res = v
		}
	}
	if res != 10 {
		return res
	}
	a, b := min(nums1), min(nums2)
	if a > b {
		return a + (b * 10)
	} else {
		return b + (a * 10)
	}
}

func min(nums []int) int {
	var res = nums[0]
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}
