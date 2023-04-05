package problem2439

/*
You are given a 0-indexed array nums comprising of n non-negative integers.
In one operation, you must:
    Choose an integer i such that 1 <= i < n and nums[i] > 0.
    Decrease nums[i] by 1.
    Increase nums[i - 1] by 1.
Return the minimum possible value of the maximum integer of nums after performing any number of operations.
*/

func minimizeArrayValue(nums []int) int {
	var sum, res, i int64
	for i = 0; i < int64(len(nums)); i++ {
		sum += int64(nums[i])
		res = max(res, (sum+i)/(i+1))
	}
	return int(res)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
