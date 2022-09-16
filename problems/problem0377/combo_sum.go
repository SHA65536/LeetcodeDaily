package problem0377

/*
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.
The test cases are generated so that the answer can fit in a 32-bit integer.
*/

func combinationSum4Naive(nums []int, target int) int {
	var res int
	if target == 0 {
		return 1
	}
	for _, num := range nums {
		if target >= num {
			res += combinationSum4Naive(nums, target-num)
		}
	}
	return res
}

func combinationSum4(nums []int, target int) int {
	return solveCache(make(map[int]int), nums, target)
}

func solveCache(cache map[int]int, nums []int, target int) int {
	var res int
	if target == 0 {
		return 1
	}
	if val, ok := cache[target]; ok {
		return val
	}
	for _, num := range nums {
		if target >= num {
			res += solveCache(cache, nums, target-num)
		}
	}
	cache[target] = res
	return res
}
