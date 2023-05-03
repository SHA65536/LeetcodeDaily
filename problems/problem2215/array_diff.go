package problem2215

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
    answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
    answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.
*/

func findDifference(nums1 []int, nums2 []int) [][]int {
	var res = make([][]int, 2)
	var set1 = map[int]struct{}{}
	var set2 = map[int]struct{}{}

	for i := range nums1 {
		set1[nums1[i]] = struct{}{}
	}

	for i := range nums2 {
		set2[nums2[i]] = struct{}{}
	}

	for k := range set1 {
		if _, ok := set2[k]; !ok {
			res[0] = append(res[0], k)
		}
	}

	for k := range set2 {
		if _, ok := set1[k]; !ok {
			res[1] = append(res[1], k)
		}
	}

	return res
}
