package problem2366

/*
You are given a 0-indexed integer array nums.
In one operation you can replace any element of the array with any two elements that sum to it.
For example, consider nums = [5,6,7]. In one operation, we can replace nums[1] with 2 and 4 and convert nums to [5,2,4,7].
Return the minimum number of operations to make an array that is sorted in non-decreasing order.
*/

func minimumReplacement(nums []int) int64 {
	var res int64

	// Loop from the end to the start
	var prev = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		// splitInto the the number of times we need to split
		// nums[i] to make an ascending order
		splitInto := nums[i] / prev

		// If it cannot be split perfectly, we need an additional split
		if nums[i]%prev != 0 {
			splitInto++
		}

		// Update the previous element (remember that we split it)
		prev = nums[i] / splitInto

		res += int64(splitInto - 1)
	}
	return res
}
