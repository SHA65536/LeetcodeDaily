package problem0046

/*
Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.
*/

func permute(nums []int) [][]int {
	var res [][]int

	permuteHelper(nums, 0, &res)

	return res
}

func permuteHelper(nums []int, start int, res *[][]int) {
	if start >= len(nums) {
		*res = append(*res, copySlice(nums))
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permuteHelper(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func copySlice(in []int) []int {
	res := make([]int, len(in))
	copy(res, in)
	return res
}

func permuteAlloc(nums []int) [][]int {
	var res = make([][]int, 0, factorial(len(nums)))

	permuteHelper(nums, 0, &res)

	return res
}

func factorial(n int) int {
	var res int = 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}
