package problem2553

/*
Given an array of positive integers nums,
return an array answer that consists of the digits of each integer in nums after separating them in the same order they appear in nums.
To separate the digits of an integer is to get all the digits it has in the same order.
For example, for the integer 10921, the separation of its digits is [1,0,9,2,1].
*/

func separateDigits(nums []int) []int {
	var res = make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		var n = nums[i]
		for n > 0 {
			res = append(res, n%10)
			n /= 10
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
