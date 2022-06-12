package problem1695

/*
https://leetcode.com/problems/maximum-erasure-value

You are given an array of positive integers nums and want to erase a subarray containing unique elements.
The score you get by erasing the subarray is equal to the sum of its elements.

Return the maximum score you can get by erasing exactly one subarray.

An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).
*/

func maximumUniqueSubarray(nums []int) int {
	var maxSum, curSum int
	var curUnique = map[int]bool{}
	for start, end := 0, 0; end < len(nums); end++ {
		for curUnique[nums[end]] {
			curSum -= nums[start]
			delete(curUnique, nums[start])
			start++
		}
		curSum += nums[end]
		curUnique[nums[end]] = true
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
