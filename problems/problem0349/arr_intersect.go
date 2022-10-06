package problem0349

/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.
*/

func intersection(nums1 []int, nums2 []int) []int {
	var set = map[int]bool{}
	var res = []int{}
	for i := range nums1 {
		set[nums1[i]] = true
	}
	for i := range nums2 {
		if set[nums2[i]] {
			res = append(res, nums2[i])
			delete(set, nums2[i])
		}
	}
	return res
}
