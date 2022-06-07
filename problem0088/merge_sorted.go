package problem0088

/*
https://leetcode.com/problems/merge-sorted-array/

You are given two integer arrays nums1 and nums2,
sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	var curM, curN = m - 1, n - 1
	for curM+curN >= -1 {
		if curM == -1 {
			nums1[curN] = nums2[curN]
			curN--
		} else if curN == -1 {
			curM--
		} else if nums1[curM] > nums2[curN] {
			nums1[curM+curN+1] = nums1[curM]
			curM--
		} else {
			nums1[curM+curN+1] = nums2[curN]
			curN--
		}
	}
}
