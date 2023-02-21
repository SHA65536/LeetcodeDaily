package problem0540

/*
You are given a sorted array consisting of only integers where every element appears exactly twice,
except for one element which appears exactly once.
Return the single element that appears only once.
Your solution must run in O(log n) time and O(1) space.
*/

func singleNonDuplicate(nums []int) int {
	var start, end, mid int
	// Basically the pairs left of the pivot start at
	// even indexes, and to the right of the pivot they start
	// at odd indexes
	end = len(nums) - 1
	// Binary search till found
	for end > start {
		// Find middle
		mid = (start + end) / 2
		if mid%2 == 0 {
			// If we're in an even position, check the element ahead
			if nums[mid] == nums[mid+1] {
				// If it's the same, the pivot is to the right
				start = mid + 1
			} else {
				// If it's different, the pivot is to the left
				end = mid
			}
		} else {
			// If we're in an odd position, check the element before
			if nums[mid] == nums[mid-1] {
				// If it's the same, the pivot is to the right
				start = mid + 1
			} else {
				// If it's different, the pivot is to the left
				end = mid
			}
		}
	}
	return nums[start]
}
