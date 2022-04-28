package problem0004

/*
https://leetcode.com/problems/median-of-two-sorted-arrays/

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var idx1, idx2, half, last1, last2 int
	even := (len(nums1)+len(nums2))%2 == 0
	half = (len(nums1) + len(nums2)) / 2
	for idx1+idx2 <= half {
		last1 = last2
		if idx1 >= len(nums1) {
			last2 = nums2[idx2]
			idx2++
		} else if idx2 >= len(nums2) {
			last2 = nums1[idx1]
			idx1++
		} else {
			if nums1[idx1] > nums2[idx2] {
				last2 = nums2[idx2]
				idx2++
			} else {
				last2 = nums1[idx1]
				idx1++
			}
		}
	}
	if even {
		return float64(last1+last2) / 2
	}
	return float64(last2)
}
