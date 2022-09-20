package problem0718

/*
Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.
*/

func findLength(nums1 []int, nums2 []int) int {
	var res int
	var cache [][]int = make([][]int, len(nums1)+1)
	for i := range cache {
		cache[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				cache[i][j] = 0
			}
			if cache[i][j] > res {
				res = cache[i][j]
			}
		}
	}
	return res
}
