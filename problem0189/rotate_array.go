package problem0189

/*
https://leetcode.com/problems/rotate-array/

Given an array, rotate the array to the right by k steps, where k is non-negative.
*/
func rotate(nums []int, k int) {
	var temp = make([]int, len(nums))
	copy(temp, nums)
	k = k % len(nums)
	if k == 0 {
		return
	}
	for idx := 0; idx < len(nums); idx++ {
		nums[(idx+k)%len(nums)] = temp[idx]
	}
}
