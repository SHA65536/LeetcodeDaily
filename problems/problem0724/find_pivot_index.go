package problem0724

/*
Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of
the index is equal to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left.
This also applies to the right edge of the array.
Return the leftmost pivot index. If no such index exists, return -1.
*/

func pivotIndex(nums []int) int {
	var rightSum, leftSum int
	// Finding right side sum
	for i := range nums {
		rightSum += nums[i]
	}
	// Finding pivot
	for i := range nums {
		rightSum -= nums[i] // Finding sums without current number
		if leftSum == rightSum {
			return i
		}
		// Moving current number to the left
		leftSum += nums[i]
	}
	return -1
}
