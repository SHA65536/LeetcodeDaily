package problem0456

/*
Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i],
nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise, return false.
*/

func find132patternNaive(nums []int) bool {
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] >= nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[k] && nums[k] < nums[j] {
					return true
				}
			}
		}
	}
	return false
}

const MINVAL = -1000000001

func find132pattern(nums []int) bool {
	var kNum int = MINVAL
	var numStack = make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < kNum {
			return true
		}
		for len(numStack) > 0 && nums[i] > numStack[len(numStack)-1] {
			kNum = numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
		}
		numStack = append(numStack, nums[i])
	}
	return false
}
