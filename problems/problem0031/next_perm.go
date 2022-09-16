package problem0031

import "sort"

/*
A permutation of an array of integers is an arrangement of its members into a sequence or linear order.
For example, for arr = [1,2,3], the following are considered permutations of arr: [1,2,3], [1,3,2], [3,1,2], [2,3,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer.
More formally, if all the permutations of the array are sorted in one container according to their lexicographical order,
then the next permutation of that array is the permutation that follows it in the sorted container.
If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).
For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.
The replacement must be in place and use only constant extra memory.
*/

func nextPermutation(nums []int) {
	var switchPos, minPos int
	minPos = -1
	// Identifying where the switch needs to happen
	for switchPos = len(nums) - 1; switchPos > 0; switchPos-- {
		if nums[switchPos] > nums[switchPos-1] {
			break
		}
	}
	switchPos--
	// If no switch, means we are at the last permutation
	if switchPos != -1 {
		// Finding the minimal number that is still larger than the switching number
		for i := len(nums) - 1; i > switchPos; i-- {
			if nums[i] > nums[switchPos] && (minPos == -1 || nums[i] < nums[minPos]) {
				minPos = i
			}
		}
		// Switching the numbers
		nums[switchPos], nums[minPos] = nums[minPos], nums[switchPos]
	}
	// Sorting the rest of the array after the switch
	sort.Ints(nums[switchPos+1:])
}
