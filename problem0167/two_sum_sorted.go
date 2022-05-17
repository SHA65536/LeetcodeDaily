package problem0167

/*
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.
*/

func twoSum(numbers []int, target int) []int {
	var start, end, val int
	end = len(numbers) - 1
	for end > start {
		val = numbers[start] + numbers[end]
		if val == target {
			return []int{start + 1, end + 1}
		}
		if val > target {
			end--
		} else {
			start++
		}
	}
	return []int{}
}
