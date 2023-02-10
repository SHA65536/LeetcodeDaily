package problem1413

/*
Given an array of integers nums, you start with an initial positive value startValue.
In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
Return the minimum positive value of startValue such that the step by step sum is never less than 1.
*/

func minStartValue(nums []int) int {
	var min, sum int
	// Finding the minimal sum for all the steps
	for _, n := range nums {
		sum += n
		if sum < min {
			min = sum
		}
	}
	// Returning the opposite to make sure it's 1 at the lowest step
	return 1 - min
}
