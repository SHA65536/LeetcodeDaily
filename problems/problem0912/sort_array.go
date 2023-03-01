package problem0912

/*
Given an array of integers nums, sort the array in ascending order and return it.
You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.
*/

// This is a merge sort
func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	// Split the array in 2
	a, b := nums[:len(nums)/2], nums[len(nums)/2:]
	// Sort each array recursively
	a = sortArray(a)
	b = sortArray(b)
	// Merge both arrays together
	return mergeArray(a, b)
}

// Merge two sorted arrays together
func mergeArray(a, b []int) []int {
	var res = make([]int, 0, len(a)+len(b))
	// Loop through both arrays
	var aidx, bidx int
	for aidx < len(a) && bidx < len(b) {
		// Pick the lower array
		if a[aidx] < b[bidx] {
			// Add to res
			res = append(res, a[aidx])
			aidx++
		} else {
			res = append(res, b[bidx])
			bidx++
		}
	}
	// Take the slack from the array if there is any
	for aidx < len(a) {
		res = append(res, a[aidx])
		aidx++
	}
	for bidx < len(b) {
		res = append(res, b[bidx])
		bidx++
	}
	return res
}
