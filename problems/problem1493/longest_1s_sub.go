package problem1493

/*
Given a binary array nums, you should delete one element from it.
Return the size of the longest non-empty subarray containing only 1's in the resulting array.
Return 0 if there is no such subarray.
*/

func longestSubarray(nums []int) int {
	var res, cur, prev int
	var hasZero bool
	for i := range nums {
		if nums[i] == 1 {
			cur++
			res = max(res, cur+prev)
		} else {
			prev = cur
			cur = 0
			hasZero = true
		}
	}
	if !hasZero {
		return len(nums) - 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
