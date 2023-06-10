package problem1802

/*
You are given three positive integers: n, index, and maxSum.
You want to construct an array nums (0-indexed) that satisfies the following conditions:
    nums.length == n
    nums[i] is a positive integer where 0 <= i < n.
    abs(nums[i] - nums[i+1]) <= 1 where 0 <= i < n-1.
    The sum of all the elements of nums does not exceed maxSum.
    nums[index] is maximized.
Return nums[index] of the constructed array.
Note that abs(x) equals x if x >= 0, and -x otherwise.
*/

func maxValue(n, idx, maxSum int) int {
	maxSum -= n
	var low, high = 0, maxSum
	for low < high {
		mid := (low + high + 1) / 2
		if testMax(n, idx, mid) <= maxSum {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low + 1
}

func testMax(n, idx, val int) int {
	temp := max(val-idx, 0)
	res := (val + temp) * (val - temp + 1) / 2
	temp = max(val-((n-1)-idx), 0)
	res += (val + temp) * (val - temp + 1) / 2
	return res - val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
